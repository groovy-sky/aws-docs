# DeleteUser

Deletes a user from the user pool.

## Request Syntax

```nohighlight

{
   "AuthenticationType": "string",
   "UserName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AuthenticationType](#API_DeleteUser_RequestSyntax)**

The authentication type for the user. You must specify USERPOOL.

Type: String

Valid Values: `API | SAML | USERPOOL | AWS_AD`

Required: Yes

**[UserName](#API_DeleteUser_RequestSyntax)**

The email address of the user.

###### Note

Users' email addresses are case-sensitive.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/deleteuser.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/deleteuser.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteUsageReportSubscription

DescribeAppBlockBuilderAppBlockAssociations

All content copied from https://docs.aws.amazon.com/.
