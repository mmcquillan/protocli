# protocli

This is a CLI prototyping tool using the [Matcher Format](https://github.com/mmcquillan/matcher) with response variable substitution.


## Install

On the mac, you can install via brew as:

`brew install mmcquillan/tools/protocli`

Launch by passing in a config file:

`protocli <config>`


## Commands

The only native command inside protocli is `?` which lists all possible command matches.


## Config

The configuration is a YAML based file. Examples can be found in the [here](examples).

```
---
prompt: "> "
commands:
-
  command: "do version"
  response: "0.2.0"
-
  command: "do plan <file>"
  responses:
  -
    response: "Planning ${file}..."
  -
    response: "Complete"
    delay: 2
-
  command: "do apply <file>"
  responses:
  -
    response: "Applying ${file}..."
  -
    response: "Error: applying ${file}"
    color: "red"
    delay: 3
-
  command: "do plan [...]"
  response: "do plan <file>"
-
  command: "do apply [...]"
  response: "do apply <file>"
-
  command: "do [...]"
  response: "do <plan|apply> <file>"
```

### prompt
- Description: A prompt to show in the protocli interaction
- Default: `> `
- Type: string

### default
- Description: A default response when no command is matched
- Default: none
- Type: string

### commands
- Description: A list of commands to respond to
- Default: none
- Type: list

### command
- Description: A command to respond to, using the [Matcher Format](https://github.com/mmcquillan/matcher)
- Default: none
- Type: string

### response
- Description: A response to show when the command is written (you can also list responses)
- Default: none
- Type: string

### color
- Description: Color to print the response in (red | green | yellow | blue | magenta | cyan | white)
- Default: white
- Type: string

### delay
- Description: Number of seconds to delay before showing response
- Default: 0
- Type: int


## Releases

v0.2.0
- Added custom prompts
- Added proper exit handling
- Added config command line paramater
- Added output color
- Added dynamic config reload
- Updated matcher library for short flags
- Added delay option to response
- Added default resonse for no match

v0.1.0
- Initial Release
