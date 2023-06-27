package functionaltests

import (
	"ports_microservice/functionaltests/clientsuites"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CreatePortTestSuite struct {
	clientsuites.GrpcClientSuite
}

func (s *CreatePortTestSuite) SetupSuite() {
	s.GrpcClientSuite.SetupSuite()
}

func TestInit(t *testing.T) {
	suite.Run(t, new(CreatePortTestSuite))
}

func (s *CreatePortTestSuite) Test_CreatePort() {
	req := generateCreatePortReq()
	resp, err := s.CreatePort(req)
	if err != nil {
		s.T().Fatalf("Failed to create port %v", err)
	}

	assert.Equal(s.T(), req, resp, "Ports are not equal")
}
