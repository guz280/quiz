How to run the test:
        1. run quizapi: quiz\quizapi> go run main.go
        2. There are 2 options to run the quiz/quizcli>: 
                i. By entering startquiz (this will take you through all the process asking for answer ids as inputs. 
                                         Finally it will give you results for your answers with statistics).
                        quiz/quizcli> go run main.go startquiz 
                ii. By running the calls individually. 
                    This is a list of functions available:
                        a. Get List of question Ids: 
                                        go run main.go questions

                        b. Get Question & List of Possible Answers (Last digit is the question Id): 
                                        go run main.go question 2               

                        c. Post Results & get good answers, bad answers, statistics data as return.
                                        go run main.go results 1 1 1 1 6 7 8
        

***********************************************************************************************
***********************************************************************************************
***********************************************************************************************


In order to enable the program to be more user friendly the program could be developed further by:
        - Validating inputs in startquiz
        - The code assumes the api is running on localhost, but ideally this would be overidable 
                from a command line parameter.
        - The cobra commands, question questions and result are left there as a testing tool. 
                They would be removed in reality.
        - The statistics algorithm accuracy could be sharpened


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





