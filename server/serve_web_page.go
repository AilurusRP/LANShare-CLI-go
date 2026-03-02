package server

import (
	"fmt"
	"lanshare/info"
	"net/http"
	"strings"
)

func ServeWebpage(webPath string) {
	if !strings.HasSuffix(webPath, "/") {
		webPath = webPath + "/"
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(webPath+"static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, webPath+"index.html")
	})

	fmt.Println("The LANShare webpage is now running,")
	fmt.Printf("visit http://%s%s/ to use LANShare-web. \n", info.IP, info.PORT)
	fmt.Print("Make sure your device is in the same LAN with LANShare-CLI.\n\n")
}
