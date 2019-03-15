package main

import (
	"fmt"

	vinput "github.com/multiverse-os/vinput"
	os "github.com/multiverse-os/vinput/libs/os"
)

func main() {
	fmt.Println("[ portal-gun ]: vinput ( virtual keyboard user input )")
	fmt.Println("===============================================================================")
	fmt.Println("  The `portal-gun` virtual keyboard provides low level access to provision")
	fmt.Println("portals and completely automate the entire configuration, image building, ")
	fmt.Println("testing.                                                                  ")
	fmt.Println("===============================================================================")
	fmt.Println("-------------------------------------------------------------------------------")
	os.RunAsRoot()
	keyboard := vinput.NewKeyboard()
	defer keyboard.RemoveDevice()
}
