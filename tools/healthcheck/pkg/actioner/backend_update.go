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

package actioner

/*
BackendAction Actioner Params:
-------------------------------------------------------
name                value
-------------------------------------------------------

-------------------------------------------------------
*/

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/iqiyi/dpvs/tools/healthcheck/pkg/comm"
	"github.com/iqiyi/dpvs/tools/healthcheck/pkg/types"
	"github.com/iqiyi/dpvs/tools/healthcheck/pkg/utils"
)

var _ ActionMethod = (*BackendAction)(nil)

const backendActionerName = "BackendUpdate"

func init() {
	registerMethod(backendActionerName, &BackendAction{})
}

type BackendAction struct {
	name      string
	apiServer string
}

func (a *BackendAction) Act(signal types.State, timeout time.Duration,
	data ...interface{}) (interface{}, error) {
	if timeout <= 0 {
		return nil, fmt.Errorf("zero timeout on actioner %s", a.name)
	}
	if len(data) < 1 {
		return nil, fmt.Errorf("%s missing backend data", a.name)
	}
	vs, ok := data[0].(*comm.VirtualServer)
	if !ok || vs == nil || len(vs.RSs) == 0 {
		return nil, fmt.Errorf("invalid backend data for %s", a.name)
	}

	glog.V(7).Infof("starting %s actioner %s ...", backendActionerName, a.name)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	newVS, err := comm.UpdateCheckState(a.apiServer, vs, ctx)
	if err != nil {
		glog.Errorf("%s actioner %s (VS: %v) failed: %v", backendActionerName, a.name, *vs, err)
	} else if newVS != nil {
		glog.Warningf("%s actioner %s (VS: %v) outdated and returned newVS %v",
			backendActionerName, a.name, *vs, newVS)
	} else {
		glog.V(6).Infof("%s actioner %s (VS %v) succeed", backendActionerName, a.name, *vs)
	}

	return newVS, err
}

func (a *BackendAction) validate(params map[string]string) error {
	unsupported := make([]string, 0, len(params))
	for param, _ := range params {
		switch param {
		default:
			unsupported = append(unsupported, param)
		}
	}
	if len(unsupported) > 0 {
		return fmt.Errorf("unsupported action params: %s", strings.Join(unsupported, ","))
	}
	return nil
}

func (a *BackendAction) create(target *utils.L3L4Addr, params map[string]string,
	extras ...interface{}) (ActionMethod, error) {
	actioner := &BackendAction{name: target.String()}

	if len(extras) > 0 {
		if apiServer, ok := extras[0].(string); ok {
			actioner.apiServer = apiServer
		}
	}

	if len(actioner.apiServer) == 0 {
		return nil, fmt.Errorf("%s actioner misses dpvs api server config", backendActionerName)
	}

	if err := a.validate(params); err != nil {
		return nil, fmt.Errorf("%s actioner param validation failed: %v", backendActionerName, err)
	}

	return actioner, nil
}
