package gowyre_test

import (
	"context"
	"testing"

	"github.com/elysiumbridge/gowyre"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(GowyreTestSuite))
}

type GowyreTestSuite struct {
	suite.Suite
	secret string
	env    string
}

func (s *GowyreTestSuite) SetupSuite() {
	s.secret = ""
	s.env = "test"
}

func (s *GowyreTestSuite) TestAccesDenied() {
	c, err := gowyre.NewClient(s.secret, s.env, "", nil)
	s.Nil(err)

	_, err = c.CreateWallet(context.Background(), &gowyre.CreateWallet{Name: "user:x"})
	s.Require().NotNil(err)
	s.Require().ErrorAs(err, &gowyre.Error{})
	s.Equal(401, err.(gowyre.Error).StatusCode)
	s.Equal("Invalid Session", err.Error())
}
