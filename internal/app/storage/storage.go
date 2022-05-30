package storage

type Storage interface {
	Save(url string) (int64, error)
	Get(id int64) (string, error)
}
