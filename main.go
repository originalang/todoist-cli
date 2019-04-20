package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/originalang/togoist"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "project",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "list projects",
					Action: func(c *cli.Context) error {
						dat, _ := ioutil.ReadFile("config")
						client := togoist.NewClient(string(dat))
						client.Sync()

						w := new(tabwriter.Writer)
						w.Init(os.Stdout, 8, 8, 0, '\t', 0)
						defer w.Flush()

						fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "Id", "Name", "Indent", "Favorite")
						fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "--", "----", "------", "--------")

						for _, proj := range client.Projects {
							fmt.Fprintf(w, "\n %v\t%s\t%v\t%v\t", proj.Id, proj.Name, proj.Indent, proj.IsFavorite)
						}

						return nil
					},
				},

				{
					Name:  "add",
					Usage: "add a new project",
					Action: func(c *cli.Context) error {
						dat, _ := ioutil.ReadFile("config")
						client := togoist.NewClient(string(dat))
						client.Sync()
						client.AddProject(c.String("name"), c.Int("indent"))

						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name, n",
							Value: "New Project",
						},
						cli.IntFlag{
							Name:  "indent, i",
							Value: 1,
						},
					},
				},

				{
					Name:  "delete",
					Usage: "delete a project",
					Action: func(c *cli.Context) error {
						dat, _ := ioutil.ReadFile("config")
						client := togoist.NewClient(string(dat))
						client.Sync()

						ids := []int64{c.Int64("id")}
						client.DeleteProjects(ids)

						return nil
					},
					Flags: []cli.Flag{
						cli.Int64Flag{
							Name: "id",
						},
					},
				},
			},
		},

		{
			Name: "item",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "list all items under a project",
					Action: func(c *cli.Context) error {
						dat, _ := ioutil.ReadFile("config")
						client := togoist.NewClient(string(dat))
						client.Sync()

						w := new(tabwriter.Writer)
						w.Init(os.Stdout, 8, 8, 0, '\t', 0)
						defer w.Flush()

						fmt.Fprintf(w, "\n %s\t%s\t", "Id", "Content")
						fmt.Fprintf(w, "\n %s\t%s\t", "--", "-------")

						for _, itm := range client.Items {
							if (c.Int64("projectid") == itm.ProjectId) {
								fmt.Fprintf(w, "\n %v\t%s\t", itm.Id, itm.Content)
							}
						}

						return nil
					},
					Flags: []cli.Flag{
						cli.Int64Flag{
							Name: "projectid, p",
						},
					},
				},

				{
					Name:  "add",
					Usage: "add an item to a project",
					Action: func(c *cli.Context) error {
						dat, _ := ioutil.ReadFile("config")
						client := togoist.NewClient(string(dat))
						client.Sync()
						
						client.AddItem(c.Int64("projectid"), c.String("content"), c.Int("indent"))

						return nil
					},
					Flags: []cli.Flag{
						cli.Int64Flag{
							Name: "projectid, p",
						},
						cli.StringFlag{
							Name: "content, c",
						},
						cli.IntFlag{
							Name:  "indent, i",
							Value: 1,
						},
					},
				},

				{
					Name:  "delete",
					Usage: "delete an item from a project",
					Action: func(c *cli.Context) error {
						dat, _ := ioutil.ReadFile("config")
						client := togoist.NewClient(string(dat))
						client.Sync()
						
						ids := []int64{c.Int64("id")}
						client.DeleteItems(ids)

						return nil
					},
					Flags: []cli.Flag{
						cli.Int64Flag{
							Name: "id",
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}