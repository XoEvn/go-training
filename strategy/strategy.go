package strategy

type StorageStrategy interface {
	Save(name string, data []byte) error
}
