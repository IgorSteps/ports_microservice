package portsrest_test

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"ports_microservice/internal/adapters/portsrest"
	"ports_microservice/internal/domain/entities"
	mocks "ports_microservice/mocks/adapters"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type GetPortsTestSuite struct {
	facade *mocks.PortFacade
}

func NewCreateTestPortSuite() *GetPortsTestSuite {
	return &GetPortsTestSuite{
		facade: &mocks.PortFacade{},
	}
}

func Test_Init(t *testing.T) {
	t.Run("Test_GetPortsAdapter", NewCreateTestPortSuite().TestGetPorts)
}

func (s *GetPortsTestSuite) TestGetPorts(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testPorts := NewTestPortsSlice(5)
	testError := errors.New("A Test Error For Getting ports")

	testCases := []*GetPortsTestCase{
		NewGetPortsTestCase("Successful ports get", nil, testPorts, func() {
			s.facade.EXPECT().GetPorts().Return(testPorts, nil).Once()
		}),
		NewGetPortsTestCase("Failed ports get", testError, nil, func() {
			s.facade.EXPECT().GetPorts().Return(nil, testError).Once()
		}),
	}
	adapter := portsrest.NewPortRestApi(s.facade)

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		r := gin.Default()
		ctx := gin.CreateTestContextOnly(w, r)

		tc.ExpectedMocks()

		// -----
		// ACT
		// -----
		adapter.GetPorts(ctx)

		// -------
		// ASSERT
		// -------
		// Need to unmarshal returned JSON into Go slice.
		var returnedPorts []*entities.Port
		json.Unmarshal(w.Body.Bytes(), &returnedPorts)

		assert.Equal(t, tc.ExpectedPorts, returnedPorts, "Port returned is not equal to test port")
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
