# protocli

This is a CLI prototyping tool using the [Matcher Format](https://github.com/mmcquillan/matcher) with response variable substitution.

## Install

On the mac, you can install via brew as.

`brew install mmcquillan/tools/protocli`



## Config

The configuration is a YAML based file, like the one here:

```
---
prompt: "> "
commands:
-
  command: "version"
  response: "0.1"
-
  command: "do plan <file>"
  response: "Planning ${file}...
	\nComplete"
```

### prompt
- Description: A prompt to show in the protocli interaction
- Default: `> `
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
- Description: A response to show when the command is written
- Default: none
- Type: string


## Releases

v0.2.0
- Added custom prompts


v0.1.0
- Initial Release
