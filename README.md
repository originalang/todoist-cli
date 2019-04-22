# todoist CLI

## Installation

**Note:** Ensure that the Go [workspace](https://golang.org/doc/code.html#Workspaces) is set up before proceeding.

Run these commands:

```
$ go get github.com/originalang/todoist-cli
$ cd $GOPATH/src/github.com/originalang/todoist-cli
```
Once inside the downloaded directory, run the go installation command:

```
$ go install   
```
This will save an executable file in the `$GOPATH/bin` directory. Make sure that this directory has been added to your `$PATH` variable to ensure that you can use the `todoist-cli` command from any location.

## Configuration

1. Login to your [todoist](https://todoist.com/) account
2. Navigate to [this](https://todoist.com/prefs/integrations) page and follow the instructions to issue a new API token
3. Create a new environment variable called `TODOIST_TOKEN` and set it to your API token

## Usage

For detailed usage instructions, use the ```-h``` flag with the relevant command

### Project Commands

Retrieve a list of user projects:

```
$ todoist-cli project list 
```

Add a new project:

```
$ todoist-cli project add -n "New Project"
```

Delete a project:

```
$ todoist-cli project delete -n "Project_Name"
```

### Item Commands

The todoist Sync API refers to tasks as items, so this project does the same.

Retrieve and display a list of items for a particular project. These will be displayed hierarchically with the correct indent.

```
$ todoist-cli item list -n "Project_Name"
```

Add an item to a project. If no project is specified, the item is added to the "Inbox". The indent can be specified using the ```-i``` flag (1, 2, 3, or 4).

```
$ todoist-cli item add -n "Project_Name" -c "New item: Do a task" -d "tomorrow @ 1pm"
```

Delete a project. The item ID is needed to execute this command. This is displayed next to the item when you list items for a project.

```
$ todoist-cli item delete --id 294573453
```

Complete a project. The ID is needed for this command as well. If the item is a subtask, it will be completed and displayed with a check next to it. If it is a main item (no indent), it will be completed and moved to history.

```
$ todoist-cli item complete --id 294573453
```
