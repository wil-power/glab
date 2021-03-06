package commands

import (
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
	"glab/internal/git"
	"log"
)

var issueListCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   `List project issues`,
	Long:    ``,
	Aliases: []string{"ls"},
	Args:    cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var state string
		if lb, _ := cmd.Flags().GetBool("all"); lb {
			state = "all"
		} else if lb, _ := cmd.Flags().GetBool("closed"); lb {
			state = "closed"
		} else {
			state = "opened"
		}

		l := &gitlab.ListProjectIssuesOptions{
			State: gitlab.String(state),
		}
		if lb, _ := cmd.Flags().GetString("label"); lb != "" {
			label := gitlab.Labels{
				lb,
			}
			l.Labels = label
		}
		if lb, _ := cmd.Flags().GetString("milestone"); lb != "" {
			l.Milestone = gitlab.String(lb)
		}
		if lb, _ := cmd.Flags().GetBool("confidential"); lb {
			l.Confidential = gitlab.Bool(lb)
		}

		gitlabClient, repo := git.InitGitlabClient()

		issues, _, err := gitlabClient.Issues.ListProjectIssues(repo, l)
		if err != nil {
			log.Fatal(err)
		}
		displayAllIssues(issues)

	},
}

func init() {
	issueListCmd.Flags().StringP("label", "l", "", "Filter issue by label <name>")
	issueListCmd.Flags().StringP("milestone", "", "", "Filter issue by milestone <id>")
	issueListCmd.Flags().BoolP("all", "a", false, "Get all issues")
	issueListCmd.Flags().BoolP("closed", "c", false, "Get only closed issues")
	issueListCmd.Flags().BoolP("opened", "o", false, "Get only opened issues")
	issueListCmd.Flags().BoolP("confidential", "", false, "Filter by confidential issues")
	issueCmd.AddCommand(issueListCmd)
}
