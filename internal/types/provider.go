package types

type Provider interface {
	Summarize(string) string
}
