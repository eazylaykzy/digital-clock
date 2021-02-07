package main

import (
	"fmt"
	"time"
)

func main() {
	type placeholder [5]string
	blinker := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for {

		now := time.Now()
		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			blinker,
			digits[minute/10], digits[minute%10],
			blinker,
			digits[second/10], digits[second%10],
		}
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == blinker && second%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
		fmt.Print("\033[H\033[2J")
	}

	//fmt.Println(now.Clock())
}
