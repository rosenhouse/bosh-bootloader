package commands

import (
	"fmt"
	"io"
	"strings"

	"github.com/pivotal-cf-experimental/bosh-bootloader/storage"
)

const USAGE = `
Usage:
  bbl [GLOBAL OPTIONS] COMMAND [OPTIONS]

Global Options:
  --help    [-h] "print usage"
  --version [-v] "print version"

  --aws-access-key-id     "AWS AccessKeyID value"
  --aws-secret-access-key "AWS SecretAccessKey value"
  --aws-region            "AWS Region"
  --state-dir             "Directory that stores the state.json"

Commands:
  help                                          "print usage"
  version                                       "print version"
  unsupported-deploy-bosh-on-aws-for-concourse  "deploys a BOSH Director on AWS"
  destroy [--no-confirm]                        "tears down a BOSH Director environment on AWS"
`

type Usage struct {
	stdout io.Writer
}

func NewUsage(stdout io.Writer) Usage {
	return Usage{stdout}
}

func (u Usage) Execute(globalFlags GlobalFlags, subcommandFlags []string, state storage.State) (storage.State, error) {
	u.Print()
	return state, nil
}

func (u Usage) Print() {
	fmt.Fprint(u.stdout, strings.TrimSpace(USAGE))
}
