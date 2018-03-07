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
	"github.com/MottainaiCI/mottainai-server/pkg/agentconn"
	"github.com/MottainaiCI/mottainai-server/pkg/client"

	setting "github.com/MottainaiCI/mottainai-server/pkg/settings"

	agenttasks "github.com/MottainaiCI/mottainai-server/pkg/tasks"
	"github.com/MottainaiCI/mottainai-server/pkg/utils"
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	anagent "github.com/mudler/anagent"
	"github.com/urfave/cli"
)

var Agent = cli.Command{
	Name:        "agent",
	Usage:       "Start agent",
	Description: `Mottainai agent`,
	Action:      runAgent,
	Flags: []cli.Flag{
		stringFlag("config, c", "custom/conf/agent.yml", "Custom configuration file path"),
	},
}

func runAgent(c *cli.Context) error {
	setting.GenDefault()
	agent := anagent.New()
	if c.IsSet("config") {
		setting.LoadFromFileEnvironment(c.String("config"))
	}

	rabbit, m_error := agentconn.NewMachineryServer()
	if m_error != nil {
		panic(m_error)
	}

	th := agenttasks.DefaultTaskHandler()
	th.RegisterTasks(rabbit)
	agent.Map(th)
	ID := utils.GenID()
	log.INFO.Println("Worker ID: " + ID)

	worker := rabbit.NewWorker(ID, setting.Configuration.AgentConcurrency)
	Register(ID)
	// agent.TimerSeconds(int64(200), true, func(l *corelog.Logger) {
	// 		Register(ID)
	// 	})

	go func(w *machinery.Worker, a *anagent.Anagent) {
		agent.Map(w)
		agent.Start()
	}(worker, agent)

	return worker.Launch()
}

func Register(ID string) {
	fetcher := client.NewClient()
	fetcher.RegisterNode(ID)
}
