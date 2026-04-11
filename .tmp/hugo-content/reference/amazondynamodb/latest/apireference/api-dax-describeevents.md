# DescribeEvents

Returns events related to DAX clusters and parameter groups. You can
obtain events specific to a particular DAX cluster or parameter group by
providing the name as a parameter.

By default, only the events occurring within the last 24 hours are returned;
however, you can retrieve up to 14 days' worth of events if necessary.

## Request Syntax

```nohighlight

{
   "Duration": number,
   "EndTime": number,
   "MaxResults": number,
   "NextToken": "string",
   "SourceName": "string",
   "SourceType": "string",
   "StartTime": number
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[Duration](#API_dax_DescribeEvents_RequestSyntax)**

The number of minutes' worth of events to retrieve.

Type: Integer

Required: No

**[EndTime](#API_dax_DescribeEvents_RequestSyntax)**

The end of the time interval for which to retrieve events, specified in ISO 8601
format.

Type: Timestamp

Required: No

**[MaxResults](#API_dax_DescribeEvents_RequestSyntax)**

The maximum number of results to include in the response. If more results exist
than the specified `MaxResults` value, a token is included in the response so
that the remaining results can be retrieved.

The value for `MaxResults` must be between 20 and 100.

Type: Integer

Required: No

**[NextToken](#API_dax_DescribeEvents_RequestSyntax)**

An optional token returned from a prior request. Use this token for pagination of
results from this action. If this parameter is specified, the response includes only
results beyond the token, up to the value specified by
`MaxResults`.

Type: String

Required: No

**[SourceName](#API_dax_DescribeEvents_RequestSyntax)**

The identifier of the event source for which events will be returned. If not
specified, then all sources are included in the response.

Type: String

Required: No

**[SourceType](#API_dax_DescribeEvents_RequestSyntax)**

The event source to retrieve events for. If no value is specified, all events are
returned.

Type: String

Valid Values: `CLUSTER | PARAMETER_GROUP | SUBNET_GROUP`

Required: No

**[StartTime](#API_dax_DescribeEvents_RequestSyntax)**

The beginning of the time interval to retrieve events for, specified in ISO 8601
format.

Type: Timestamp

Required: No

## Response Syntax

```nohighlight

{
   "Events": [
      {
         "Date": number,
         "Message": "string",
         "SourceName": "string",
         "SourceType": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Events](#API_dax_DescribeEvents_ResponseSyntax)**

An array of events. Each element in the array represents one event.

Type: Array of [Event](api-dax-event.md) objects

**[NextToken](#API_dax_DescribeEvents_ResponseSyntax)**

Provides an identifier to allow retrieval of paginated results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Two or more incompatible parameters were specified.

HTTP Status Code: 400

**InvalidParameterValueException**

The value for a parameter is invalid.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/describeevents.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/describeevents.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/describeevents.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/describeevents.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/describeevents.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/describeevents.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/describeevents.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/describeevents.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/describeevents.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/describeevents.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDefaultParameters

DescribeParameterGroups

All content copied from https://docs.aws.amazon.com/.
