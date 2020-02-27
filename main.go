// Copyright (c) Jargon Author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import "github.com/imdevin567/jargon/cmd"

// These values will be injected into these variables at the build time.
var Version, GitCommit string

func main() {
	if err := cmd.Execute(Version, GitCommit); err != nil {
		panic(err)
	}
}
