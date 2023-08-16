///////////////////////////////////////////////////////////////////////////////
//
// Copyright (c) 2019-present Apulis Technology (Shenzhen) Incorporated. All Rights Reserved
//
//
// Distributed under the MIT License (http://opensource.org/licenses/MIT)
//
///////////////////////////////////////////////////////////////////////////////

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "user-auth",
	Short: "user-auth",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute the current command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
