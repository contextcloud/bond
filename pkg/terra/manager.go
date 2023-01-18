package terra

type Manager interface{}

type manager struct{}

func NewManager() (Manager, error) {
	return &manager{}, nil
}
