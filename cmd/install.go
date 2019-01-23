package cmd

import (
	"errors"
	"strings"

	"github.com/bmuschko/lets-gopher-exercise/remote"
	"github.com/bmuschko/lets-gopher-exercise/templ"
	"github.com/bmuschko/lets-gopher-exercise/utils"
	"github.com/spf13/cobra"
)

var repoUrl string

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installs a template from a URL",
	Run:   installTemplate,
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.PersistentFlags().StringVar(&repoUrl, "url", "", "template URL")
	installCmd.MarkFlagRequired("url")
}

func installTemplate(cmd *cobra.Command, args []string) {
	err := utils.CreateDir(templ.TemplateDir)
	utils.CheckIfError(err)
	install(repoUrl)
}

func install(repoUrl string) {
	var repo remote.Repository

	if strings.HasSuffix(repoUrl, ".git") {
		repo = &remote.GitRepo{repoUrl, templ.TemplateDir}
	} else {
		errors.New("Currently templates can only be installed from a Git repository")
	}

	repo.Install()
}
