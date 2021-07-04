# Mock API

Rest API mocking and intercepting in seconds. Replace the endpoint in the code and you are ready. It's that simple!

- [Installation and Setup](#installation-and-setup)
- [Create A Mock API](#create-a-mock-api)
    * [Request](#request)
    * [Response](#response)
        + [Created 201](#created-201)
- [Read a Mock API config](#read-a-mock-api-config)
    * [Request](#request-1)
- [Update a Mock API Config](#update-a-mock-api-config)
    * [Patch Request](#patch-request)
- [Delete A Mock API config](#delete-a-mock-api-config)
    * [Request](#request-2)
    * [Response](#response-1)
- [How to use your mock config](#how-to-use-your-mock-config)
    * [Post Request](#post-request)
    * [Post Response](#post-response)
    * [Delete Request](#delete-request)

## Installation and Setup

1. Clone the project
2. Import and Sync go.mod dependencies

The project contains a manifest.yaml file for cloud-foundry users.

1. Login to your cloud-foundry org and space.
2. Execute `cf push` from the repository directory.

## Create A Mock API

When your mock-api application is running, you can start using it and set up a new config.

### Request

`POST` request to `https://<host>/mockConfig`
The request body should contain the following JSON template:

```json
{
  "post": {
    "status": 201,
    "body": {
      "name": "tal.yaakov@sap.com"
    }
  },
  "delete": {
    "status": 200,
    "body": {
      "name": "tal.yaakov@sap.com"
    }
  }
}
```

- `postResponseStatus` - (int64) the expected response status of your `post` requests
- `deleteResponseStatus` - (int64)  the expected response status of your `delete` requests
- `postResponseBody` - the expected response body of your `post` requests (can be a json format)

### Response

#### Created 201

```json
{
  "id": "32c079dc-a8e1-43d2-bbd0-cb3b2f18d3f8",
  "post": {
    "status": 201,
    "body": {
      "name": "tal.yaakov@sap.com"
    }
  },
  "delete": {
    "status": 200,
    "body": {
      "name": "tal.yaakov@sap.com"
    }
  }
}
```

- `id` - (uuid) the unique id of your config, will be used for mocking your requests to the client and for setup/delete
- `<method>` (`post`/`delete`) - (json object) contains the config (`status` and `body`) of the configured method.
- `status` - (int64) the expected response status of your `post`/`delete` requests
- `body` - the expected response body of your `post`/`delete` requests (can be a json format)

## Read a Mock API config

### Request

`GET` request to `https://<host>/mockConfig?id=<uuid>`
`uuid` - is the unique id of your mock config

**NOTE!**

You can get all configurations in your app memory using the `all=true` flag.

Example:
`GET` request to `https://<host>/mockConfig?all=true`

## Update a Mock API Config

After creating a mock api config, you can change its settings.

This can be used when running a dynamic test that its configurations need to be changed on runtime.

### Patch Request

- `PATCH` request to `https://<host>/?id=<uuid>`
- `id` - (uuid) the unique id of your config, which received on the `Build A Mock API` step The request body should
  contain the following JSON template:

```json
{
  "post": {
    "status": 201,
    "body": {
      "name": "tal.yaakov@sap.com"
    }
  },
  "delete": {
    "status": 400,
    "body": {
      "name": "tal.yaakov@sap.com"
    }
  }
}
```

- `<method>` (`post`/`delete`) - (json object) contains the config (`status` and `body`) of the configured method.
- `status` - (int64) the expected response status of your `post`/`delete` requests
- `body` - the expected response body of your `post`/`delete` requests (can be a json format)

## Delete A Mock API config

When finishing with the mock api config, remember to delete it in order to avoid overloading your application in-memory.

### Request

`DELETE` request to `https://<host>/mockConfig?id=<uuid>`
`uuid` - is the unique id of your mock config

**NOTE!**

You can delete all configurations in your app memory using the `all=true` flag.

Example:
`DELETE` request to `https://<host>/mockConfig?all=true`

### Response

`OK 200` on success

## How to use your mock config

### Post Request

- `POST` request to `https://<host>/?id=<uuid>`
- `id` - (uuid) the unique id of your config, which received on the `Build A Mock API` step

### Post Response

- `Status Code`: should be as defined earlier.
- The request body should contain the `postResponseBody` defined earlier:

```json
{
  "name": "tal.yaakov@sap.com"
}
```

### Delete Request

- `DELETE` request to `https://<host>/?id=<uuid>`
- `id` - (uuid) the unique id of your config, which received on the `Build A Mock API` step.
