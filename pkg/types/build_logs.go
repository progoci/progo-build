package types

// BuildLogs represent the message sent to the Loom service to store logs for
// a build.
type BuildLogs struct {
	BuildID  string
	TaskUUID string
	CmdID    string
	Cmd      string
	Logs     string
	// Whether this corresponds to the first logs of the current command.
	First bool
}
