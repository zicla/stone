package input

import (
	"bufio"
	"os"
	"fmt"
)

func GetConsoleInputText() string {

	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (print a '*' in a single line to stop):")
	var lines string
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			//enter * to stop
			if line[0] == '*' {
				break
			}
		}
		lines = lines + line + "\n"
	}

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	if len(lines) == 0 {
		return ""
	}

	return lines
}
