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


# Contact info 

Name
Company -- another table with company data. 
ProfileImage -- same talbe
Email -- same talbe
BirthDate -- same talbe
PhoneNumber (Work, Personal) (could therefore have more than 1, ability to add multiple ones) -- have to be on different tables, attached via contact ID
Address (Should have state or city. [this is an assumption]) -- Table with States, Table with City

# DB
- Contacts (contactID, name, companyID, profileImage, email, birthDate)
- Companies (companyID, companyName, street (with number), stateID, cityID)
- Phones (phoneID, contactID, phone, typeID)



