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
- Create the MVP first : Expose an endpoint, handle it through services.
  Create contactInfo. [Ok, missing validations]
  Retrieve that contactInfo by id [Ok, missing validations and correct return]
  Delete that contactInfo  [Ok, validating if exists before deleting]
  Update that contactInfo [Ok, seems to be working just fine]
  Search that contactInfo by email or phone number [Tested by email, seems to be working fine]
   Make two separate search endpoints.
   Search that contactInfo by company or city. [Ok, works great]
- Polish design
  Add validation
- Documentation 
  just a .me explaining on each folder
  Swagger [NO TIME]
- Unit tests  
- Docker

- Check code.

# DB
- Contacts (contactID, name, companyID, profileImage, email, birthDate, addressID)
- Companies (companyID, companyName, addressID)
- Phones (phoneID, contactID, phone, typeID)
- States (stateID, name)
- Cities (cityID, stateID, name)
- Address (addressID, street, number, stateId, cityId)
- PhoneTypes (typeID, name)



TODO: review all annotations, add descriptions to funcs and structs