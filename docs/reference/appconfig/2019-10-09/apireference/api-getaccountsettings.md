# GetAccountSettings

Returns information about the status of the `DeletionProtection`
parameter.

## Request Syntax

```

GET /settings HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

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

**[DeletionProtection](#API_GetAccountSettings_ResponseSyntax)**

A parameter to configure deletion protection. Deletion protection prevents a user from
deleting a configuration profile or an environment if AWS AppConfig has called either
[GetLatestConfiguration](api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](api-getconfiguration.md) for the
configuration profile or from the environment during the specified interval. The default
interval for `ProtectionPeriodInMinutes` is 60.

Type: [DeletionProtectionSettings](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_DeletionProtectionSettings.html) object

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appconfig-2019-10-09/GetAccountSettings)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appconfig-2019-10-09/GetAccountSettings)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteHostedConfigurationVersion

GetApplication
