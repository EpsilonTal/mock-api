# Mock API
Rest API mocking and intercepting in seconds.
Replace the endpoint in the code and you are ready. It's that simple!

- [Mock API](#mock-api)
    * [Installation and Setup](#installation-and-setup)
    * [Build A Mock API](#build-a-mock-api)
        + [Request](#request)
        + [Response](#response)
            - [Created 201](#created-201)
    * [Delete A Mock API](#delete-a-mock-api)
        + [Request](#request-1)
        + [Response](#response-1)
    * [How to use](#how-to-use)
        + [Post Request](#post-request)
        + [Post Response](#post-response)
    
## Installation and Setup
1. Clone the project
2. Import and Sync go.mod dependencies

The project contains a manifest.yaml file for cloud-foundry users.
1. Login to your cloud-foundry org and space.
2. Execute `cf push` from the repository directory.

## Build A Mock API

When your mock-api application is running, you can start using it and set up a new config.

### Request
`POST` request to `https://<host>/mockConfig`
The request body should contain the following JSON template:
```json
{
    "postResponseStatus": 201,
    "deleteResponseStatus": 200,
    "postResponseBody": {
        "name": "tal.yaakov@sap.com"
    }
}
```
`postResponseStatus` - (int64) the expected response status of your `post` requests
`deleteResponseStatus` - (int64)  the expected response status of your `delete` requests
`postResponseBody` - the expected response body of your `post` requests (can be a json format)

### Response

#### Created 201
```json
{
    "id": "32c079dc-a8e1-43d2-bbd0-cb3b2f18d3f8",
    "postResponseStatus": 201,
    "postResponseBody": {
        "name": "tal.yaakov@sap.com"
    },
    "deleteResponseStatus": 200
}
```
`id` - (uuid) the unique id of your config, will be used for mocking your requests to the client and for setup/delete
`postResponseStatus` - (int64) the expected response status of your `post` requests
`deleteResponseStatus` - (int64)  the expected response status of your `delete` requests
`postResponseBody` - the expected response body of your `post` requests (can be a json format)

## Delete A Mock API
When finishing with the mock api config, remember to delete it in order to avoid overloading your application inmemory.

### Request
`DELETE` request to `https://<host>/mockConfig?id=<uuid>`
`uuid` - is the unique id of your mock config

### Response
`OK 200` on success

## How to use

### Post Request
`POST` request to `https://<host>/?id=<uuid>`
`id` - (uuid) the unique id of your config, which recieved on the `Build A Mock API` step

### Post Response
`Status Code`: should be as defined eralier.
The request body should contain the `postResponseBody` defined earlier:
```json
{
    "name": "tal.yaakov@sap.com"
}
```
### Delete Request
`POST` request to `https://<host>/?id=<uuid>`
`id` - (uuid) the unique id of your config, which recieved on the `Build A Mock API` step

The request body should contain the `postResponseBody` defined earlier:
```json
{
    "name": "tal.yaakov@sap.com"
}
```
