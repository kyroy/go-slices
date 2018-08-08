// Copyright 2018 Dennis Kuhnert
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/kyroy/go-slices/generator/internal"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type sliceDef struct {
	Timestamp time.Time
	Package   string `yaml:"package"`
	SliceType string `yaml:"slice_type"`
	Type      string `yaml:"type"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("wrong args", os.Args)
		os.Exit(1)
	}
	types, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("failed to read types: %v\n", err)
		os.Exit(1)
	}

	var defs []sliceDef
	if err := yaml.Unmarshal(types, &defs); err != nil {
		fmt.Printf("failed to decode definitions: %v\n", err)
		os.Exit(1)
	}
	t := time.Now()
	for _, d := range defs {
		d.Timestamp = t
		if err := generateType(d); err != nil {
			fmt.Printf("failed to generate type %q: %v\n", d.Type, err)
			os.Exit(1)
		}
	}
}

func generateType(d sliceDef) error {
	if err := os.MkdirAll(d.Package, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create dir %q: %v", d.Package, err)
	}
	f, err := os.Create(path.Join(d.Package, "type.go"))
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()

	if err = internal.TypeTemplate.Execute(f, d); err != nil {
		return err
	}
	return nil
}