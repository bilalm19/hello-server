package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	webHandler(response, request)

	got := response.Body.String()

	sysInfo, err := getSysInfo()
	if err != nil {
		t.Fatal(err)
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

	title := "<h1>Hello World</h1>"
	architecture := "<div><b>Architecture:</b> " + string(arch) + "</div>"
	opsys := "<div><b>Operating System:</b> " + string(os) + "</div>"

	want := title + architecture + opsys

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
