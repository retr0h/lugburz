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

package cmd

import (
	"fmt"
	"strings"

	"github.com/retr0h/lugburz/resource"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	filename string
	r        resource.Resource
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource by filename",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkRequiredFlags(cmd.Flags())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		err := r.UnmarshalYAMLFile(filename)
		if err != nil {
			panic(err)
		}
		return nil
	},
}

func checkRequiredFlags(flags *pflag.FlagSet) error {
	missingFlagNames := []string{}

	flags.VisitAll(func(flag *pflag.Flag) {
		requiredAnnotation := flag.Annotations[cobra.BashCompOneRequiredFlag]
		if len(requiredAnnotation) == 0 {
			return
		}

		flagRequired := requiredAnnotation[0] == "true"

		if flagRequired && !flag.Changed {
			missingFlagNames = append(missingFlagNames, flag.Name)
		}
	})

	if len(missingFlagNames) > 0 {
		return fmt.Errorf("Required flag/flags: \"%s\" has/have not been set", strings.Join(missingFlagNames, "\", \""))
	}

	return nil
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename to use to create the resource")
	createCmd.MarkFlagRequired("filename")
}
