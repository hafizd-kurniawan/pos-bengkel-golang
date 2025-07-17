package usecase

import (
	"boilerplate/internal/repository"
	"boilerplate/internal/usecase/implementations"
	"boilerplate/internal/usecase/interfaces"
)

// UsecaseManager contains all usecase interfaces
type UsecaseManager struct {
	// Foundation & Security
	User       interfaces.UserUsecase
	Outlet     interfaces.OutletUsecase
	Role       interfaces.RoleUsecase
	Permission interfaces.PermissionUsecase

	// Customer & Vehicle
	Customer        interfaces.CustomerUsecase
	CustomerVehicle interfaces.CustomerVehicleUsecase

	// Add other usecases as they are implemented
}

// NewUsecaseManager creates a new usecase manager with all usecases
func NewUsecaseManager(repo *repository.RepositoryManager) *UsecaseManager {
	return &UsecaseManager{
		// Foundation & Security
		User:   implementations.NewUserUsecase(repo),
		Outlet: implementations.NewOutletUsecase(repo),
		// Role:       implementations.NewRoleUsecase(repo),
		// Permission: implementations.NewPermissionUsecase(repo),

		// Customer & Vehicle
		Customer:        implementations.NewCustomerUsecase(repo),
		CustomerVehicle: implementations.NewCustomerVehicleUsecase(repo),

		// Add other usecases as they are implemented
	}
}