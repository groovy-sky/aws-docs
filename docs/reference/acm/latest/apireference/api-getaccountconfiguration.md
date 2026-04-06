# GetAccountConfiguration

Returns the account configuration options associated with an AWS account.

## Response Syntax

```nohighlight

{
   "ExpiryEvents": {
      "DaysBeforeExpiry": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ExpiryEvents](#API_GetAccountConfiguration_ResponseSyntax)**

Expiration events configuration options associated with the AWS account.

Type: [ExpiryEventsConfiguration](api-expiryeventsconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have access required to perform this action.

HTTP Status Code: 400

**ThrottlingException**

The request was denied because it exceeded a quota.

**throttlingReasons**

One or more reasons why the request was throttled.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/acm-2015-12-08/GetAccountConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/acm-2015-12-08/GetAccountConfiguration)

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/getaccountconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/acm-2015-12-08/getaccountconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/getaccountconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/acm-2015-12-08/getaccountconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/acm-2015-12-08/getaccountconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/acm-2015-12-08/getaccountconfiguration.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/acm-2015-12-08/GetAccountConfiguration)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/getaccountconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExportCertificate

GetCertificate
