package models

type Entity interface {
	Prepare()
	Validate() error
}
