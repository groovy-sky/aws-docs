This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping

The `AWS::Lambda::EventSourceMapping` resource creates a mapping between an event source and
an AWS Lambda function. Lambda reads items from the event source and triggers the function.

For details about each event source type, see the following topics. In particular, each of the topics
describes the required and optional parameters for the specific event source.

- [Configuring a Dynamo DB stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-dynamodb-eventsourcemapping)

- [Configuring a Kinesis stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-eventsourcemapping)

- [Configuring an SQS queue as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-eventsource)

- [Configuring an MQ broker as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-eventsourcemapping)

- [Configuring MSK as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html)

- [Configuring Self-Managed Apache Kafka as an event source](https://docs.aws.amazon.com/lambda/latest/dg/kafka-smaa.html)

- [Configuring Amazon DocumentDB as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-documentdb.html)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::EventSourceMapping",
  "Properties" : {
      "AmazonManagedKafkaEventSourceConfig" : AmazonManagedKafkaEventSourceConfig,
      "BatchSize" : Integer,
      "BisectBatchOnFunctionError" : Boolean,
      "DestinationConfig" : DestinationConfig,
      "DocumentDBEventSourceConfig" : DocumentDBEventSourceConfig,
      "Enabled" : Boolean,
      "EventSourceArn" : String,
      "FilterCriteria" : FilterCriteria,
      "FunctionName" : String,
      "FunctionResponseTypes" : [ String, ... ],
      "KmsKeyArn" : String,
      "LoggingConfig" : LoggingConfig,
      "MaximumBatchingWindowInSeconds" : Integer,
      "MaximumRecordAgeInSeconds" : Integer,
      "MaximumRetryAttempts" : Integer,
      "MetricsConfig" : MetricsConfig,
      "ParallelizationFactor" : Integer,
      "ProvisionedPollerConfig" : ProvisionedPollerConfig,
      "Queues" : [ String, ... ],
      "ScalingConfig" : ScalingConfig,
      "SelfManagedEventSource" : SelfManagedEventSource,
      "SelfManagedKafkaEventSourceConfig" : SelfManagedKafkaEventSourceConfig,
      "SourceAccessConfigurations" : [ SourceAccessConfiguration, ... ],
      "StartingPosition" : String,
      "StartingPositionTimestamp" : Number,
      "Tags" : [ Tag, ... ],
      "Topics" : [ String, ... ],
      "TumblingWindowInSeconds" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::EventSourceMapping
Properties:
  AmazonManagedKafkaEventSourceConfig:
    AmazonManagedKafkaEventSourceConfig
  BatchSize: Integer
  BisectBatchOnFunctionError: Boolean
  DestinationConfig:
    DestinationConfig
  DocumentDBEventSourceConfig:
    DocumentDBEventSourceConfig
  Enabled: Boolean
  EventSourceArn: String
  FilterCriteria:
    FilterCriteria
  FunctionName: String
  FunctionResponseTypes:
    - String
  KmsKeyArn: String
  LoggingConfig:
    LoggingConfig
  MaximumBatchingWindowInSeconds: Integer
  MaximumRecordAgeInSeconds: Integer
  MaximumRetryAttempts: Integer
  MetricsConfig:
    MetricsConfig
  ParallelizationFactor: Integer
  ProvisionedPollerConfig:
    ProvisionedPollerConfig
  Queues:
    - String
  ScalingConfig:
    ScalingConfig
  SelfManagedEventSource:
    SelfManagedEventSource
  SelfManagedKafkaEventSourceConfig:
    SelfManagedKafkaEventSourceConfig
  SourceAccessConfigurations:
    - SourceAccessConfiguration
  StartingPosition: String
  StartingPositionTimestamp: Number
  Tags:
    - Tag
  Topics:
    - String
  TumblingWindowInSeconds: Integer

```

## Properties

`AmazonManagedKafkaEventSourceConfig`

Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.

_Required_: No

_Type_: [AmazonManagedKafkaEventSourceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-amazonmanagedkafkaeventsourceconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BatchSize`

The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation
(6 MB).

- **Amazon Kinesis** – Default 100. Max 10,000.

- **Amazon DynamoDB Streams** – Default 100. Max 10,000.

- **Amazon Simple Queue Service** – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.

- **Amazon Managed Streaming for Apache Kafka** – Default 100. Max 10,000.

- **Self-managed Apache Kafka** – Default 100. Max 10,000.

- **Amazon MQ (ActiveMQ and RabbitMQ)** – Default 100. Max 10,000.

- **DocumentDB** – Default 100. Max 10,000.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BisectBatchOnFunctionError`

(Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry. The default value is false.

###### Note

When using `BisectBatchOnFunctionError`, check the `BatchSize` parameter in the `OnFailure` destination message's metadata. The `BatchSize`
could be greater than 1 since Lambda consolidates failed messages metadata when writing to the `OnFailure` destination.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationConfig`

(Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) A configuration object that specifies the destination of an event after Lambda processes it.

_Required_: No

_Type_: [DestinationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-destinationconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentDBEventSourceConfig`

Specific configuration settings for a DocumentDB event source.

_Required_: No

_Type_: [DocumentDBEventSourceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-documentdbeventsourceconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

When true, the event source mapping is active. When false, Lambda pauses polling and invocation.

Default: True

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventSourceArn`

The Amazon Resource Name (ARN) of the event source.

- **Amazon Kinesis** – The ARN of the data stream or a stream consumer.

- **Amazon DynamoDB Streams** – The ARN of the stream.

- **Amazon Simple Queue Service** – The ARN of the queue.

- **Amazon Managed Streaming for Apache Kafka** – The ARN of the cluster or the ARN of the VPC connection (for [cross-account event source mappings](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#msk-multi-vpc)).

- **Amazon MQ** – The ARN of the broker.

- **Amazon DocumentDB** – The ARN of the DocumentDB change stream.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.*)`

_Minimum_: `12`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterCriteria`

An object that defines the filter criteria that
determine whether Lambda should process an event. For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).

_Required_: No

_Type_: [FilterCriteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-filtercriteria.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionName`

The name or ARN of the Lambda function.

###### Name formats

- **Function name** – `MyFunction`.

- **Function ARN** – `arn:aws:lambda:us-west-2:123456789012:function:MyFunction`.

- **Version or Alias ARN** – `arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD`.

- **Partial ARN** – `123456789012:function:MyFunction`.

The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64
characters in length.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:(aws[a-zA-Z-]*)?:lambda:)?((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST(\.PUBLISHED)?|[a-zA-Z0-9-_]+))?`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionResponseTypes`

(Kinesis, DynamoDB Streams, and SQS) A list of current response type enums applied to the event source mapping.

Valid Values: `ReportBatchItemFailures`

_Required_: No

_Type_: Array of String

_Allowed values_: `ReportBatchItemFailures`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the AWS Key Management Service (AWS KMS) customer managed key that Lambda
uses to encrypt your function's [filter criteria](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics).

_Required_: No

_Type_: String

_Pattern_: `(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()`

_Minimum_: `12`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingConfig`

(Amazon MSK, and self-managed Apache Kafka only) The logging configuration for your event source. For more information, see [Event source mapping logging](https://docs.aws.amazon.com/lambda/latest/dg/esm-logging.html).

_Required_: No

_Type_: [LoggingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-loggingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBatchingWindowInSeconds`

The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.

**Default (Kinesis, DynamoDB, Amazon SQS event sources)**: 0

**Default (Amazon MSK, Kafka, Amazon MQ, Amazon DocumentDB event sources)**: 500 ms

**Related setting:** For Amazon SQS event sources, when you set `BatchSize`
to a value greater than 10, you must set `MaximumBatchingWindowInSeconds` to at least 1.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRecordAgeInSeconds`

(Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) Discard records older than the specified age. The default value is -1,
which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.

###### Note

The minimum valid value for maximum record age is 60s. Although values less than 60 and greater than -1 fall within the parameter's absolute range, they are not allowed

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `604800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRetryAttempts`

(Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) Discard records after the specified number of retries. The default value is -1,
which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsConfig`

The metrics configuration for your event source. For more information, see [Event source mapping metrics](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics-types.html#event-source-mapping-metrics).

_Required_: No

_Type_: [MetricsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-metricsconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelizationFactor`

(Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard. The default value is 1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedPollerConfig`

(Amazon SQS, Amazon MSK, and self-managed Apache Kafka only) The provisioned mode configuration for the event source.
For more information, see [provisioned mode](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-provisioned-mode).

_Required_: No

_Type_: [ProvisionedPollerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Queues`

(Amazon MQ) The name of the Amazon MQ broker destination queue to consume.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1000 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingConfig`

This property is for Amazon SQS event sources only. You cannot use `ProvisionedPollerConfig` while using `ScalingConfig`.
These options are mutually exclusive. To remove the scaling configuration, pass an empty value.

_Required_: No

_Type_: [ScalingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-scalingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedEventSource`

The self-managed Apache Kafka cluster for your event source.

_Required_: No

_Type_: [SelfManagedEventSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-selfmanagedeventsource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SelfManagedKafkaEventSourceConfig`

Specific configuration settings for a self-managed Apache Kafka event source.

_Required_: No

_Type_: [SelfManagedKafkaEventSourceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-selfmanagedkafkaeventsourceconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceAccessConfigurations`

An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.

_Required_: No

_Type_: Array of [SourceAccessConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html)

_Minimum_: `1`

_Maximum_: `22`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartingPosition`

The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB.

- **LATEST** \- Read only new records.

- **TRIM\_HORIZON** \- Process all available records.

- **AT\_TIMESTAMP** \- Specify a time from which to start reading records.

_Required_: No

_Type_: String

_Pattern_: `(LATEST|TRIM_HORIZON|AT_TIMESTAMP)+`

_Minimum_: `6`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartingPositionTimestamp`

With `StartingPosition` set to `AT_TIMESTAMP`, the time from which to start
reading, in Unix time seconds. `StartingPositionTimestamp` cannot be in the future.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags to add to the event source mapping.

###### Note

You must have the `lambda:TagResource`, `lambda:UntagResource`,
and `lambda:ListTags` permissions for your [IAM principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the CloudFormation stack. If you
don't have these permissions, there might be unexpected behavior with stack-level tags
propagating to the resource during resource creation and update.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-eventsourcemapping-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topics`

The name of the Kafka topic.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `249 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TumblingWindowInSeconds`

(Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds indicates no tumbling window.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the mapping's ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EventSourceMappingArn`

The Amazon Resource Name (ARN) of the event source mapping.

`Id`

The event source mapping's ID.

## Examples

### Event Source Mapping

Create an event source mapping that reads events from Amazon Kinesis and invokes a Lambda function in the
same template.

#### JSON

```json

"EventSourceMapping": {
    "Type": "AWS::Lambda::EventSourceMapping",
    "Properties": {
        "EventSourceArn": {
            "Fn::Join": [
                "",
                [
                    "arn:aws:kinesis:",
                    {
                        "Ref": "AWS::Region"
                    },
                    ":",
                    {
                        "Ref": "AWS::AccountId"
                    },
                    ":stream/",
                    {
                        "Ref": "KinesisStream"
                    }
                ]
            ]
        },
        "FunctionName": {
            "Fn::GetAtt": [
                "LambdaFunction",
                "Arn"
            ]
        },
        "StartingPosition": "TRIM_HORIZON"
    }
}
```

#### YAML

```yaml

MyEventSourceMapping:
  Type: AWS::Lambda::EventSourceMapping
  Properties:
    EventSourceArn:
      Fn::Join:
        - ""
        -
          - "arn:aws:kinesis:"
          -
            Ref: "AWS::Region"
          - ":"
          -
            Ref: "AWS::AccountId"
          - ":stream/"
          -
            Ref: "KinesisStream"
    FunctionName:
      Fn::GetAtt:
        - "LambdaFunction"
        - "Arn"
    StartingPosition: "TRIM_HORIZON"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OnSuccess

AmazonManagedKafkaEventSourceConfig
