This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kinesis::Stream

Creates a Kinesis stream that captures and transports data records that are emitted
from data sources. For information about creating streams, see [CreateStream](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_CreateStream.html) in the Amazon Kinesis API Reference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Kinesis::Stream",
  "Properties" : {
      "DesiredShardLevelMetrics" : [ String, ... ],
      "MaxRecordSizeInKiB" : Integer,
      "Name" : String,
      "RetentionPeriodHours" : Integer,
      "ShardCount" : Integer,
      "StreamEncryption" : StreamEncryption,
      "StreamModeDetails" : StreamModeDetails,
      "Tags" : [ Tag, ... ],
      "WarmThroughputMiBps" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::Kinesis::Stream
Properties:
  DesiredShardLevelMetrics:
    - String
  MaxRecordSizeInKiB: Integer
  Name: String
  RetentionPeriodHours: Integer
  ShardCount: Integer
  StreamEncryption:
    StreamEncryption
  StreamModeDetails:
    StreamModeDetails
  Tags:
    - Tag
  WarmThroughputMiBps: Integer

```

## Properties

`DesiredShardLevelMetrics`

A list of shard-level metrics in properties to enable enhanced monitoring mode.

_Required_: No

_Type_: Array of String

_Maximum_: `7`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRecordSizeInKiB`

The maximum record size of a single record in kibibyte (KiB) that you can write to, and read from a stream.

_Required_: No

_Type_: Integer

_Minimum_: `1024`

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the Kinesis stream. If you don't specify a name, AWS
CloudFormation generates a unique physical ID and uses that ID for the stream name. For
more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RetentionPeriodHours`

The number of hours for the data records that are stored in shards to remain
accessible. The default value is 24. For more information about the stream retention
period, see [Changing the Data Retention\
Period](https://docs.aws.amazon.com/streams/latest/dev/kinesis-extended-retention.html) in the Amazon Kinesis Developer Guide.

_Required_: No

_Type_: Integer

_Minimum_: `24`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShardCount`

The number of shards that the stream uses. For greater provisioned throughput,
increase the number of shards.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamEncryption`

When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack
template and updating your stack disables encryption.

_Required_: No

_Type_: [StreamEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesis-stream-streamencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamModeDetails`

Specifies the capacity mode to which you want to set your data stream. Currently, in
Kinesis Data Streams, you can choose between an **on-demand** capacity mode and a **provisioned** capacity mode for your data streams.

_Required_: No

_Type_: [StreamModeDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesis-stream-streammodedetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
For information about constraints for this property, see [Tag Restrictions](https://docs.aws.amazon.com/streams/latest/dev/tagging.html#tagging-restrictions)
in the _Amazon Kinesis Developer Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesis-stream-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmThroughputMiBps`

The target warm throughput in MB/s that the stream should be scaled to handle. This represents the throughput capacity that will be immediately available for write operations.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you specify an AWS::Kinesis::Stream resource as an argument to the
`Ref` function, AWS CloudFormation returns the stream
name (physical ID).

For more information about using the Ref function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for the `Arn` attribute.

For more information about using Fn::GetAtt, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Arn`

The Amazon resource name (ARN) of the Kinesis stream, such as
`arn:aws:kinesis:us-east-2:123456789012:stream/mystream`.

## Examples

### Create a Stream

The following example creates a `Stream` resource that uses three
shards, sets a seven-day retention period, and specifies the KMS key for
server-side encryption.

#### JSON

```json

"MyStream": {
    "Type": "AWS::Kinesis::Stream",
    "Properties": {
        "Name": "MyKinesisStream",
        "RetentionPeriodHours" : 168,
        "ShardCount": 3,
        "StreamEncryption": {
            "EncryptionType": "KMS",
            "KeyId": "!Ref myKey"
            },
        "Tags": [ {
            "Key": "Environment",
            "Value": "Production" } ]
        }
}
```

#### YAML

```yaml

MyStream:
    Type: AWS::Kinesis::Stream
    Properties:
        Name: MyKinesisStream
        RetentionPeriodHours: 168
        ShardCount: 3
        StreamEncryption:
            EncryptionType: KMS
            KeyId: !Ref myKey
        Tags:
            -
                Key: Environment Value:
                Production
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Kinesis::ResourcePolicy

StreamEncryption
