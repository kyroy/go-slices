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
	"bytes"
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

type testDef struct {
	Timestamp time.Time
	Package   string
	Type      string
	Map       []struct {
		Name string
		In   string
		Func string
		Out  string
	}
	Filter []struct {
		Name string
		In   string
		Func string
		Out  string
	}
	Reduce []struct {
		Name    string
		In      string
		Func    string
		Neutral string
		Out     string
	}
	Unique []struct {
		Name string
		In   string
		Out  string
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("wrong args", os.Args)
		os.Exit(1)
	}
	typeDefs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("failed to read types: %v\n", err)
		os.Exit(1)
	}
	testDefs, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Printf("failed to read tests: %v\n", err)
		os.Exit(1)
	}

	var defs []sliceDef
	decoder := yaml.NewDecoder(bytes.NewReader(typeDefs))
	decoder.SetStrict(true)
	if err := decoder.Decode(&defs); err != nil {
		fmt.Printf("failed to decode definitions: %v\n", err)
		os.Exit(1)
	}
	var tests map[string]testDef
	decoder = yaml.NewDecoder(bytes.NewReader(testDefs))
	decoder.SetStrict(true)
	if err := decoder.Decode(&tests); err != nil {
		fmt.Printf("failed to decode tests: %v\n", err)
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
	for pkg, d := range tests {
		d.Timestamp = t
		d.Package = pkg
		if err := generateTest(pkg, d); err != nil {
			fmt.Printf("failed to generate tests %q: %v\n", pkg, err)
			os.Exit(1)
		}
	}
}

func generateType(d sliceDef) error {
	if err := os.MkdirAll(d.Package, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create dir %q: %v", d.Package, err)
	}
	// type.go
	f, err := os.Create(path.Join(d.Package, "type.go"))
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()
	return internal.TypeTemplate.Execute(f, d)
}

func generateTest(pkg string, d testDef) error {
	// package_test.go
	t, err := os.Create(path.Join(pkg, pkg+"_test.go"))
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer t.Close()
	return internal.PackageTestTemplate.Execute(t, d)
}
