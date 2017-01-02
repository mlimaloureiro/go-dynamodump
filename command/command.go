package command

// DumpRestoreCommand is the type with the CLI command
type DumpRestoreCommand struct {
	RestoreConfig config.Restore
	DumpConfig config.Dump
}

func NewCommand() *DumpRestoreCommand {
	return &DumpRestoreCommand{}
}
