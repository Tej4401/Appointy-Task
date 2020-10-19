# Appointy-Task

1. Sending a get request of the form meeting/<id here> will retrieve all the meetings in the databse with given id.
  ![image](https://user-images.githubusercontent.com/49359890/96423426-f261ac80-1216-11eb-9974-d86fb5e53c33.png)
2. Sending a get request of the form meeting/?start=<start time here>&end=<end time here>’ will retrieve all the meetings in the databse within given timeframe.
  ![image](https://user-images.githubusercontent.com/49359890/96423605-2b018600-1217-11eb-8555-6f9833629d97.png)
3. Sending a get request of the form meeting/?participant=<email id>’ will retrieve all the meetings in the databse of a given participant.
  ![image](https://user-images.githubusercontent.com/49359890/96423713-58e6ca80-1217-11eb-8348-f1ae4654c1b8.png)
4. Sending a post request to /meeting will create a meeting with given details in database
  ![image](https://user-images.githubusercontent.com/49359890/96423855-89c6ff80-1217-11eb-9348-a376415bd69a.png)
  
Below, I have tried to mention what all work I have completed till now and what is remaining. Due to some connectivity issues with DB, I got stuck and lost time thus I was not able to complete the logic for avoding overlapping of meetings and logic for maintaing all the meetings of a participant but I have implemented the basic working of all routes. In a perfectly managed DB, the get requests would work just fine and retrieve required data.

Anyways, it was a fun and learning experience in working on this task provided by Appointy. Learning a new language in one day was never going to be easy, but the difficulty brought with it some really good challenges which were intriguing and fun to solve. Looking forward to working on more such challenges.

Work Completed:

i)Schedule a meeting
Should be a POST request
Use JSON request body
URL should be ‘/meetings’
Must return the meeting in JSON format
ii)Get a meeting using id
Should be a GET request
Id should be in the url parameter
URL should be ‘/meeting/<id here>’
Must return the meeting in JSON format
iii)List all meetings within a time frame
Should be a GET request
URL should be ‘/meetings?start=<start time here>&end=<end time here>’
Must return a an array of meetings in JSON format that are within the time range
iv)List all meetings of a participant
Should be a GET request
URL should be ‘/meetings?participant=<email id>’
Must return a an array of meetings in JSON format that have the participant received in the email within the time range
  
Work not completed till now:

Meetings should not be overlapped i.e. one participant (uniquely identified by email) should not have 2 or more meetings with RSVP Yes with any overlap between their times.




  
