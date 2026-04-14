package parser

func ParseFlags(args []string) (map[string]bool, []string) {
	flags := make(map[string]bool)
	var params []string

	for _, arg := range args {
		if len(arg) > 1 && arg[0] == '-' {
			for _, ch := range arg[1:] {
				flags[string(ch)] = true
			}
		} else {
			params = append(params, arg)
		}
	}

	return flags, params
}