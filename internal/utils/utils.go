package utils

import (
	"fmt"
	"log"
	"os/exec"
	"ssh_connections_manager/internal/utils/colors"
)


var (
	ColorRedBold = "31"
)

func GetPrompt() string {
    user, err := exec.Command("whoami").Output()
    if err != nil {
        log.Println("Error fetching user:", err)
        return "$ "
    }

    host, err := exec.Command("hostname").Output()
    if err != nil {
        log.Println("Error fetching host:", err)
        return "$ "
    }

    cwd, err := exec.Command("pwd").Output()
    if err != nil {
        log.Println("Error fetching current directory:", err)
        return "$ "
    }

    return fmt.Sprintf("%s@%s:~%s", user[:len(user) - 1], host[:len(host) - 1], cwd[:len(cwd) - 1])
}

func GetUserName() string {
	user, err := exec.Command("whoami").Output()
    if err != nil {
        log.Println("Error fetching user:", err)
        return "$ "
    }

	return colors.ColorString(string(user[:len(user) - 1]), colors.ColorGreenBold)
}


func GetHost() string {
	host, err := exec.Command("hostname").Output()
    if err != nil {
        log.Println("Error fetching host:", err)
        return "$ "
    }

	return colors.ColorString(string(host[:len(host) - 1]), colors.ColorGreenBold)
}

func GetPwd() string {
	cwd, err := exec.Command("pwd").Output()
    if err != nil {
        log.Println("Error fetching current directory:", err)
        return "$ "
    }

	return colors.ColorString(string(cwd[:len(cwd) - 1]), colors.ColorGreenBold)

}