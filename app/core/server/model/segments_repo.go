package model

type SegmentsRepository interface {
	Create(segment string) error
	Delete(segment string) error
}
