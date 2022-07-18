# go-monnify
A monnify library written in golang to enable the interaction with the monnify API 

# Please ensure to create issues in this repo if :
- You encounter any error while using this package and that issue would be attended to immediately.

# Get Started
- In other to use this package, you need to first create an account with monnify via https://app.monnify.com/create-account 
- After your account have been successfully created, locate the developer option on the bottom left of your dashboard to get your:
1. Api key
2. Secret Key
3. Contract code

# Installation
To install this monnify package, you need to install [Go](https://golang.org/) and set your Go workspace first.
1. You can use the below Go command to install go-monnify
```sh
$ go get -u github.com/hisyntax/go-monnify
```
2. Import it in your code:
```sh
import "github.com/hisyntax/go-monnify"
```
## Note : All methods in this package returns three (3) things:
- [x] The object of the response
- [x] An int (status code)
- [x] An error (if any)
