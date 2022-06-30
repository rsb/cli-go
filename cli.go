// Package cli allows you to integrate cli interactions into your application
package cli

import (
	flag "github.com/rsb/pflag"
)

// FParseErrWhitelist configures Flag parse errors to be ignored
type FParseErrWhitelist flag.ParseErrorsWhitelist

// ControlUsageFn is the function signature for the usage closure.
type ControlUsageFn func(*Cmd) error

// ControlFlagErrorFn is a function signature to allow user to control when
// the parsing of a flag returns an error
type ControlFlagErrorFn func(*Cmd, error) error

// ControlHelpFn is a function signature to allow users to control help
type ControlHelpFn func(*Cmd, []string)

// GlobalNormalizeFlagFn defined the signature for the global normalization
// function that can be used on every pflag set and children commands
type GlobalNormalizeFlagFn func(f *flag.FlagSet, name string) flag.NormalizedName
