package sinks

import "errors"

type SyslogConfig struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
	Tag     string `yaml:"tag"`
}

func NewSyslogSink(config *SyslogConfig) (Sink, error) {
	return nil, errors.New("syslog sink is not supported on windows")
}
