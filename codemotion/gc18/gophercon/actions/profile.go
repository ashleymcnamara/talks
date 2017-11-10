package actions

import (
	"fmt"

	proto "github.com/bketelsen/talks/codemotion/gc18/services/profile-srv/proto/record"

	uproto "github.com/bketelsen/talks/codemotion/gc18/services/user-srv/proto/account"
	"github.com/gobuffalo/buffalo"
	micro "github.com/micro/go-micro"
)

func ProfileGet(c buffalo.Context) error {
	service := micro.NewService(micro.Name("buffalo.user.client"))

	ucl := uproto.NewAccountClient("gophercon.srv.user", service.Client())
	creq := &uproto.SearchRequest{
		Username: c.Session().Get("current_user_id").(string),
		Limit:    10,
		Offset:   0,
	}

	crsp, err := ucl.Search(c, creq)
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
	prof := proto.NewRecordClient("gophercon.srv.profile", service.Client())
	req := &proto.ReadRequest{
		Id: id,
	}
	rsp, err := prof.Read(c, req)
	if err != nil {
		fmt.Println("Houston we have a problem", err)
		c.Flash().Add("danger", "That didn't work")
		return c.Redirect(302, "/login")
	}
	fmt.Println("Response", rsp)
	c.Set("profile", rsp.GetProfile())
	return c.Render(200, r.HTML("user/profile.html"))
}
