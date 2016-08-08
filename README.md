# toggl

This is a toggle command line utility. It is written in GO.

## How to use

### Show statistic

    toggl account

This command prints information about your timezone, the number of workspaces you are subscribed to and the number of tags you have. Appending `--time` shows the currently running time entry ( if one )

### Managing projects

With `toggl project` you can run command onh projects. You can use one of the following commands:

* list
* create
* delete

#### list

List shows a list of available projects, for each a line. The number in brackets is the project ID. A sample looks like this:

    Project (2345234)

#### create

To create a new project you have to provide the project name with the `--name` flag.

#### delete

To delete a project you have to provide the project name with the `--name` flag.

### Tracking time

Use the verb `time` to handle starting and stopping rime tracking.

#### start

You have to provide the `--desc` flag as a description and you can provide the `--project` flag to set a project you are working on

#### stop

Just call stop to stop working

## Acknowledgements

This project uses go-toggl from https://github.com/jason0x43/go-toggl by downloading and changing the code. See `get_toggl` script

## History

### TBD

* Initial version
