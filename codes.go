//ansi escape sequences can be used to control cursor location, color, font styling in terminals
//we are benefitting it to create custom effects and animations
//for more info check : https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
//thanks to above repo maintainers, this package mainly derivied from that.

package ansi

import "fmt"

const (
	ESC = "\x1b"
	CSI = "\x9B" //Control Sequence Introducer:  ESC [
	DCS = "\x90" //Device Control String:  ESC P
	OSC = "\x9D" //Operating System Command: ESC ]
)

func Csi(code string) {
	fmt.Printf("%s%s", CSI, code)
}

func Esc(code string) {
	fmt.Printf("%s%s", ESC, code)
}
