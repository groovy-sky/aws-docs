This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kinesis::StreamConsumer

Use the AWS CloudFormation
`AWS::Kinesis::StreamConsumer` resource to register a consumer with a
Kinesis data stream. The consumer you register can then call [SubscribeToShard](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_SubscribeToShard.html)
to receive data from the stream using enhanced fan-out, at a rate of up to 2 MiB per
second for every shard you subscribe to. This rate is unaffected by the total number of
consumers that read from the same stream.

You can register up to 20 consumers per stream. However, you can request a limit
increase using the [Kinesis Data Streams limits\
form](https://console.aws.amazon.com/support/v1?). A given consumer can only be registered with one stream at a time.

For more information, see [Using Consumers\
with Enhanced Fan-Out](https://docs.aws.amazon.com/streams/latest/dev/introduction-to-enhanced-consumers.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Kinesis::StreamConsumer",
  "Properties" : {
      "ConsumerName" : String,
      "StreamARN" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Kinesis::StreamConsumer
Properties:
  ConsumerName: String
  StreamARN: String
  Tags:
    - Tag

```

## Properties

`ConsumerName`

The name of the consumer is something you choose when you register the
consumer.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StreamARN`

The ARN of the stream with which you registered the consumer.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws.*:kinesis:.*:\d{12}:stream/\S+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of tags to be added to a specified Kinesis resource. A tag consists of a required key and an optional value. You can specify up to 50 tag key-value pairs.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesis-streamconsumer-tag.html)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of an `AWS::Kinesis::StreamConsumer`
resource to the intrinsic Ref function, the function returns the consumer ARN. For
example ARN formats, see [Example\
ARNs](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arns-syntax).

For more information about using the Ref function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using Fn::GetAtt, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`ConsumerARN`

When you register a consumer, Kinesis Data Streams generates an ARN for it. You
need this ARN to be able to call [SubscribeToShard](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_SubscribeToShard.html).

If you delete a consumer and then create a new one with the same name, it won't
have the same ARN. That's because consumer ARNs contain the creation timestamp. This is
important to keep in mind if you have IAM policies that reference consumer ARNs.

`ConsumerCreationTimestamp`

The time at which the consumer was created.

`ConsumerName`

The name you gave the consumer when you registered it.

`ConsumerStatus`

A consumer can't read data while in the `CREATING` or
`DELETING` states.

`StreamARN`

The ARN of the data stream with which the consumer is registered.

## Examples

### Register a Consumer with a Kinesis Data Stream

#### JSON

```json

{
    "Parameters": {
        "TestStreamARN": {
            "Type": "String" },
        "TestConsumerName": {
            "Type": "String" } },
    "Resources": {
        "StreamConsumer": {
            "Type": "AWS::Kinesis::StreamConsumer",
            "Properties": {
                "StreamARN": { "Ref" : TestStreamARN },
                "ConsumerName": { "Ref" : TestConsumerName }
                }
        }
   }
}
```

#### YAML

```yaml

    Parameters:
        TestStreamARN:
            Type: String
        TestConsumerName:
            Type: String

    Resources: StreamConsumer:
        Type: "AWS::Kinesis::StreamConsumer"
        Properties:
            StreamARN: !Ref TestStreamARN
            ConsumerName: !Ref TestConsumerName
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WarmThroughputObject

Tag
