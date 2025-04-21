package power

import "os/exec"

func Shutdown() error {
	args := []string{"/s", "/t", "0"}
	return exec.Command("shutdown", args...).Start()
}

func Restart() error {
	args := []string{"/r", "/t", "0"}
	return exec.Command("shutdown", args...).Start()
}
