# TeamWork


This is the team work application each directory has a detailed to do that should help understand what this program wiil do,
this is just a weekend project type of deal so it will take some time to build out.

Main next project/goal
- connect front end socket with the backend
- work on developing chat rooms 


June 21st update
- Decided to use Ent as ORM instead of GORM good choice
- Ran first migration with autoupdate mirgration
- Todos
-   Create Functions for CRUD capabilities on Users and messages
-   Created edges for users and messages
-   Create todos for users and messages
-   Create edges for todos

June 30th update
- Partially dealing with the ability to create a user
- Need to add Validation sequence in the handler method for creating a user
- clean up UI with HTMX for the front end create user form
- Deal with validation if a user is not unique and process that error from server to UI

July 6th
- Cleaned up the DAL capability for creating a user and validating whether a username and password already exists
- Need to connect the user finder and usercreation function
- Add animation for dealing with authentication errors
- Need to give ability to user to create chatrooms after registering
