# UpdateAccountSettings

Updates the value of the `DeletionProtection` parameter.

## Request Syntax

```nohighlight

PATCH /settings HTTP/1.1
Content-type: application/json

{
   "DeletionProtection": {
      "Enabled": boolean,
      "ProtectionPeriodInMinutes": number
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[DeletionProtection](#API_UpdateAccountSettings_RequestSyntax)**

A parameter to configure deletion protection. Deletion protection prevents a user from
deleting a configuration profile or an environment if AWS AppConfig has called either
[GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](api-getconfiguration.md) for the
configuration profile or from the environment during the specified interval. The default
interval for `ProtectionPeriodInMinutes` is 60.

Type: [DeletionProtectionSettings](api-deletionprotectionsettings.md) object

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "DeletionProtection": {
      "Enabled": boolean,
      "ProtectionPeriodInMinutes": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DeletionProtection](#API_UpdateAccountSettings_ResponseSyntax)**

A parameter to configure deletion protection. Deletion protection prevents a user from
deleting a configuration profile or an environment if AWS AppConfig has called either
[GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](api-getconfiguration.md) for the
configuration profile or from the environment during the specified interval. The default
interval for `ProtectionPeriodInMinutes` is 60.

Type: [DeletionProtectionSettings](api-deletionprotectionsettings.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The input fails to satisfy the constraints specified by an AWS service.

**Details**

Detailed information about the input that failed to satisfy the constraints specified by
a call.

HTTP Status Code: 400

**InternalServerException**

There was an internal failure in the AWS AppConfig service.

HTTP Status Code: 500

## Examples

### Example

This example illustrates one usage of UpdateAccountSettings.

#### Sample Request

```

PATCH /settings HTTP/1.1
Host: appconfig.us-east-1.amazonaws.com
Accept-Encoding: identity
Content-Type: application/json
User-Agent: aws-cli/2.17.58 md/awscrt#0.21.2 ua/2.0 os/windows#10 md/arch#amd64 lang/python#3.12.6 md/pyimpl#CPython cfg/retry-mode#standard md/installer#exe md/prompt#off md/command#appconfig.update-account-settings
X-Amz-Date: 20241001T190010Z
Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20241001/us-east-1/appconfig/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature=39c3b3042cd2aEXAMPLE
Content-Length: 41

{"DeletionProtection": {"Enabled": true}}
```

#### Sample Response

```

{
	"DeletionProtection": {
		"Enabled": true,
		"ProtectionPeriodInMinutes": 60
	}
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appconfig-2019-10-09/updateaccountsettings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/updateaccountsettings.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UntagResource

UpdateApplication
