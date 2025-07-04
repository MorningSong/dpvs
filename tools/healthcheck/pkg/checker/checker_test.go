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
	"flag"
	"os"
	"testing"

	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	// To support test args, run test with params like:
	// `go test -v . -args -logtostderr=true -v=9`
	// or test a specific method ping like:
	// `go test -v -run TestPingChecker --args -logtostderr=true -v=9`
	flag.Parse()

	rc := m.Run()
	glog.Flush()
	os.Exit(rc)
}
