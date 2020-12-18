package spacemath

import "os/exec"

func open(filePath string) error {
	cmd := exec.Command("open", filePath)
	return cmd.Start()
}
