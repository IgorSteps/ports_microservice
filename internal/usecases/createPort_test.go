package usecases_test

import (
	"context"
	"errors"
	"ports_microservice/internal/domain/entities"
	"ports_microservice/internal/usecases"
	mocks "ports_microservice/mocks/domain/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CreatePortTestSuite struct {
	portStore *mocks.PortStore
}

func NewCreateTestPortSuite() *CreatePortTestSuite {
	return &CreatePortTestSuite{
		portStore: &mocks.PortStore{},
	}
}

func Test_Init(t *testing.T) {
	t.Run("TestCreatePort", NewCreateTestPortSuite().TestCreatePort)
}

func (s *CreatePortTestSuite) TestCreatePort(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testPort := NewTestPort()
	testError := errors.New("A Test Error For Port Insert")

	testCases := []*TestCase{
		NewTestCase("Successful port creation", nil, testPort, func() {
			s.portStore.EXPECT().Insert(testPort).Return(nil).Once()
		}),
		NewTestCase("Failed port creation", testError, nil, func() {
			s.portStore.EXPECT().Insert(testPort).Return(testError).Once()
		}),
	}
	usecase := usecases.NewCreatePort(s.portStore)

	for _, tc := range testCases {
		ctx := context.Background()
		tc.ExpectedMocks()

		// -----
		// ACT
		// -----
		port, err := usecase.Execute(ctx, testPort)

		// -------
		// ASSERT
		// -------
		assert.Equal(t, tc.ExepcetedError, err, "Error returned is not equal to test error")
		assert.Equal(t, tc.ExpectedPort, port, "Port returned is not equal to test port")
	}
}

// --------------
// HELPERS BELOW
// --------------

type TestCase struct {
	Name           string
	ExepcetedError error
	ExpectedPort   *entities.Port
	ExpectedMocks  func()
}

func NewTestCase(name string,
	expectedErr error,
	expectedPort *entities.Port,
	expectedMocks func(),
) *TestCase {
	return &TestCase{
		Name:           name,
		ExepcetedError: expectedErr,
		ExpectedPort:   expectedPort,
		ExpectedMocks:  expectedMocks,
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
