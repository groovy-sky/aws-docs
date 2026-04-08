# ListLogAnomalyDetectors

Retrieves a list of the log anomaly detectors in the account.

## Request Syntax

```nohighlight

{
   "filterLogGroupArn": "string",
   "limit": number,
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[filterLogGroupArn](#API_ListLogAnomalyDetectors_RequestSyntax)**

Use this to optionally filter the results to only include anomaly detectors that are
associated with the specified log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**[limit](#API_ListLogAnomalyDetectors_RequestSyntax)**

The maximum number of items to return. If you don't specify a value, the default
maximum value of 50 items is used.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[nextToken](#API_ListLogAnomalyDetectors_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "anomalyDetectors": [
      {
         "anomalyDetectorArn": "string",
         "anomalyDetectorStatus": "string",
         "anomalyVisibilityTime": number,
         "creationTimeStamp": number,
         "detectorName": "string",
         "evaluationFrequency": "string",
         "filterPattern": "string",
         "kmsKeyId": "string",
         "lastModifiedTimeStamp": number,
         "logGroupArnList": [ "string" ]
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[anomalyDetectors](#API_ListLogAnomalyDetectors_ResponseSyntax)**

An array of structures, where each structure in the array contains information about one
anomaly detector.

Type: Array of [AnomalyDetector](api-anomalydetector.md) objects

**[nextToken](#API_ListLogAnomalyDetectors_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/listloganomalydetectors.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/listloganomalydetectors.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListIntegrations

ListLogGroups
