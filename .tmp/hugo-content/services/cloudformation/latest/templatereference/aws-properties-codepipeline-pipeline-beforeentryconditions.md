This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline BeforeEntryConditions

The conditions for making checks for entry to a stage. For more information about
conditions, see [Stage conditions](../../../codepipeline/latest/userguide/stage-conditions.md)
and [How do\
stage conditions work?](../../../codepipeline/latest/userguide/concepts-how-it-works-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditions" : [ Condition, ... ]
}

```

### YAML

```yaml

  Conditions:
    - Condition

```

## Properties

`Conditions`

The conditions that are configured as entry conditions.

_Required_: No

_Type_: Array of [Condition](aws-properties-codepipeline-pipeline-condition.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArtifactStoreMap

BlockerDeclaration

All content copied from https://docs.aws.amazon.com/.
