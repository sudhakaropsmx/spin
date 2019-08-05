// Copyright (c) 2019, OpsMx, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/sudhakaropsmx/spinmx/cmd"
	"github.com/spinnaker/spin/util"
)

func main() {
	if err := cmd.Execute(os.Stdout); err != nil {
		if util.UI != nil {
			util.UI.Error(err.Error())
		} else {
			fmt.Fprintf(os.Stderr, "\n%v\n", err)
		}
		os.Exit(1)
	}
}