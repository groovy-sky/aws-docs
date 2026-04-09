# CreateGateway

Creates a backup gateway. After you create a gateway, you can associate it with a server
using the `AssociateGatewayToServer` operation.

## Request Syntax

```nohighlight

{
   "ActivationKey": "string",
   "GatewayDisplayName": "string",
   "GatewayType": "string",
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

**[ActivationKey](#API_BGW_CreateGateway_RequestSyntax)**

The activation key of the created gateway.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[0-9a-zA-Z\-]+`

Required: Yes

**[GatewayDisplayName](#API_BGW_CreateGateway_RequestSyntax)**

The display name of the created gateway.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `[a-zA-Z0-9-]*`

Required: Yes

**[GatewayType](#API_BGW_CreateGateway_RequestSyntax)**

The type of created gateway.

Type: String

Valid Values: `BACKUP_VM`

Required: Yes

**[Tags](#API_BGW_CreateGateway_RequestSyntax)**

A list of up to 50 tags to assign to the gateway. Each tag is a key-value pair.

Type: Array of [Tag](api-bgw-tag.md) objects

Required: No

## Response Syntax

```nohighlight

{
   "GatewayArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[GatewayArn](#API_BGW_CreateGateway_ResponseSyntax)**

The Amazon Resource Name (ARN) of the gateway you create.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 180.

Pattern: `arn:(aws|aws-cn|aws-us-gov):backup-gateway(:[a-zA-Z-0-9]+){3}\/[a-zA-Z-0-9]+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

The operation did not succeed because an internal error occurred. Try again later.

**ErrorCode**

A description of which internal error occured.

HTTP Status Code: 500

**ThrottlingException**

TPS has been limited to protect against intentional or unintentional
high request volumes.

**ErrorCode**

Error: TPS has been limited to protect against intentional or unintentional
high request volumes.

HTTP Status Code: 400

**ValidationException**

The operation did not succeed because a validation error occurred.

**ErrorCode**

A description of what caused the validation error.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for Python](../../../goto/boto3/backup-gateway-2021-01-01/creategateway.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-gateway-2021-01-01/creategateway.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssociateGatewayToServer

DeleteGateway

All content copied from https://docs.aws.amazon.com/.
