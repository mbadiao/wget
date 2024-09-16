package main
/* 

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	cmd := "/bin/wget"
	args := os.Args[1:]
	err := execCommand(cmd, args...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func execCommand(cmd string, args ...string) error {
	fullArgs := append([]string{cmd}, args...)
	env := os.Environ()

	stdin := os.Stdin.Fd()
	stdout := os.Stdout.Fd()
	stderr := os.Stderr.Fd()

	pid, err := syscall.ForkExec(cmd, fullArgs, &syscall.ProcAttr{
		Dir:   "",
		Env:   env,
		Files: []uintptr{stdin, stdout, stderr},
		Sys:   nil,
	})

	if err != nil {
		return fmt.Errorf("erreur lors de la création du processus : %w", err)
	}

	var wstatus syscall.WaitStatus
	_, err = syscall.Wait4(pid, &wstatus, 0, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de l'attente du processus : %w", err)
	}

	if wstatus.ExitStatus() != 0 {
		return fmt.Errorf("la commande a échoué avec le code de sortie : %d", wstatus.ExitStatus())
	}

	return nil
}
*/