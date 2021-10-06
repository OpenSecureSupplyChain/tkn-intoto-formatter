# tkn-intoto-formatter

A common library to convert any tekton resource to intoto attestation format. 

In this library, the idea is to support intoto attestation conversion for all tekton resources including

- Pipeline
- Task
- Pipelinerun
- Taskrun

Currently, the conversion is supported for `Task` and `Pipeline` to intoto layout.

## Usage

Currently, `tkn-intoto-formatter` is also made available as a stand-alone CLI. You can follow following procedure to install CLI:

```
git clone 
cd tkn-intoto-formatter
make
```

This should create an executable binary `tkn-intoto-formatter`. You can add this binary to your local PATH.

### Quick start

1. Get Help

```
% tkn-intoto-formatter -h
tkn-intoto-formatter is tool to manage various attestation functions, including
		conversion to intoto format and comparisons.

Usage:
  tkn-intoto-formatter [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  convert     converts yaml tekton spec to specified format
  help        Help about any command
  version     tkn-intoto-formatter version

Flags:
      --config string   config file (default is $HOME/.tkn-intoto-formatter.yaml)
  -h, --help            help for tkn-intoto-formatter
  -t, --toggle          Help message for toggle

Use "tkn-intoto-formatter [command] --help" for more information about a command.
```

2. Convert static tekton resources to intoto format

```
% tkn-intoto-formatter convert -i sample-pipeline/task-bom.yaml -f ./task-bom-attest.json
```

## WIP

1. Support Pipelinerun and Taskrun 
2. Try to capture attestation flow from event source -> pipeline -> tasks
