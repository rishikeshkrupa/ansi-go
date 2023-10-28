//ansi escape sequences can be used to control cursor location, color, font styling in terminals
//we are benefitting it to create custom effects and animations
//for more info check : https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
//thanks to above repo maintainers, this package mainly derivied from that.

package ansi

import "fmt"

/* Note : Tried using sequences other than ESC didn't work (\x9B) in the terminals
 * I've developed this with hence switching to ESC representation(\ESC [) */

const (
	ESC = "\x1b"
	CSI = ESC + "["  //Control Sequence Introducer: "\x9B"
	DCS = ESC + " P" //Device Control String:  "\x90"
	OSC = ESC + "]"  //Operating System Command: "\x9D"
)

func Esc(code string) {
	fmt.Printf("%s%s", ESC, code)
}

func Csi(code string) {
	fmt.Printf("%s%s", CSI, code)
}

func Dcs(code string) {
	fmt.Printf("%s%s", DCS, code)
}

func Osc(code string) {
	fmt.Printf("%s%s", OSC, code)
}
