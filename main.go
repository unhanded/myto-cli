package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/unhanded/myto-cli/internal/util"
	"github.com/unhanded/mytomake/pkg/mytomake"
)

func main() {
	var filepath string
	var dataDir string
	flag.StringVar(&filepath, "f", "", "the selected myto file")
	flag.StringVar(&dataDir, "d", "$HOME/.myto", "myto data directory")

	flag.CommandLine.Usage = func() {
		util.Banner()

		fmt.Print("\tUsage: myto -f <file>\n\n")
		fmt.Print("\tExample: myto -f ./message.myto\n\n")
	}

	flag.Parse()

	myto := mytomake.NewMyto(dataDir)

	if filepath == "" {
		b, bErr := io.ReadAll(os.Stdin)
		if bErr != nil {
			panic(bErr)
		}
		res, err := myto.FromBytes(b)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", res)
	} else {
		res, err := myto.FromFile(filepath)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", res)
	}
}
