package dictionary

//go:generate rm -f ./dictionary.pb.go
//go:generate go run generator/generator.go -o ./dictionary.pb.go

var (
	words map[string]bool
)

func Have(target string) bool {
	return words[target]
}
