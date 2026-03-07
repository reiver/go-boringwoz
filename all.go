package boringwoz

var (
	allAdjectives []string
	allSurnames   []string
)

var (
	maxNameLength int
)

func init() {
	allAdjectives = append(allAdjectives, wozAdjectives[:]...)
	allAdjectives = append(allAdjectives, vanAdjectives[:]...)

	allSurnames = append(allSurnames, wozSurnames[:]...)
	allSurnames = append(allSurnames, vanSurnames[:]...)

	{
		{
			var max int

			for _, str := range allAdjectives {
				// We calculate the string length using len() because all characters in the 'adjectives' and 'surnames' are ASCII (and thus 1 byte long).
				length := len(str)
				if max < length {
					max = length
				}
			}

			maxNameLength += max
		}

		maxNameLength += len("_")

		{
			var max int

			for _, str := range allSurnames {
				// We calculate the string length using len() because all characters in the 'adjectives' and 'surnames' are ASCII (and thus 1 byte long).
				length := len(str)
				if max < length {
					max = length
				}
			}

			maxNameLength += max
		}

		maxNameLength += len("_")

		maxNameLength += 2

		if maxNameLength < 32 {
			maxNameLength = 32
		}
	}
}
