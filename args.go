package cli

type PositionalArgs func(cmd *Cmd, args []string) error
