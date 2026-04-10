This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule Target

Targets are the resources to be invoked when a rule is triggered. For a complete list of
services and resources that can be set as a target, see [PutTargets](../../../../reference/eventbridge/latest/apireference/api-puttargets.md).

If you are setting the event bus of another account as the target, and that account
granted permission to your account through an organization instead of directly by the account
ID, then you must specify a `RoleArn` with proper permissions in the
`Target` structure. For more information, see [Sending and\
Receiving Events Between AWS Accounts](../../../eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.md) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppSyncParameters" : AppSyncParameters,
  "Arn" : String,
  "BatchParameters" : BatchParameters,
  "DeadLetterConfig" : DeadLetterConfig,
  "EcsParameters" : EcsParameters,
  "HttpParameters" : HttpParameters,
  "Id" : String,
  "Input" : String,
  "InputPath" : String,
  "InputTransformer" : InputTransformer,
  "KinesisParameters" : KinesisParameters,
  "RedshiftDataParameters" : RedshiftDataParameters,
  "RetryPolicy" : RetryPolicy,
  "RoleArn" : String,
  "RunCommandParameters" : RunCommandParameters,
  "SageMakerPipelineParameters" : SageMakerPipelineParameters,
  "SqsParameters" : SqsParameters
}

```

### YAML

```yaml

  AppSyncParameters:
    AppSyncParameters
  Arn: String
  BatchParameters:
    BatchParameters
  DeadLetterConfig:
    DeadLetterConfig
  EcsParameters:
    EcsParameters
  HttpParameters:
    HttpParameters
  Id: String
  Input: String
  InputPath: String
  InputTransformer:
    InputTransformer
  KinesisParameters:
    KinesisParameters
  RedshiftDataParameters:
    RedshiftDataParameters
  RetryPolicy:
    RetryPolicy
  RoleArn: String
  RunCommandParameters:
    RunCommandParameters
  SageMakerPipelineParameters:
    SageMakerPipelineParameters
  SqsParameters:
    SqsParameters

```

## Properties

`AppSyncParameters`

Contains the GraphQL operation to be parsed and executed, if the event target is an
AWS AppSync API.

_Required_: No

_Type_: [AppSyncParameters](aws-properties-events-rule-appsyncparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Arn`

The Amazon Resource Name (ARN) of the target.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BatchParameters`

If the event target is an AWS Batch job, this contains the job definition, job
name, and other parameters. For more information, see [Jobs](../../../batch/latest/userguide/jobs.md) in the _AWS Batch_
_User Guide_.

_Required_: No

_Type_: [BatchParameters](aws-properties-events-rule-batchparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeadLetterConfig`

The `DeadLetterConfig` that defines the target queue to send dead-letter queue
events to.

_Required_: No

_Type_: [DeadLetterConfig](aws-properties-events-rule-deadletterconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcsParameters`

Contains the Amazon ECS task definition and task count to be used, if the event target is
an Amazon ECS task. For more information about Amazon ECS tasks, see [Task\
Definitions](../../../amazonecs/latest/developerguide/task-defintions.md) in the _Amazon EC2 Container Service Developer_
_Guide_.

_Required_: No

_Type_: [EcsParameters](aws-properties-events-rule-ecsparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpParameters`

Contains the HTTP parameters to use when the target is a API Gateway endpoint or
EventBridge ApiDestination.

If you specify an API Gateway API or EventBridge ApiDestination as a target,
you can use this parameter to specify headers, path parameters, and query string keys/values
as part of your target invoking request. If you're using ApiDestinations, the corresponding
Connection can also have these values configured. In case of any conflicting keys, values from
the Connection take precedence.

_Required_: No

_Type_: [HttpParameters](aws-properties-events-rule-httpparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the target within the specified rule. Use this ID to reference the target when
updating the rule. We recommend using a memorable and unique string.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Input`

Valid JSON text passed to the target. In this case, nothing from the event itself is
passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data\
Interchange Format](http://www.rfc-editor.org/rfc/rfc7159.txt).

_Required_: No

_Type_: String

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputPath`

The value of the JSONPath that is used for extracting part of the matched event when
passing it to the target. You may use JSON dot notation or bracket notation. For more
information about JSON paths, see [JSONPath](http://goessner.net/articles/JsonPath).

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputTransformer`

Settings to enable you to provide custom input to a target based on certain event data.
You can extract one or more key-value pairs from the event and then use that data to send
customized input to the target.

_Required_: No

_Type_: [InputTransformer](aws-properties-events-rule-inputtransformer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisParameters`

The custom parameter you can use to control the shard assignment, when the target is a
Kinesis data stream. If you do not include this parameter, the default is to use the
`eventId` as the partition key.

_Required_: No

_Type_: [KinesisParameters](aws-properties-events-rule-kinesisparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftDataParameters`

Contains the Amazon Redshift Data API parameters to use when the target is a Amazon Redshift cluster.

If you specify a Amazon Redshift Cluster as a Target, you can use this to specify
parameters to invoke the Amazon Redshift Data API ExecuteStatement based on EventBridge events.

_Required_: No

_Type_: [RedshiftDataParameters](aws-properties-events-rule-redshiftdataparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryPolicy`

The retry policy configuration to use
for the dead-letter queue.

_Required_: No

_Type_: [RetryPolicy](aws-properties-events-rule-retrypolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is
triggered. If one rule triggers multiple targets, you can use a different IAM role for each
target.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunCommandParameters`

Parameters used when you are using the rule to invoke Amazon EC2 Run Command.

_Required_: No

_Type_: [RunCommandParameters](aws-properties-events-rule-runcommandparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SageMakerPipelineParameters`

Contains the SageMaker AI Model Building Pipeline parameters to start execution of a
SageMaker AI Model Building Pipeline.

If you specify a SageMaker AI Model Building Pipeline as a target, you can use this
to specify parameters to start a pipeline execution based on EventBridge events.

_Required_: No

_Type_: [SageMakerPipelineParameters](aws-properties-events-rule-sagemakerpipelineparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqsParameters`

Contains the message group ID to use when the target is an Amazon SQS fair or FIFO queue.

If you specify a fair or FIFO queue as a target, the queue must have content-based
deduplication enabled.

_Required_: No

_Type_: [SqsParameters](aws-properties-events-rule-sqsparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Target with KinesisParameters](#aws-properties-events-rule-target--examples--Target_with_KinesisParameters)

- [Target with EcsParameters](#aws-properties-events-rule-target--examples--Target_with_EcsParameters)

### Target with KinesisParameters

The following snippet creates a Kinesis data stream target.

#### JSON

```json

"MyEventsRule": {
    "Type": "AWS::Events::Rule",
    "Properties": {
        "Description": "Events Rule with KinesisParameters",
        "EventPattern": {
            "source": [
                "aws.ec2"
            ]
        },
        "RoleArn": {
            "Fn::GetAtt": [
                "EventsInvokeKinesisTargetRole",
                "Arn"
            ]
        },
        "ScheduleExpression": "rate(5 minutes)",
        "State": "ENABLED",
        "Targets": [
            {
                "Arn": {
                    "Fn::GetAtt": [
                        "MyFirstStream",
                        "Arn"
                    ]
                },
                "Id": "Id123",
                "RoleArn": {
                    "Fn::GetAtt": [
                        "EventsInvokeKinesisTargetRole",
                        "Arn"
                    ]
                },
                "KinesisParameters": {
                    "PartitionKeyPath": "$"
                }
            }
        ]
    }
}
```

#### YAML

```yaml

MyEventsRule:
  Type: AWS::Events::Rule
  Properties:
    Description: Events Rule with KinesisParameters
    EventPattern:
      source:
        - aws.ec2
    RoleArn: !GetAtt
      - EventsInvokeKinesisTargetRole
      - Arn
    ScheduleExpression: rate(5 minutes)
    State: ENABLED
    Targets:
      - Arn: !GetAtt
          - MyFirstStream
          - Arn
        Id: Id123
        RoleArn: !GetAtt
          - EventsInvokeKinesisTargetRole
          - Arn
        KinesisParameters:
          PartitionKeyPath: $
```

### Target with EcsParameters

The following snippet creates an Amazon ECS task target.

#### JSON

```json

"MyEventsRule": {
  "Type": "AWS::Events::Rule",
  "Properties": {
      "Description": "Events Rule with EcsParameters",
      "EventPattern": {
          "source": [
              "aws.ec2"
          ],
          "detail-type": [
              "EC2 Instance State-change Notification"
          ],
          "detail": {
              "state": [
                  "stopping"
              ]
          }
      },
      "ScheduleExpression": "rate(15 minutes)",
      "State": "DISABLED",
      "Targets": [
          {
              "Arn": {
                  "Fn::GetAtt": [
                      "MyCluster",
                      "Arn"
                  ]
              },
              "RoleArn": {
                  "Fn::GetAtt": [
                      "ECSTaskRole",
                      "Arn"
                  ]
              },
              "Id": "Id345",
              "EcsParameters": {
                  "TaskCount": 1,
                  "TaskDefinitionArn": {
                      "Ref": "MyECSTask"
                  }
              }
          }
      ]
  }
}
```

#### YAML

```yaml

MyEventsRule:
  Type: AWS::Events::Rule
  Properties:
    Description: Events Rule with EcsParameters
    EventPattern:
      source:
        - aws.ec2
      detail-type:
        - EC2 Instance State-change Notification
      detail:
        state:
          - stopping
    ScheduleExpression: rate(15 minutes)
    State: DISABLED
    Targets:
      - Arn: !GetAtt
          - MyCluster
          - Arn
        RoleArn: !GetAtt
          - ECSTaskRole
          - Arn
        Id: Id345
        EcsParameters:
          TaskCount: 1
          TaskDefinitionArn: !Ref MyECSTask
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
