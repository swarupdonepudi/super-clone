package cmd

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"os"
)

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Open Issues on SCM Web Interface",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remoteUrl := util.GetGitRemoteUrl(repo)
			repoUrl := url.Parse(remoteUrl)
			if repoUrl.GetIssuesUrl() != "" {
				open.Run(repoUrl.GetIssuesUrl())
			} else {
				println(fmt.Sprintf("SCM Provider %s does not support Issues", repoUrl.ScmProvider))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)
}
