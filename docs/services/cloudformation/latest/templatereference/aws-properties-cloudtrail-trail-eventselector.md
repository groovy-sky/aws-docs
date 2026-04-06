This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Trail EventSelector

Use event selectors to further specify the management and data event settings for your
trail. By default, trails created without specific event selectors will be configured to
log all read and write management events, and no data events. When an event occurs in your
account, CloudTrail evaluates the event selector for all trails. For each trail, if
the event matches any event selector, the trail processes and logs the event. If the event
doesn't match any event selector, the trail doesn't log the event.

You can configure up to five event selectors for a trail.

You cannot apply both event selectors and advanced event selectors to a trail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataResources" : [ DataResource, ... ],
  "ExcludeManagementEventSources" : [ String, ... ],
  "IncludeManagementEvents" : Boolean,
  "ReadWriteType" : String
}

```

### YAML

```yaml

  DataResources:
    - DataResource
  ExcludeManagementEventSources:
    - String
  IncludeManagementEvents: Boolean
  ReadWriteType: String

```

## Properties

`DataResources`

CloudTrail supports data event logging for Amazon S3 objects in standard S3 buckets, AWS Lambda functions, and Amazon DynamoDB tables with basic event selectors.
You can specify up to 250 resources for an individual event selector, but the total number
of data resources cannot exceed 250 across all event selectors in a trail. This limit does
not apply if you configure resource logging for all data events.

For more information, see [Data\
Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) and [Limits in AWS CloudTrail](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) in the _AWS CloudTrail User_
_Guide_.

###### Note

To log data events for all other resource types including objects stored in
[directory buckets](../../../s3/latest/userguide/directory-buckets-overview.md), you must use [AdvancedEventSelectors](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedEventSelector.html). You must also
use `AdvancedEventSelectors` if you want to filter on the `eventName` field.

_Required_: Conditional

_Type_: Array of [DataResource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-trail-dataresource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeManagementEventSources`

An optional list of service event sources from which you do not want management events
to be logged on your trail. In this release, the list can be empty (disables the filter),
or it can filter out AWS Key Management Service or Amazon RDS Data API events by
containing `kms.amazonaws.com` or `rdsdata.amazonaws.com`. By
default, `ExcludeManagementEventSources` is empty, and AWS KMS and
Amazon RDS Data API events are logged to your trail. You can exclude management
event sources only in Regions that support the event source.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeManagementEvents`

Specify if you want your event selector to include management events for your
trail.

For more information, see [Management Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html) in the _AWS CloudTrail User_
_Guide_.

By default, the value is `true`.

The first copy of management events is free. You are charged for additional copies of
management events that you are logging on any subsequent trail in the same Region. For more
information about CloudTrail pricing, see [AWS CloudTrail Pricing](https://aws.amazon.com/cloudtrail/pricing).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadWriteType`

Specify if you want your trail to log read-only events, write-only events, or all. For
example, the EC2 `GetConsoleOutput` is a read-only API operation and
`RunInstances` is a write-only API operation.

By default, the value is `All`.

_Required_: No

_Type_: String

_Allowed values_: `All | ReadOnly | WriteOnly`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataResource

InsightSelector
