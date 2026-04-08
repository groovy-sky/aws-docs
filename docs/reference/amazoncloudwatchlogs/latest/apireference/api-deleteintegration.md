# DeleteIntegration

Deletes the integration between CloudWatch Logs and OpenSearch Service. If your
integration has active vended logs dashboards, you must specify `true` for the
`force` parameter, otherwise the operation will fail. If you delete the
integration by setting `force` to `true`, all your vended logs
dashboards powered by OpenSearch Service will be deleted and the data that was on them will no
longer be accessible.

## Request Syntax

```nohighlight

{
   "force": boolean,
   "integrationName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[force](#API_DeleteIntegration_RequestSyntax)**

Specify `true` to force the deletion of the integration even if vended logs
dashboards currently exist.

The default is `false`.

Type: Boolean

Required: No

**[integrationName](#API_DeleteIntegration_RequestSyntax)**

The name of the integration to delete. To find the name of your integration, use [ListIntegrations](api-listintegrations.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/deleteintegration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deleteintegration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteIndexPolicy

DeleteLogAnomalyDetector
