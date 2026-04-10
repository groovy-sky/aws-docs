This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe

Specifies a pipe. Amazon EventBridge Pipes connect event sources to targets and reduces the need for specialized knowledge and integration code.

###### Note

As an aid to help you jumpstart developing CloudFormation templates, the EventBridge console
enables you to create templates from the existing pipes in your account. For more information, see [Generate an\
CloudFormation template from EventBridge Pipes](../../../eventbridge/latest/userguide/pipes-generate-template.md) in the _Amazon EventBridge User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pipes::Pipe",
  "Properties" : {
      "Description" : String,
      "DesiredState" : String,
      "Enrichment" : String,
      "EnrichmentParameters" : PipeEnrichmentParameters,
      "KmsKeyIdentifier" : String,
      "LogConfiguration" : PipeLogConfiguration,
      "Name" : String,
      "RoleArn" : String,
      "Source" : String,
      "SourceParameters" : PipeSourceParameters,
      "Tags" : {Key: Value, ...},
      "Target" : String,
      "TargetParameters" : PipeTargetParameters
    }
}

```

### YAML

```yaml

Type: AWS::Pipes::Pipe
Properties:
  Description: String
  DesiredState: String
  Enrichment: String
  EnrichmentParameters:
    PipeEnrichmentParameters
  KmsKeyIdentifier: String
  LogConfiguration:
    PipeLogConfiguration
  Name: String
  RoleArn: String
  Source: String
  SourceParameters:
    PipeSourceParameters
  Tags:
    Key: Value
  Target: String
  TargetParameters:
    PipeTargetParameters

```

## Properties

`Description`

A description of the pipe.

_Required_: No

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DesiredState`

The state the pipe should be in.

_Required_: No

_Type_: String

_Allowed values_: `RUNNING | STOPPED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enrichment`

The ARN of the enrichment resource.

_Required_: No

_Type_: String

_Pattern_: `^$|arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-]+):([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.+)$`

_Minimum_: `0`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnrichmentParameters`

The parameters required to set up enrichment on your pipe.

_Required_: No

_Type_: [PipeEnrichmentParameters](aws-properties-pipes-pipe-pipeenrichmentparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyIdentifier`

The identifier of the AWS KMS
customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt pipe data. The identifier can be the key
Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

To update a pipe that is using the default AWS owned key to use a customer managed key instead, or update a pipe that is using a customer managed key to use a
different customer managed key, specify a customer managed key identifier.

To update a pipe that is using a customer managed key to use the default AWS owned key, specify an empty string.

For more information, see [Managing keys](../../../kms/latest/developerguide/getting-started.md) in the _AWS Key Management Service_
_Developer Guide_.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The logging configuration settings for the pipe.

_Required_: No

_Type_: [PipeLogConfiguration](aws-properties-pipes-pipe-pipelogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the pipe.

_Required_: No

_Type_: String

_Pattern_: `^[\.\-_A-Za-z0-9]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the role that allows the pipe to send data to the target.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z0-9+=,.@\-_/]+$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The ARN of the source resource.

_Required_: Yes

_Type_: String

_Pattern_: `^smk://(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]):[0-9]{1,5}|arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-]+):([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.+)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceParameters`

The parameters required to set up a source for your pipe.

_Required_: No

_Type_: [PipeSourceParameters](aws-properties-pipes-pipe-pipesourceparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of key-value pairs to associate with the pipe.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The ARN of the target resource.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-]+):([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.+)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetParameters`

The parameters required to set up a target for your pipe.

For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](../../../eventbridge/latest/userguide/eb-pipes-event-target.md) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [PipeTargetParameters](aws-properties-pipes-pipe-pipetargetparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the pipe that was created by the
request.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the pipe.

`CreationTime`

The time the pipe was created.

`CurrentState`

The state the pipe is in.

`LastModifiedTime`

When the pipe was last updated, in [ISO-8601 format](https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).

`StateReason`

The reason the pipe is in its current state.

## Examples

- [Create a Pipe with an enrichment](#aws-resource-pipes-pipe--examples--Create_a_Pipe_with_an_enrichment)

- [Create a pipe with an event filter](#aws-resource-pipes-pipe--examples--Create_a_pipe_with_an_event_filter)

### Create a Pipe with an enrichment

Create a Pipe with an Amazon SQS source, an API Gateway enrichment, and a Step Functions state machine target.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "TestPipe": {
      "Type": "AWS::Pipes::Pipe",
      "Properties": {
        "Name": "PipeCfnExample",
        "RoleArn": "arn:aws:iam::123456789123:role/Pipe-Dev-All-Targets-Dummy-Execution-Role",
        "Source": "arn:aws:sqs:us-east-1:123456789123:pipeDemoSource",
        "Enrichment": "arn:aws:execute-api:us-east-1:123456789123:53eo2i89p9/*/POST/pets",
        "Target": "arn:aws:states:us-east-1:123456789123:stateMachine:PipeTargetStateMachine"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  TestPipe:
    Type: AWS::Pipes::Pipe
    Properties:
      Name:  PipeCfnExample
      RoleArn: arn:aws:iam::123456789123:role/Pipe-Dev-All-Targets-Dummy-Execution-Role
      Source: arn:aws:sqs:us-east-1:123456789123:pipeDemoSource
      Enrichment: arn:aws:execute-api:us-east-1:123456789123:53eo2i89p9/*/POST/pets
      Target: arn:aws:states:us-east-1:123456789123:stateMachine:PipeTargetStateMachine
```

### Create a pipe with an event filter

The following example:

- Provisions a DynamoDB table and associated data stream to act as the pipe source, and a Amazon SQS queue for the pipe target.

- Provisions an IAM execution role for the pipe that defines the necessary permissions to access both the
source and target.

- Creates a pipe that connects the DynamoDB stream source to the Amazon SQS queue target.

- Within the pipe, defines an event filter with an event pattern that selects events where `eventname` is
`INSERT` or `MODIFY`.

###### Note

Be aware that you will be billed for the AWS resources used if you create a stack
from this template.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",

 "Description" : "EventBridge Pipe template example. Provisions a pipe, along with a DynamoDB stream as the pipe source and an SQS queue as the pipe target. Also provisions an execution role that contains the necessary permissions to access both the source and target. Once provisioned, the pipe receives events from the DynamoDB data stream, applies a filter, and sends matching events on to an SQS Queue. You will be billed for the Amazon resources used if you create a stack from this template.",

  "Parameters" : {
    "SourceTableName" : {
      "Type" : "String",
      "Default" : "pipe-example-source",
      "Description" : "Specify the name of the table to provision as the pipe source, or accept the default."
    },
  "TargetQueueName" : {
    "Type" : "String",
    "Default" : "pipe-example-target",
    "Description" : "Specify the name of the queue to provision as the pipe target, or accept the default."
  },
    "PipeName" : {
      "Type" : "String",
      "Default" : "pipe-with-filtering-example",
      "Description" : "Specify the name of the table to provision as the pipe source, or accept the default."
    }
},
  "Resources": {
    "PipeSourceDynamoDBTable": {
      "Type": "AWS::DynamoDB::Table",
      "Properties": {
        "AttributeDefinitions": [{
            "AttributeName": "Album",
            "AttributeType": "S"
          },
          {
            "AttributeName": "Artist",
            "AttributeType": "S"
          }

        ],
        "KeySchema": [{
            "AttributeName": "Album",
            "KeyType": "HASH"

          },
          {
            "AttributeName": "Artist",
            "KeyType": "RANGE"
          }
        ],
        "ProvisionedThroughput": {
          "ReadCapacityUnits": 10,
          "WriteCapacityUnits": 10
        },
        "StreamSpecification": {
          "StreamViewType": "NEW_AND_OLD_IMAGES"
        },
        "TableName": { "Ref" : "SourceTableName" }
      }
    },
    "PipeTargetQueue": {
      "Type": "AWS::SQS::Queue",
      "Properties": {
        "QueueName": { "Ref" : "TargetQueueName" }
      }
    },
    "PipeTutorialPipeRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [{
            "Effect": "Allow",
            "Principal": {
              "Service": "pipes.amazonaws.com"
            },
            "Action": "sts:AssumeRole",
            "Condition": {
              "StringLike": {
                "aws:SourceArn": {
	"Fn::Join": [
		"",
                [
                  "arn:",
                  { "Ref": "AWS::Partition" },
                  ":pipes:",
                  { "Ref": "AWS::Region" },
                  ":",
                  { "Ref": "AWS::AccountId" },
                  ":pipe/",
                  { "Ref": "PipeName" }
               ]
	]
       },
                "aws:SourceAccount": { "Ref" : "AWS::AccountId" }
              }
            }
          }]
        },
        "Description" : "EventBridge Pipe template example. Execution role that grants the pipe the permissions necessary to send events to the specified pipe.",
        "Path": "/",
        "Policies": [{
            "PolicyName": "SourcePermissions",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [{
                "Effect": "Allow",
                "Action": [
                  "dynamodb:DescribeStream",
                  "dynamodb:GetRecords",
                  "dynamodb:GetShardIterator",
                  "dynamodb:ListStreams"
                ],
                "Resource": [
                  { "Fn::GetAtt" : [ "PipeSourceDynamoDBTable", "StreamArn" ] }
                ]
              }]
            }
          },
          {
            "PolicyName": "TargetPermissions",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [{
                "Effect": "Allow",
                "Action": [
                  "sqs:SendMessage"
                ],
                "Resource": [
                  { "Fn::GetAtt" : [ "PipeTargetQueue", "Arn" ] }
                ]
              }]
            }
          }
        ]
      }
  },
    "PipeWithFiltering": {
      "Type": "AWS::Pipes::Pipe",
      "Properties": {
        "Description" : "EventBridge Pipe template example. Pipe that receives events from a DynamoDB stream, applies a filter, and sends matching events on to an SQS Queue.",
        "Name": { "Ref" : "PipeName" },
        "RoleArn": {"Fn::GetAtt" : ["PipeTutorialPipeRole", "Arn"] },
        "Source": { "Fn::GetAtt" : [ "PipeSourceDynamoDBTable", "StreamArn" ] },
        "SourceParameters": {
          "DynamoDBStreamParameters" : {
            "StartingPosition" : "LATEST"
         },
        "FilterCriteria" : {
          "Filters" : [ {
            "Pattern" : "{ \"eventName\": [\"INSERT\", \"MODIFY\"] }"
         }]
        }
        },
        "Target": { "Fn::GetAtt" : [ "PipeTargetQueue", "Arn" ] }
      }
    }
  }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  EventBridge Pipe template example. Provisions a pipe, along with a DynamoDB
  stream as the pipe source and an SQS queue as the pipe target. Also provisions
  an execution role that contains the necessary permissions to access both the
  source and target. Once provisioned, the pipe receives events from the
  DynamoDB data stream, applies a filter, and sends matching events on to an SQS
  Queue. You will be billed for the Amazon resources used if you create a stack
  from this template.
Parameters:
  SourceTableName:
    Type: String
    Default: pipe-example-source
    Description: >-
      Specify the name of the table to provision as the pipe source, or accept
      the default.
  TargetQueueName:
    Type: String
    Default: pipe-example-target
    Description: >-
      Specify the name of the queue to provision as the pipe target, or accept
      the default.
  PipeName:
    Type: String
    Default: pipe-with-filtering-example
    Description: >-
      Specify the name of the table to provision as the pipe source, or accept
      the default.
Resources:
  PipeSourceDynamoDBTable:
    Type: 'AWS::DynamoDB::Table'
    Properties:
      AttributeDefinitions:
        - AttributeName: Album
          AttributeType: S
        - AttributeName: Artist
          AttributeType: S
      KeySchema:
        - AttributeName: Album
          KeyType: HASH
        - AttributeName: Artist
          KeyType: RANGE
      ProvisionedThroughput:
        ReadCapacityUnits: 10
        WriteCapacityUnits: 10
      StreamSpecification:
        StreamViewType: NEW_AND_OLD_IMAGES
      TableName:
        Ref: SourceTableName
  PipeTargetQueue:
    Type: 'AWS::SQS::Queue'
    Properties:
      QueueName:
        Ref: TargetQueueName
  PipeTutorialPipeRole:
    Type: 'AWS::IAM::Role'
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service: pipes.amazonaws.com
            Action: 'sts:AssumeRole'
            Condition:
              StringLike:
                'aws:SourceArn':
                  'Fn::Join':
                    - ''
                    - - 'arn:'
                      - Ref: 'AWS::Partition'
                      - ':pipes:'
                      - Ref: 'AWS::Region'
                      - ':'
                      - Ref: 'AWS::AccountId'
                      - ':pipe/'
                      - Ref: PipeName
                'aws:SourceAccount':
                  Ref: 'AWS::AccountId'
      Description: >-
        EventBridge Pipe template example. Execution role that grants the pipe
        the permissions necessary to send events to the specified pipe.
      Path: /
      Policies:
        - PolicyName: SourcePermissions
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - 'dynamodb:DescribeStream'
                  - 'dynamodb:GetRecords'
                  - 'dynamodb:GetShardIterator'
                  - 'dynamodb:ListStreams'
                Resource:
                  - 'Fn::GetAtt':
                      - PipeSourceDynamoDBTable
                      - StreamArn
        - PolicyName: TargetPermissions
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - 'sqs:SendMessage'
                Resource:
                  - 'Fn::GetAtt':
                      - PipeTargetQueue
                      - Arn
  PipeWithFiltering:
    Type: 'AWS::Pipes::Pipe'
    Properties:
      Description: >-
        EventBridge Pipe template example. Pipe that receives events from a
        DynamoDB stream, applies a filter, and sends matching events on to an
        SQS Queue.
      Name:
        Ref: PipeName
      RoleArn:
        'Fn::GetAtt':
          - PipeTutorialPipeRole
          - Arn
      Source:
        'Fn::GetAtt':
          - PipeSourceDynamoDBTable
          - StreamArn
      SourceParameters:
        DynamoDBStreamParameters:
          StartingPosition: LATEST
        FilterCriteria:
          Filters:
            - Pattern: '{ "eventName": ["INSERT", "MODIFY"] }'
      Target:
        'Fn::GetAtt':
          - PipeTargetQueue
          - Arn

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon EventBridge Pipes

AwsVpcConfiguration

All content copied from https://docs.aws.amazon.com/.
