package types

type StorageBucket interface {
	Add(string, string)
	Del(...string) int
	Find(...string) map[string]Value
	Indexes(string) []string
}
