# GetUser

Describes the universally unique identifier (UUID) associated with a local user in a
data source.

## Request Syntax

```nohighlight

GET /applications/applicationId/users/userId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_GetUser_RequestSyntax)**

The identifier of the application connected to the user.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[userId](#API_GetUser_RequestSyntax)**

The user email address attached to the user.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "userAliases": [
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

**[userAliases](#API_GetUser_ResponseSyntax)**

A list of user aliases attached to a user.

Type: Array of [UserAlias](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_UserAlias.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qbusiness-2023-11-27/GetUser)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qbusiness-2023-11-27/GetUser)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/GetUser)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qbusiness-2023-11-27/GetUser)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/GetUser)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qbusiness-2023-11-27/GetUser)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qbusiness-2023-11-27/GetUser)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qbusiness-2023-11-27/GetUser)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qbusiness-2023-11-27/GetUser)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/GetUser)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetRetriever

GetWebExperience
