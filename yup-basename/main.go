package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/basename"
)

const (
	flagSuffix   = "suffix"
	flagMultiple = "multiple"
	flagZero     = "zero"
)

func main() {
	app := &cli.App{
		Name:  "basename",
		Usage: "strip directory and suffix from filenames",
		UsageText: `basename NAME [SUFFIX]
   basename OPTION... NAME...

   Print NAME with any leading directory components removed.
   If specified, also remove a trailing SUFFIX.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagSuffix,
				Aliases: []string{"s"},
				Usage:   "remove a trailing SUFFIX; implies -a",
			},
			&cli.BoolFlag{
				Name:    flagMultiple,
				Aliases: []string{"a"},
				Usage:   "support multiple arguments and treat each as a NAME",
			},
			&cli.BoolFlag{
				Name:    flagZero,
				Aliases: []string{"z"},
				Usage:   "end each output line with NUL, not newline",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "basename: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.IsSet(flagSuffix) {
		params = append(params, Suffix(c.String(flagSuffix)))
	}
	if c.Bool(flagMultiple) {
		params = append(params, Multiple)
	}
	if c.Bool(flagZero) {
		params = append(params, Zero)
	}

	// Create and execute the basename command
	cmd := Basename(params...)
	return gloo.Run(cmd)
}
