package storage

type (
	Storage interface {
		AddCandidate() (string, error)
	}
)