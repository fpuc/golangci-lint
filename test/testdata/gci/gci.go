//args: -Egci
//config_path: testdata/configs/gci.yml
package gci

import (
	"fmt"

	"github.com/fpuc/golangci-lint/pkg/config"
	"github.com/pkg/errors"
)

func GoimportsLocalTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = errors.New("")
}
