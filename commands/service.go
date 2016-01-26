package commands

import (
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine"
)


func cmdServiceLs(c CommandLine, api libmachine.API) error {
	log.Info("Executing:  service ls")
	return nil
}
func cmdServiceStart(c CommandLine, api libmachine.API) error {
	log.Info("Executing:  service start")
	return nil
}
func cmdServiceCreate(c CommandLine, api libmachine.API) error {
	log.Info("Executing:  service create")
	return nil
}
func cmdServiceUp(c CommandLine, api libmachine.API) error {
	log.Info("Executing:  service up")
	return nil
}
func cmdServiceStop(c CommandLine, api libmachine.API) error {
	log.Info("Executing:  service stop")
	return nil
}
func cmdServiceRemove(c CommandLine, api libmachine.API) error {
	log.Info("Executing:  service rm")
	return nil
}
func cmdServiceScale(c CommandLine, api libmachine.API) error {
	log.Info("Executing:  service scale")
	return nil
}
