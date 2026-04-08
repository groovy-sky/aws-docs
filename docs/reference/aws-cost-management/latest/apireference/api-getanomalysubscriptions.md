# GetAnomalySubscriptions

Retrieves the cost anomaly subscription objects for your account. You can filter using a
list of cost anomaly monitor Amazon Resource Names (ARNs).

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "MonitorArn": "string",
   "NextPageToken": "string",
   "SubscriptionArnList": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_GetAnomalySubscriptions_RequestSyntax)**

The number of entries a paginated response contains.

Type: Integer

Required: No

**[MonitorArn](#API_GetAnomalySubscriptions_RequestSyntax)**

Cost anomaly monitor ARNs.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: No

**[NextPageToken](#API_GetAnomalySubscriptions_RequestSyntax)**

The token to retrieve the next set of results. AWS provides the token when
the response from a previous call has more results than the maximum page size.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 8192.

Pattern: `[\S\s]*`

Required: No

**[SubscriptionArnList](#API_GetAnomalySubscriptions_RequestSyntax)**

A list of cost anomaly subscription ARNs.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: No

## Response Syntax

```nohighlight

{
   "AnomalySubscriptions": [
      {
         "AccountId": "string",
         "Frequency": "string",
         "MonitorArnList": [ "string" ],
         "Subscribers": [
            {
               "Address": "string",
               "Status": "string",
               "Type": "string"
            }
         ],
         "SubscriptionArn": "string",
         "SubscriptionName": "string",
         "Threshold": number,
         "ThresholdExpression": {
            "And": [
               "Expression"
            ],
            "CostCategories": {
               "Key": "string",
               "MatchOptions": [ "string" ],
               "Values": [ "string" ]
            },
            "Dimensions": {
               "Key": "string",
               "MatchOptions": [ "string" ],
               "Values": [ "string" ]
            },
            "Not": "Expression",
            "Or": [
               "Expression"
            ],
            "Tags": {
               "Key": "string",
               "MatchOptions": [ "string" ],
               "Values": [ "string" ]
            }
         }
      }
   ],
   "NextPageToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[AnomalySubscriptions](#API_GetAnomalySubscriptions_ResponseSyntax)**

A list of cost anomaly subscriptions that includes the detailed metadata for each one.

Type: Array of [AnomalySubscription](api-anomalysubscription.md) objects

**[NextPageToken](#API_GetAnomalySubscriptions_ResponseSyntax)**

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

**UnknownSubscriptionException**

The cost anomaly subscription does not exist for the account.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ce-2017-10-25/getanomalysubscriptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/getanomalysubscriptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetAnomalyMonitors

GetApproximateUsageRecords
