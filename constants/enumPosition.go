package constants

type Position int

const (
	Student Position = iota + 1
	Teacher
	Admin
)