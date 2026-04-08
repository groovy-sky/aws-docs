# GetIntegration

Returns information about one integration between CloudWatch Logs and OpenSearch Service.

## Request Syntax

```nohighlight

{
   "integrationName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[integrationName](#API_GetIntegration_RequestSyntax)**

The name of the integration that you want to find information about. To find the name of
your integration, use [ListIntegrations](api-listintegrations.md)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

## Response Syntax

```nohighlight

{
   "integrationDetails": { ... },
   "integrationName": "string",
   "integrationStatus": "string",
   "integrationType": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[integrationDetails](#API_GetIntegration_ResponseSyntax)**

A structure that contains information about the integration configuration. For an
integration with OpenSearch Service, this includes information about OpenSearch Service
resources such as the collection, the workspace, and policies.

Type: [IntegrationDetails](api-integrationdetails.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

**[integrationName](#API_GetIntegration_ResponseSyntax)**

The name of the integration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[\.\-_/#A-Za-z0-9]+`

**[integrationStatus](#API_GetIntegration_ResponseSyntax)**

The current status of this integration.

Type: String

Valid Values: `PROVISIONING | ACTIVE | FAILED`

**[integrationType](#API_GetIntegration_ResponseSyntax)**

The type of integration. Integrations with OpenSearch Service have the type
`OPENSEARCH`.

Type: String

Valid Values: `OPENSEARCH`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getintegration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getintegration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getintegration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getintegration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getintegration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getintegration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getintegration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getintegration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getintegration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getintegration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetDeliverySource

GetLogAnomalyDetector
