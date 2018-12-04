package provisioning

import (
	"fmt"
	"os/exec"
	"strings"
)

func runBASHCmd(cmd string) {
	cmds := strings.Fields(cmd)

	maincmd := cmds[0]
	var args []string
	if len(cmds[1:]) == 0 {
		args = []string{}
	} else if len(cmds[1:]) == 1 {
		args = []string{cmds[1]}
	} else {
		args = cmds[1:]
	}
	out, err := exec.Command(maincmd, strings.Join(args, " ")).Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output: %v", string(out))
}
