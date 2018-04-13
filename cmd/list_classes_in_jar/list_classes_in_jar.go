// Copyright 2018 The Jadep Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Program list_classes_in_jar consumes .jar files and prints a sorted list of Java classes they contains to stdout.
// Jade consumes the result in order to avoid looking for dependencies for built-in class names.
//
// The program takes the .jar file names on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/bazelbuild/tools_jvm_autodeps/jadep/jadeplib"
	"github.com/bazelbuild/tools_jvm_autodeps/jadep/listclassesinjar"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s jar1 ...", os.Args[0])
	}

	var classes []string
	seen := make(map[jadeplib.ClassName]bool)
	for i := 1; i < len(os.Args); i++ {
		cls, err := listclassesinjar.List(os.Args[i])
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range cls {
			if !seen[c] {
				seen[c] = true
				classes = append(classes, string(c))
			}
		}
	}
	sort.Strings(classes)
	for _, p := range classes {
		fmt.Println(p)
	}
}
