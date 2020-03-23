# Objective
Develop CRUD implementation for Solstice Challenge

# Deliverables
- Domain: Contact info
- Crud operations
- search by email or phone (Phone could return more than 1 value.)
- retrieve all records from the same state or city (obvious pagination here.)

ContactInfo: 
Name
Company
ProfileImage
Email
BirthDate
PhoneNumber (Work, Personal) (could therefore have more than 1, ability to add multiple ones)
Address (Should have state or city. [this is an assumption])

# Further Specification
- Unit test for at least 1 endpoint 
- Make all assumptions needed and document them. Make validations.
- Documentation on how to set the server up and running
- This should be deploy-ready (for AWS or Azure. See how to dockerize)

# Steps
- NO AUTH (says nothing about this)
- Use gorsk implementation for setting up a connection with a PostgreSql db. (Maybe add to the documentation why an SQL db was selected. IS an SQL db needed?)
  Check SQL vs NoSQL first.   
- Create the MVP first : Expose an endpoint, handle it through services, save it on db.
  Retrieve that contactInfo
  Search that contactInfo
  Delete that contactInfo
  Do the searchBy.. stuff
- Polish design
  Add validation
- Documentation 
  just a .me explaining on each folder
  Swagger
- Unit tests  
- Docker

# DB
- Contacts (contactID, name, companyID, profileImage, email, birthDate, addressID)
- Companies (companyID, companyName, addressID)
- Phones (phoneID, contactID, phone, typeID)
- States (stateID, name)
- Cities (cityID, stateID, name)
- Address (addressID, street, number, stateId, cityId)
- PhoneTypes (typeID, name)



TODO: review all annotations, add descriptions to funcs and structs