This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetStateMachineParameters

The parameters for using a Step Functions state machine as a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InvocationType" : String
}

```

### YAML

```yaml

  InvocationType: String

```

## Properties

`InvocationType`

Specify whether to invoke the Step Functions state machine synchronously or
asynchronously.

- `REQUEST_RESPONSE` (default) - Invoke synchronously. For more
information, see [StartSyncExecution](../../../../reference/step-functions/latest/apireference/api-startsyncexecution.md) in the _AWS Step Functions API_
_Reference_.

###### Note

`REQUEST_RESPONSE` is not supported for `STANDARD` state
machine workflows.

- `FIRE_AND_FORGET` \- Invoke asynchronously. For more information, see
[StartExecution](../../../../reference/step-functions/latest/apireference/api-startexecution.md) in the _AWS Step Functions API_
_Reference_.

For more information, see [Invocation\
types](../../../eventbridge/latest/userguide/eb-pipes.md#pipes-invocation) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `REQUEST_RESPONSE | FIRE_AND_FORGET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetSqsQueueParameters

PipeTargetTimestreamParameters

All content copied from https://docs.aws.amazon.com/.
