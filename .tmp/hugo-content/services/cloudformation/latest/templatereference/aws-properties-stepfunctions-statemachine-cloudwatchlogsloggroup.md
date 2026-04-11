This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachine CloudWatchLogsLogGroup

Defines a CloudWatch log group.

###### Note

For more information see [Standard Versus\
Express Workflows](../../../step-functions/latest/dg/concepts-standard-vs-express.md) in the AWS Step Functions Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupArn" : String
}

```

### YAML

```yaml

  LogGroupArn: String

```

## Properties

`LogGroupArn`

The ARN of the the CloudWatch log group to which you want your logs emitted to. The ARN
must end with `:*`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::StepFunctions::StateMachine

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
