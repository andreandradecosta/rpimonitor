package rpimonitor

//Device implemented by components that read hw data.
type Device interface {
	StatusReader
	SampleReader
}
