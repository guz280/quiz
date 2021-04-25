
In order to enable the program to be more user friendly the program could be developed further by:
        - Validating inputs in startquiz
        - The code assumes the api is running on localhost, but ideally this would be overidable 
                from a command line parameter.
        - The cobra commands, question questions and result are left there as a testing tool. 
                They would be removed in reality.


To run quizapi rn as follows:
quiz\quizapi> go run main.go

***********************************************************************************************
***********************************************************************************************
***********************************************************************************************

Urls:
Get all Question Ids: http://localhost:1000/questions
Get Question And Answers: http://localhost:1000/question?id=2
Post Answer for questions: http://localhost:1000/results
        - array length has to be 7, since we have 7 Question & Answers prepared
        - payload example:- {"answers":[{"questionId":1,"answerId":3},{"questionId":2,"answerId":3},{"questionId":3,"answerId":1},{"questionId":4,"answerId":1},{"questionId":5,"answerId":1},{"questionId":6,"answerId":1},{"questionId":7,"answerId":3}]}

***********************************************************************************************
***********************************************************************************************
***********************************************************************************************




