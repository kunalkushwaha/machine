package commands

import (
	//	"os"

	"github.com/codegangsta/cli"
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/log"
	//"github.com/docker/machine/libmachine/libcompose"
	cliApp "github.com/docker/libcompose/cli/app"
	dockerApp "github.com/docker/libcompose/cli/docker/app"
)

var (
	sharedServiceFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "file,f",
			Usage:  "Specify an alternate compose file (default: docker-compose.yml)",
			Value:  "docker-compose.yml",
			EnvVar: "COMPOSE_FILE",
		},
		cli.StringFlag{
			Name:  "project-name,p",
			Usage: "Specify an alternate project name (default: directory name)",
		},
		/*		cli.BoolFlag{
					Name:  "tls",
					Usage: "Use TLS; implied by --tlsverify",
				},
				cli.BoolFlag{
					Name:   "tlsverify",
					Usage:  "Use TLS and verify the remote",
					EnvVar: "DOCKER_TLS_VERIFY",
				},
				cli.StringFlag{
					Name:  "tlscacert",
					Usage: "Trust certs signed only by this CA",
				},
				cli.StringFlag{
					Name:  "tlscert",
					Usage: "Path to TLS certificate file",
					EnvVar: "DOCKER_CERT_PATH",
				},
				cli.StringFlag{
					Name:  "tlskey",
					Usage: "Path to TLS key file",
					EnvVar: "DOCKER_CERT_PATH",
				},
				cli.StringFlag{
					Name:  "configdir",
					Usage: "Path to docker config dir, default ${HOME}/.docker",
				},
		*/
	}
)

func cmdServiceLs(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service ls")
	return nil
}

func cmdServicePs(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service ps")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectPs(p, context)
	return nil
}

func cmdServiceBuild(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service build")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectBuild(p, context)
	return nil
}

func cmdServicePort(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service port")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectPort(p, context)
	return nil
}

func cmdServicePull(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service pull")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectPull(p, context)
	return nil
}

func cmdServiceRestart(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service restart")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectRestart(p, context)
	return nil
}

func cmdServiceLog(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service logs")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectLog(p, context)
	return nil
}

func cmdServiceKill(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service kill")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectKill(p, context)
	return nil
}

func cmdServicePause(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service pause")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectPause(p, context)
	return nil
}

func cmdServiceUnpause(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service unpause")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectUnpause(p, context)
	return nil
}

func cmdServiceStart(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service start")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectStart(p, context)
	return nil
}

func cmdServiceCreate(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service create")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectCreate(p, context)
	return nil
}

func cmdServiceUp(context *cli.Context, api libmachine.API) error {

	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectUp(p, context)
	return nil
}

func cmdServiceStop(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service stop")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectDown(p, context)
	return nil
}
func cmdServiceRemove(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service rm")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectDelete(p, context)
	return nil
}

func cmdServiceScale(context *cli.Context, api libmachine.API) error {
	log.Info("Executing:  service scale")
	factory := &dockerApp.ProjectFactory{}

	p, err := factory.Create(context)
	if err != nil {
		log.Info("Failed to read project: %v", err)
	}

	cliApp.ProjectScale(p, context)
	return nil
}
