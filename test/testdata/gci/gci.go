//args: -Egci
//config: linters-settings.gci.local-prefixes=github.com/fpuc/golangci-lint
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
