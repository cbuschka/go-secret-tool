package cli

import (
	"fmt"
	"github.io/cbuschka/go-secret-tool/internal/backend"
)

func ListSecrets() error {
	b, err := backend.GetBackend()
	if err != nil {
		return err
	}
	defer b.Close()

	secrets, err := b.ListSecrets()
	if err != nil {
		return err
	}

	for index, secret := range secrets {
		fmt.Printf("%d %s\n", index, secret.Name())
	}

	return nil
}
