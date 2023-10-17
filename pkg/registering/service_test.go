package registering

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockDbRepo struct {
	mock.Mock
}

func (m *mockDbRepo) RegisterUser(user User) (id uuid.UUID, err error) {
	args := m.Called(user)
	return args.Get(0).(uuid.UUID), args.Error(1)
}
func (m *mockDbRepo) UserExists(user User) (unique bool, err error) {
	args := m.Called(user)
	return args.Bool(0), args.Error(1)
}

type RegisteringTestSuite struct {
	suite.Suite
	RegisteringInterface
	*mockDbRepo
}

func (s *RegisteringTestSuite) SetupTest() {
	// create mocks and initialize service for testing
	s.mockDbRepo = new(mockDbRepo)
	s.RegisteringInterface = NewRegisteringService(
		s.mockDbRepo,
	)
}
func (s *RegisteringTestSuite) TearDownTest() {
	s.mockDbRepo.AssertExpectations(s.T())
}
func TestRegisteringService(t *testing.T) {
	suite.Run(t, new(RegisteringTestSuite))
}

func (s *RegisteringTestSuite) TestUserRegisterationSuccess() {

	payload := fakeUser()
	userID := newUUID()
	s.mockDbRepo.On("RegisterUser", payload).Return(userID, nil)
	s.mockDbRepo.On("UserExists", payload).Return(false, nil)
	id, err := s.RegisteringInterface.RegisterUser(payload)
	s.Nil(err)
	s.True(*id == userID)
}

func fakeUser() User {
	return User{FirstName: "stk", LastName: "himself", Address: "The STK's", Username: "stk", Email: "stk@gmail.com", PhoneNumber: "+23784949384943"}
}
