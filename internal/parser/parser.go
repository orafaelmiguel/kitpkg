package parser

func Parse(input string) []string {
	var args []string
	var current string
	inQuotes := false
	var quoteChar byte

	for i := 0; i < len(input); i++ {
		c := input[i]

		switch c {

		case '"', '\'':
			if inQuotes && c == quoteChar {
				inQuotes = false
			} else if !inQuotes {
				inQuotes = true
				quoteChar = c
			} else {
				current += string(c)
			}

		case ' ':
			if inQuotes {
				current += " "
			} else if current != "" {
				args = append(args, current)
				current = ""
			}

		default:
			current += string(c)
		}
	}

	if current != "" {
		args = append(args, current)
	}

	return args
}