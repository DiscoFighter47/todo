package inmemory

// BlackList ...
type BlackList struct {
	repo map[string]struct{}
}

// NewBlackList ...
func NewBlackList() *BlackList {
	return &BlackList{
		repo: map[string]struct{}{},
	}
}

// AddKey ...
func (list *BlackList) AddKey(key string) error {
	list.repo[key] = struct{}{}
	return nil
}

// Contains ...
func (list *BlackList) Contains(key string) (bool, error) {
	_, ok := list.repo[key]
	return ok, nil
}
