This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket PartitionedPrefix

Amazon S3 keys for log objects are partitioned in the following format:

`[DestinationPrefix][SourceAccountId]/[SourceRegion]/[SourceBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]`

PartitionedPrefix defaults to EventTime delivery when server access logs are delivered.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PartitionDateSource" : String
}

```

### YAML

```yaml

  PartitionDateSource: String

```

## Properties

`PartitionDateSource`

Specifies the partition date source for the partitioned prefix. `PartitionDateSource` can
be `EventTime` or `DeliveryTime`.

For `DeliveryTime`, the time in the log file names corresponds to the delivery time for
the log files.

For `EventTime`, The logs delivered are for a specific day only. The year, month, and
day correspond to the day on which the event occurred, and the hour, minutes and seconds are set to 00
in the key.

_Required_: No

_Type_: String

_Allowed values_: `EventTime | DeliveryTime`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OwnershipControlsRule

PublicAccessBlockConfiguration
