# GetInsightSelectors

Describes the settings for the Insights event selectors that you configured for your
trail or event data store. `GetInsightSelectors` shows if CloudTrail Insights logging is enabled
and which Insights types are configured with corresponding event categories. If you run
`GetInsightSelectors` on a trail or event data store that does not have Insights events enabled,
the operation throws the exception `InsightNotEnabledException`

Specify either the `EventDataStore` parameter to get Insights event selectors for an event data store,
or the `TrailName` parameter to the get Insights event selectors for a trail. You cannot specify these parameters together.

For more information, see [Working with CloudTrail Insights](../../../../services/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.md) in the _AWS CloudTrail User Guide_.

## Request Syntax

```nohighlight

{
   "EventDataStore": "string",
   "TrailName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[EventDataStore](#API_GetInsightSelectors_RequestSyntax)**

Specifies the ARN (or ID suffix of the ARN) of the event data store for which you want to get Insights
selectors.

You cannot use this parameter with the `TrailName` parameter.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

Required: No

**[TrailName](#API_GetInsightSelectors_RequestSyntax)**

Specifies the name of the trail or trail ARN. If you specify a trail name, the string
must meet the following requirements:

- Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
(\_), or dashes (-)

- Start with a letter or number, and end with a letter or number

- Be between 3 and 128 characters

- Have no adjacent periods, underscores or dashes. Names like
`my-_namespace` and `my--namespace` are not valid.

- Not be in IP address format (for example, 192.168.5.4)

If you specify a trail ARN, it must be in the format:

`arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail`

You cannot use this parameter with the `EventDataStore` parameter.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "EventDataStoreArn": "string",
   "InsightsDestination": "string",
   "InsightSelectors": [
      {
         "EventCategories": [ "string" ],
         "InsightType": "string"
      }
   ],
   "TrailARN": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[EventDataStoreArn](#API_GetInsightSelectors_ResponseSyntax)**

The ARN of the source event data store that enabled Insights events.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

**[InsightsDestination](#API_GetInsightSelectors_ResponseSyntax)**

The ARN of the destination event data store that logs Insights events.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

**[InsightSelectors](#API_GetInsightSelectors_ResponseSyntax)**

Contains the Insights types that are enabled on a trail or event data store. It also specifies the event categories on which a particular Insight type is enabled.
`ApiCallRateInsight` and `ApiErrorRateInsight` are valid Insight
types.The EventCategory field can specify `Management` or `Data` events or both. For event data store, you can log Insights for management events only.

Type: Array of [InsightSelector](api-insightselector.md) objects

**[TrailARN](#API_GetInsightSelectors_ResponseSyntax)**

The Amazon Resource Name (ARN) of a trail for which you want to get Insights
selectors.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CloudTrailARNInvalidException**

This exception is thrown when an operation is called with an ARN that is not valid.

The following is the format of a trail ARN: `arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail`

The following is the format of an event data store ARN:
`arn:aws:cloudtrail:us-east-2:123456789012:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE`

The following is the format of a dashboard ARN: `arn:aws:cloudtrail:us-east-1:123456789012:dashboard/exampleDash`

The following is the format of a channel ARN:
`arn:aws:cloudtrail:us-east-2:123456789012:channel/01234567890`

HTTP Status Code: 400

**InsightNotEnabledException**

If you run `GetInsightSelectors` on a trail or event data store that does not have Insights
events enabled, the operation throws the exception
`InsightNotEnabledException`.

HTTP Status Code: 400

**InvalidParameterCombinationException**

This exception is thrown when the combination of parameters provided is not
valid.

HTTP Status Code: 400

**InvalidParameterException**

The request includes a parameter that is not valid.

HTTP Status Code: 400

**InvalidTrailNameException**

This exception is thrown when the provided trail name is not valid. Trail names must
meet the following requirements:

- Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
(\_), or dashes (-)

- Start with a letter or number, and end with a letter or number

- Be between 3 and 128 characters

- Have no adjacent periods, underscores or dashes. Names like
`my-_namespace` and `my--namespace` are not valid.

- Not be in IP address format (for example, 192.168.5.4)

HTTP Status Code: 400

**NoManagementAccountSLRExistsException**

This exception is thrown when the management account does not have a service-linked
role.

HTTP Status Code: 400

**OperationNotPermittedException**

This exception is thrown when the requested operation is not permitted.

HTTP Status Code: 400

**ThrottlingException**

This exception is thrown when the request rate exceeds the limit.

HTTP Status Code: 400

**TrailNotFoundException**

This exception is thrown when the trail with the given name is not found.

HTTP Status Code: 400

**UnsupportedOperationException**

This exception is thrown when the requested operation is not supported.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-2013-11-01/getinsightselectors.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/getinsightselectors.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetImport

GetQueryResults

All content copied from https://docs.aws.amazon.com/.
