package boringwoz

var (
	allAdjectives []string
	allSurnames   []string
)

func init() {
	allAdjectives = append(allAdjectives, wozAdjectives[:]...)
	allAdjectives = append(allAdjectives, vanAdjectives[:]...)

	allSurnames = append(allSurnames, wozSurnames[:]...)
	allSurnames = append(allSurnames, vanSurnames[:]...)
}
