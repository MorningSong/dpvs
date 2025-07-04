// /*
// Copyright 2025 IQiYi Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

package checker

import (
	"net"
	"testing"
	"time"

	"github.com/iqiyi/dpvs/tools/healthcheck/pkg/utils"
)

var http_targets = []utils.L3L4Addr{
	{net.ParseIP("192.168.88.30"), 80, utils.IPProtoTCP},
	{net.ParseIP("192.168.88.30"), 443, utils.IPProtoTCP},
	{net.ParseIP("2001::30"), 80, utils.IPProtoTCP},
	{net.ParseIP("2001::30"), 443, utils.IPProtoTCP},

	// control group of proxy protocol
	{net.ParseIP("192.168.88.30"), 8002, utils.IPProtoTCP},
	{net.ParseIP("2001::30"), 8002, utils.IPProtoTCP},
}

var http_proxy_proto_targets = []utils.L3L4Addr{
	{net.ParseIP("192.168.88.30"), 8002, utils.IPProtoTCP},
	{net.ParseIP("2001::30"), 8002, utils.IPProtoTCP},
}

var http_url_targets = []string{
	"http://www.baidu.com",
	"https://www.baidu.com",
	"http://www.iqiyi.com",
	"https://www.iqiyi.com",
	"https://www.google.com",
}

func TestHttpChecker(t *testing.T) {
	timeout := 2 * time.Second

	for _, target := range http_targets {
		params := map[string]string{
			"response-codes": "200-299",
		}
		checker, err := (&HTTPChecker{}).create(params)
		if err != nil {
			t.Fatalf("Failed to create http checker %v: %v", target, err)
		}

		state, err := checker.Check(&target, timeout)
		if err != nil {
			t.Errorf("Failed to execute http checker %v: %v", target, err)
		} else {
			t.Logf("[ HTTP ] %v ==> %v", target, state)
		}
	}

	for _, target := range http_proxy_proto_targets {
		params := map[string]string{
			"response-codes": "200-299",
			"proxy-protocol": "v1",
		}
		checker, err := (&HTTPChecker{}).create(params)
		if err != nil {
			t.Fatalf("Failed to create http checker %v: %v", target, err)
		}

		// Proxy Protocol v1 tests
		state, err := checker.Check(&target, timeout)
		if err != nil {
			t.Errorf("Failed to execute http checker %v: %v", target, err)
		} else {
			t.Logf("[ HTTP(PPv1) ] %v ==> %v", target, state)
		}

		// Proxy Protocol v2 tests
		params["proxy-protocol"] = "v2"
		checker, err = (&HTTPChecker{}).create(params)
		if err != nil {
			t.Fatalf("Failed to create http checker %v: %v", target, err)
		}

		state, err = checker.Check(&target, timeout)
		if err != nil {
			t.Errorf("Failed to execute http checker %v: %v", target, err)
		} else {
			t.Logf("[ HTTP(PPv2) ] %v ==> %v", target, state)
		}
	}

	for _, target := range http_url_targets {
		params := map[string]string{
			"uri":            target,
			"response-codes": "200",
		}
		checker, err := (&HTTPChecker{}).create(params)
		if err != nil {
			t.Fatalf("Failed to create http checker %v: %v", target, err)
		}

		state, err := checker.Check(&utils.L3L4Addr{}, timeout)
		if err != nil {
			t.Errorf("Failed to execute http checker %v: %v", target, err)
		} else {
			t.Logf("[ HTTP ] %v ==> %v", target, state)
		}
	}
}
