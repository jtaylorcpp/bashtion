package provisioning

import (
	"fmt"
	"io/ioutil"
)

func ProvisionFromFile(filepath string) {
	cmdfile, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	cmd := CommandFromYamlString(cmdfile)

	fmt.Printf("parsed command: %v\n", cmd)

	switch cmd.Kind {
	case "script":
		switch cmd.Type {
		case "bash":
			runBASHCmd(cmd.Content)
		default:
			fmt.Printf("unkown script type: %s\n", cmd.Type)
		}
	default:
		fmt.Printf("unknown kind: %s\n", cmd.Kind)
	}

}
