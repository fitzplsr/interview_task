package parser

type Parser interface {
	GetTargetCount() (int64, error)
	GetSourceName() string
}
