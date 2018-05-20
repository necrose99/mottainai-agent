/*

Copyright (C) 2017-2018  Ettore Di Giacinto <mudler@gentoo.org>

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
	"github.com/MottainaiCI/mottainai-server/pkg/mottainai"

	"github.com/urfave/cli"
)

var Health = cli.Command{
	Name:        "health",
	Usage:       "Start HealthCheck service",
	Description: `Mottainai Agent Healthcheck`,
	Action:      runHealthCheck,
	Flags: []cli.Flag{
		stringFlag("config, c", "custom/conf/agent.yml", "Custom configuration file path"),
		boolFlag("oneshot, o", "Execute once"),
	},
}

func runHealthCheck(c *cli.Context) {
	m := mottainai.NewAgent()
	var config string
	if c.IsSet("config") {
		config = c.String("config")
	}
	if c.IsSet("oneshot") {
		m.HealthCheckSetup(config)
		m.HealthClean()
		return
	}

	m.HealthCheckRun(config)
}
