package models

type AvailableServices struct {
	Services []Service
}

/*
	Used for appending channel to Channels array inside AvailableChannel datastructure
*/
func (MyAvailableSevice *AvailableServices) AddService(MyService Service) {
	MyAvailableSevice.Services = append(MyAvailableSevice.Services, MyService)
}

/*
	Used for getting list of all available services
*/
func getListOfServices() AvailableServices {
	service1 := Service{
		Name:  "LearnEnglish",
		Price: 100,
	}
	service2 := Service{
		Name:  "LearnHindi",
		Price: 200,
	}
	var services AvailableServices
	services.AddService(service1)
	services.AddService(service2)
	return services

}

/*
	Used for searching whether a service exists
*/
func (availableServices *AvailableServices) DoesServiceExist(svc string) bool {
	var Exists bool
	Exists = false

	for i := 0; i < len(availableServices.Services); i++ {
		if svc == availableServices.Services[i].Name {
			Exists = true
			break
		}
	}
	return Exists
}

/*
	Used for getting the service
*/
func (availableServices *AvailableServices) GetService(svc string) Service {
	var Exists Service

	for i := 0; i < len(availableServices.Services); i++ {
		if svc == availableServices.Services[i].Name {
			Exists = availableServices.Services[i]
			break
		}
	}
	return Exists
}
