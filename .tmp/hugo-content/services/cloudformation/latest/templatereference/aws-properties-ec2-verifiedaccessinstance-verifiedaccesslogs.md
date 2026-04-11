This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessInstance VerifiedAccessLogs

Describes the options for Verified Access logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogs,
  "IncludeTrustContext" : Boolean,
  "KinesisDataFirehose" : KinesisDataFirehose,
  "LogVersion" : String,
  "S3" : S3
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogs
  IncludeTrustContext: Boolean
  KinesisDataFirehose:
    KinesisDataFirehose
  LogVersion: String
  S3:
    S3

```

## Properties

`CloudWatchLogs`

CloudWatch Logs logging destination.

_Required_: No

_Type_: [CloudWatchLogs](aws-properties-ec2-verifiedaccessinstance-cloudwatchlogs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeTrustContext`

Indicates whether to include trust data sent by trust providers in the logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisDataFirehose`

Kinesis logging destination.

_Required_: No

_Type_: [KinesisDataFirehose](aws-properties-ec2-verifiedaccessinstance-kinesisdatafirehose.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogVersion`

The logging version.

Valid values: `ocsf-0.1` \| `ocsf-1.0.0-rc.2`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Amazon S3 logging options.

_Required_: No

_Type_: [S3](aws-properties-ec2-verifiedaccessinstance-s3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VerifiedAccessTrustProvider

All content copied from https://docs.aws.amazon.com/.
