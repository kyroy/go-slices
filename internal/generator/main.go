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
	"github.com/kyroy/go-slices/internal/generator/templates"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type definition struct {
	Types    []sliceDef
	Examples map[string]exampleDef
	Tests    *testDef
}

type sliceDef struct {
	Package string
	Type    string
}

type exampleDef struct {
	In  string
	Out string
}

type testDef struct {
	Map []struct {
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
	Intersect []struct {
		Name string
		In   string
		More string
		Out  string
	}
	Contains []struct {
		Name string
		In   string
		Elem string
		Out  string
	}
	IndexOf []struct {
		Name string
		In   string
		Elem string
		Out  string
	}
	Find []struct {
		Name  string
		In    string
		Func  string
		Out   string
		Found bool
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("wrong args", os.Args)
		os.Exit(1)
	}

	defs, err := definitionFromFolder(os.Args[1])
	if err != nil {
		fmt.Printf("failed to read types: %v\n", err)
		os.Exit(1)
	}

	for _, d := range defs {
		if err := generateType(d); err != nil {
			fmt.Printf("failed to generate type: %v\n", err)
			os.Exit(1)
		}
		if err := generateExamples(d); err != nil {
			fmt.Printf("failed to generate examples: %v\n", err)
			os.Exit(1)
		}
		if err := generateTests(d); err != nil {
			fmt.Printf("failed to generate tests: %v\n", err)
			os.Exit(1)
		}
	}
}

func definitionFromFolder(dir string) ([]definition, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("could not open directory: %v", err)
	}
	var defs []definition
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		f, err := os.Open(path.Join(dir, f.Name()))
		if err != nil {
			return nil, fmt.Errorf("could not open file %q: %v", f.Name(), err)
		}
		var def definition
		decoder := yaml.NewDecoder(f)
		decoder.SetStrict(true)
		if err := decoder.Decode(&def); err != nil {
			return nil, fmt.Errorf("failed to decode definition %q: %v", f.Name(), err)
		}
		defs = append(defs, def)
	}
	return defs, nil
}

func generateType(d definition) error {
	for _, t := range d.Types {
		if err := os.MkdirAll(t.Package, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create dir %q: %v", t.Package, err)
		}
		// type.go
		f, err := os.Create(path.Join(t.Package, "type.go"))
		if err != nil {
			return fmt.Errorf("could not open file: %v", err)
		}
		defer f.Close()
		if err := internal.TypeTemplate.Execute(f, t); err != nil {
			return fmt.Errorf("failed to generate type %q: %v", t.Type, err)
		}
	}
	return nil
}

func generateExamples(d definition) error {
	// examples_test.go
	for _, t := range d.Types {
		if len(d.Examples) == 0 {
			continue
		}
		f, err := os.Create(path.Join(t.Package, "examples_test.go"))
		if err != nil {
			return fmt.Errorf("could not open file: %v", err)
		}
		defer f.Close()
		old := make(map[string]string)
		for fn, test := range d.Examples {
			old[fn] = test.In
			test.In = strings.Replace(test.In, "<TYPE>", t.Type, -1)
			d.Examples[fn] = test
		}
		x := struct {
			Package  string
			Type     string
			Examples map[string]exampleDef
		}{
			Package:  t.Package,
			Type:     t.Type,
			Examples: d.Examples,
		}
		if err := internal.ExampleTemplate.Execute(f, x); err != nil {
			return fmt.Errorf("failed to generate examples %q: %v", t.Type, err)
		}
		for fn, test := range d.Examples {
			test.In = old[fn]
			d.Examples[fn] = test
		}
	}
	return nil
}

func generateTests(d definition) error {
	// package_test.go
	for _, t := range d.Types {
		if d.Tests == nil {
			continue
		}
		f, err := os.Create(path.Join(t.Package, t.Package+"_test.go"))
		if err != nil {
			return fmt.Errorf("could not open file: %v", err)
		}
		defer f.Close()
		x := struct {
			Package string
			Type    string
			Tests   *testDef
		}{
			Package: t.Package,
			Type:    t.Type,
			Tests:   d.Tests,
		}
		if err := internal.PackageTestTemplate.Execute(f, x); err != nil {
			return fmt.Errorf("failed to generate tests %q: %v", t.Type, err)
		}
	}
	return nil
}
