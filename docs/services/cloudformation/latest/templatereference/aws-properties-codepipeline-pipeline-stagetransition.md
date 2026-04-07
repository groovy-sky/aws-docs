This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline StageTransition

The name of the pipeline in which you want to disable the flow of artifacts from
one stage to another.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Reason" : String,
  "StageName" : String
}

```

### YAML

```yaml

  Reason: String
  StageName: String

```

## Properties

`Reason`

The reason given to the user that a stage is disabled, such as waiting for manual
approval or manual tests. This message is displayed in the pipeline console
UI.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9!@ \(\)\.\*\?\-]+`

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StageName`

The name of the stage where you want to disable the inbound or outbound transition
of artifacts.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StageDeclaration

SuccessConditions
