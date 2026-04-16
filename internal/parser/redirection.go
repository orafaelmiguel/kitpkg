package parser

import "strings"

type Redirection struct {
	Command string
	File    string
	Append  bool
}

func ParseRedirection(input string) Redirection {
	if strings.Contains(input, ">>") {
		parts := strings.SplitN(input, ">>", 2)
		return Redirection{
			Command: strings.TrimSpace(parts[0]),
			File:    strings.TrimSpace(parts[1]),
			Append:  true,
		}
	}

	if strings.Contains(input, ">") {
		parts := strings.SplitN(input, ">", 2)
		return Redirection{
			Command: strings.TrimSpace(parts[0]),
			File:    strings.TrimSpace(parts[1]),
			Append:  false,
		}
	}

	return Redirection{
		Command: input,
	}
}