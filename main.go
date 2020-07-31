package main

import (
	"log"
	"net/http"
	"tweet_olivia/credential"
	"tweet_olivia/routes"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
)

func main() {

	cred := credential.GetTwitterAPIKeys()

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store
	goth.UseProviders(
		twitter.New(cred.TwitterAPIKeys.ConsumerKey, cred.TwitterAPIKeys.ConsumerSecret, cred.TwitterAPIKeys.CallbackURL),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
	)
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
