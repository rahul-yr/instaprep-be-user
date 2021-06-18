# instaprep-be-user

## Services


###     Mobile Task API DETAILS


    <!-- pre requiste -->

    @POST   ->   Must be cached for every 3 days (for first few days...)
        Get all Domains first
        Get all Question Levels
        Get all Practice Test types
    

    <!-- learning_paths -->

    @POST   ->  Get all learning_paths based on domain id provided
        @inputs >> domain_id
        retuns only the published learning_paths along with the subject_names, subject_ids


    <!-- topics -->

    @POST   ->  Get all the topics related to particular subject_id
        @inputs >> subject_id
        returns all the topic_names, topic_ids


    <!-- Questions  -->
    
    @POST   ->  Get all the questions based on level selected
            @inputs >>  learning_path_id,subject_id, topic_id, level_id
            @optional inputs    >>  starts_at(question_id)

                senario 1 : level_id = ('basic', 'intermediate')
                        free don't worry since basic, intermediate
                            return only 25 records
                
                senario 2 : level_id = ('advanced','coding','realtime')
                        check in Auth Token whether domain_id(is there as part of products) == learning_path_id.domain_id
                            means purchased the requested domain
                                allow access to other question levels
                                else deny

    <!-- Theory -->

    @POST   ->  Get all the Theory based on content_type('public' , 'premium', 'additional') selected
            @inputs >>  learning_path_id,subject_id, topic_id, theory_id, content_type   >> verify all arfe correct
            @optional inputs    >>  page_count

                senario 1 : content_type = ('public')
                        free don't worry since 'public' is free
                            return the theory
                
                senario 1 : content_type = ('premium','additional' )
                        check in Auth Token whether domain_id(is there as part of products) == learning_path_id.domain_id
                            means purchased the requested domain
                                allow access to premium content_type
                                else deny
    
    <!-- Test -->
    @POST   ->  Start the test based on learning_path_id, practice_test_type_id, subject_id (if subject_id == null means He wants to take test based on learning_path else subject)
            @inputs >>  learning_path_id,subject_id, practice_test_type_id
            @optional inputs    >>  subject_id


                senario 1 : If premium user 
                            Get including 25 premium questions (random questions)
                
                senario 2 : else 
                            Get 25 free questions only

                Authetication_check : 
                            check in Auth Token whether domain_id(is there as part of products) == learning_path_id.domain_id means purchased the requested domain
                            allow access to premium else only free