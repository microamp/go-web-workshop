// Copyright 2016 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	baseURL := "https://http-methods.appspot.com/Hungary/"
	client := http.DefaultClient

	v := url.Values{}
	v.Set("v", "true")

	fmt.Println(fmt.Sprintf("%s%s", baseURL, v.Encode()))

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", baseURL, v.Encode()), nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	log.Printf(res.Status)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("reading response body failed: %v", err)
	}

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		log.Printf(line)
	}
}
