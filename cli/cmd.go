package main

import "github.com/spf13/cobra"

// The root command serves as a base for all subcommands.
var root = &cobra.Command{
	Use:   `foodunit`,
	Short: `FoodUnit CLI`,
	Run:   rootHandler,
}

// `offers` runs an API request for the active offers, fetches the data
// and prints the result.
var offers = &cobra.Command{
	Use:   `offers`,
	Short: `Lists all active offers`,
	Run:   offerHandler,
}

// `dishes` displays the menu for a given supplier. The supplier ID may
// be obtained by running `offers` before for example.
var dishes = &cobra.Command{
	Use:   `dishes`,
	Short: `Lists all dishes for a given supplier`,
	Run:   dishesHandler,
}

// `supplier` fetches all important data for a given supplier. Its ID
// can be retrieved by running `offers` for example.
var supplier = &cobra.Command{
	Use:   `supplier`,
	Short: `Displays basic data for a given supplier`,
	Run:   supplierHandler,
}

// Execute creates all necessary flags for each command and then attaches
// the commands to the root command. The root command is then executed.
func Execute() error {
	dishes.Flags().String("supplier", "", `The ID of the supplier`)
	supplier.Flags().String("supplier", "", `The ID of the supplier`)

	root.AddCommand(offers)
	root.AddCommand(dishes)
	root.AddCommand(supplier)

	err := root.Execute()
	return err
}
