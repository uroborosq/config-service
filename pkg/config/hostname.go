package config

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

type HostnameService struct {
}

func NewHostnameService() *HostnameService {
	return &HostnameService{}
}

func (s *HostnameService) GetHostname() (string, error) {
	var builder strings.Builder
	file, err := os.Open("/etc/hostname")
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	return builder.String(), scanner.Err()
}

func (s *HostnameService) UpdateHostname(newHostname string) error {
	return exec.Command("hostnamectl", "set-hostname", newHostname).Run()
}
