# CreateUser

Creates a new user in the user pool.

## Request Syntax

```nohighlight

{
   "AuthenticationType": "string",
   "FirstName": "string",
   "LastName": "string",
   "MessageAction": "string",
   "UserName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AuthenticationType](#API_CreateUser_RequestSyntax)**

The authentication type for the user. You must specify USERPOOL.

Type: String

Valid Values: `API | SAML | USERPOOL | AWS_AD`

Required: Yes

**[FirstName](#API_CreateUser_RequestSyntax)**

The first name, or given name, of the user.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `^[A-Za-z0-9_\-\s]+$`

Required: No

**[LastName](#API_CreateUser_RequestSyntax)**

The last name, or surname, of the user.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `^[A-Za-z0-9_\-\s]+$`

Required: No

**[MessageAction](#API_CreateUser_RequestSyntax)**

The action to take for the welcome email that is sent to a user after the user is created in the user pool. If you specify SUPPRESS, no email is sent. If you specify RESEND, do not specify the first name or last name of the user. If the value is null, the email is sent.

###### Note

The temporary password in the welcome email is valid for only 7 days. If users don’t set their passwords within 7 days, you must send them a new welcome email.

Type: String

Valid Values: `SUPPRESS | RESEND`

Required: No

**[UserName](#API_CreateUser_RequestSyntax)**

The email address of the user.

###### Note

Users' email addresses are case-sensitive. During login, if they specify an email address that doesn't use the same capitalization as the email address specified when their user pool account was created, a "user does not exist" error message displays.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidAccountStatusException**

The resource cannot be created because your AWS account is suspended. For assistance, contact AWS Support.

**Message**

The error message in the exception.

HTTP Status Code: 400

**InvalidParameterCombinationException**

Indicates an incorrect combination of parameters, or a missing parameter.

**Message**

The error message in the exception.

HTTP Status Code: 400

**LimitExceededException**

The requested limit exceeds the permitted limit for an account.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceAlreadyExistsException**

The specified resource already exists.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/createuser.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/createuser.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/createuser.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/createuser.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/createuser.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/createuser.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/createuser.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/createuser.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/createuser.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/createuser.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateUsageReportSubscription

DeleteAppBlock
