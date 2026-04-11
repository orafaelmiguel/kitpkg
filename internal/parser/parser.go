package parser

func Parse(input string) []string {
	var args []string
	var current string
	inQuotes := false

	for i := 0; i < len(input); i++ {
		char := input[i]

		switch char {
		case '"':
			inQuotes = !inQuotes

		case ' ':
			if inQuotes {
				current += string(char)
			} else if current != "" {
				args = append(args, current)
				current = ""
			}

		default:
			current += string(char)
		}
	}

	if current != "" {
		args = append(args, current)
	}

	return args
}