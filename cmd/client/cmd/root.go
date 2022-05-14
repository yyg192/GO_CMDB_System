package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	vers         bool
	ossProvider  string
	aliAccessID  string
	aliAccessKey string
)

var RootCmd = &cobra.Command{
	Use:   "cloud-station-cli",        //One Line usage message
	Short: "cloud-station-cli 文件中转服务", //Short is the short description shown in the "help" output
	Long:  "cloud-station-cli ...",    //Long is the long message shown in the 'help<this-command>' output
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println("0.0.1")
			return nil
		}
		return errors.New("no flags find")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil { //执行的是RunE
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&ossProvider, "oss_provider", "p", "aliCloud", "the oss provider [aliCloud/tencentCloud/minio]")
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the cloud-station-cli version")
}
