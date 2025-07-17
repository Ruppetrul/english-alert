package main

import "errors"

type ServiceConfig struct {
	ServiceName      string
	Description      string
	ExecPath         string
	WorkingDirectory string
}

const appName = "english-alert"

func NewServiceConfig(desc string, execPatch string, directory string) (*ServiceConfig, error) {
	if execPatch == "" {
		return nil, errors.New("execPatch is required")
	}

	if directory == "" {
		return nil, errors.New("directory is required")
	}

	if desc == "" {
		desc = appName
	}

	return &ServiceConfig{
		ServiceName:      appName,
		Description:      desc,
		ExecPath:         execPatch,
		WorkingDirectory: directory,
	}, nil
}
