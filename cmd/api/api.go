package api

import (
	"github.com/spf13/cobra"
)

func init() {
	Api.AddCommand(cloud)
}

var Api = &cobra.Command{
	Use:   "api",
	Short: "Use an IONOS API",
}
