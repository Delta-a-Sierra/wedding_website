package config

import "fmt"

type Config struct {
	Log                Log
	PresentationMethod PresentationMethod
}

type PresentationMethod string

const HTMXPM PresentationMethod = "htmx"

func (p PresentationMethod) Validate() error {
	switch p {
	case HTMXPM:
		return nil
	default:
		return NewValidationError(fmt.Sprintf("invalid presentation method provided, method %s not supported", p))
	}
}

type Log struct {
	OutputFile string
	Level      string
}
