# DescribeUserStackAssociations

Retrieves a list that describes the UserStackAssociation objects. You must specify either or both of the following:

- The stack name

- The user name (email address of the user associated with the stack) and the authentication type for the user

## Request Syntax

```nohighlight

{
   "AuthenticationType": "string",
   "MaxResults": number,
   "NextToken": "string",
   "StackName": "string",
   "UserName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AuthenticationType](#API_DescribeUserStackAssociations_RequestSyntax)**

The authentication type for the user who is associated with the stack. You must specify USERPOOL.

Type: String

Valid Values: `API | SAML | USERPOOL | AWS_AD`

Required: No

**[MaxResults](#API_DescribeUserStackAssociations_RequestSyntax)**

The maximum size of each page of results.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 500.

Required: No

**[NextToken](#API_DescribeUserStackAssociations_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[StackName](#API_DescribeUserStackAssociations_RequestSyntax)**

The name of the stack that is associated with the user.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[UserName](#API_DescribeUserStackAssociations_RequestSyntax)**

The email address of the user who is associated with the stack.

###### Note

Users' email addresses are case-sensitive.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "UserStackAssociations": [
      {
         "AuthenticationType": "string",
         "SendEmailNotification": boolean,
         "StackName": "string",
         "UserName": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_DescribeUserStackAssociations_ResponseSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If there are no more pages, this value is null.

Type: String

Length Constraints: Minimum length of 1.

**[UserStackAssociations](#API_DescribeUserStackAssociations_ResponseSyntax)**

The UserStackAssociation objects.

Type: Array of [UserStackAssociation](api-userstackassociation.md) objects

Array Members: Minimum number of 1 item. Maximum number of 25 items.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Indicates an incorrect combination of parameters, or a missing parameter.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/describeuserstackassociations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/describeuserstackassociations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeUsers

DisableUser

All content copied from https://docs.aws.amazon.com/.
