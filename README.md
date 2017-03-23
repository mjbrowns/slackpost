# slackpost

A highly portable utility to post a message to a slack channel.

## Synopsis

This utility, written in go, sends a message to a slack channel by constructing
a message object containing all availabl slack incoming webhook API capabilities

## Usage

**slackpost [options] [command] [arguments]**

#### Available Commands:

| Command | Description |
| --------- | ------- |
|  addfield |   Add fields to the current attachment|
|  attach  |    Attach a text string to the message|
|  clean   |    Cleans up the message object|
|  help    |    Help about any command|
|  init    |  Initializes the message object|
|  send    |   Sends the message|
|  version |    Show version|

#### Global Flags

*Global flags are usable with any subcommand*

| Short | Long | Description |
| - | - | - |
|     | --config |     config file (default is $HOME/.slackpost.json)
|  -m | --message-id | message id number.  Can be anything.  Defaults to PID of caller
|  -v | --verbose  | Enable verbose messages
|  -w | --workdir |    working directory to use to store message objects

### init

Initializes the message object

####Usage:
**slackpost init \<webhook> \<message> [flags]**

#### Arguments
| Name | Value |
| --------- | ------- |
|  webhook  | URL of slack channel incoming webook |
|  message |  Message to send |

#### Flags

| Short | Long | Description |
| - | - | - |
|  -c| --channel |  Name of channel to send to (ignored if passing a channel hook)|
|  -i| --icon |     Icon name to post with the message|
|  -I| --icon-url | URL of Icon to post with the message|
|  -u| --username | User Name to associate with the message|

#### Example:
```
  slackpost init https://slack.foo "This is my message"
```

### attach

Attach a text string to the message.

	This function will add a string as an attachment to the message.
	There can be unlimited attachments to a message, and each attachment has many
	options available to customize how the attachment will appear.


#### Usage:

**slackpost attach \<text> [flags]**

#### Arguments

| Name | Value |
| --------- | ------- |
|  text | string message to attach |

#### Flags

| Short | Long | Description |
| - | - | - |
|  -a| --author |      Name of author
|  -I| --author-icon | URL to a 16x16px icon to display by the author's name
|  -A| --author-link | URL to hyperlink to the author's name
|  -c| --color |       Name of color to use with this attachment
|  -B| --fallback |    Fallback string to use for clients that can't display the message
|  -f| --footer |      Text that appears as a footer to the attachment.
|  -F| --footer-icon | URL of an icon to display with the footer.
|  -i| --image |       URL of image to display inside the attachment.
|  -p| --pretext |     String to display above the message attachment block
|  -b| --thumb |       URL of a thumbnail to display to the top-right of the attachment.
|  -d| --timestamp |   Epoch-time date/time stamp for the attachment.
|  -t | --title | String to display as the title of the attachment.  Displays as larger, bolder text|
|  -T | --title-link | URL to hyperlink to the attachment title.|

#### Example:
```
slackpost attach "$(<logfile.txt)" -c "#abcdef" -a "yourname"
```

### addfield

Add fields to the current attachment.
```
  Message atachments can have multiple fields (see slack message documentation).
  This function adds a few field to the most recently added attachment.
```
#### Usage

**slackpost addfield \<title> \<value> [flags]**

#### Arguments

| Name | Value |
| --------- | ------- |
| title | Title string associated with this field |
| value | Value string associated with this field |

#### Flags

| Short | Long | Description |
| - | - | - |
| -s | `--short` | If set, indicates that the title/value pair is short enough to be displayed vertically with other fields |

#### Example
```
slackpost addfield AssignedTo "Jane Doe"
```

### send

Sends the constructed message to the slack channel

#### Usage

**slackpost send [flags]**

### clean

Cleans up the temporary file(s) used to create the message object.

	This function is normally performed automatically after sending the object,
	but in some cases you may need to clean up the message manually.

#### Usage

**slackpost clean [flags]**

#### Flags

| Short | Long | Description |
| - | - | - |
|  -a | --all | If set, will clean up any detected messages| 
| | --force |  When using -a|--all this is required to actually delete files|
