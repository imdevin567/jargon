// Copyright (c) Jargon Author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import "github.com/imdevin567/jargon/cmd"

// Version ...
var Version string = "1.0.0"

func main() {
	if err := cmd.Execute(Version); err != nil {
		panic(err)
	}
}
