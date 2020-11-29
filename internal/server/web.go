package server

import (
	"fmt"
	"hello-server/internal/logger"
	"net/http"
	"syscall"
)

// HelloServer is designed to give more flexibility in programming the HTTP server.
type HelloServer struct {
	HTTPServer *http.Server
}

// New creates and returns a HelloServer
func New() HelloServer {
	http.DefaultServeMux = new(http.ServeMux)
	return HelloServer{
		&http.Server{
			Addr: ":7100",
		},
	}
}

// StartHelloServer starts the HelloServer server.
func (server *HelloServer) StartHelloServer() error {
	http.HandleFunc("/", webHandler)
	logger.Logger.Info("Server listening on port 7100")
	return server.HTTPServer.ListenAndServe()
}

func webHandler(writer http.ResponseWriter, request *http.Request) {
	sysInfo, err := getSysInfo()
	if err != nil {
		logger.Logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "<h1>Internal Server Error</h1>")
		return
	}
	arch := make([]byte, 0, len(sysInfo.Machine))
	for _, i := range sysInfo.Machine {
		if i != 0 {
			arch = append(arch, byte(i))
		}
	}

	os := make([]byte, 0, len(sysInfo.Sysname))
	for _, i := range sysInfo.Sysname {
		if i != 0 {
			os = append(os, byte(i))
		}
	}

	if err != nil {
		logger.Logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "<h1>Internal Server Error</h1>")
		return
	}
	writer.WriteHeader(http.StatusOK)
	title := "<h1>Hello World</h1>"
	architecture := "<div><b>Architecture:</b> " + string(arch) + "</div>"
	opsys := "<div><b>Operating System:</b> " + string(os) + "</div>"
	fmt.Fprintf(writer, title+architecture+opsys)
}

func getSysInfo() (syscall.Utsname, error) {
	// Get architecture
	utsname := syscall.Utsname{}
	err := syscall.Uname(&utsname)
	if err != nil {
		return utsname, err
	}

	return utsname, nil
}
