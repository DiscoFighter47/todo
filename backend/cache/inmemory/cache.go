package inmemory

// Cache ...
type Cache struct {
	*BlackList
}

// NewCache ...
func NewCache() *Cache {
	return &Cache{
		NewBlackList(),
	}
}
