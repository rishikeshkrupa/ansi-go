package ansi

import "fmt"

// moves cursor to home position (0, 0)
func HomePosition() {
	Csi("H")
}

func MoveToLineColumn() {
	//moves cursor to line #, column #
	Csi("{line};{column}H")
}

func UpColumns(n int) {
	//moves cursor up # lines
	Csi(fmt.Sprintf("%d%s", n, "A"))
}

func DownColumns(n int) {
	//moves cursor down # lines
	Csi(fmt.Sprintf("%d%s", n, "B"))
}

func RightColumns(n int) {
	//moves cursor right # columns
	Csi(fmt.Sprintf("%d%s", n, "C"))
}

func LeftColumns(n int) {
	//moves cursor left # columns
	Csi(fmt.Sprintf("%d%s", n, "D"))
}

func BeginningOfNextLinesDown(n int) {
	//moves cursor to beginning of next line, # lines down
	Csi(fmt.Sprintf("%d%s", n, "E"))
}

func BeginningOfPreviousLinesUp(n int) {
	//moves cursor to beginning of previous line, # lines up
	Csi(fmt.Sprintf("%d%s", n, "F"))
}

func Column(n int) {
	//moves cursor to column #
	Csi(fmt.Sprintf("%d%s", n, "G"))
}

func RequestPos() {
	//request cursor position (reports as ESC[#;#R)
	Csi("6n")
}

func OneLineUp() {
	//moves cursor one line up, scrolling if needed
	Esc("M")
}

func SavePosDEC() {
	//save cursor position (DEC)
	Esc("7")
}

func RestorePosDEC() {
	//restores the cursor to the last saved position (DEC)
	Esc("8")
}

func SavePosSCO() {
	//save cursor position (SCO)
	Csi("s")
}

func RestorePosSCO() {
	//restores the cursor to the last saved position (SCO)
	Csi("u")
}
