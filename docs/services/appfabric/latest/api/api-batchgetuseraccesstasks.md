# BatchGetUserAccessTasks

Gets user access details in a batch request.

This action polls data from the tasks that are kicked off by the
`StartUserAccessTasks` action.

## Request Syntax

```nohighlight

POST /useraccess/batchget HTTP/1.1
Content-type: application/json

{
   "appBundleIdentifier": "string",
   "taskIdList": [ "string" ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[appBundleIdentifier](#API_BatchGetUserAccessTasks_RequestSyntax)**

The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle
to use for the request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**[taskIdList](#API_BatchGetUserAccessTasks_RequestSyntax)**

The tasks IDs to use for the request.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "userAccessResultsList": [
      {
         "app": "string",
         "email": "string",
         "resultStatus": "string",
         "taskError": {
            "errorCode": "string",
            "errorMessage": "string"
         },
         "taskId": "string",
         "tenantDisplayName": "string",
         "tenantId": "string",
         "userFirstName": "string",
         "userFullName": "string",
         "userId": "string",
         "userLastName": "string",
         "userStatus": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[userAccessResultsList](#API_BatchGetUserAccessTasks_ResponseSyntax)**

Contains a list of user access results.

Type: Array of [UserAccessResultItem](api-useraccessresultitem.md) objects

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

- [AWS Command Line Interface V2](../../../goto/cli2/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for Python](../../../goto/boto3/appfabric-2023-05-19/batchgetuseraccesstasks.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/batchgetuseraccesstasks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

ConnectAppAuthorization

All content copied from https://docs.aws.amazon.com/.
