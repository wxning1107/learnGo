package scrapefunc

import (
	"log"
	"os/exec"
)

func ExecScript(script string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", script, "script port status")
	output, err := cmd.Output()
	//output, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func ScrapePortStatus() map[string]map[string]float64 {
	shell := `netstat -an |grep 80`
	_, err := ExecScript(shell)
	if err != nil {
		log.Printf("Scrape portStatus error: %v\n", err)
		return nil
	}
	//fmt.Println(1111)
	//fmt.Printf("端口信息：%v\n", data)

	result := make(map[string]map[string]float64)
	return result
}
