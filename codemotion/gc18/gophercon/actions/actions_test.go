package actions_test

import (
	"testing"

	"github.com/bketelsen/talks/codemotion/gc18/gophercon/actions"
	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(actions.App())}
	suite.Run(t, as)
}
