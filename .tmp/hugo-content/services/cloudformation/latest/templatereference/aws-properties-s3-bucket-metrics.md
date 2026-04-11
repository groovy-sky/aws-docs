This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket Metrics

A container specifying replication metrics-related settings enabling replication metrics and
events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventThreshold" : ReplicationTimeValue,
  "Status" : String
}

```

### YAML

```yaml

  EventThreshold:
    ReplicationTimeValue
  Status: String

```

## Properties

`EventThreshold`

A container specifying the time threshold for emitting the
`s3:Replication:OperationMissedThreshold` event.

_Required_: No

_Type_: [ReplicationTimeValue](aws-properties-s3-bucket-replicationtimevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Specifies whether the replication metrics are enabled.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](../userguide/aws-properties-s3-bucket.md#aws-properties-s3-bucket--examples)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataTableEncryptionConfiguration

MetricsConfiguration

All content copied from https://docs.aws.amazon.com/.
