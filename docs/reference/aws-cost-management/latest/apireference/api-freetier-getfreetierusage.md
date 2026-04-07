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

Type: [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_freetier_Expression.html) object

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

Type: Array of [FreeTierUsage](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_freetier_FreeTierUsage.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/freetier-2023-09-07/GetFreeTierUsage)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/freetier-2023-09-07/GetFreeTierUsage)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetAccountPlanState

ListAccountActivities
