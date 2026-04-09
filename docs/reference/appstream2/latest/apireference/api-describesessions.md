# DescribeSessions

Retrieves a list that describes the streaming sessions for a specified stack and fleet. If a UserId is provided for the stack and fleet,
only streaming sessions for that user are described. If an authentication type is not provided,
the default is to authenticate users using a streaming URL.

## Request Syntax

```nohighlight

{
   "AuthenticationType": "string",
   "FleetName": "string",
   "InstanceId": "string",
   "Limit": number,
   "NextToken": "string",
   "StackName": "string",
   "UserId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AuthenticationType](#API_DescribeSessions_RequestSyntax)**

The authentication method. Specify `API` for a user
authenticated using a streaming URL or `SAML` for a SAML federated user.
The default is to authenticate users using a streaming URL.

Type: String

Valid Values: `API | SAML | USERPOOL | AWS_AD`

Required: No

**[FleetName](#API_DescribeSessions_RequestSyntax)**

The name of the fleet. This value is case-sensitive.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

Required: Yes

**[InstanceId](#API_DescribeSessions_RequestSyntax)**

The identifier for the instance hosting the session.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[Limit](#API_DescribeSessions_RequestSyntax)**

The size of each page of results. The default value is 20 and the maximum value is 50.

Type: Integer

Required: No

**[NextToken](#API_DescribeSessions_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[StackName](#API_DescribeSessions_RequestSyntax)**

The name of the stack. This value is case-sensitive.

Type: String

Pattern: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

Required: Yes

**[UserId](#API_DescribeSessions_RequestSyntax)**

The user identifier (ID). If you specify a user ID, you must also specify the authentication type.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 128.

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "Sessions": [
      {
         "AuthenticationType": "string",
         "ConnectionState": "string",
         "FleetName": "string",
         "Id": "string",
         "InstanceDrainStatus": "string",
         "InstanceId": "string",
         "MaxExpirationTime": number,
         "NetworkAccessConfiguration": {
            "EniId": "string",
            "EniIpv6Addresses": [ "string" ],
            "EniPrivateIpAddress": "string"
         },
         "StackName": "string",
         "StartTime": number,
         "State": "string",
         "UserId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_DescribeSessions_ResponseSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If there are no more pages, this value is null.

Type: String

Length Constraints: Minimum length of 1.

**[Sessions](#API_DescribeSessions_ResponseSyntax)**

Information about the streaming sessions.

Type: Array of [Session](api-session.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Indicates an incorrect combination of parameters, or a missing parameter.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/describesessions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/describesessions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/describesessions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/describesessions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/describesessions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/describesessions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/describesessions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/describesessions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/describesessions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/describesessions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeImages

DescribeSoftwareAssociations

All content copied from https://docs.aws.amazon.com/.
