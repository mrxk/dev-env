package cmd

import (
	"github.com/mrxk/dev-env/pkg/constants"
	"github.com/spf13/pflag"
)

func typesFromFlags(flags *pflag.FlagSet) ([]string, error) {
	types, err := addOptionNameIfPresent([]string{}, flags, constants.AllOption)
	if err != nil {
		return nil, err
	}
	types, err = addOptionNameIfPresent(types, flags, constants.ConnectOption)
	if err != nil {
		return nil, err
	}
	types, err = addOptionNameIfPresent(types, flags, constants.RunOption)
	if err != nil {
		return nil, err
	}
	types, err = addOptionNameIfPresent(types, flags, constants.SpawnOption)
	if err != nil {
		return nil, err
	}
	if len(types) == 0 {
		return []string{constants.ConnectOption}, nil
	}
	return types, nil
}

func addOptionNameIfPresent(types []string, flags *pflag.FlagSet, option string) ([]string, error) {
	value, err := flags.GetBool(option)
	if err != nil {
		return nil, err
	}
	if value {
		types = append(types, option)
	}
	return types, nil
}
