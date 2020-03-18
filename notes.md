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
- Create the MVP first : Expose an endpoint, handle it through services, save it on db. Retrieve data, delete data, search data. 
- Focus on CRUD, correct Design and validation
- Next, on Unit tests and documentation (Swagger best, fist only a readme)
- Last, see about how to use / implement docker. 



