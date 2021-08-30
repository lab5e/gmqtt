// +build ignore

package main

import (
	"bytes"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var tmpl = `//go:generate sh -c "cd ../../ && go run plugin_generate.go"
// generated by plugin_generate.go; DO NOT EDIT

package main

import (
    {{- range $index, $element := .}}
    _ "{{$element}}"
    {{- end}}
)
`

const (
	pluginFile = "./cmd/gmqttd/plugins.go"
	pluginCfg  = "plugin_imports.yml"
	importPath = "github.com/lab5e/gmqtt/plugin"
)

type ymlCfg struct {
	Packages []string `yaml:"packages"`
}

func main() {
	b, err := ioutil.ReadFile(pluginCfg)
	if err != nil {
		log.Fatalf("ReadFile error %s", err)
		return
	}

	var cfg ymlCfg
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatalf("Unmarshal error: %s", err)
		return
	}
	t, err := template.New("plugin_gen").Parse(tmpl)
	if err != nil {
		log.Fatalf("Parse template error: %s", err)
		return
	}

	for k, v := range cfg.Packages {
		if !strings.Contains(v, "/") {
			cfg.Packages[k] = importPath + "/" + v
		}
	}

	if err != nil && err != io.EOF {
		log.Fatalf("read error: %s", err)
		return
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, cfg.Packages)
	if err != nil {
		log.Fatalf("excute template error: %s", err)
		return
	}
	rs, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("format error: %s", err)
		return
	}
	err = ioutil.WriteFile(pluginFile, rs, 0666)
	if err != nil {
		log.Fatalf("writeFile error: %s", err)
		return
	}
	return
}
