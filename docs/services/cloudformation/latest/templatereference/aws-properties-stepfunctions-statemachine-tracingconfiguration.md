This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachine TracingConfiguration

Selects whether or not the state machine's AWS X-Ray tracing is enabled. To configure
your state machine to send trace data to X-Ray, set `Enabled` to
`true`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean

```

## Properties

`Enabled`

When set to `true`, X-Ray tracing is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagsEntry

AWS::StepFunctions::StateMachineAlias

All content copied from https://docs.aws.amazon.com/.
