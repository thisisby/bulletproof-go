package module

// IModule defines the interface for a module in the application.
// Modules are responsible for managing controllers and services and can include other submodules.
// Implementers of this interface must provide methods to get the module's name, controllers, and services.
type IModule interface {
	GetName() string
	GetControllers() []any
	GetServices() []any
}

// Module represents a module in the application.
// It holds the name, a list of controllers, services, and any submodules.
// The module can be used to organize and group related functionality.
type Module struct {
	Name        string
	submodules  []IModule
	Controllers []interface{}
	Services    []interface{}
}

func NewModule(name string, controllers []interface{}, services []interface{}, submodules []IModule) *Module {
	return &Module{
		Name:        name,
		Controllers: controllers,
		Services:    services,
		submodules:  submodules,
	}
}

func (m *Module) GetName() string {
	return m.Name
}

func (m *Module) GetControllers() []interface{} {
	return m.Controllers
}

func (m *Module) GetServices() []interface{} {
	return m.Services
}
