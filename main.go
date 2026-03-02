package main

import (
	"flag"
	"fmt"
	"lanshare/info"
	"lanshare/server"
	"os/user"
)

func main() {
	info.GetLocalIP()
	fmt.Printf("Your current desktop local IP address: %s\n\n", info.IP)

	qrString, err := generateQRCode(info.IP)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(qrString)

	// If you need to run LANShare-web,
	// the path to the web resources should be passed from the command line.
	// Example:
	//     $ LANShare-CLI --path /var/www/LANShare-web
	webPath := flag.String("path", "", "the path of the LANShare-web directory")
	flag.Parse()
	if *webPath != "" {
		server.ServeWebpage(*webPath)
	} else {
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
		server.ServeWebpage(currentUser.HomeDir + "/LANShare-web")
	}
	server.StartServer()
}
