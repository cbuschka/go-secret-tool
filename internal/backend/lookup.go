package backend

func GetBackend() (Backend, error) {
	b, err := newLinuxBackend()
	if err != nil {
		return nil, err
	}
	return Backend(b), nil
}
