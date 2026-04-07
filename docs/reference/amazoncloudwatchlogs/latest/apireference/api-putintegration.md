# PutIntegration

Creates an integration between CloudWatch Logs and another service in this account.
Currently, only integrations with OpenSearch Service are supported, and currently you can have
only one integration in your account.

Integrating with OpenSearch Service makes it possible for you to create curated vended
logs dashboards, powered by OpenSearch Service analytics. For more information, see [Vended log\
dashboards powered by Amazon OpenSearch Service](../../../../services/amazoncloudwatch/latest/logs/cloudwatchlogs-opensearch-dashboards.md).

You can use this operation only to create a new integration. You can't modify an existing
integration.

## Request Syntax

```nohighlight

{
   "integrationName": "string",
   "integrationType": "string",
   "resourceConfig": { ... }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[integrationName](#API_PutIntegration_RequestSyntax)**

A name for the integration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

**[integrationType](#API_PutIntegration_RequestSyntax)**

The type of integration. Currently, the only supported type is
`OPENSEARCH`.

Type: String

Valid Values: `OPENSEARCH`

Required: Yes

**[resourceConfig](#API_PutIntegration_RequestSyntax)**

A structure that contains configuration information for the integration that you are
creating.

Type: [ResourceConfig](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ResourceConfig.html) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

## Response Syntax

```nohighlight

{
   "integrationName": "string",
   "integrationStatus": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[integrationName](#API_PutIntegration_ResponseSyntax)**

The name of the integration that you just created.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[\.\-_/#A-Za-z0-9]+`

**[integrationStatus](#API_PutIntegration_ResponseSyntax)**

The status of the integration that you just created.

After you create an integration, it takes a few minutes to complete. During this time,
you'll see the status as `PROVISIONING`.

Type: String

Valid Values: `PROVISIONING | ACTIVE | FAILED`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/PutIntegration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/PutIntegration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/PutIntegration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/PutIntegration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/PutIntegration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/PutIntegration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/PutIntegration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/PutIntegration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/PutIntegration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/PutIntegration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutIndexPolicy

PutLogEvents
