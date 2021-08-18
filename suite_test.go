package gowyre_test

import (
	"context"
	"testing"

	"github.com/elysiumbridge/gowyre"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

type WalletTestSuite struct {
	suite.Suite
	secret string
	env    string
}

func (s *WalletTestSuite) SetupSuite() {
	s.secret = ""
	s.env = "test"
}

func (s *WalletTestSuite) TestAccesDenied() {
	c, err := gowyre.NewClient(s.secret, s.env, "", nil)
	s.Nil(err)

	_, err = c.CreateWallet(context.Background(), &gowyre.CreateWallet{Name: "user:x"})
	s.Require().NotNil(err)
	s.ErrorIs(err, gowyre.ErrAccessDenied)
}
