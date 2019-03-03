package cmd

import (
	"fmt"
	"github.com/bmuschko/lets-gopher/templ"
	"github.com/spf13/cobra"
	"io"
)

type templateInstallCmd struct {
	name string
	URL  string
	out  io.Writer
	home templ.Home
}

func newTemplateInstallCmd(out io.Writer) *cobra.Command {
	add := &templateInstallCmd{out: out}

	cmd := &cobra.Command{
		Use:   "install [URL] [NAME]",
		Short: "installs a template from a URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkArgsLength(len(args), "the url of the template archive", "name for the template"); err != nil {
				return err
			}

			add.URL = args[0]
			add.name = args[1]
			add.home = templ.LetsGopherSettings.Home
			return add.run()
		},
	}
	return cmd
}

func (a *templateInstallCmd) run() error {
	downloader := &templ.TemplateDownloader{Home: templ.LetsGopherSettings.Home, Getter: templ.NewHTTPGetter()}
	templateZIP, err := downloader.DownloadTo(a.URL, a.name)

	if err != nil {
		return nil
	}

	if err := addTemplate(a.name, templateZIP, a.home); err != nil {
		return err
	}
	fmt.Fprintf(a.out, "%q has been added to your templates\n", a.name)
	return nil
}

func addTemplate(name string, templateZIP string, home templ.Home) error {
	f, err := templ.LoadTemplatesFile(home.TemplatesFile())
	if err != nil {
		return err
	}

	if f.Has(name) {
		return fmt.Errorf("template with name (%s) already exists, please specify a different name", name)
	}

	c := templ.Template{
		Name:        name,
		ArchivePath: templateZIP,
	}
	f.Update(&c)

	return f.WriteFile(home.TemplatesFile(), 0644)
}