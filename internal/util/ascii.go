package util

import "fmt"

func Banner() error {
	lines := []string{
		" __   __  __   __  _______  _______ ",
		"|  |_|  ||  | |  ||       ||       |",
		"|       ||  |_|  ||_     _||   _   |",
		"|       ||       |  |   |  |  | |  |",
		"|       ||_     _|  |   |  |  |_|  |",
		"| ||_|| |  |   |    |   |  |       |",
		"|_|   |_|  |___|    |___|  |_______|",
	}

	fmt.Print("\n")

	for _, line := range lines {
		fmt.Print(line + "\n")
	}

	fmt.Print("\n")

	fmt.Print("Fictious data generation\n")
	fmt.Print("---   ---   ---  --  --  --  -- - - -\n\n")

	return nil
}
