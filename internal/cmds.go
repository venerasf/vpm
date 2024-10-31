package internal

type Suggest struct {
	Text        string
	Description string
}

func Commands() []Suggest {
	return []Suggest {
		{Text: "search", 	Description: "Search for scripts with a pattern"},
		{Text: "install",	Description: "Install a script"},
		{Text: "sync",		Description: "Synchronize with remote repository"},
		{Text: "verify",	Description: "Verify package signature"},
	}
}

func Usage(cmds []string) string {
	return `
Venera Package Manager:
	'vpm' Call vpm commands.
	'vpm search <pattern>' List matching scripts.
	'vpm install <script>' Install a script.
		Usage:'vpm install /path/to/the/script.lua'
	'vpm sync' Synchronize local repository with remote repository.
	'vpm verify' Verify the signature of the configured remote repository.
`
}

func EntrypointCMD() string {
	return "vpm"
}