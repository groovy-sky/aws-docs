# StartUserAccessTasks

Starts the tasks to search user access status for a specific email address.

The tasks are stopped when the user access status data is found. The tasks are
terminated when the API calls to the application time out.

## Request Syntax

```nohighlight

POST /useraccess/start HTTP/1.1
Content-type: application/json

{
   "appBundleIdentifier": "string",
   "email": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[appBundleIdentifier](#API_StartUserAccessTasks_RequestSyntax)**

The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle
to use for the request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**[email](#API_StartUserAccessTasks_RequestSyntax)**

The email address of the target user.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 320.

Pattern: ``[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*``

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "userAccessTasksList": [
      {
         "app": "string",
         "error": {
            "errorCode": "string",
            "errorMessage": "string"
         },
         "taskId": "string",
         "tenantId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[userAccessTasksList](#API_StartUserAccessTasks_ResponseSyntax)**

Contains a list of user access task information.

Type: Array of [UserAccessTaskItem](api-useraccesstaskitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You are not authorized to perform this operation.

HTTP Status Code: 403

**InternalServerException**

The request processing has failed because of an unknown error, exception, or failure
with an internal server.

**retryAfterSeconds**

The period of time after which you should retry your request.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource does not exist.

**resourceId**

The resource ID.

**resourceType**

The resource type.

HTTP Status Code: 404

**ThrottlingException**

The request rate exceeds the limit.

**quotaCode**

The code for the quota exceeded.

**retryAfterSeconds**

The period of time after which you should retry your request.

**serviceCode**

The code of the service.

HTTP Status Code: 429

**ValidationException**

The request has invalid or missing parameters.

**fieldList**

The field list.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for Python](../../../goto/boto3/appfabric-2023-05-19/startuseraccesstasks.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/startuseraccesstasks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartIngestion

StopIngestion

All content copied from https://docs.aws.amazon.com/.
