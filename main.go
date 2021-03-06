package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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
						token := os.Getenv("TODOIST_TOKEN")
						client := togoist.NewClient(string(token))
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
						token := os.Getenv("TODOIST_TOKEN")
						client := togoist.NewClient(string(token))
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
						token := os.Getenv("TODOIST_TOKEN")
						client := togoist.NewClient(string(token))
						client.Sync()

						id, err := togoist.GetProjectId(client, c.String("name"))

						if err != nil {
							return err
						}

						ids := []int64{id}
						client.DeleteProjects(ids)

						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "name, n",
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
						token := os.Getenv("TODOIST_TOKEN")
						client := togoist.NewClient(string(token))
						client.Sync()

						id, err := togoist.GetProjectId(client, c.String("name"))

						if err != nil {
							return err
						}

						for _, itm := range client.Items {
							if id == itm.ProjectId {
								if itm.Indent == 1 {
									fmt.Printf("\n %s %s [%v]", "⌲", itm.Content, itm.Id)
								} else {
									if itm.Checked == 1 {
										fmt.Printf("\n    %s└── [X] %s [%v]", strings.Repeat("\t", itm.Indent-2), itm.Content, itm.Id)
									} else {
										fmt.Printf("\n    %s└── [ ] %s [%v]", strings.Repeat("\t", itm.Indent-2), itm.Content, itm.Id)
									}
								}

							}
						}

						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "name, n",
						},
					},
				},

				{
					Name:  "add",
					Usage: "add an item to a project",
					Action: func(c *cli.Context) error {
						token := os.Getenv("TODOIST_TOKEN")
						client := togoist.NewClient(string(token))
						client.Sync()

						projId, err := togoist.GetProjectId(client, c.String("projectname"))

						if err != nil {
							return err
						}

						client.AddItem(projId, c.String("content"), c.Int("indent"), c.String("duedate"))

						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "projectname, n",
							Value: "Inbox",
						},
						cli.StringFlag{
							Name: "content, c",
						},
						cli.IntFlag{
							Name:  "indent, i",
							Value: 1,
						},
						cli.StringFlag{
							Name:  "duedate, d",
							Value: "tomorrow",
						},
					},
				},

				{
					Name:  "delete",
					Usage: "delete an item from a project",
					Action: func(c *cli.Context) error {
						token := os.Getenv("TODOIST_TOKEN")
						client := togoist.NewClient(string(token))
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

				{
					Name:  "complete",
					Usage: "complete an item in a project",
					Action: func(c *cli.Context) error {
						token := os.Getenv("TODOIST_TOKEN")
						client := togoist.NewClient(string(token))
						client.Sync()

						parentId := c.Int64("id")
						children := togoist.GetChildrenIds(client, parentId)

						if children != nil  {
							client.CompleteItems(children, true)
							client.CompleteItems([]int64{parentId}, false)
						} else {
							ids := []int64{c.Int64("id")}
							client.CompleteItems(ids, false)
						}

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
