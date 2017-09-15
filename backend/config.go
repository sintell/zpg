package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func ReadConfig(conf interface{}, name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(file)
	dec.Decode(conf)

	return nil
}

type DBConfig struct {
	PgAddr string `json:"pg_addr,omitempty"`
}

type ServerConfig struct {
	DBConfig
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}

func (sc *ServerConfig) Addr() string {
	port := strconv.Itoa(sc.Port)
	return strings.Join([]string{sc.Host, port}, ":")
}
