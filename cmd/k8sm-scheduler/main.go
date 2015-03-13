/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// apiserver is the main api server and master for the cluster.
// it is responsible for serving the cluster management API.
package main

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
	"github.com/mesosphere/kubernetes-mesos/pkg/scheduler/service"
	"github.com/mesosphere/kubernetes-mesos/pkg/version/verflag"
	"github.com/spf13/pflag"
)

func main() {
	s := service.NewSchedulerServer()
	s.AddFlags(pflag.CommandLine)

	util.InitFlags()
	util.InitLogs()
	defer util.FlushLogs()

	verflag.PrintAndExitIfRequested()

	s.Run(nil, pflag.CommandLine.Args())
}
