// Copyright 2014 Brett Slatkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mapWrapper

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

const pointerHeaderTemplate = `// generated by map-wrapper -- DO NOT EDIT
package {{.Package}}

import (
    "encoding/json"
    "reflect"
)
`

var (
	generatedPointerHeaderTemplate = template.Must(template.New("render").Parse(pointerHeaderTemplate))
	generatedTemplate = template.Must(template.New("render").Parse(keyObjectValuePointerTemplate))
)

type GeneratedType struct {
	Key string
	KeyTitle string
	Value string
}

func getRenderedPath(inputPath string) (string, error) {
	if !strings.HasSuffix(inputPath, ".go") {
		return "", fmt.Errorf("Input path %s doesn't have .go extension", inputPath)
	}
	trimmed := strings.TrimSuffix(inputPath, ".go")
	dir, file := filepath.Split(trimmed)
	return filepath.Join(dir, fmt.Sprintf("%s_map.go", file)), nil
}

type generateTemplateData struct {
	Package string
	Types     []GeneratedType
}

func render(w io.Writer, packageName string, types []GeneratedType) error {
	hErr := generatedPointerHeaderTemplate.Execute(w, generateTemplateData{Package: packageName})
	if hErr != nil {
		return hErr
	}

	return generatedTemplate.Execute(w, generateTemplateData{packageName, types})
}
