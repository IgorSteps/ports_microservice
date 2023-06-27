package datastore_test

import (
	"errors"
	"ports_microservice/internal/adapters/datastore"
	models "ports_microservice/internal/adapters/datastore/model"
	"ports_microservice/internal/domain/entities"
	mocks "ports_microservice/mocks/adapters/datastore"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DatastoreTestSuite struct {
	dbWrapper *mocks.DBWrapper
}

func NewDatastoreTestSuite() *DatastoreTestSuite {
	return &DatastoreTestSuite{
		dbWrapper: &mocks.DBWrapper{},
	}
}

func Test_Init(t *testing.T) {
	t.Run("Test_Insert_Datastore", NewDatastoreTestSuite().TestInsertPort)
	t.Run("Test_GetMultiple_Datastore", NewDatastoreTestSuite().TestGetMultiplePorts)
}

func (s *DatastoreTestSuite) TestInsertPort(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testPort := NewTestPort()
	schemaTestPort := models.ConvertEntityToDB(testPort)
	testError := errors.New("A Test Error For inserting port")

	testCases := []*DatastoreTestCase{
		NewDatastoreTestCase("Successful port insert", nil, nil, nil, func() {
			s.dbWrapper.EXPECT().Create(schemaTestPort).Return(s.dbWrapper).Once()
			s.dbWrapper.EXPECT().Error().Return(nil).Once()
		}),
		NewDatastoreTestCase("Failed port insert", testError, nil, nil, func() {
			s.dbWrapper.EXPECT().Create(schemaTestPort).Return(s.dbWrapper).Once()
			s.dbWrapper.EXPECT().Error().Return(testError).Once()
		}),
	}

	adapter := datastore.NewPortDataStore(s.dbWrapper)

	for _, tc := range testCases {
		tc.ExpectedMocks()

		// -----
		// ACT
		// -----
		err := adapter.Insert(testPort)

		// -------
		// ASSERT
		// -------
		assert.Equal(t, tc.ExepcetedError, err, "Error returned is not equal to test error")
	}
}

func (s *DatastoreTestSuite) TestGetMultiplePorts(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testPorts := NewTestPortsSlice(5)
	// A model to be passed to dbWrapper.Find()
	var pp []*models.PortSchema
	// Ports of schema type to be get from dbWrapper.Find()
	var schemaTestPorts []*models.PortSchema
	for _, port := range testPorts {
		schemaTestPorts = append(schemaTestPorts, models.ConvertEntityToDB(port))
	}

	testError := errors.New("A Test Error For getting multiple ports")

	// @Note:
	// Need to run '.Run()' callback to be able to modify ports variable in GetMultiple that
	// gets populated from dbWrapper.Find() call. Since the DB isn't running for UTs that call
	// would return an empty slice which we don't want.
	testCases := []*DatastoreTestCase{
		NewDatastoreTestCase("Successful ports getMultiple", nil, nil, testPorts, func() {
			s.dbWrapper.EXPECT().Find(&pp).Return(s.dbWrapper).
				Run(func(value interface{}) {
					arg := value.(*[]*models.PortSchema)
					// Set them to be schema ports.
					*arg = append(*arg, schemaTestPorts...)
				}).Once()
			s.dbWrapper.EXPECT().Error().Return(nil).Once()
		}),
		NewDatastoreTestCase("Failed ports getMultiple", testError, nil, nil, func() {
			s.dbWrapper.EXPECT().Find(&pp).Return(s.dbWrapper).
				Run(func(value interface{}) {
					arg := value.(*[]*models.PortSchema)
					// Set them to be schema ports.
					*arg = append(*arg, schemaTestPorts...)
				}).Once()
			s.dbWrapper.EXPECT().Error().Return(testError).Once()
		}),
	}

	adapter := datastore.NewPortDataStore(s.dbWrapper)

	for _, tc := range testCases {
		tc.ExpectedMocks()

		// -----
		// ACT
		// -----
		ports, err := adapter.GetMultiple()

		// -------
		// ASSERT
		// -------
		assert.Equal(t, tc.ExepcetedError, err, "Error returned is not equal to test error")
		assert.Equal(t, tc.ExpectedPorts, ports, "Ports not equal")
	}
}

// --------------
// HELPERS BELOW
// --------------

type DatastoreTestCase struct {
	Name           string
	ExepcetedError error
	ExpectedPort   *entities.Port
	ExpectedPorts  []*entities.Port
	ExpectedMocks  func()
}

func NewDatastoreTestCase(
	name string,
	expectedErr error,
	expectedPort *entities.Port,
	expectedPorts []*entities.Port,
	expectedMocks func(),
) *DatastoreTestCase {
	return &DatastoreTestCase{
		Name:           name,
		ExepcetedError: expectedErr,
		ExpectedPort:   expectedPort,
		ExpectedPorts:  expectedPorts,
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

func NewTestPortsSlice(num int) []*entities.Port {
	var ports []*entities.Port
	for i := 0; i < num; i++ {
		ports = append(ports, NewTestPort())
	}

	return ports
}
