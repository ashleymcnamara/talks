package actions

import (
	"fmt"

	mware "github.com/bketelsen/talks/codemotion/gc18/gophercon-ot/middleware"
	proto "github.com/bketelsen/talks/codemotion/gc18/services-ot/profile-srv/proto/record"

	uproto "github.com/bketelsen/talks/codemotion/gc18/services-ot/user-srv/proto/account"
	"github.com/gobuffalo/buffalo"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	mot "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

func ProfileGet(c buffalo.Context) error {
	service := micro.NewService(micro.Name("buffalo.user.client"))
	//cl := client.NewClient(client.Wrap(mot.NewClientWrapper(Tracer)))
	client.DefaultClient = client.NewClient(
		client.Wrap(
			mot.NewClientWrapper(Tracer)),
	)

	ctx := mware.MetadataContext(c)

	ucl := uproto.NewAccountClient("gophercon.srv.userot", service.Client())
	creq := &uproto.SearchRequest{
		Username: c.Session().Get("current_user_id").(string),
		Limit:    10,
		Offset:   0,
	}

	crsp, err := ucl.Search(ctx, creq)
	if err != nil {
		fmt.Println("Houston we have a problem", err)
		c.Flash().Add("danger", "That didn't work")
		return c.Redirect(302, "/login")
	}
	var id string
	id = ""
	if len(crsp.Users) > 0 {
		user := crsp.Users[0]
		id = user.Id
		fmt.Println("Found id:", id)
	}
	fmt.Println("Searching for id:", id)
	prof := proto.NewRecordClient("gophercon.srv.profileot", service.Client())
	req := &proto.ReadRequest{
		Id: id,
	}
	rsp, err := prof.Read(ctx, req)
	if err != nil {
		fmt.Println("Houston we have a problem", err)
		c.Flash().Add("danger", "That didn't work")
		return c.Redirect(302, "/login")
	}
	fmt.Println("Response", rsp)
	c.Set("profile", rsp.GetProfile())
	return c.Render(200, r.HTML("user/profile.html"))
}
