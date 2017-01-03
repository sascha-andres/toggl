# toggl

This is a toggle command line utility. It is written in GO.

[![Go Report Card](https://goreportcard.com/badge/github.com/sascha-andres/toggl)](https://goreportcard.com/report/github.com/sascha-andres/toggl) [![Build Status](https://travis-ci.org/sascha-andres/toggl.svg?branch=master)](https://travis-ci.org/sascha-andres/toggl) [![codebeat badge](https://codebeat.co/badges/66adec90-4ccb-4d6a-94c6-484f11bd4c2a)](https://codebeat.co/projects/github-com-sascha-andres-toggl)

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

This project uses go-toggl from https://github.com/jason0x43/go-toggl

_Note:_ To disable logging I forked the library and created a pull request. Until this pull request has been accepted, I am using my fork located at https://github.com/sascha-andres/go-toggl

## Code ##

Code is open source under the Apache 2.0 License. You can obtain it at https://github.com/sascha-andres/toggl

If you want to contribute feel free to open an issue

## Help

Feel free to contact me by creating an issue on https://github.com/sascha-andres/toggl/issues.
You can connect to me using Twitter at https://twitter.com/livingit_de.

## History

|Version|Authors|Description|
|---|---|---|
|20170103|Sascha Andres|Forked library|
|20161215|Sascha Andres|Initial|
