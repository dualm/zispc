package zispc

type ProcessData interface {
	Encode() ([]byte, error)
}
