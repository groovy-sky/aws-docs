This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule DynamoDBv2Action

Describes an action to write to a DynamoDB table.

This DynamoDB action writes each attribute in the message payload into it's own
column in the DynamoDB table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PutItem" : PutItemInput,
  "RoleArn" : String
}

```

### YAML

```yaml

  PutItem:
    PutItemInput
  RoleArn: String

```

## Properties

`PutItem`

Specifies the DynamoDB table to which the message data will be written. For
example:

`{ "dynamoDBv2": { "roleArn": "aws:iam:12341251:my-role" "putItem": { "tableName":
            "my-table" } } }`

Each attribute in the message payload will be written to a separate column in the
DynamoDB database.

_Required_: No

_Type_: [PutItemInput](aws-properties-iot-topicrule-putiteminput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants access to the DynamoDB table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AWS SDK for C++](https://sdk.amazonaws.com/cpp/api/LATEST/class_aws_1_1_io_t_1_1_model_1_1_dynamo_d_bv2_action.html).

- [AWS SDK for Go](../../../../reference/sdk-for-go/api/service/iot.md#DynamoDBv2Action).

- [AWS SDK for Java](../../../../reference/awsjavasdk/latest/javadoc/com/amazonaws/services/iot/model/dynamodbv2action.md).

- [AWS SDK for Ruby V2](../../../../reference/sdkforruby/api/aws/iot/types/dynamodbv2action.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDBAction

ElasticsearchAction

All content copied from https://docs.aws.amazon.com/.
