api contains the transport, the platform, the service and the api initializer. 
The transport contains also the handlers, which parse the data to the lower layers. 
The service implementation is on contact.go, and its interface on service.go
The data CRUD operations are later executed on /platform/contact.go