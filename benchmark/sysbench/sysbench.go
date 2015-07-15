package main

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckDockerfile() {
	// confirm the Dockerfile exsists
	f, err := os.Open("./Dockerfile")
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("Dockerfile not exist!\n")
		return
	}
	fmt.Printf("Dockerfile exist!\n")
	defer f.Close()
}

func BuildDockerfile() {
	// excute build process
	cmd := exec.Command("sudo", "docker", "build", "-no-cache=true", "--tag=\"open-container-sysbench\"", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(output))
}

func SysbenchCommand(testparameter string) {
	// excute sysbench command in docker container
	cmd := exec.Command("sudo", "docker", "run", "-t", "open-container-sysbench", "bash", "-c", "sysbench "+testparameter)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(output))
}

func main() {
	CheckDockerfile()
	BuildDockerfile()
	cputest := "--num-threads=4 --max-requests=1000 --test=cpu run"
	memorytest := "--num-threads=4 --max-requests=1000 --test=memory run"
	fileiotestprepare := "--file-total-size=500M --test=fileio prepare"
	fileiotestrun := "--file-total-size=500M  --file-test-mode=seqwr --test=fileio run"
	fileiotestclean := "--file-total-size=500M --test=fileio cleanup"
	SysbenchCommand(cputest)
	SysbenchCommand(memorytest)
	SysbenchCommand(fileiotestprepare)
	SysbenchCommand(fileiotestrun)
	SysbenchCommand(fileiotestclean)
}
