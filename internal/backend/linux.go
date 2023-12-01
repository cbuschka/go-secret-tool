package backend

import (
	"r00t2.io/gosecret"
)

type linuxBackend struct {
	service *gosecret.Service
}

type linuxSecret struct {
	name  string
	value string
}

const (
	collectionName string = "login"
)

func (b *linuxBackend) ListSecrets() ([]Secret, error) {

	collection, err := b.service.GetCollection(collectionName)
	if err != nil {
		return nil, err
	}

	itemResults, err := collection.Items()
	if err != nil {
		return nil, err
	}

	secrets := []Secret{}
	for _, result := range itemResults {
		if result.Secret.ContentType == "text/plain" {
			bs, err := result.Secret.Value.MarshalJSON()
			if err != nil {
				return nil, err
			}
			value := string(bs)
			secrets = append(secrets, Secret(&linuxSecret{name: result.LabelName, value: value}))
		}
	}

	return secrets, nil
}

func (b *linuxBackend) Close() error {
	return b.service.Close()
}

func newLinuxBackend() (*linuxBackend, error) {
	service, err := gosecret.NewService()
	if err != nil {
		return nil, err
	}

	return &linuxBackend{service: service}, nil
}

func (s *linuxSecret) Name() string {
	return s.name
}

func (s *linuxSecret) Value() string {
	return s.value
}
