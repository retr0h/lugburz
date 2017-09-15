// Copyright (c) 2017 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

// Package resource TODO
package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

type containersEnvEntry struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type containersEntry struct {
	// Name of the container to start.
	Name string `yaml:"name"`
	// Image name and tag used to create the container.
	Image string `yaml:"image"`
	// Command to execute when the container starts.
	Command string `yaml:"command"`
	// List containing containersEnvEntry struct.
	Env []containersEnvEntry `yaml:"env"`
}

// Resource holds the configuration passed from the --filename flag.
// All fields are required unless otherwise specified
type Resource struct {
	// Lugbúrz API version to use.
	APIVersion string `yaml:"apiVersion"`
	// NOTE(retr0h): Will probably remove this.
	Kind string `yaml:"kind"`
	Spec struct {
		// List containing containersEntry struct.
		Containers []containersEntry `yaml:"containers"`
	} `yaml:"spec"`
}

// UnmarshalYAML decodes the first YAML document found within the data byte
// slice, passes the string through a generic YAML-to-JSON converter, validate,
// provide the resulting JSON to json.Unmarshal, and assigns the decoded values to
// Resource.
// TODO(retr0h): returns the object
func (r *Resource) UnmarshalYAML(data []byte) error {
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}

	err = r.validate(jsonData)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, r)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalYAMLFile reads the file named by filename and passes the source
// data byte slice to UnmarshalYAML for decoding.
func (r *Resource) UnmarshalYAMLFile(filename string) error {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = r.UnmarshalYAML([]byte(source))
	if err != nil {
		return err
	}
	return nil
}

// Validate the the data byte slice against the v1 JSON schema.
func (r *Resource) validate(data []byte) error {
	// Load the JSON schema bindata for validation.
	schema, err := Asset("resource/resource_schema_v1.json")
	if err != nil {
		return err
	}

	schemaLoader := gojsonschema.NewBytesLoader(schema)
	documentLoader := gojsonschema.NewBytesLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			log.Error(fmt.Sprintf("The document is not valid - %s.", desc))
		}
		return errors.New("Invalid YAML provided")
	}
	return nil
}
