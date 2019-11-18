package routes

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joaopandolfi/blackwhale/configurations"
)

type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func public(r *mux.Router) {
	r.PathPrefix(configurations.Configuration.StaticPath).
		Handler(http.StripPrefix(strings.TrimRight(configurations.Configuration.StaticPath, "/"),
			http.FileServer(FileSystem{http.Dir(configurations.Configuration.StaticDir)})))
}
