/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/ArkTree/fabric/bccsp/factory"
	"github.com/ArkTree/fabric/notinternal/peer/chaincode"
	"github.com/ArkTree/fabric/notinternal/peer/channel"
	"github.com/ArkTree/fabric/notinternal/peer/common"
	"github.com/ArkTree/fabric/notinternal/peer/lifecycle"
	"github.com/ArkTree/fabric/notinternal/peer/node"
	"github.com/ArkTree/fabric/notinternal/peer/snapshot"
	"github.com/ArkTree/fabric/notinternal/peer/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The main command describes the service and
// defaults to printing the help message.
var mainCmd = &cobra.Command{Use: "peer"}

func main() {
	// For environment variables.
	viper.SetEnvPrefix(common.CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Define command-line flags that are valid for all peer commands and
	// subcommands.
	mainFlags := mainCmd.PersistentFlags()

	mainFlags.String("logging-level", "", "Legacy logging level flag")
	viper.BindPFlag("logging_level", mainFlags.Lookup("logging-level"))
	mainFlags.MarkHidden("logging-level")

	cryptoProvider := factory.GetDefault()

	mainCmd.AddCommand(version.Cmd())
	mainCmd.AddCommand(node.Cmd())
	mainCmd.AddCommand(chaincode.Cmd(nil, cryptoProvider))
	mainCmd.AddCommand(channel.Cmd(nil))
	mainCmd.AddCommand(lifecycle.Cmd(cryptoProvider))
	mainCmd.AddCommand(snapshot.Cmd(cryptoProvider))

	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
