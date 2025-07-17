package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func enableService(config *ServiceConfig) error {
	if err := createSystemdFile(config, serviceTemplate, "service"); err != nil {
		return fmt.Errorf("error creating systemd service file: %v", err)
	}

	if err := createSystemdFile(config, timerTemplate, "timer"); err != nil {
		return fmt.Errorf("error creating systemd timer file: %v", err)
	}

	commands := []string{
		"systemctl daemon-reload",
		"systemctl enable " + config.ServiceName + ".timer",
		"systemctl start " + config.ServiceName + ".timer",
	}

	for _, cmd := range commands {
		if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
			return fmt.Errorf("timer command failed: %s, error: %v", cmd, err)
		}
	}

	return nil
}

func disableService(config *ServiceConfig) error {
	commands := []string{
		"systemctl disable " + config.ServiceName + ".timer",
	}

	for _, cmd := range commands {
		if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
			return fmt.Errorf("timer disable command failed: %s, error: %v", cmd, err)
		}
	}

	return nil
}

func createSystemdFile(config *ServiceConfig, fileTemplate string, extension string) error {
	serviceTmpl, err := template.New("service").Parse(fileTemplate)
	if err != nil {
		return fmt.Errorf("service template error: %v", err)
	}

	serviceFile, err := os.Create("/etc/systemd/system/" + config.ServiceName + "." + extension)
	if err != nil {
		return fmt.Errorf("create service file error: %v", err)
	}

	if err := serviceTmpl.Execute(serviceFile, config); err != nil {
		return fmt.Errorf("service template execution error: %v", err)
	}
	return nil
}
