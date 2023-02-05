package cfg_test

import (
	"fmt"
	"testing"

	"github.com/grin-ch/grin-auth/cfg"
)

func TestInitConfig(t *testing.T) {
	cfg.InitConfig()
	fmt.Printf("%+v", cfg.Config)
}
