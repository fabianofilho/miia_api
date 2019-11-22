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
            apgar5 : "", // valor bruto
            peso : "", // valor bruto
            idade_mae : "", //  valor bruto
            consultas : "", // 1,2,3,4,9
            gestacao : "", // 1,2,3,4,5,6,9
            sit_conjugal_mae : "", // 1,2,3,4,5,9
            sexo : "", // m,f,i
            anomalia : "", // 1,2,9
            gravidez : "", //1,2,3,9
            parto: "" // 1,2,9
        }
    ]
}
@return
{
    success:(bool),
    result: [<probNao>,<probSim>]
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
            input: {
                apgar5 : "", // valor bruto
                peso : "", // valor bruto
                idade_mae : "", //  valor bruto
                consultas : "", // 1,2,3,4,9
                gestacao : "", // 1,2,3,4,5,6,9
                sit_conjugal_mae : "", // 1,2,3,4,5,9
                sexo : "", // m,f,i
                anomalia : "", // 1,2,9
                gravidez : "", //1,2,3,9
                parto: "" // 1,2,9
            },
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
            input: {
                apgar5 : "", // valor bruto
                peso : "", // valor bruto
                idade_mae : "", //  valor bruto
                consultas : "", // 1,2,3,4,9
                gestacao : "", // 1,2,3,4,5,6,9
                sit_conjugal_mae : "", // 1,2,3,4,5,9
                sexo : "", // m,f,i
                anomalia : "", // 1,2,9
                gravidez : "", //1,2,3,9
                parto: "" // 1,2,9
            },
            result: (RESULTADO DO FABIANUS)
        }
}
```