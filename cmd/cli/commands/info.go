package commands

import (
	"bytes"
	"github.com/dimiro1/banner"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	RootCmd.AddCommand(infoCmd)
}

const message = `
GoVersion: {{ .GoVersion }}
GOOS: {{ .GOOS }}
GOARCH: {{ .GOARCH }}
NumCPU: {{ .NumCPU }}
GOPATH: {{ .GOPATH }}
GOROOT: {{ .GOROOT }}
Compiler: {{ .Compiler }}
ENV: {{ .Env "GOPATH" }}
Now: {{ .Now "Monday, 2 Jan 2006" }}
`

var infoCmd = &cobra.Command{
	Use: "info",
	Run: func(cmd *cobra.Command, args []string) {
		banner.Init(os.Stdout, true, true, bytes.NewBufferString(message))
	},
}
