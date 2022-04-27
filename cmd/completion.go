/*
Copyright © 2022 kubetrail.io authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"strings"

	"github.com/kubetrail/bip32/pkg/app"
	"github.com/spf13/cobra"
)

var longCompletionCmd = `To load completions:

Bash:

  $ source <(appName completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ appName completion bash > /etc/bash_completion.d/appName
  # macOS:
  $ appName completion bash > /usr/local/etc/bash_completion.d/appName

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ appName completion zsh > "${fpath[1]}/_appName"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ appName completion fish | source

  # To load completions for each session, execute once:
  $ appName completion fish > ~/.config/fish/completions/appName.fish

PowerShell:

  PS> appName completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> appName completion powershell > appName.ps1
  # and source this file from your PowerShell profile.
`

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion",
	Long: strings.ReplaceAll(
		longCompletionCmd,
		"appName",
		app.Name),
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
