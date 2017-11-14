package cmd

import (
	"github.com/MottainaiCI/mottainai-server/pkg/agentconn"
	setting "github.com/MottainaiCI/mottainai-server/pkg/settings"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/mudler/mottainai-server/pkg/utils"

	agenttasks "github.com/MottainaiCI/mottainai-server/pkg/tasks"
	"github.com/urfave/cli"
)

var Agent = cli.Command{
	Name:        "agent",
	Usage:       "Start agent",
	Description: `Mottainai agent`,
	Action:      runAgent,
}

func runAgent(c *cli.Context) error {
	setting.GenDefault()

	rabbit, m_error := agentconn.NewMachineryServer()
	if m_error != nil {
		panic(m_error)
	}
	agenttasks.RegisterTasks(rabbit)
	ID := utils.GenID()
	log.INFO.Println("Worker ID: " + ID)

	worker := rabbit.NewWorker(ID, setting.AgentConcurrency)

	return worker.Launch()
}
