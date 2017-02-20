// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package entry implements the GRR client entry point.
package entry

import (
	"log"

	"github.com/google/grr_go_client/client"
)

// Flags is a container for command line flags.
type Flags struct {
	// TODO: Handle the ConfigPath flag.
	ConfigPath *string
}

// Entry is the GRR client entry point.
func Entry(*Flags) {
	cl, err := client.New()
	if err != nil {
		log.Fatalf("Unable to initialize the client: %v", err)
	}

	if err := cl.Run(); err != nil {
		log.Fatalf("The client errored out: %v", err)
	}
}
