# GetAuthorizationToken

Retrieves an authorization token. An authorization token represents your IAM
authentication credentials. You can use it to access any Amazon ECR registry that your IAM
principal has access to. The authorization token is valid for 12 hours. This API requires
the `ecr-public:GetAuthorizationToken` and
`sts:GetServiceBearerToken` permissions.

## Response Syntax

```nohighlight

{
   "authorizationData": {
      "authorizationToken": "string",
      "expiresAt": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[authorizationData](#API_GetAuthorizationToken_ResponseSyntax)**

An authorization token data object that corresponds to a public registry.

Type: [AuthorizationData](api-authorizationdata.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

HTTP Status Code: 500

**UnsupportedCommandException**

The action isn't supported in this Region.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-public-2020-10-30/getauthorizationtoken.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-public-2020-10-30/getauthorizationtoken.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeRepositories

GetRegistryCatalogData

All content copied from https://docs.aws.amazon.com/.
