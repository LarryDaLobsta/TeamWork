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


August 3rd
Have the login and create user page just need icons for different social media logins
Can create users and validate that request, need to go through and add sessions to users that can be implemented whether a user is returning or just created an account. Then probably want to create the login dashboard that a user can see after successfully logging in. Some point need to deal with error handling for the users.


Then deal with chatrooms, messages, spreadsheets and more....
December 1st 2025
Cleaned up middleware issues and also added a starter lobby.
The next thing to work on is logging in a user or having a user create account
Then show a dashboard of projects/ chat rooms a user is a part of
Clean up how the chat rooms look add the ability to upload files, emojis, and etc.

December 2nd 2025
Refactored and cleaned up connection string 
Need to clean up chatroom.go with chatgpt suggestions and the handler to reflect new chatroom additions
Then need to test and make sure everything still works
Then create login and join ability for users

December 3rd 2025
Need to put all the modularized code together
Test that it works
then add create login and join ability for users


December 10th 2025
Modularized code does work now working on
login and join look bll/auth_signup.go, dal/user_dal.go, and server.go
    -- These are putting together the first round of logic for me in terms of creating a user just need some clean up there and also in the server.go file as well



December 13th
Cleaning up changes
-- finish debugging so application can run
-- then test out new additions for valdiations with test cases
-- then move on to cleaning up the front end for a user sign up


Dec 27th 2025
Cleaned up bll issues everything runs 
-- need to clean up createuser.html so that the form can look correct with headers, footers, and etc
-- clean up handler logic. You will see on line 255 and 256 where handler logic needs to be added.
-- Finish out the rest of the crud for users and go from there


Dec 28th 2025
Cleaning up newUserPostHandler in the server file
-- form looks good just need to clean up the backend to handle the request
-- line 44 is where you should start and link everything up with the handler rappers and what the handler should pass to the bll.

Dec 29th 2025
Server runs and on successful addition user can be created, just gotcha cases to clean up and front end to fix.
ToDo:
	•	Extract signup form into a reusable partial template (HTMX-friendly)
	•	Extract navigation bar into a shared layout partial
	•	Handle database unique-constraint errors (email / username) and map them to user-facing validation messages
	•	Improve validation flow to return form partials on error (no full-page re-render with HTMX)


Jan 3rd, 2026
    - Cleaned up route handling for htmx and handlers for the createuser
    - it runs
    - issue with redirect need to change logindashboard to go like fashion like how I did the createuser page. This will allow me to call the file in the render like I did with post and get for crate user


Jan 7th, 2026
    - Routing is a lot cleaner with the breakup of explicit blocks in go templates. Rather than global templates that caused a lot of bleeding over.
    - Next that needs to be done is go full crud, work on middle ware, sessions and cookies then shared layouts and make front end beautiful.

January 10th, 2026
    - cleaned up middleware authentication, flow, and re-direction for authenticated and unauthenticated users
    - Lets create the rest of CRUD for user, then create edit, delete form, edit and delete routes and handlers

January 10th, 2026 later on
    - added partial and full view for the edit and delete user. 
    - need to build out backend routes, handlers, BLL, and DAL. Plus any helpers
    - add buttons to user account page


January 31st, 2026
    - begin building out end to end for the update user
    - look at auth_edit_user.go need to build bll 
    - then after build out route and handler
    - then hook up to front end
    - test



February 2nd, 2026
    - Started creating helpers in helper folder for the edit form use those for edit user bll 
    - look in chatgpt chat for flow of how to use them and where. 
    - Hint cant use nil for VM we are checking each field and returning empty strings in the error field 
    - The front end will only show errors for non emtyp error strings