package vinput

import (
	"fmt"

	uinput "github.com/multiverse-os/vinput/libs/uinput"
)

type VirtualInput struct {
	Keyboard     uinput.Keyboard
	Mouse        uinput.Mouse
	InputProcess string // Process user input is being entered into, like GnomeTerminal
}

func NewKeyboard() uinput.Keyboard {
	keyboard, err := uinput.Keyboard.New("portal-gun-v0.1.0")
	if err != nil {
		fmt.Println("[error] Failed to create `portal-gun-v0.1.0` keyboard: ", err)

	}
	fmt.Println("kbd: ", keyboard)
	fmt.Println("Pressing key [30:Key A]:", keyboard.PressKey(30))
	// MAIN LOOP
	//ticker := time.NewTicker(20 * time.Millisecond)
	//tick := time.Tick()
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick...")
	//	}
	//}
	return keyboard
}
