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

package resource

import (
	"bytes"
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

var resource Resource

func captureLogOutput(fn func()) string {
	var buffer bytes.Buffer

	log.SetOutput(&buffer)
	fn()
	log.SetOutput(os.Stderr)

	return buffer.String()
}

func TestValidate(t *testing.T) {
	var data = `
---
apiVersion: v1
kind: Resource
spec:
  containers:
    - name: busybox
      image: busybox
      command: sleep infinity & wait
      env:
        - name: FOO
          value: bar
        - name: FOO_BAR
          value: baz
`

	jsonData, _ := yaml.YAMLToJSON([]byte(data))
	err := resource.validate([]byte(jsonData))
	assert.NoError(t, err)
}

func TestValidateAPIVersionReturnsErrorAndLogsValidationError(t *testing.T) {
	var data = `
---
apiVersion: v2
kind: resource
spec:
  containers:
    - name: 123
      image: 123
      command: 123
      env:
        - name: foo
          value: invalid
        - name: FOO-BAR
          value: invalid
        - name: foo-bar
          value: invalid
        - name: 123_456
          value: invalid
      container_key: invalid
root_key: invalid
`

	logOutput := captureLogOutput(func() {
		jsonData, _ := yaml.YAMLToJSON([]byte(data))
		err := resource.validate([]byte(jsonData))
		assert.Error(t, err)
	})

	msgs := []string{
		"The document is not valid - root_key: Additional property root_key is not allowed.",
		"The document is not valid - apiVersion: Does not match pattern '^v1$'",
		"The document is not valid - kind: Does not match pattern '^Resource$'.",
		"The document is not valid - container_key: Additional property container_key is not allowed.",
		"The document is not valid - spec.containers.0.name: Invalid type. Expected: string, given: integer.",
		"The document is not valid - spec.containers.0.image: Invalid type. Expected: string, given: integer.",
		"The document is not valid - spec.containers.0.command: Invalid type. Expected: string, given: integer.",
		"The document is not valid - spec.containers.0.env.0.name: Does not match pattern '^[A-Z0-9_]*$'.",
		"The document is not valid - spec.containers.0.env.1.name: Does not match pattern '^[A-Z0-9_]*$'.",
		"The document is not valid - spec.containers.0.env.2.name: Does not match pattern '^[A-Z0-9_]*$'.",
		"The document is not valid - spec.containers.0.env.3.name: Invalid type. Expected: string, given: integer.",
	}
	for _, msg := range msgs {
		assert.Contains(t, logOutput, msg)
	}
}
