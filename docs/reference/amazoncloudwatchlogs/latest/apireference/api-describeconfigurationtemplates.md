# DescribeConfigurationTemplates

Use this operation to return the valid and default values that are used when creating
delivery sources, delivery destinations, and deliveries. For more information about
deliveries, see [CreateDelivery](api-createdelivery.md).

## Request Syntax

```nohighlight

{
   "deliveryDestinationTypes": [ "string" ],
   "limit": number,
   "logTypes": [ "string" ],
   "nextToken": "string",
   "resourceTypes": [ "string" ],
   "service": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[deliveryDestinationTypes](#API_DescribeConfigurationTemplates_RequestSyntax)**

Use this parameter to filter the response to include only the configuration templates that
apply to the delivery destination types that you specify here.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 4 items.

Valid Values: `S3 | CWL | FH | XRAY`

Required: No

**[limit](#API_DescribeConfigurationTemplates_RequestSyntax)**

Use this parameter to limit the number of configuration templates that are returned in the
response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[logTypes](#API_DescribeConfigurationTemplates_RequestSyntax)**

Use this parameter to filter the response to include only the configuration templates that
apply to the log types that you specify here.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w]*`

Required: No

**[nextToken](#API_DescribeConfigurationTemplates_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[resourceTypes](#API_DescribeConfigurationTemplates_RequestSyntax)**

Use this parameter to filter the response to include only the configuration templates that
apply to the resource types that you specify here.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w-_]*`

Required: No

**[service](#API_DescribeConfigurationTemplates_RequestSyntax)**

Use this parameter to filter the response to include only the configuration templates that
apply to the AWS service that you specify here.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w_-]*`

Required: No

## Response Syntax

```nohighlight

{
   "configurationTemplates": [
      {
         "allowedActionForAllowVendedLogsDeliveryForResource": "string",
         "allowedFieldDelimiters": [ "string" ],
         "allowedFields": [
            {
               "mandatory": boolean,
               "name": "string"
            }
         ],
         "allowedOutputFormats": [ "string" ],
         "allowedSuffixPathFields": [ "string" ],
         "defaultDeliveryConfigValues": {
            "fieldDelimiter": "string",
            "recordFields": [ "string" ],
            "s3DeliveryConfiguration": {
               "enableHiveCompatiblePath": boolean,
               "suffixPath": "string"
            }
         },
         "deliveryDestinationType": "string",
         "logType": "string",
         "resourceType": "string",
         "service": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[configurationTemplates](#API_DescribeConfigurationTemplates_ResponseSyntax)**

An array of objects, where each object describes one configuration template that matches
the filters that you specified in the request.

Type: Array of [ConfigurationTemplate](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ConfigurationTemplate.html) objects

**[nextToken](#API_DescribeConfigurationTemplates_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 400

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/DescribeConfigurationTemplates)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/DescribeConfigurationTemplates)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeAccountPolicies

DescribeDeliveries
