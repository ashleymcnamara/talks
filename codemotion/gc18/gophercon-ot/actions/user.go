package actions

import (
	"fmt"
	"time"

	mware "github.com/bketelsen/talks/codemotion/gc18/gophercon-ot/middleware"
	proto "github.com/bketelsen/talks/codemotion/gc18/services-ot/user-srv/proto/account"
	"github.com/gobuffalo/buffalo"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	mot "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

func UserLogin(c buffalo.Context) error {

	return c.Render(200, r.HTML("user/login.html"))
}

func UserLogout(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("info", "You were logged out")
	return c.Render(200, r.HTML("index.html"))
}

// UserAuthenticate default implementation.
func UserAuthenticate(c buffalo.Context) error {
	service := micro.NewService(micro.Name("buffalo.user.client"))
	//cl := client.NewClient(client.Wrap(mot.NewClientWrapper(Tracer)))
	client.DefaultClient = client.NewClient(
		client.Wrap(
			mot.NewClientWrapper(Tracer)),
	)
	ctx := mware.MetadataContext(c)
	user := proto.NewAccountClient("gophercon.srv.userot", service.Client())
	req := &proto.LoginRequest{
		Username: "bketelsen",
		Email:    "bketelsen@gmail.com",
		Password: "password1",
	}
	rsp, err := user.Login(ctx, req)
	if err != nil {
		fmt.Println("Houston we have a problem", err)
		c.Flash().Add("danger", "That didn't work")
		return c.Redirect(302, "/login")
	}
	fmt.Println("Response", rsp)
	c.Session().Set("session", rsp.Session.Id)
	c.Session().Set("current_user_id", rsp.Session.Username)

	c.Flash().Add("info", "welcome "+rsp.Session.Username)
	return c.Render(200, r.HTML("user/authenticate.html"))
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		fmt.Println("SetCurrentUser")
		if uid := c.Session().Get("current_user_id"); uid != nil {
			fmt.Println("Setting current user")
			fmt.Println("UID:", uid)
			c.Set("current_user", "SET")
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		fmt.Println("AUTHORIZE")
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		sid := c.Session().Get("session")
		if sid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		sess := sid.(string)

		service := micro.NewService(micro.Name("buffalo.user.client"))
		ctx := mware.MetadataContext(c)
		session := proto.NewAccountClient("gophercon.srv.userot", service.Client())
		req := &proto.ReadSessionRequest{
			SessionId: sess,
		}
		rsp, err := session.ReadSession(ctx, req)
		if err != nil {
			fmt.Println("Houston we have a problem", err)
			return c.Error(500, err)
		}
		fmt.Println("Response", rsp)
		c.Session().Set("session", rsp.Session.Id)
		c.Session().Set("current_user_id", rsp.Session.Username)
		expire := time.Unix(rsp.Session.Expires, 0)
		if err != nil {
			fmt.Println("date parse error")
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(500, "/")
		}
		if expire.Before(time.Now()) {
			c.Flash().Add("danger", "Session Expired!")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
