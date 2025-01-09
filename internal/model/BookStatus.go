package model

type BookStatusEnum string

const (
	StatusToRead    BookStatusEnum = "to_read"
	StatusReading   BookStatusEnum = "reading"
	StatusRead      BookStatusEnum = "read"
	StatusAbandoned BookStatusEnum = "abandoned"
)
