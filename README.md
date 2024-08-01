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




July 13th
- want to think of things for the ui considering I can add and update users apprpopriately
- UI: 
    login: add or update
    login dashboard:
        - projects
            - chat rooms
        - role of the logged in the user 
            - roles of all users in a project
        - calendars for dates for the projects, timelines and overall todos for assigned users
    - This will be a good start for now....
- finish CRUD abilities
- create a status function for handling messages to send back to the user when CRUD capabilities are being used.



July 19th 
- Cleaned up the create new user and update a user functionality
- need to test that out, then build delete and create query function to load a user
- last thing create dashobard for user when they log in



July 21st 
- committing to chat room development
- Things to do 
    - need to clean up how the hub will be implemented in the server.go file
    - clean up the handling of a new web socket connection when someone is joining a room
    - then break out the chatroom.go file so that there is some decoupling with the different models and associated methods
    - add database implementation to save chatrooms, chatmessages, log chat messages, etc

July 21st night edition
- clean up the rest of the StartServer file in the chatroom module
- then clean up the implementation in the main server.go file





July 23rd 
- need to fix views that will not run and go air issues


July 24th 
- need to figure out how to deal with returning JSON from  the createnewroom and in the future the Joinroom handler function.
- I am thinking the best way to do deal with this is see how to access the context and its stuff
- if not return the error/json as seen fit after the request has been validated


July 26th 
- Able to create and join chatrooms need to be able to broadcast message about a user joining the server. 
- Maybe add functions to be able to grab all clients are get all rooms or get a certain room
- Then need to work on the front end probably.


July 31st 
- Front End Todos
--> Clean up Home Page to log a user in or have them create an account
--> Create User webpage that will alllow them to have a dashboard for projects
    --> Under porjects it should have rooms listed built around assignments 
    --> Show current user role in response to project
--> Chat room should be able to show status of the project and have panels for tracking project status
