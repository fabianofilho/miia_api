# miia_api
MIIA api


## Routes

### Login
```
@url /rest/login
@post
@type form-url-encoded (form normal)
@data
{
    username: (string),
    password: (string)
}
@return
{
    success: (bool),
    token: (string),
    institution: (string)
}
@session
A2%!#23dad#32$
```

### Logout
```
@url /rest/logout
@post @get
@return
{
    success: (bool),
}
@destroy @session
A2%!#23dad#32$
```

### Predict
```
@url /rest/predict/new
@post
@header
{
    token: (string),
}
@data
{
    version: "1.0",
    data:[
        {
            PAYLOAD DO FABIANUS
        }
    ]
}
@return
{
    success:(bool),
    result: (RESULTADO DO FABIANUS)
}
```

### Get All Predicts
```
@url /rest/predict/get/all
@post
@header
{
    token: (string),
}
@return
{
    success:(bool),
    data: [
        {
            input: (PAYLOAD DO FABIANUS),
            result: (RESULTADO DO FABIANUS)
        }
    ]
}
```

### Get One Predic
```
@url /rest/predict/get/{id:[0-9]+}
@post
@header
{
    token: (string),
}
@return
{
    success:(bool),
    data:{
            input: (PAYLOAD DO FABIANUS),
            result: (RESULTADO DO FABIANUS)
        }
}
```