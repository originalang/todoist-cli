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
