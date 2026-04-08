# GetFreeTierUsage

Returns a list of all Free Tier usage objects that match your filters.

## Request Syntax

```nohighlight

{
   "filter": {
      "And": [
         "Expression"
      ],
      "Dimensions": {
         "Key": "string",
         "MatchOptions": [ "string" ],
         "Values": [ "string" ]
      },
      "Not": "Expression",
      "Or": [
         "Expression"
      ]
   },
   "maxResults": number,
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[filter](#API_freetier_GetFreeTierUsage_RequestSyntax)**

An expression that specifies the conditions that you want each `FreeTierUsage` object
to meet.

Type: [Expression](api-freetier-expression.md) object

Required: No

**[maxResults](#API_freetier_GetFreeTierUsage_RequestSyntax)**

The maximum number of results to return in the response. `MaxResults` means
that there can be up to the specified number of values, but there might be fewer results based
on your filters.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[nextToken](#API_freetier_GetFreeTierUsage_RequestSyntax)**

The pagination token that indicates the next set of results to retrieve.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8192.

Pattern: `[\S\s]*`

Required: No

## Response Syntax

```nohighlight

{
   "freeTierUsages": [
      {
         "actualUsageAmount": number,
         "description": "string",
         "forecastedUsageAmount": number,
         "freeTierType": "string",
         "limit": number,
         "operation": "string",
         "region": "string",
         "service": "string",
         "unit": "string",
         "usageType": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[freeTierUsages](#API_freetier_GetFreeTierUsage_ResponseSyntax)**

The list of Free Tier usage objects that meet your filter expression.

Type: Array of [FreeTierUsage](api-freetier-freetierusage.md) objects

**[nextToken](#API_freetier_GetFreeTierUsage_ResponseSyntax)**

The pagination token that indicates the next set of results to retrieve.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8192.

Pattern: `[\S\s]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

An unexpected error occurred during the processing of your request.

HTTP Status Code: 500

**ThrottlingException**

The request was denied due to request throttling.

HTTP Status Code: 400

**ValidationException**

The input fails to satisfy the constraints specified by an AWS service.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for Python](../../../../services/goto/boto3/freetier-2023-09-07/getfreetierusage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/freetier-2023-09-07/getfreetierusage.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetAccountPlanState

ListAccountActivities
