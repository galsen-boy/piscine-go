package piscine

func ConcatParams(args []string) string {
	answer := args[0]
	for i := 1; i < len(args); i++ {
		answer += "\n" + args[i]
	}
	return answer
}
