# GetAnomalies

Retrieves all of the cost anomalies detected on your account during the time period that's
specified by the `DateInterval` object. Anomalies are available for up to 90
days.

## Request Syntax

```nohighlight

{
   "DateInterval": {
      "EndDate": "string",
      "StartDate": "string"
   },
   "Feedback": "string",
   "MaxResults": number,
   "MonitorArn": "string",
   "NextPageToken": "string",
   "TotalImpact": {
      "EndValue": number,
      "NumericOperator": "string",
      "StartValue": number
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DateInterval](#API_GetAnomalies_RequestSyntax)**

Assigns the start and end dates for retrieving cost anomalies. The returned anomaly object
will have an `AnomalyEndDate` in the specified time range.

Type: [AnomalyDateInterval](api-anomalydateinterval.md) object

Required: Yes

**[Feedback](#API_GetAnomalies_RequestSyntax)**

Filters anomaly results by the feedback field on the anomaly object.

Type: String

Valid Values: `YES | NO | PLANNED_ACTIVITY`

Required: No

**[MaxResults](#API_GetAnomalies_RequestSyntax)**

The number of entries a paginated response contains.

Type: Integer

Required: No

**[MonitorArn](#API_GetAnomalies_RequestSyntax)**

Retrieves all of the cost anomalies detected for a specific cost anomaly monitor Amazon
Resource Name (ARN).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: No

**[NextPageToken](#API_GetAnomalies_RequestSyntax)**

The token to retrieve the next set of results. AWS provides the token when
the response from a previous call has more results than the maximum page size.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 8192.

Pattern: `[\S\s]*`

Required: No

**[TotalImpact](#API_GetAnomalies_RequestSyntax)**

Filters anomaly results by the total impact field on the anomaly object. For example, you
can filter anomalies `GREATER_THAN 200.00` to retrieve anomalies, with an estimated
dollar impact greater than 200.

Type: [TotalImpactFilter](api-totalimpactfilter.md) object

Required: No

## Response Syntax

```nohighlight

{
   "Anomalies": [
      {
         "AnomalyEndDate": "string",
         "AnomalyId": "string",
         "AnomalyScore": {
            "CurrentScore": number,
            "MaxScore": number
         },
         "AnomalyStartDate": "string",
         "DimensionValue": "string",
         "Feedback": "string",
         "Impact": {
            "MaxImpact": number,
            "TotalActualSpend": number,
            "TotalExpectedSpend": number,
            "TotalImpact": number,
            "TotalImpactPercentage": number
         },
         "MonitorArn": "string",
         "RootCauses": [
            {
               "Impact": {
                  "Contribution": number
               },
               "LinkedAccount": "string",
               "LinkedAccountName": "string",
               "Region": "string",
               "Service": "string",
               "UsageType": "string"
            }
         ]
      }
   ],
   "NextPageToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Anomalies](#API_GetAnomalies_ResponseSyntax)**

A list of cost anomalies.

Type: Array of [Anomaly](api-anomaly.md) objects

**[NextPageToken](#API_GetAnomalies_ResponseSyntax)**

The token to retrieve the next set of results. AWS provides the token when
the response from a previous call has more results than the maximum page size.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 8192.

Pattern: `[\S\s]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidNextTokenException**

The pagination token is invalid. Try again without a pagination token.

HTTP Status Code: 400

**LimitExceededException**

You made too many calls in a short period of time. Try again later.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ce-2017-10-25/getanomalies.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ce-2017-10-25/getanomalies.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/getanomalies.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ce-2017-10-25/getanomalies.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/getanomalies.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ce-2017-10-25/getanomalies.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ce-2017-10-25/getanomalies.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ce-2017-10-25/getanomalies.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ce-2017-10-25/getanomalies.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/getanomalies.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeCostCategoryDefinition

GetAnomalyMonitors
