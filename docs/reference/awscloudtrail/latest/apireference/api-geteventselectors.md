# GetEventSelectors

Describes the settings for the event selectors that you configured for your trail. The
information returned for your event selectors includes the following:

- If your event selector includes read-only events, write-only events, or all
events. This applies to management events, data events, and network activity events.

- If your event selector includes management events.

- If your event selector includes network activity events, the event sources
for which you are logging network activity events.

- If your event selector includes data events, the resources on which you are
logging data events.

For more information about logging management, data, and network activity events, see the following topics
in the _AWS CloudTrail User Guide_:

- [Logging management events](../../../../services/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.md)

- [Logging data events](../../../../services/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md)

- [Logging network activity events](../../../../services/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.md)

## Request Syntax

```nohighlight

{
   "TrailName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[TrailName](#API_GetEventSelectors_RequestSyntax)**

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

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "AdvancedEventSelectors": [
      {
         "FieldSelectors": [
            {
               "EndsWith": [ "string" ],
               "Equals": [ "string" ],
               "Field": "string",
               "NotEndsWith": [ "string" ],
               "NotEquals": [ "string" ],
               "NotStartsWith": [ "string" ],
               "StartsWith": [ "string" ]
            }
         ],
         "Name": "string"
      }
   ],
   "EventSelectors": [
      {
         "DataResources": [
            {
               "Type": "string",
               "Values": [ "string" ]
            }
         ],
         "ExcludeManagementEventSources": [ "string" ],
         "IncludeManagementEvents": boolean,
         "ReadWriteType": "string"
      }
   ],
   "TrailARN": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[AdvancedEventSelectors](#API_GetEventSelectors_ResponseSyntax)**

The advanced event selectors that are configured for the trail.

Type: Array of [AdvancedEventSelector](api-advancedeventselector.md) objects

**[EventSelectors](#API_GetEventSelectors_ResponseSyntax)**

The event selectors that are configured for the trail.

Type: Array of [EventSelector](api-eventselector.md) objects

**[TrailARN](#API_GetEventSelectors_ResponseSyntax)**

The specified trail ARN that has the event selectors.

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

**TrailNotFoundException**

This exception is thrown when the trail with the given name is not found.

HTTP Status Code: 400

**UnsupportedOperationException**

This exception is thrown when the requested operation is not supported.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-2013-11-01/geteventselectors.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/geteventselectors.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetEventDataStore

GetImport

All content copied from https://docs.aws.amazon.com/.
