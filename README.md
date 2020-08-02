## Overview

This repo tries to implement some of the Azure AD OAuth2.0 flows in Go. There are several examples out there in DOT.NET/DOT.NET CORE, Python, JavaScript etc. but none in Go.

So it would be a nice addition for Gophers to tryout the same in Go and probably later build a MSAL for Go :) 

Currently, the project is in-progress and just finished two **proof-of-concepts** :

*   [OAuth2.0 authorization code flow](https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-auth-code-flow). 

*   [OAuth2.0 device code flow](https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-device-code)


## How to run

Navigate to any of the folders starting with sm-azuread-go* and simply run go run main.go and follow onscreen instruction from terminal.

