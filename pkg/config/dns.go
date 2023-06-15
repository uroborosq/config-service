package config

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"golang.org/x/exp/slices"
)

var UnknownOutputErr = errors.New("can't parse networkmanager output")

type DnsService struct {
	nmConnection string
}

func NewDnsService(nmConnection string) *DnsService {
	return &DnsService{nmConnection: nmConnection}
}

func (s *DnsService) GetServerList() ([]string, error) {
	servers, err := s.parseServerList()
	if err != nil {
		return nil, err
	}
	result := make([]string, len(servers))
	for i, server := range servers {
		result[i] = string(server)
	}
	return result, nil
}

func (s *DnsService) AddServer(address string) error {
	servers, err := s.parseServerList()
	if err != nil {
		return err
	}
	addressBytes := []byte(address)
	for i := range servers {
		if slices.Equal(servers[i], addressBytes) {
			return nil
		}
	}
	return s.applyServiceList(append(servers, addressBytes))
}

func (s *DnsService) RemoveServer(address string) error {
	servers, err := s.parseServerList()
	if err != nil {
		return err
	}
	addressBytes := []byte(address)
	for i := range servers {
		if slices.Equal(servers[i], addressBytes) {
			servers = append(servers[:i], servers[i+1:]...)
			break
		}
	}
	return s.applyServiceList(servers)
}

func (s *DnsService) parseServerList() ([][]byte, error) {
	output, err := exec.Command("nmcli", "-f", "ipv4.dns", "c", "show", s.nmConnection).Output()
	if err != nil {
		fmt.Println(string(err.(*exec.ExitError).Stderr))
		return nil, err
	}
	fields := bytes.Fields(output)
	if len(fields) == 2 && string(fields[1]) == "--" {
		return [][]byte{}, nil
	} else if len(fields) >= 2 {
		return bytes.Split(fields[1], []byte{','}), nil
	}
	return nil, UnknownOutputErr
}

func (s *DnsService) applyServiceList(servers [][]byte) error {
	var builder strings.Builder
	for i := range servers {
		builder.Write(servers[i])
		if i != len(servers)-1 {
			builder.WriteByte(',')
		}
	}
	return exec.Command("nmcli", "c", "modify", s.nmConnection, "ipv4.dns", builder.String()).Run()
}
