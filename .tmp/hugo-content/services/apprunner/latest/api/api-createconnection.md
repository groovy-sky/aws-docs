# CreateConnection

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Create an AWS App Runner connection resource. App Runner requires a connection resource when you create App Runner services that access private repositories from
certain third-party providers. You can share a connection across multiple services.

A connection resource is needed to access GitHub and Bitbucket repositories. Both require
a user interface approval process through the App Runner console before you can use the
connection.

## Request Syntax

```nohighlight

{
   "ConnectionName": "string",
   "ProviderType": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ConnectionName](#API_CreateConnection_RequestSyntax)**

A name for the new connection. It must be unique across all App Runner connections for the AWS account in the AWS Region.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 32.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

Required: Yes

**[ProviderType](#API_CreateConnection_RequestSyntax)**

The source repository provider.

Type: String

Valid Values: `GITHUB | BITBUCKET`

Required: Yes

**[Tags](#API_CreateConnection_RequestSyntax)**

A list of metadata items that you can associate with your connection resource. A tag is a key-value pair.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Syntax

```nohighlight

{
   "Connection": {
      "ConnectionArn": "string",
      "ConnectionName": "string",
      "CreatedAt": number,
      "ProviderType": "string",
      "Status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Connection](#API_CreateConnection_ResponseSyntax)**

A description of the App Runner connection that's created by this request.

Type: [Connection](api-connection.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ServiceQuotaExceededException**

App Runner can't create this resource. You've reached your account quota for this resource type.

For App Runner per-resource quotas, see [AWS App Runner endpoints and quotas](../../../../general/latest/gr/apprunner.md) in the
_AWS General Reference_.

HTTP Status Code: 400

## Examples

### Create a GitHub connection

This example illustrates how to create a connection to a private GitHub code repository. The connection status after a successful call is
`PENDING_HANDSHAKE`. This is because an authentication handshake with the provider still hasn't happened. Complete the handshake using
the App Runner console. For more information, see [Managing App Runner connections](../dg/manage-connections.md) in
the _AWS App Runner Developer Guide_.

#### Sample Request

```json

$ aws apprunner create-connection --cli-input-json "`cat`"
{
  "ConnectionName": "my-github-connection",
  "ProviderType": "GITHUB"
}
```

#### Sample Response

```json

{
  "Connection": {
    "ConnectionArn": "arn:aws:apprunner:us-east-1:123456789012:connection/my-github-connection",
    "ConnectionName": "my-github-connection",
    "Status": "PENDING_HANDSHAKE",
    "CreatedAt": "2020-11-03T00:32:51Z",
    "ProviderType": "GITHUB"
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/createconnection.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/createconnection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateAutoScalingConfiguration

CreateObservabilityConfiguration

All content copied from https://docs.aws.amazon.com/.
