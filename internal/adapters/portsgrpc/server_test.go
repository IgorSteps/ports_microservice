package portsgrpc_test

import (
	"context"
	"errors"
	port_pb "ports_microservice/external/proto/ports"
	"ports_microservice/internal/adapters/portsgrpc"
	"ports_microservice/internal/domain/entities"
	mocks "ports_microservice/mocks/adapters"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CreatePortTestSuite struct {
	facade *mocks.PortFacade
}

func NewCreatePortTestSuite() *CreatePortTestSuite {
	return &CreatePortTestSuite{
		facade: &mocks.PortFacade{},
	}
}

func Test_Init(t *testing.T) {
	t.Run("Test_CreatePortAdapter", NewCreatePortTestSuite().TestCreatePorts)
}

func (s *CreatePortTestSuite) TestCreatePorts(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testProtoPortReq := NewTestPortReq()
	testEntityPort := portsgrpc.ConvertProtoToEntity(testProtoPortReq)
	testProtoPortResp := portsgrpc.ConvertEntityToProto(testEntityPort)

	testError := errors.New("A Test Error For creating port")
	ctx := context.Background()

	testCases := []*CreatePortTestCase{
		NewCreatePortTestCase("Successful ports get", nil, testProtoPortResp, func() {
			s.facade.EXPECT().CreatePort(ctx, testEntityPort).Return(testEntityPort, nil).Once()
		}),
		NewCreatePortTestCase("Failed ports get", testError, nil, func() {
			s.facade.EXPECT().CreatePort(ctx, testEntityPort).Return(nil, testError).Once()
		}),
	}

	adapter := portsgrpc.NewPortGrpc(s.facade)

	for _, tc := range testCases {
		ctx := context.Background()
		tc.ExpectedMocks()

		// -----
		// ACT
		// -----
		resp, err := adapter.CreatePort(ctx, NewTestPortReq())

		// -------
		// ASSERT
		// -------
		assert.Equal(t, tc.ExepcetedError, err, "Error returned is not equal to test error")
		assert.Equal(t, tc.ExpectedPort, resp, "Port response returned is not equal to test port")

	}
}

// --------------
// HELPERS BELOW
// --------------

type CreatePortTestCase struct {
	Name           string
	ExepcetedError error
	ExpectedPort   *port_pb.CreatePortResponse
	ExpectedMocks  func()
}

func NewCreatePortTestCase(
	name string,
	expectedErr error,
	expectedPort *port_pb.CreatePortResponse,
	expectedMocks func(),
) *CreatePortTestCase {
	return &CreatePortTestCase{
		Name:           name,
		ExepcetedError: expectedErr,
		ExpectedPort:   expectedPort,
		ExpectedMocks:  expectedMocks,
	}
}

func NewTestPortReq() *port_pb.CreatePortRequest {
	return &port_pb.CreatePortRequest{
		Name:        "name",
		City:        "city",
		Country:     "country",
		Alias:       []string{"alias", "aliases"},
		Regions:     []string{"regions", "regions"},
		Coordinates: []float32{23, 123},
		Province:    "province",
		Timezone:    "timezone",
	}
}

func NewTestPort() *entities.Port {
	return &entities.Port{
		Name:        "name",
		City:        "city",
		Country:     "country",
		Alias:       []string{"alias", "aliases"},
		Regions:     []string{"regions", "regions"},
		Coordinates: []float32{23, 123},
		Province:    "province",
		Timezone:    "timezone",
	}
}
