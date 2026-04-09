# ListIntegrations

Returns a list of integrations between CloudWatch Logs and other services in this
account. Currently, only one integration can be created in an account, and this integration
must be with OpenSearch Service.

## Request Syntax

```nohighlight

{
   "integrationNamePrefix": "string",
   "integrationStatus": "string",
   "integrationType": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[integrationNamePrefix](#API_ListIntegrations_RequestSyntax)**

To limit the results to integrations that start with a certain name prefix, specify that
name prefix here.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**[integrationStatus](#API_ListIntegrations_RequestSyntax)**

To limit the results to integrations with a certain status, specify that status
here.

Type: String

Valid Values: `PROVISIONING | ACTIVE | FAILED`

Required: No

**[integrationType](#API_ListIntegrations_RequestSyntax)**

To limit the results to integrations of a certain type, specify that type here.

Type: String

Valid Values: `OPENSEARCH`

Required: No

## Response Syntax

```nohighlight

{
   "integrationSummaries": [
      {
         "integrationName": "string",
         "integrationStatus": "string",
         "integrationType": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[integrationSummaries](#API_ListIntegrations_ResponseSyntax)**

An array, where each object in the array contains information about one CloudWatch Logs
integration in this account.

Type: Array of [IntegrationSummary](api-integrationsummary.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/listintegrations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/listintegrations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/listintegrations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/listintegrations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/listintegrations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/listintegrations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/listintegrations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/listintegrations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/listintegrations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/listintegrations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAnomalies

ListLogAnomalyDetectors

All content copied from https://docs.aws.amazon.com/.
