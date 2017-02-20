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

// Package client implements the GRR client class.
package client

import (
	"errors"
)

// Client is the GRR client class.
type Client struct{}

// New instantiates Client.
func New() (*Client, error) {
	return &Client{}, nil
}

// Run is the main run method of the GRR client.
// This method does not return unless there's been an error.
func (c *Client) Run() error {
	return errors.New("not yet implemented") // TODO
}
