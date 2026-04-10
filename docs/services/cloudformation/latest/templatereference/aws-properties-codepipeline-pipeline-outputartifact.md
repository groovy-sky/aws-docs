This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline OutputArtifact

Represents information about the output of an action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Files" : [ String, ... ],
  "Name" : String
}

```

### YAML

```yaml

  Files:
    - String
  Name: String

```

## Properties

`Files`

The files that you want to associate with the output artifact that will be exported
from the compute action.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the output of an artifact, such as "My App".

The output artifact name must exactly match the input artifact declared for a
downstream action. However, the downstream action's input artifact does not have to be the
next action in strict sequence from the action that provided the output artifact. Actions in
parallel can declare different output artifacts, which are in turn consumed by different
following actions.

Output artifact names must be unique within a pipeline.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\-]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputArtifact

PipelineTriggerDeclaration

All content copied from https://docs.aws.amazon.com/.
