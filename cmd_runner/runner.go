package cmd_runner

import (
	"fmt"
	"sort"

	"github.com/centurylinkcloud/clc-go-cli/auth"
	"github.com/centurylinkcloud/clc-go-cli/command_loader"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/formatter_provider"
	"github.com/centurylinkcloud/clc-go-cli/model_adjuster"
	"github.com/centurylinkcloud/clc-go-cli/model_loader"
	"github.com/centurylinkcloud/clc-go-cli/model_validator"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"github.com/centurylinkcloud/clc-go-cli/parser"
	"github.com/centurylinkcloud/clc-go-cli/state"
)

func Run(args []string) string {
	if len(args) == 0 {
		return help()
	}
	if len(args) == 1 && args[0] == "--help" {
		return help()
	}
	cmdArg := ""
	optionArgs := args[1:]
	if len(args) >= 2 {
		cmdArg = args[1]
		optionArgs = args[2:]
	}
	resource, err := command_loader.LoadResource(args[0])
	if err != nil {
		return err.Error()
	}
	cmd, err := command_loader.LoadCommand(resource, cmdArg)
	if err != nil {
		if cmdArg == "--help" {
			return command_loader.GetCommandsWithDescriptions(resource)
		}
		return err.Error()
	}
	if cmd.Command() == "" {
		optionArgs = args[1:]
	}
	parsedArgs, err := parser.ParseArguments(optionArgs)
	if err != nil {
		return err.Error()
	}
	yes, filename, err := options.AreToBeTakenFromFile(parsedArgs)
	if err != nil {
		return err.Error()
	}
	if yes {
		parsedArgs, err = state.ArgumentsFromJSON(filename)
		if err != nil {
			return err.Error()
		}
	}
	yes, err = options.AreToBeSaved(parsedArgs)
	if err != nil {
		return err.Error()
	}
	if yes {
		output, err := state.ArgumentsToJSON(parsedArgs, cmd.InputModel())
		if err != nil {
			return err.Error()
		}
		return output
	}
	options, err := options.ExtractFrom(parsedArgs)
	if err != nil {
		return err.Error()
	}
	if options.Help {
		return cmd.ShowHelp()
	}
	conf, err := config.LoadConfig()
	if err != nil {
		return err.Error()
	}
	if cmd.Resource() == "login" {
		return cmd.(*commands.Login).Login(options, conf)
	}
	err = model_loader.LoadModel(parsedArgs, cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	datacenter.ApplyDefault(cmd.InputModel(), conf)
	err = model_validator.ValidateModel(cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	err = model_adjuster.ApplyDefaultBehaviour(cmd.InputModel())
	if err != nil {
		return err.Error()
	}
	if cmd.IsOffline() {
		res, err := cmd.ExecuteOffline()
		if err != nil {
			return err.Error()
		}
		return res
	}
	cn, err := auth.AuthenticateCommand(options, conf)
	if err != nil {
		return err.Error()
	}
	err = model_adjuster.InferID(cmd.InputModel(), cn)
	if err != nil {
		return err.Error()
	}
	err = cmd.Execute(cn)
	if err != nil {
		return err.Error()
	}
	err = state.SaveLastResult(cmd.OutputModel())
	if err != nil {
		return err.Error()
	}
	f, err := formatter_provider.GetOutputFormatter(options, conf)
	if err != nil {
		return err.Error()
	}
	outputModel := cmd.OutputModel()
	if messagePtr, ok := outputModel.(*string); ok {
		return *messagePtr
	}
	detyped, err := parser.ConvertToMapOrSlice(outputModel)
	if err != nil {
		return err.Error()
	}
	if options.Filter != "" {
		filtered, err := parser.ParseFilter(detyped, options.Filter)
		if err != nil {
			return err.Error()
		} else if filtered == nil {
			return "No results found for the given filter."
		} else {
			detyped = filtered
		}
	}
	if options.Query != "" {
		queried, err := parser.ParseQuery(detyped, options.Query)
		if err != nil {
			return err.Error()
		} else if queried == nil {
			return "No results found for the given query."
		} else {
			detyped = queried
		}
	}
	output, err := f.FormatOutput(detyped)
	if err != nil {
		return err.Error()
	}
	return output
}

func help() string {
	res := `Usage: clc <resource> [<command>] [options and parameters]

clc <resource> --help
    get a list of all available commands for the given
    resource if any or to get a direct resource description
clc <resource> <command> --help
    get a command description and a list of all available
    parameters for the given command

Available resources:

`
	resources := command_loader.GetResources()
	sort.Strings(resources)
	for _, rsr := range resources {
		res += fmt.Sprintf("\t%v\n", rsr)
	}
	return res
}
