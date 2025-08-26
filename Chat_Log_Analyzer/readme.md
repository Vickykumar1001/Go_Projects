Question: Chat Log Analyzer


Problem Statement:
    Write a chat log processor that:
        1. Extracts the top 10 most used words. 
        2. Identifies most active users.
        3. Calculates average message length per user.
        4. Detects spammers (users sending messages too fast).

implementation: 
  step 1 reading the file using os.Read function and returning it as string
  step 2 processing the string and extracting timestamp, user and message and storing it in slice of MessageDetail struct
  step 3 Extracts the top 10 most used words: using map to store freq of each word and then converting it to slice to sort and return top 10.
  step 4 Identifies most active users.: using map to store freq of each user and then converting it to slice to sort and return top 5.\
  step 5 Calculates average message length per user: counting the length of message of each user and freq of total msg of that user and dividing to get avg length
 step 6: Detects spammers (users sending messages too fast).
How to run: Unzip the file and run `./main.exe`
