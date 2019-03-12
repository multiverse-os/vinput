package main

import (
	"fmt"

	uinput "github.com/multiverse-os/vinput/uinput"
)

// TODO: Need to ber able to generally know:
// 1) What process is user input being entered into? (Like is it gnome-terminal? etc)
// 2) What has been entered since the last <enter> or <carriage-return> (to
//    check for types)
// 3) Status of various toggles? (CapsLock, what keys are being held down,
//    NumLock, etc)
// 4) What stdout/stderror has been produced since the last <enter> or
//    <carriage-return>
// 5) Informationa about the machine the input is connected to (preferably, we
//    would have a logical K(v)M switch, and be able to rotate between machines.
// 6) In `uinput` lib, setup macros for the common key combos like
//    CTRL+ALT+DELTE or ALT+TAB. Make it easy to define more on-the-fly.
// 7) Track avaialble options in current window we can <TAB> between, what
//    things can be clicked and what their x/y coordinates are.
// 8) Generate realsitic USB id numbers, switch between different real keymaps
// 9) Support unicode
// 10) Support API with ability to do:
//     a) Type("ls -lah")
//     b) AltTab("window-name")

//

// TODO: Need to be able to ask the mouse:
// 1) Where the current location is?
// 2) What window it has focused?
// 3) What the cursor type is?
// 4) Color underneath cursor?
// 5) If highlighting text, what is highlighted?

type VirtualInput struct {
	Keyboard     vKeyboard
	Mouse        vMouse
	InputProcess string // Process user input is being entered into
}

func NewKeyboard() vKeyboard {

	RunAsRoot()

	keyboard, err := uinput.NewKeyboard("/dev/uinput", []byte("portal-gun-v0.1.0"))
	if err != nil {
		fmt.Println("[error] Failed to create `portal-gun-v0.1.0` keyboard: ", err)

	}
	fmt.Println("kbd: ", keyboard)
	fmt.Println("Pressing key [30:Key A]:", keyboard.KeyPress(30))

	// MAIN LOOP
	//ticker := time.NewTicker(20 * time.Millisecond)
	//tick := time.Tick()
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick...")
	//	}
	//}

}
