//golangcitest:args -Egci
//golangcitest:config_path testdata/configs/gci.yml
package gci

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fpuc/golangci-lint/pkg/config"
)

func GoimportsLocalTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = errors.New("")
}
