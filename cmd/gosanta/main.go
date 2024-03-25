package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gkits/gosanta"
	"github.com/gkits/gosanta/cmd/gosanta/config"
	flag "github.com/spf13/pflag"
)

func main() {
	shuff := gosanta.NewShuffler()

	cfg, err := config.LoadCfg()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fPrint := flag.BoolP("print", "p", false, "print the assignments to std-out")
	fSend := flag.BoolP("send", "s", false, "send e-mails to the participants containing the name of their assigned person")
	fFile := flag.StringP("file", "f", "", "read in the participants from a file")
	flag.Parse()
	args := flag.Args()

	if *fFile != "" {

	} else {

	}

	for _, arg := range args {
		split := strings.Split(arg, ":")
		if len(split) != 2 {
			fmt.Println("")
			os.Exit(1)
		}

		if err := shuff.AddParticipant(gosanta.Participant{Name: split[0], Email: split[1]}); err != nil {
			fmt.Println("")
			os.Exit(1)
		}
	}
	fmt.Println(shuff)

	mapping := shuff.Shuffle()

	if *fPrint {
		fmt.Println(mapping)
	}

	if *fSend {
		fmt.Println("trying to send the assignments via e-mail")
	}
}
