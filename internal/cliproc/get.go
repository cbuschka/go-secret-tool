package cliproc

import (
	"fmt"
	"github.io/cbuschka/go-secret-tool/internal/backend"
)

func GetSecret(name string) error {
	b, err := backend.GetBackend()
	if err != nil {
		return err
	}
	defer b.Close()

	secrets, err := b.ListSecrets()
	if err != nil {
		return err
	}

	for _, secret := range secrets {
		if name == secret.Name() {
			fmt.Printf("%s", secret.Value())
			return nil
		}
	}

	return fmt.Errorf("secret not found")
}
