package internal

var (
	id uint64
)

func GetNextID() uint64 {
	id++
	return id
}
