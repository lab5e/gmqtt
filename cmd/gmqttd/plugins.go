//go:generate sh -c "cd ../../ && go run plugin_generate.go"
// generated by plugin_generate.go; DO NOT EDIT

package main

import (
	_ "github.com/lab5e/gmqtt/plugin/admin"
	_ "github.com/lab5e/gmqtt/plugin/auth"
	_ "github.com/lab5e/gmqtt/plugin/federation"
	_ "github.com/lab5e/gmqtt/plugin/prometheus"
)
