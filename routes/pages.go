package routes

import "github.com/gorilla/mux"

// PageController - Interface for controllers that will render pages
type PageController interface {
	Precompile()
}

func precompilePages() {

}

func pages(r *mux.Router) {

}
