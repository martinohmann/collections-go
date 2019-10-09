package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinohmann/collections-go/internal/codegen"
	"github.com/spf13/cobra"
)

// NewCommand creates the collections-gen command.
func NewCommand() *cobra.Command {
	config := codegen.NewDefaultConfig()

	cmd := &cobra.Command{
		Use:           "collections-gen [flags] <output-file>",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := config.Validate()
			if err != nil {
				return err
			}

			err = config.Complete(args)
			if err != nil {
				return err
			}

			return Run(config)
		},
	}

	config.AddFlags(cmd)

	return cmd
}

// Run runs the code generator with given configuration.
func Run(config *codegen.Config) error {
	buf, err := codegen.Generate(config, codegen.Template)
	if err != nil {
		fmt.Fprintln(os.Stderr, string(buf))
		return err
	}

	if config.OutputFile == "-" {
		fmt.Fprintln(os.Stdout, string(buf))
		return nil
	}

	return ioutil.WriteFile(config.OutputFile, buf, 0666)
}

func main() {
	cmd := NewCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
