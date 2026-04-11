# DeleteConnection

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Delete an AWS App Runner connection. You must first ensure that there are no running App Runner services that use this connection. If there are any, the
`DeleteConnection` action fails.

## Request Syntax

```nohighlight

{
   "ConnectionArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ConnectionArn](#API_DeleteConnection_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner connection that you want to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

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

**[Connection](#API_DeleteConnection_ResponseSyntax)**

A description of the App Runner connection that this request just deleted.

Type: [Connection](api-connection.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource doesn't exist for the specified Amazon Resource Name (ARN) in your AWS account.

HTTP Status Code: 400

## Examples

### Delete a connection

This example illustrates deleting an App Runner connection. The connection status after a successful call is `DELETED`. This is because the
connection is no longer available.

#### Sample Request

```json

$ aws apprunner delete-connection --cli-input-json "`cat`"
{
  "ConnectionArn": "arn:aws:apprunner:us-east-1:123456789012:connection/my-github-connection"
}
```

#### Sample Response

```json

{
  "Connection": {
    "ConnectionArn": "arn:aws:apprunner:us-east-1:123456789012:connection/my-github-connection",
    "ConnectionName": "my-github-connection",
    "Status": "DELETED",
    "CreatedAt": "2020-11-03T00:32:51Z",
    "ProviderType": "GITHUB"
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/deleteconnection.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/deleteconnection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAutoScalingConfiguration

DeleteObservabilityConfiguration

All content copied from https://docs.aws.amazon.com/.
