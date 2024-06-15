package io_manager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(any) error
}
