package main

import (
	"log"

	"github.com/georgi-georgiev/blunder"
	"github.com/georgi-georgiev/blunder/cmd/blunder/gen"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "blunder",
	Short:   "Blunder: An elegant toolkit for http error responses.",
	Long:    `Blunder: An elegant toolkit for http error responses.`,
	Version: blunder.Version,
}

func init() {
	rootCmd.AddCommand(gen.CmdGen)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
