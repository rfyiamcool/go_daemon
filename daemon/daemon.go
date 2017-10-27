package daemon

import (
	"log"
	"os"
	"os/exec"
)

func DaemonInit(mark *bool) {
	if *mark {
		args := os.Args[1:]
		i := 0
		for ; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		log.Println("start success: [PID]", cmd.Process.Pid)
		// 退出父进程
		os.Exit(0)
	}
}

