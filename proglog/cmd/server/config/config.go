package config

import (
	"encoding/json"
	"fmt"
)

type ServerConfiguration struct {
	Server Server `json:"server"`
}

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (s *ServerConfiguration) GetServerUrl() string {
	return fmt.Sprintf("%s:%d",
		s.Server.Host,
		s.Server.Port,
	)
}

func LoadServerConfiguration(configResource []byte) ServerConfiguration {
	var serverConfiguration ServerConfiguration
	err := json.Unmarshal(configResource, &serverConfiguration)
	if err != nil {
		panic(err)
	}

	return serverConfiguration
}
