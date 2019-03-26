package services

import "os/exec"

func RunInSystem(script string, args []string) (output []byte, err error) {
	osCommand := exec.Command(script, args...)
	output, err = osCommand.Output()
	if err != nil {
		return nil, err
	}
	return output, nil
}
