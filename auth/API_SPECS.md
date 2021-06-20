# auth service

### API details

    status = false means failure case (type of bool)
    success response - 200
    error case 404, 403 

    @POST /login
    @inputs >>  email
    @output 
        message, status



    @POST /verify-otp
    @inputs >>  email, otp
    @output 
        token, message, status


    @POST /verify-token
    @inputs >>  token
    @output 
        status

    @POST /refresh-token
    @inputs >>  token
    @output 
        token, message, status


    @POST /logout
    @inputs >>  token
    @output 
        status