package cmd

import (
	"time"

	"github.com/MottainaiCI/mottainai-server/pkg/agentconn"
	"github.com/MottainaiCI/mottainai-server/pkg/client"
	setting "github.com/MottainaiCI/mottainai-server/pkg/settings"

	"github.com/MottainaiCI/mottainai-server/pkg/utils"
	"github.com/RichardKnop/machinery/v1/log"

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
	stopRegisterTimer := utils.RecurringTimer(func() { Register(ID) }, 360*time.Second)
	defer func() { stopRegisterTimer <- true }()

	return worker.Launch()
}

func Register(ID string) {
	fetcher := client.NewClient()
	fetcher.GetOptions("/nodes/register", map[string]string{
		"key":    setting.AgentKey,
		"nodeid": ID,
	})
}
