# User service

### API Specs


    status = false means failure case (type of bool)
    success response - 200
    error case 404, 403 


    @POST /question-level/all
    @inputs
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name}, {id, name},{id, name}], status


    @POST /learning-path/all
    @inputs
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name, description}, {id, name, description},{id, name, description}], status


    @POST /subject/all
    @inputs     >>  learning_path_id
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name, description, question_level_ids}, {id, name, description, question_level_ids},{id, name, description, question_level_ids}], status


    @POST /topic/all
    @inputs     >>  subject_id
    @output 
        json
        failure : 404
                 error , status
        success : 200
            results :[{id, name, description}, {id, name, description},{id, name, description}], status


    @POST /question/all
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