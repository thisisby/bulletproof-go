package service

// Service defines the interface for all services in the application.
// Implementations of this interface can provide specific functionalities
// required by different parts of the application.
type Service any

// BaseService provides a default implementation of the Service interface.
// This struct can be embedded in other service implementations to inherit
// common functionalities or to be extended with custom methods.
type BaseService struct{}

// PerformTask is a method of BaseService that performs a generic task.
// This method can be overridden by services embedding BaseService to provide
// specific behavior or functionality.
func (bs *BaseService) PerformTask() {

}
