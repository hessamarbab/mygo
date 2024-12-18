package models

type Filler interface {
	Fill() Filler
}
