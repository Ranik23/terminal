package errors

import "fmt"


var (
	ErrAlreadyRoot = fmt.Errorf("already at the root directory")
	ErrInvalidDirectory = fmt.Errorf("invalid directory")
	ErrNotDirectory = fmt.Errorf("not a directory")
	ErrFileStat =  fmt.Errorf("failed to get info about the file")
)