package routes

import (
	"fmt"
	"net/http"
	"sort"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/auth/{provider}/callback", TwitterLogin).Methods("GET")
	r.HandleFunc("/logout/{provider}", TwitterLogout).Methods("GET")
	r.HandleFunc("/auth/{provider}", NoReAuthentication).Methods("GET")
	r.HandleFunc("/", IndexPage).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// r.HandleFunc("/{username}", middleware.AuthRequired(userGetHandler)).Methods("GET")
	return r
}

func TwitterLogin(res http.ResponseWriter, req *http.Request) {

	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.ParseFiles("templates/success.html")

	t.Execute(res, user)
}

func TwitterLogout(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func NoReAuthentication(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		t, _ := template.ParseFiles("templates/success.html")
		t.Execute(res, gothUser)
	} else {
		gothic.BeginAuthHandler(res, req)
	}
}

func IndexPage(res http.ResponseWriter, req *http.Request) {

	m := make(map[string]string)
	m["twitter"] = "Twitter"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(res, providerIndex)
}
