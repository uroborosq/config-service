package config

import (
	"os"
	"os/exec"
)

type HostnameService struct {
}

func NewHostnameService() *HostnameService {
	return &HostnameService{}
}

func (s *HostnameService) GetHostname() (string, error) {
	return os.Hostname()
}

func (s *HostnameService) UpdateHostname(newHostname string) error {
	return exec.Command("hostnamectl", "set-hostname", newHostname).Run()
}
