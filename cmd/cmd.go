package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wangyi12358/ratelctl/cmd/initial"
	"github.com/wangyi12358/ratelctl/pkg/version"
	"os"
	"runtime"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "ratelctl",
		Short: "ratelctl 用于快速生成项目模板",
	}

	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s", version.BuildVersion,
		runtime.GOOS, runtime.GOARCH)
	rootCmd.AddCommand(initial.Cmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
