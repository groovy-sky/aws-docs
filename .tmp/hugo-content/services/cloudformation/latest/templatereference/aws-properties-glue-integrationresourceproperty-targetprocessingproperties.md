This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::IntegrationResourceProperty TargetProcessingProperties

The structure used to define the resource properties associated with the integration target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionName" : String,
  "EventBusArn" : String,
  "KmsArn" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  ConnectionName: String
  EventBusArn: String
  KmsArn: String
  RoleArn: String

```

## Properties

`ConnectionName`

The AWS Glue network connection to configure the AWS Glue job running in the customer VPC.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBusArn`

The ARN of an Eventbridge event bus to receive the integration status notification.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsArn`

The ARN of the KMS key used for encryption.

_Required_: No

_Type_: String

_Pattern_: `arn:aws:kms:.*:[0-9]+:.*`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role to access the AWS Glue database.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:iam:.*:[0-9]+:.*`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Glue::Job

All content copied from https://docs.aws.amazon.com/.
