This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SNS::Topic

The `AWS::SNS::Topic` resource creates a topic to which notifications can be
published.

###### Note

One account can create a maximum of 100,000 standard topics and 1,000 FIFO topics.
For more information, see [Amazon SNS endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/sns.html) in the _AWS General Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SNS::Topic",
  "Properties" : {
      "ArchivePolicy" : Json,
      "ContentBasedDeduplication" : Boolean,
      "DataProtectionPolicy" : Json,
      "DeliveryStatusLogging" : [ LoggingConfig, ... ],
      "DisplayName" : String,
      "FifoThroughputScope" : String,
      "FifoTopic" : Boolean,
      "KmsMasterKeyId" : String,
      "SignatureVersion" : String,
      "Subscription" : [ Subscription, ... ],
      "Tags" : [ Tag, ... ],
      "TopicName" : String,
      "TracingConfig" : String
    }
}

```

### YAML

```yaml

Type: AWS::SNS::Topic
Properties:
  ArchivePolicy: Json
  ContentBasedDeduplication: Boolean
  DataProtectionPolicy: Json
  DeliveryStatusLogging:
    - LoggingConfig
  DisplayName: String
  FifoThroughputScope: String
  FifoTopic: Boolean
  KmsMasterKeyId: String
  SignatureVersion: String
  Subscription:
    - Subscription
  Tags:
    - Tag
  TopicName: String
  TracingConfig: String

```

## Properties

`ArchivePolicy`

The `ArchivePolicy` determines the number of days Amazon SNS retains
messages in FIFO topics. You can set a retention period ranging from 1 to 365 days. This
property is only applicable to FIFO topics; attempting to use it with standard topics will
result in a creation failure.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentBasedDeduplication`

`ContentBasedDeduplication` enables deduplication of messages based on their
content for FIFO topics. By default, this property is set to false. If you create a FIFO
topic with `ContentBasedDeduplication` set to false, you must provide a
`MessageDeduplicationId` for each `Publish` action. When set to
true, Amazon SNS automatically generates a `MessageDeduplicationId`
using a SHA-256 hash of the message body (excluding message attributes). You can optionally
override this generated value by specifying a `MessageDeduplicationId` in the
`Publish` action. Note that this property only applies to FIFO topics; using
it with standard topics will cause the creation to fail.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataProtectionPolicy`

The body of the policy document you want to use for this topic.

You can only add one policy per topic.

The policy must be in JSON string format.

Length Constraints: Maximum length of 30,720.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryStatusLogging`

The `DeliveryStatusLogging` configuration enables you to log the delivery
status of messages sent from your Amazon SNS topic to subscribed endpoints with the
following supported delivery protocols:

- HTTP

- Amazon Kinesis Data Firehose

- AWS Lambda

- Platform application endpoint

- Amazon Simple Queue Service

Once configured, log entries are sent to Amazon CloudWatch Logs.

_Required_: No

_Type_: Array of [LoggingConfig](aws-properties-sns-topic-loggingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name to use for an Amazon SNS topic with SMS subscriptions. The
display name must be maximum 100 characters long, including hyphens (-), underscores (\_),
spaces, and tabs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FifoThroughputScope`

Specifies the throughput quota and deduplication behavior to apply for the FIFO topic.
Valid values are `Topic` or `MessageGroup`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FifoTopic`

Set to true to create a FIFO topic.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsMasterKeyId`

The ID of an AWS managed customer master key (CMK) for Amazon SNS
or a custom CMK. For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms). For
more examples, see `
                            KeyId
                        ` in the _AWS Key Management Service API Reference_.

This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SignatureVersion`

The signature version corresponds to the hashing algorithm used while creating the
signature of the notifications, subscription confirmations, or unsubscribe confirmation
messages sent by Amazon SNS. By default, `SignatureVersion` is set to
`1`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subscription`

The Amazon SNS subscriptions (endpoints) for this topic.

###### Important

If you specify the `Subscription` property in the
`AWS::SNS::Topic` resource and it creates an associated subscription
resource, the associated subscription is not deleted when the
`AWS::SNS::Topic` resource is deleted.

_Required_: No

_Type_: [Array](aws-properties-sns-topic-subscription.md) of [Subscription](aws-properties-sns-topic-subscription.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of tags to add to a new topic.

###### Note

To be able to tag a topic on creation, you must have the
`sns:CreateTopic` and `sns:TagResource`
permissions.

_Required_: No

_Type_: Array of [Tag](aws-properties-sns-topic-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicName`

The name of the topic you want to create. Topic names must include only uppercase and
lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256
characters long. FIFO topic names must end with `.fifo`.

If you don't specify a name, CloudFormation generates a unique physical ID and uses
that ID for the topic name. For more information, see [Name\
type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

###### Important

If you specify a name, you can't perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TracingConfig`

Tracing mode of an Amazon SNS topic. By default `TracingConfig` is
set to `PassThrough`, and the topic passes through the tracing header it
receives from an Amazon SNS publisher to its subscriptions. If set to
`Active`, Amazon SNS will vend X-Ray segment data to topic owner
account if the sampled flag in the tracing header is true.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the topic ARN, for example:
`arn:aws:sns:us-east-1:123456789012:mystack-mytopic-NZJ5JSMVGFIE`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TopicArn`

Returns the ARN of an Amazon SNS topic.

`TopicName`

Returns the name of an Amazon SNS topic.

## Examples

### An Amazon SNS topic with two Amazon SQS queue subscriptions

#### JSON

```json

{
  "MySNSTopic": {
    "Type": "AWS::SNS::Topic",
    "Properties": {
      "Subscription": [
        {
          "Endpoint": {
            "Fn::GetAtt": [
              "MyQueue1",
              "Arn"
            ]
          },
          "Protocol": "sqs"
        },
        {
          "Endpoint": {
            "Fn::GetAtt": [
              "MyQueue2",
              "Arn"
            ]
          },
          "Protocol": "sqs"
        }
      ],
      "TopicName": "SampleTopic",
      "ArchivePolicy": {
      "MessageRetentionPeriod": "7"
      }
    }
  },
  "Outputs": {
    "TopicArn": {
      "Description": "The ARN of the created SNS topic",
      "Value": {
        "Ref": "MySNSTopic"
      }
    },
    "TopicName": {
      "Description": "The name of the created SNS topic",
      "Value": "SampleTopic"
    }
  }
}
```

#### YAML

```yaml

MySNSTopic:
  Type: AWS::SNS::Topic
  Properties:
    Subscription:
      - Endpoint:
          Fn::GetAtt:
            - "MyQueue1"
            - "Arn"
        Protocol: "sqs"
      - Endpoint:
          Fn::GetAtt:
            - "MyQueue2"
            - "Arn"
        Protocol: "sqs"
    TopicName: "SampleTopic"
    ArchivePolicy:
      MessageRetentionPeriod: "7"

Outputs:
  TopicArn:
    Description: The ARN of the created SNS topic
    Value: !Ref MySNSTopic
  TopicName:
    Description: The name of the created SNS topic
    Value: "SampleTopic"

```

## See also

- [Using an CloudFormation template to create a topic that sends messages to Amazon SQS queues](https://docs.aws.amazon.com/sns/latest/dg/SendMessageToSQS.cloudformation.html) in the _Amazon SNS Developer Guide_

- The [Using CloudFormation](https://docs.aws.amazon.com/sns/latest/dg/fifo-topic-code-examples.html#fifo-topic-cfn) code example for FIFO topics in the _Amazon SNS Developer Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SNS::Subscription

LoggingConfig
