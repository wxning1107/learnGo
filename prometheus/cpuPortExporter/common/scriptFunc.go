package common

import "os/exec"

func ExeScript(script string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", script, "script port status")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
