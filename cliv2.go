package ecspresso

import (
	"fmt"

	"github.com/alecthomas/kong"
)

func ParseCLIv2(args []string) (string, *CLIOptions, error) {
	// compatible with v1
	if len(args) == 0 || len(args) > 0 && args[0] == "help" {
		args = []string{"--help"}
	}

	var opts CLIOptions
	parser, err := kong.New(&opts, kong.Vars{"version": Version})
	if err != nil {
		return "", nil, fmt.Errorf("failed to new kong: %w", err)
	}
	c, err := parser.Parse(args)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse args: %w", err)
	}
	sub := c.Command()

	for _, envFile := range opts.Envfile {
		if err := ExportEnvFile(envFile); err != nil {
			return sub, &opts, fmt.Errorf("failed to load envfile: %w", err)
		}
	}

	opts.Option = &Option{
		ConfigFilePath: opts.Config,
		Debug:          opts.Debug,
		ExtStr:         opts.ExtStr,
		ExtCode:        opts.ExtCode,
	}
	switch sub {
	case "init":
		opts.Option.InitOption = opts.Init
	}
	return sub, &opts, nil
}
