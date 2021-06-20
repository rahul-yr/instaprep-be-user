# User service

### API Specs


    status = false means failure case (type of bool)
    success response - 200
    error case 404, 403 

    @POST /all-practice-test-type
    @inputs
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name}, {id, name},{id, name}], status


    @POST /all-question-level
    @inputs
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name}, {id, name},{id, name}], status


    @POST /all-learning-path
    @inputs
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name, description}, {id, name, description},{id, name, description}], status


    @POST /all-subject
    @inputs     >>  learning_path_id
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name, description, question_level_ids}, {id, name, description, question_level_ids},{id, name, description, question_level_ids}], status


    @POST /all-topic
    @inputs     >>  subject_id
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name, description}, {id, name, description},{id, name, description}], status


    @POST /all-question
    @inputs     >>  topic_id, page_num
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :{
                results : [
                    {id, question, a, b, c, d, answer, level_id},
                    {id, question, a, b, c, d, answer, level_id}
                    {id, question, a, b, c, d, answer, level_id}
                ],
                total_pages
            }
            , status