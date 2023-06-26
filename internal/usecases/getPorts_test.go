package usecases_test

import (
	"errors"
	"ports_microservice/internal/domain/entities"
	"ports_microservice/internal/usecases"
	mocks "ports_microservice/mocks/domain/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetPortsTestSuite struct {
	portStore *mocks.PortStore
}

func NewGetPortsTestPortSuite() *GetPortsTestSuite {
	return &GetPortsTestSuite{
		portStore: &mocks.PortStore{},
	}
}

func Test_GetPortsInit(t *testing.T) {
	t.Run("TestGetPorts", NewGetPortsTestPortSuite().TestGetPorts)
}

func (s *GetPortsTestSuite) TestGetPorts(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testPorts := NewTestPortsSlice(5)
	testError := errors.New("A Test Error For Getting ports")

	testCases := []*GetPortsTestCase{
		NewGetPortsTestCase("Successful ports get", nil, testPorts, func() {
			s.portStore.EXPECT().GetMultiple().Return(testPorts, nil).Once()
		}),
		NewGetPortsTestCase("Failed ports get", testError, nil, func() {
			s.portStore.EXPECT().GetMultiple().Return(nil, testError).Once()
		}),
	}
	usecase := usecases.NewGetPorts(s.portStore)

	for _, tc := range testCases {
		tc.ExpectedMocks()

		// -----
		// ACT
		// -----
		ports, err := usecase.Execute()

		// -------
		// ASSERT
		// -------
		assert.Equal(t, tc.ExepcetedError, err, "Error returned is not equal to test error")
		assert.Equal(t, tc.ExpectedPorts, ports, "Ports returned are not equal to tests port")
	}
}

// --------------
// HELPERS BELOW
// --------------

type GetPortsTestCase struct {
	Name           string
	ExepcetedError error
	ExpectedPorts  []*entities.Port
	ExpectedMocks  func()
}

func NewGetPortsTestCase(
	name string,
	expectedErr error,
	expectedPorts []*entities.Port,
	expectedMocks func(),
) *GetPortsTestCase {
	return &GetPortsTestCase{
		Name:           name,
		ExepcetedError: expectedErr,
		ExpectedPorts:  expectedPorts,
		ExpectedMocks:  expectedMocks,
	}
}

func NewTestPortsSlice(num int) []*entities.Port {
	ports := make([]*entities.Port, num)
	for i := 0; i < num; i++ {
		ports = append(ports, NewTestPort())
	}

	return ports
}
