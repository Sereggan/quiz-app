package quizrepository

import "fmt"

type RepositoryError struct {
	message string
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("Repository error happened: %d", e.message)
}
