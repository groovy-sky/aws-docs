# ListInsightsData

Returns Insights events generated on a trail that logs data events. You can list Insights events that occurred in a Region within the last 90 days.

ListInsightsData supports the following Dimensions for Insights events:

- Event ID

- Event name

- Event source

All dimensions are optional. The default number of results returned is 50, with a
maximum of 50 possible. The response includes a token that you can use to get the next page
of results.

The rate of ListInsightsData requests is limited to two per second, per account, per Region. If
this limit is exceeded, a throttling error occurs.

## Request Syntax

```nohighlight

{
   "DataType": "string",
   "Dimensions": {
      "string" : "string"
   },
   "EndTime": number,
   "InsightSource": "string",
   "MaxResults": number,
   "NextToken": "string",
   "StartTime": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DataType](#API_ListInsightsData_RequestSyntax)**

Specifies the category of events returned. To fetch Insights events, specify `InsightsEvents` as the value of `DataType`

Type: String

Valid Values: `InsightsEvents`

Required: Yes

**[Dimensions](#API_ListInsightsData_RequestSyntax)**

Contains a map of dimensions. Currently the map can contain only one item.

Type: String to string map

Map Entries: Maximum number of 1 item.

Valid Keys: `EventId | EventName | EventSource`

Value Length Constraints: Minimum length of 1. Maximum length of 2000.

Required: No

**[EndTime](#API_ListInsightsData_RequestSyntax)**

Specifies that only events that occur before or at the specified time are returned. If the specified end time is before the specified start time, an error is returned.

Type: Timestamp

Required: No

**[InsightSource](#API_ListInsightsData_RequestSyntax)**

The Amazon Resource Name(ARN) of the trail for which you want to retrieve Insights events.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

Required: Yes

**[MaxResults](#API_ListInsightsData_RequestSyntax)**

The number of events to return. Possible values are 1 through 50. The default is 50.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[NextToken](#API_ListInsightsData_RequestSyntax)**

The token to use to get the next page of results after a previous API call. This token must be passed in with the same parameters that were specified in the original call.
For example, if the original call specified a EventName as a dimension with `PutObject` as a value, the call with NextToken should include those same parameters.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 1000.

Pattern: `.*`

Required: No

**[StartTime](#API_ListInsightsData_RequestSyntax)**

Specifies that only events that occur after or at the specified time are returned. If the specified start time is after the specified end time, an error is returned.

Type: Timestamp

Required: No

## Response Syntax

```nohighlight

{
   "Events": [
      {
         "AccessKeyId": "string",
         "CloudTrailEvent": "string",
         "EventId": "string",
         "EventName": "string",
         "EventSource": "string",
         "EventTime": number,
         "ReadOnly": "string",
         "Resources": [
            {
               "ResourceName": "string",
               "ResourceType": "string"
            }
         ],
         "Username": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Events](#API_ListInsightsData_ResponseSyntax)**

A list of events returned based on the InsightSource, DataType or Dimensions specified. The events list is sorted by time. The most recent event is listed first.

Type: Array of [Event](api-event.md) objects

**[NextToken](#API_ListInsightsData_ResponseSyntax)**

The token to use to get the next page of results after a previous API call. If the token does not appear, there are no more results to return. The token must be passed in with the same parameters as the previous call.
For example, if the original call specified a EventName as a dimension with `PutObject` as a value, the call with NextToken should include those same parameters.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 1000.

Pattern: `.*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The request includes a parameter that is not valid.

HTTP Status Code: 400

**OperationNotPermittedException**

This exception is thrown when the requested operation is not permitted.

HTTP Status Code: 400

**UnsupportedOperationException**

This exception is thrown when the requested operation is not supported.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-2013-11-01/listinsightsdata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/listinsightsdata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListImports

ListInsightsMetricData

All content copied from https://docs.aws.amazon.com/.
