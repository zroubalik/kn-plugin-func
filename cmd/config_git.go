package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"knative.dev/func/pkg/config"
	fn "knative.dev/func/pkg/functions"
)

func NewConfigGitCmd(newClient ClientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git",
		Short: "Manage Git configuration of a function",
		Long: `Manage Git configuration of a function

Prints Git configuration for a function project present in
the current directory or from the directory specified with --path.
`,
		SuggestFor: []string{"gti", "Git", "Gti"},
		PreRunE:    bindEnv("path"),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return runConfigGitCmd(cmd, newClient)
		},
	}
	// Global Config
	cfg, err := config.NewDefault()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStdout(), "error loading config at '%v'. %v\n", config.File(), err)
	}

	// Function Context
	f, _ := fn.NewFunction(effectivePath())
	if f.Initialized() {
		cfg = cfg.Apply(f)
	}

	configGitSetCmd := NewConfigGitSetCmd(newClient)
	//configGitRemoveCmd := NewConfigGitRemoveCmd()

	addPathFlag(cmd)
	addVerboseFlag(cmd, cfg.Verbose)

	cmd.AddCommand(configGitSetCmd)

	return cmd
}

func runConfigGitCmd(cmd *cobra.Command, newClient ClientFactory) (err error) {
	// TODO implement, this is just a mock:
	// fmt.Printf("--------------------------- Function Git config ---------------------------\n")
	// fmt.Printf("STATUS\t PROPERTY\tVALUE\n")
	// fmt.Printf("  ✅\t Git URL:\t%s\n", f.Build.Git.URL)
	// fmt.Printf("  ❌\t Builder:\t%s\n", f.Build.Builder)
	// fmt.Printf("  \t\t\t  - Function configuration has %q but generated Pipelines have \"s2i\", please regenerate.\n\n", f.Build.Builder)
	// fmt.Printf("  ✅\t Image:  \t%s\n", f.Image)
	// fmt.Printf("\n")
	// end of mock

	return nil
}
