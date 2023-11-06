package registering_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gobase.com/base/pkg/registering"
)

type MockDatabaseRepository struct {
	registerUserFunc func(user registering.User) (uuid.UUID, error)
	userExistsFunc   func(user registering.User) (bool, error)
}

func (m *MockDatabaseRepository) RegisterUser(user registering.User) (uuid.UUID, error) {
	return m.registerUserFunc(user)
}

func (m *MockDatabaseRepository) UserExists(user registering.User) (bool, error) {
	return m.userExistsFunc(user)
}

var _ = Describe("Registering Service", func() {
	var (
		serviceInstance registering.RegisteringInterface
		mockDatabase    *MockDatabaseRepository
		user            registering.User
	)

	BeforeEach(func() {
		mockDatabase = &MockDatabaseRepository{}
		serviceInstance = registering.NewRegisteringService(mockDatabase)
		user = registering.User{
			FirstName:   "John",
			LastName:    "Doe",
			Username:    "johndoe",
			Email:       "john.doe@example.com",
			PhoneNumber: "1234567890",
		}
	})

	Describe("RegisterUser", func() {
		It("should successfully register a user", func() {
			// Mock the RegisterUser function to return a UUID and no error
			mockDatabase.registerUserFunc = func(user registering.User) (uuid.UUID, error) {
				return uuid.New(), nil
			}
			mockDatabase.userExistsFunc = func(user registering.User) (bool, error) {
				return false, nil
			}

			id, errs := serviceInstance.RegisterUser(user)
			Expect(errs).To(BeEmpty())
			Expect(id).NotTo(BeNil())
		})

		It("should handle registration errors", func() {
			mockDatabase.userExistsFunc = func(user registering.User) (bool, error) {
				return false, nil
			}
			// Mock the RegisterUser function to return an error
			mockDatabase.registerUserFunc = func(user registering.User) (uuid.UUID, error) {
				return uuid.UUID{}, fmt.Errorf("registration failed")
			}

			id, errs := serviceInstance.RegisterUser(user)
			Expect(errs).NotTo(BeEmpty())
			Expect(len(errs)).To(Equal(1))
			Expect(errs[0].Error()).To(Equal("registration failed"))
			Expect(id).To(BeNil())
		})

		It("should handle validation errors", func() {
			user = registering.User{
				// Incomplete user data to trigger validation errors
			}
			mockDatabase.userExistsFunc = func(user registering.User) (bool, error) {
				return false, nil
			}

			id, errs := serviceInstance.RegisterUser(user)
			Expect(errs).NotTo(BeEmpty())
			Expect(id).To(BeNil())
		})
	})
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registering Suite")
}
