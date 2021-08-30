package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/lab5e/gmqtt/cmd/gmqttd/command"
	_ "github.com/lab5e/gmqtt/persistence"
	_ "github.com/lab5e/gmqtt/plugin/prometheus"
	_ "github.com/lab5e/gmqtt/topicalias/fifo"
)

var (
	rootCmd = &cobra.Command{
		Use:     "gmqttd",
		Long:    "Gmqtt is a MQTT broker that fully implements MQTT V5.0 and V3.1.1 protocol",
		Version: Version,
	}
	enablePprof bool
	pprofAddr   = "127.0.0.1:6060"
)

func must(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func init() {
	configDir, err := getDefaultConfigDir()
	must(err)
	command.ConfigFile = path.Join(configDir, "gmqttd.yml")
	rootCmd.PersistentFlags().StringVarP(&command.ConfigFile, "config", "c", command.ConfigFile, "The configuration file path")
	rootCmd.AddCommand(command.NewStartCmd())
	//rootCmd.AddCommand(command.NewReloadCommand())
}

func main() {
	if enablePprof {
		go func() {
			http.ListenAndServe(pprofAddr, nil)
		}()
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
