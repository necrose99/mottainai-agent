/*

Copyright (C) 2017-2018  Ettore Di Giacinto <mudler@gentoo.org>
                         Daniele Rondina <geaaru@sabayonlinux.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	common "github.com/MottainaiCI/mottainai-agent/common"
	s "github.com/MottainaiCI/mottainai-server/pkg/settings"
	utils "github.com/MottainaiCI/mottainai-server/pkg/utils"
	viper "github.com/spf13/viper"
)

const (
	agentName = `Mottainai Agent - Task/Job Agent
Copyright (c) 2017-2018 Mottainai

`
	agentExamples = ""
)

var rootCmd = &cobra.Command{
	Short:        agentName,
	Version:      common.MAGENT_VERSION,
	Example:      agentExamples,
	Args:         cobra.OnlyValidArgs,
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		var v *viper.Viper = s.Configuration.Viper

		v.SetConfigFile(v.Get("config").(string))
		// Parse configuration file
		err = s.Configuration.Unmarshal()
		utils.CheckError(err)
	},
}

func init() {
	var pflags = rootCmd.PersistentFlags()
	pflags.StringP("config", "c", common.MAGENT_DEF_CONFFILE, "Mottainai Agent Config File")

	s.Configuration.Viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.AddCommand(
		newAgentCommand(),
		newHealtcheckCommand(),
		newPrintCommand(),
	)
}

func Execute() {
	// Start command execution
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
