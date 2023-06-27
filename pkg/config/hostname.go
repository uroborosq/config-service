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
	return os.Hostname()
}

func (s *HostnameService) UpdateHostname(newHostname string) error {
	err := exec.Command("hostnamectl", "set-hostname", newHostname).Run()
	if err != nil {
		return err
	}

	file, err := os.OpenFile("/etc/hosts", os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "127.0.1.1") {
			builder.WriteString("127.0.1.1 " + newHostname + "\n")
		} else {
			builder.WriteString(line + "\n")
		}
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.WriteString(builder.String())
	if err != nil {
		return err
	}
	return nil
}
