This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule StepFunctionsAction

Starts execution of a Step Functions state machine.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionNamePrefix" : String,
  "RoleArn" : String,
  "StateMachineName" : String
}

```

### YAML

```yaml

  ExecutionNamePrefix: String
  RoleArn: String
  StateMachineName: String

```

## Properties

`ExecutionNamePrefix`

(Optional) A name will be given to the state machine execution consisting of this
prefix followed by a UUID. Step Functions automatically creates a unique name for each state
machine execution if one is not provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants IoT permission to start execution of a state machine
("Action":"states:StartExecution").

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StateMachineName`

The name of the Step Functions state machine whose execution will be started.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [StepFunctionsAction](../../../../reference/iot/latest/apireference/api-stepfunctionsaction.md) in the _AWS IoT API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SqsAction

Tag

All content copied from https://docs.aws.amazon.com/.
