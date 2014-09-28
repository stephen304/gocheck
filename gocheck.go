/*
GoCheck provides a simple binary to detect problems and restart web servers.
Copyright (C) 2014 Stephen Smith

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see http://www.gnu.org/licenses/.
*/

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var crashes int
var url = "http://example.com/"

func main() {
	resp, err := http.Get("http://example.com/")
	for i := 1; i < 360; i++ {
		if crashes > 5 {
			os.Exit(1)
		}
		resp, err = http.Get(url)
		if err != nil {
			fmt.Println("[\u2717] Reboot server")
			nginx()
		} else {
			if resp.StatusCode != 200 {
				fmt.Printf("[\u2717] Reboot hhvm b/c status %d\n", resp.StatusCode)
				hhvm()
			} else {
				fmt.Println("[\u2713] Everything's OK")
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func nginx() {
	exec.Command("/usr/sbin/service", "nginx", "restart").Output()
	fmt.Println("[\u2713] Rebooted")
	crashes++
}

func hhvm() {
	exec.Command("/usr/sbin/service", "hhvm", "restart").Output()
	fmt.Println("[\u2713] Rebooted")
	crashes++
}
