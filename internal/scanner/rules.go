package scanner

import "regexp"

// rule defines a single threat signature
type rule struct {
	Name    string
	Pattern *regexp.Regexp
	Message string
}

// this is basic but gets the job done for a hackathon
// in a real prod env, we'd use AST parsing, but regex is fast and covers 80%
var commonRules = []rule{
	{
		Name:    "RmRf",
		Pattern: regexp.MustCompile(`rm\s+-r.*f.*`),
		Message: "Destructive file operations blocked",
	},
	{
		Name:    "ForkBomb_Bash",
		Pattern: regexp.MustCompile(`:\(\)\s*\{\s*:\|:&\s*\};:`),
		Message: "Fork bomb signature detected",
	},
	{
		Name:    "Hardcoded_AWS",
		Pattern: regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
		Message: "Possible AWS secret detected",
	},
}

var languageRules = map[string][]rule{
	"python": {
		{
			Name:    "OsSystem",
			Pattern: regexp.MustCompile(`os\.system\s*\(`),
			Message: "os.system calls are not allowed",
		},
		{
			Name:    "Subprocess",
			Pattern: regexp.MustCompile(`subprocess\.(Popen|call|run|check_output)`),
			Message: "subprocess execution is blocked",
		},
		{
			Name:    "Eval",
			Pattern: regexp.MustCompile(`eval\s*\(`),
			Message: "eval() is too dangerous",
		},
		{
			Name:    "PtySpawn",
			Pattern: regexp.MustCompile(`pty\.spawn`),
			Message: "pty spawning is blocked",
		},
	},
	"javascript": {
		{
			Name:    "ChildProcess",
			Pattern: regexp.MustCompile(`require\(['"]child_process['"]\)`),
			Message: "child_process is not allowed",
		},
		{
			Name:    "EvalJS",
			Pattern: regexp.MustCompile(`eval\s*\(`),
			Message: "eval() blocked",
		},
		{
			Name:    "FsWrite",
			Pattern: regexp.MustCompile(`fs\.writeFileSync`),
			Message: "file system writes are restricted", // read might be okay for some judges, but write is risky
		},
	},
	"go": {
		{
			Name:    "OsExec",
			Pattern: regexp.MustCompile(`os/exec`),
			Message: "os/exec import is blocked",
		},
		{
			Name:    "Syscall",
			Pattern: regexp.MustCompile(`syscall\.(Syscall|Exec)`),
			Message: "raw syscalls blocked",
		},
	},
	"cpp": {
		{
			Name:    "SystemCall",
			Pattern: regexp.MustCompile(`system\s*\(`),
			Message: "system() call blocked",
		},
		{
			Name:    "FstreamOut",
			Pattern: regexp.MustCompile(`ofstream`),
			Message: "File writing is restricted",
		},
	},
}
