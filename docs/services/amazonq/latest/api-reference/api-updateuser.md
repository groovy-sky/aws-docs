# UpdateUser

Updates a information associated with a user id.

## Request Syntax

```nohighlight

PUT /applications/applicationId/users/userId HTTP/1.1
Content-type: application/json

{
   "userAliasesToDelete": [
      {
         "dataSourceId": "string",
         "indexId": "string",
         "userId": "string"
      }
   ],
   "userAliasesToUpdate": [
      {
         "dataSourceId": "string",
         "indexId": "string",
         "userId": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_UpdateUser_RequestSyntax)**

The identifier of the application the user is attached to.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[userId](#API_UpdateUser_RequestSyntax)**

The email id attached to the user.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[userAliasesToDelete](#API_UpdateUser_RequestSyntax)**

The user aliases attached to the user id that are to be deleted.

Type: Array of [UserAlias](api-useralias.md) objects

Required: No

**[userAliasesToUpdate](#API_UpdateUser_RequestSyntax)**

The user aliases attached to the user id that are to be updated.

Type: Array of [UserAlias](api-useralias.md) objects

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "userAliasesAdded": [
      {
         "dataSourceId": "string",
         "indexId": "string",
         "userId": "string"
      }
   ],
   "userAliasesDeleted": [
      {
         "dataSourceId": "string",
         "indexId": "string",
         "userId": "string"
      }
   ],
   "userAliasesUpdated": [
      {
         "dataSourceId": "string",
         "indexId": "string",
         "userId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[userAliasesAdded](#API_UpdateUser_ResponseSyntax)**

The user aliases that have been to be added to a user id.

Type: Array of [UserAlias](api-useralias.md) objects

**[userAliasesDeleted](#API_UpdateUser_ResponseSyntax)**

The user aliases that have been deleted from a user id.

Type: Array of [UserAlias](api-useralias.md) objects

**[userAliasesUpdated](#API_UpdateUser_ResponseSyntax)**

The user aliases attached to a user id that have been updated.

Type: Array of [UserAlias](api-useralias.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**ConflictException**

You are trying to perform an action that conflicts with the current status of your
resource. Fix any inconsistencies with your resources and try again.

**message**

The message describing a `ConflictException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 409

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ServiceQuotaExceededException**

You have exceeded the set limits for your Amazon Q Business service.

**message**

The message describing a `ServiceQuotaExceededException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 402

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/updateuser.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/updateuser.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateSubscription

UpdateWebExperience

All content copied from https://docs.aws.amazon.com/.
