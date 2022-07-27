//golangcitest:args -Egci
//golangcitest:config_path testdata/configs/gci.yml
package gci

import (
	"github.com/fpuc/golangci-lint/pkg/config"
	"github.com/pkg/errors"
	"fmt"
)

func GoimportsLocalTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = errors.New("")
}
