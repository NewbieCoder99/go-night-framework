/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 00:04:54
*/
package main

import (
    "fmt"
    "os"
    "os/exec"
    "encoding/json"
)

type Pkgs struct {
	Packages []string
}

func main() {
	var pkgs Pkgs

	JsonFile, err := os.Open("packages.json");

	defer JsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	jsonParser := json.NewDecoder(JsonFile)
	jsonParser.Decode(&pkgs)

	for i := 0; i < len(pkgs.Packages); i++ {
		fmt.Println("Installing package",pkgs.Packages[i],"...")
		cmd := exec.Command("go","get",pkgs.Packages[i])
		fmt.Println("Installing successfully")
		cmd.Run()
	}
}