package actions

import (
	"errors"
	"time"

	mware "github.com/bketelsen/talks/codemotion/gc18/gophercon-ot/middleware"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	slow(c)
	ses := c.Session().Get("session")
	if ses != nil {
		ses = ses.(string)
		uname := c.Session().Get("current_user_id").(string)
		c.Session().Set("username", uname)
	}
	return c.Render(200, r.HTML("index.html"))
}

func ProtectedHandler(c buffalo.Context) error {
	ses := c.Session().Get("session").(string)
	uname := c.Session().Get("current_user_id").(string)
	c.Set("session", ses)
	c.Set("username", uname)
	c.Flash().Add("info", "You are authorized to see this page")
	return c.Render(200, r.HTML("index.html"))
}

//BadHandler returns an error
func BadHandler(c buffalo.Context) error {
	return c.Error(401, errors.New("Unauthorized!"))
}
func slow(c buffalo.Context) {
	sp := mware.ChildSpan("slow", c)
	defer sp.Finish()
	time.Sleep(10 * time.Millisecond)
}
