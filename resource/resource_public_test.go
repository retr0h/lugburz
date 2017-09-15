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

package resource_test

import (
	"path"
	"testing"

	"github.com/retr0h/lugburz/resource"
	"github.com/stretchr/testify/assert"
)

var r resource.Resource

func TestUnmarshalYAML(t *testing.T) {
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
`
	err := r.UnmarshalYAML([]byte(data))
	if assert.NoError(t, err) {
		assert.Equal(t, "v1", r.APIVersion)
		assert.Equal(t, "Resource", r.Kind)
		assert.Equal(t, "busybox", r.Spec.Containers[0].Name)
		assert.Equal(t, "busybox", r.Spec.Containers[0].Image)
		assert.Equal(t, "sleep infinity & wait", r.Spec.Containers[0].Command)
		assert.Equal(t, "FOO", r.Spec.Containers[0].Env[0].Name)
		assert.Equal(t, "bar", r.Spec.Containers[0].Env[0].Value)
	}
}

func TestUnmarshalYAMLReturnsError(t *testing.T) {
	var data = `
	foo: bar
`

	err := r.UnmarshalYAML([]byte(data))
	assert.Error(t, err)
}

func TestUnmarshalYAMLFile(t *testing.T) {
	var filename = path.Join("..", "test", "resource", "resource.yml")

	r.UnmarshalYAMLFile(filename)
	assert.NotNil(t, r.APIVersion)
}

func TestUnmarshalYAMLFileReturnsErrorWithMissingFile(t *testing.T) {
	var filename = "invalid.yml"

	err := r.UnmarshalYAMLFile(filename)
	assert.Error(t, err)
}
