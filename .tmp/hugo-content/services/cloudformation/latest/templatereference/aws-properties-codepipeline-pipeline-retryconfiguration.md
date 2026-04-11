This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline RetryConfiguration

The retry configuration specifies automatic retry for a failed stage, along with the
configured retry mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetryMode" : String
}

```

### YAML

```yaml

  RetryMode: String

```

## Properties

`RetryMode`

The method that you want to configure for automatic stage retry on stage failure. You
can specify to retry only failed action in the stage or all actions in the stage.

_Required_: No

_Type_: String

_Allowed values_: `ALL_ACTIONS | FAILED_ACTIONS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipelineTriggerDeclaration

RuleDeclaration

All content copied from https://docs.aws.amazon.com/.
