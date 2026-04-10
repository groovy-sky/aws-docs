This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster BrokerLogs

The broker logs configuration for this MSK cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogs,
  "Firehose" : Firehose,
  "S3" : S3
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogs
  Firehose:
    Firehose
  S3:
    S3

```

## Properties

`CloudWatchLogs`

Property description not available.

_Required_: No

_Type_: [CloudWatchLogs](aws-properties-msk-cluster-cloudwatchlogs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Firehose`

Details of the Kinesis Data Firehose delivery stream that is the destination for broker logs.

_Required_: No

_Type_: [Firehose](aws-properties-msk-cluster-firehose.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Details of the Amazon S3 destination for broker logs.

_Required_: No

_Type_: [S3](aws-properties-msk-cluster-s3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MSK::Cluster

BrokerNodeGroupInfo

All content copied from https://docs.aws.amazon.com/.
