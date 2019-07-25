package inmemory

// Datastore ...
type Datastore struct {
	*UserStore
}

// NewDatastore ...
func NewDatastore() *Datastore {
	return &Datastore{
		NewUserStore(),
	}
}
