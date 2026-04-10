This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule Action

Describes the actions associated with a rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudwatchAlarm" : CloudwatchAlarmAction,
  "CloudwatchLogs" : CloudwatchLogsAction,
  "CloudwatchMetric" : CloudwatchMetricAction,
  "DynamoDB" : DynamoDBAction,
  "DynamoDBv2" : DynamoDBv2Action,
  "Elasticsearch" : ElasticsearchAction,
  "Firehose" : FirehoseAction,
  "Http" : HttpAction,
  "IotAnalytics" : IotAnalyticsAction,
  "IotEvents" : IotEventsAction,
  "IotSiteWise" : IotSiteWiseAction,
  "Kafka" : KafkaAction,
  "Kinesis" : KinesisAction,
  "Lambda" : LambdaAction,
  "Location" : LocationAction,
  "OpenSearch" : OpenSearchAction,
  "Republish" : RepublishAction,
  "S3" : S3Action,
  "Sns" : SnsAction,
  "Sqs" : SqsAction,
  "StepFunctions" : StepFunctionsAction,
  "Timestream" : TimestreamAction
}

```

### YAML

```yaml

  CloudwatchAlarm:
    CloudwatchAlarmAction
  CloudwatchLogs:
    CloudwatchLogsAction
  CloudwatchMetric:
    CloudwatchMetricAction
  DynamoDB:
    DynamoDBAction
  DynamoDBv2:
    DynamoDBv2Action
  Elasticsearch:
    ElasticsearchAction
  Firehose:
    FirehoseAction
  Http:
    HttpAction
  IotAnalytics:
    IotAnalyticsAction
  IotEvents:
    IotEventsAction
  IotSiteWise:
    IotSiteWiseAction
  Kafka:
    KafkaAction
  Kinesis:
    KinesisAction
  Lambda:
    LambdaAction
  Location:
    LocationAction
  OpenSearch:
    OpenSearchAction
  Republish:
    RepublishAction
  S3:
    S3Action
  Sns:
    SnsAction
  Sqs:
    SqsAction
  StepFunctions:
    StepFunctionsAction
  Timestream:
    TimestreamAction

```

## Properties

`CloudwatchAlarm`

Change the state of a CloudWatch alarm.

_Required_: No

_Type_: [CloudwatchAlarmAction](aws-properties-iot-topicrule-cloudwatchalarmaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudwatchLogs`

Sends data to CloudWatch.

_Required_: No

_Type_: [CloudwatchLogsAction](aws-properties-iot-topicrule-cloudwatchlogsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudwatchMetric`

Capture a CloudWatch metric.

_Required_: No

_Type_: [CloudwatchMetricAction](aws-properties-iot-topicrule-cloudwatchmetricaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDB`

Write to a DynamoDB table.

_Required_: No

_Type_: [DynamoDBAction](aws-properties-iot-topicrule-dynamodbaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamoDBv2`

Write to a DynamoDB table. This is a new version of the DynamoDB action. It allows
you to write each attribute in an MQTT message payload into a separate DynamoDB
column.

_Required_: No

_Type_: [DynamoDBv2Action](aws-properties-iot-topicrule-dynamodbv2action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Elasticsearch`

Write data to an Amazon OpenSearch Service domain.

###### Note

The `Elasticsearch` action can only be used by existing rule actions.
To create a new rule action or to update an existing rule action, use the
`OpenSearch` rule action instead. For more information, see
[OpenSearchAction](../../../../reference/iot/latest/apireference/api-opensearchaction.md).

_Required_: No

_Type_: [ElasticsearchAction](aws-properties-iot-topicrule-elasticsearchaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Firehose`

Write to an Amazon Kinesis Firehose stream.

_Required_: No

_Type_: [FirehoseAction](aws-properties-iot-topicrule-firehoseaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Http`

Send data to an HTTPS endpoint.

_Required_: No

_Type_: [HttpAction](aws-properties-iot-topicrule-httpaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotAnalytics`

Sends message data to an AWS IoT Analytics channel.

_Required_: No

_Type_: [IotAnalyticsAction](aws-properties-iot-topicrule-iotanalyticsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotEvents`

Sends an input to an AWS IoT Events detector.

_Required_: No

_Type_: [IotEventsAction](aws-properties-iot-topicrule-ioteventsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotSiteWise`

Sends data from the MQTT message that triggered the rule to AWS IoT SiteWise asset
properties.

_Required_: No

_Type_: [IotSiteWiseAction](aws-properties-iot-topicrule-iotsitewiseaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Kafka`

Send messages to an Amazon Managed Streaming for Apache Kafka (Amazon MSK) or self-managed Apache Kafka cluster.

_Required_: No

_Type_: [KafkaAction](aws-properties-iot-topicrule-kafkaaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Kinesis`

Write data to an Amazon Kinesis stream.

_Required_: No

_Type_: [KinesisAction](aws-properties-iot-topicrule-kinesisaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

Invoke a Lambda function.

_Required_: No

_Type_: [LambdaAction](aws-properties-iot-topicrule-lambdaaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

Sends device location data to [Amazon Location\
Service](../../../location/latest/developerguide/welcome.md).

_Required_: No

_Type_: [LocationAction](aws-properties-iot-topicrule-locationaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenSearch`

Write data to an Amazon OpenSearch Service domain.

_Required_: No

_Type_: [OpenSearchAction](aws-properties-iot-topicrule-opensearchaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Republish`

Publish to another MQTT topic.

_Required_: No

_Type_: [RepublishAction](aws-properties-iot-topicrule-republishaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Write to an Amazon S3 bucket.

_Required_: No

_Type_: [S3Action](aws-properties-iot-topicrule-s3action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sns`

Publish to an Amazon SNS topic.

_Required_: No

_Type_: [SnsAction](aws-properties-iot-topicrule-snsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sqs`

Publish to an Amazon SQS queue.

_Required_: No

_Type_: [SqsAction](aws-properties-iot-topicrule-sqsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepFunctions`

Starts execution of a Step Functions state machine.

_Required_: No

_Type_: [StepFunctionsAction](aws-properties-iot-topicrule-stepfunctionsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timestream`

Writes attributes from an MQTT message.

_Required_: No

_Type_: [TimestreamAction](aws-properties-iot-topicrule-timestreamaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::TopicRule

AssetPropertyTimestamp

All content copied from https://docs.aws.amazon.com/.
