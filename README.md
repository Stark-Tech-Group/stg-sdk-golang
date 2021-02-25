# stg-sdk-golang
![Build-Test](https://github.com/Stark-Tech-Group/stg-sdk-golang/workflows/Go/badge.svg)
GoLang sdk for The Stark Platform

## :wrench: Setup 
This sdk project requires host operating system environment variables to be setup before use. The following environment variables are required:

| Key              	| Description          	| Example                                    |
|------------------	|----------------------	|------------------------------------------- |
| STG_SDK_API_HOST 	| The host API address 	| https://stgcapi.staging.starktechgroup.com |
| STG_SDK_API_UN   	| Your username        	| yourUserName                               |
| STG_SDK_API_PW   	| Your password        	| yourPassword                               |

:warning: Do not share your credentials

:warning: Stark Tech Group will never ask you for your password

## About Stark Tech Group
Stark Tech Group is a leader in facility optimization, aligning technology with real-world experience across a diverse portfolio of capabilities. We are a single agent source for Building Automation, Intelligence & Equipment offering a unique customer experience throughout the building lifecycle. We are a collaborative, cross-functional team working together to provide integrative, cost-effective solutions with in-house expertise for any type of building, portfolio or project.

## Examples


### CurVal
Getting the most recent value for a pointId
```go

api := starkapi.Client{}
api.Init(host)
api.Login(un, pw)

pointApi := api.PointApi

pointId := 100
curVal, err := pointApi.CurVal(pointId)
 
 fmt.Printf("cur val: %v\n", curVal.Read.Val)

```

### HisRead
Getting 1000 history reads from a pointId and unix epoch start and end
```go

api := starkapi.Client{}
api.Init(host)
api.Login(un, pw)

pointApi := api.PointApi

pointId, limit, start, end := 100, 1000, 1614024121, 1614110821
hisRead, err := pointApi.HisRead(pointId, int16(limit), int64(start), int64(end))
  
for _, his := range hisRead.His {
	fmt.Printf("his read val: %v\n", his.Val)
}
```
