package main

func RunAsRoot() {
	// Root is required because we are directly interacting with process memory
	if os.Geteuid() != 0 {
		fmt.Println("[portal-gun: virtui] Creating a new keyboard requires root access, re-running the software with sudo: `sudo !!`.")
		executable, _ := os.Executable()
		directory, binary := filepath.Split(executable)
		os.Chdir(directory)
		cmd := exec.Command("/bin/sh", "-c", "sudo ./"+binary)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			fmt.Println("[fatal error] failed to re-run the executable with sudo, exiting...")
			os.Exit(1)
		} else {
			// Exit because we just relaunched as root
			os.Exit(0)
		}
	}
}
