import (
    "golang.org/x/net/context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    newappengine "google.golang.org/appengine"
    newurlfetch "google.golang.org/appengine/urlfetch"

    "appengine"
)

var Endpoint = oauth2.Endpoint{
    AuthURL:  "https://www.facebook.com/dialog/oauth",
    TokenURL: "https://graph.facebook.com/oauth/access_token",
}

func handler(w http.ResponseWriter, r *http.Request) {
    var c appengine.Context = appengine.NewContext(r)
    c.Infof("Logging a message with the old package")

    var ctx context.Context = newappengine.NewContext(r)
    client := &http.Client{
        Transport: &oauth2.Transport{
            Source: google.AppEngineTokenSource(ctx, "scope"),
            Base:   &newurlfetch.Transport{Context: ctx},
        },
    }
    client.Get("...")
}
