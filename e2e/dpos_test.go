package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/loomnetwork/loomchain/e2e/common"
)

func TestContractDPOS(t *testing.T) {
	tests := []struct {
		name       string
		testFile   string
		validators int // TODO this is more like # of nodes than validators
		// # of validators is set in genesis params...
		accounts int
		genFile  string
		yamlFile string
	}{
		{"dpos-delegation", "dpos-delegation.toml", 4, 10, "dpos-delegation.genesis.json", "dpos-test-loom.yaml"},
		{"dpos-2", "dpos-2-validators.toml", 2, 10, "dpos.genesis.json", "dpos-test-loom.yaml"},
		{"dpos-2-r2", "dpos-2-validators.toml", 2, 10, "dpos.genesis.json", "dpos-test-loom.yaml"},
		{"dpos-4", "dpos-4-validators.toml", 4, 10, "dpos.genesis.json", "dpos-test-loom.yaml"},
		{"dpos-4-r2", "dpos-4-validators.toml", 4, 10, "dpos.genesis.json", "dpos-test-loom.yaml"},
		// {"dpos-8", "dpos-8-validators.toml", 8, 10, "dpos.genesis.json", "dpos-test-loom.yaml"},
		{"dpos-elect-time", "dpos-elect-time-2-validators.toml", 2, 10, "dpos-elect-time.genesis.json", "dpos-test-loom.yaml"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config, err := common.NewConfig(test.name, test.testFile, test.genFile, test.yamlFile, test.validators, test.accounts)
			if err != nil {
				t.Fatal(err)
			}

			binary, err := exec.LookPath("go")
			if err != nil {
				t.Fatal(err)
			}
			// required binary
			cmd := exec.Cmd{
				Dir:  config.BaseDir,
				Path: binary,
				Args: []string{binary, "build", "-tags", "evm", "-o", "example-cli", "github.com/loomnetwork/go-loom/examples/cli"},
			}
			if err := cmd.Run(); err != nil {
				t.Fatal(fmt.Errorf("fail to execute command: %s\n%v", strings.Join(cmd.Args, " "), err))
			}

			if err := common.DoRun(*config); err != nil {
				t.Fatal(err)
			}

			// pause before running the next test
			time.Sleep(500 * time.Millisecond)
		})
	}
}
