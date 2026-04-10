This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline Condition

The condition for the stage. A condition is made up of the rules and the result for
the condition. For more information about conditions, see [Stage conditions](../../../codepipeline/latest/userguide/stage-conditions.md)
and [How do\
stage conditions work?](../../../codepipeline/latest/userguide/concepts-how-it-works-conditions.md).. For more information about rules, see the [AWS CodePipeline rule reference](../../../codepipeline/latest/userguide/rule-reference.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Result" : String,
  "Rules" : [ RuleDeclaration, ... ]
}

```

### YAML

```yaml

  Result: String
  Rules:
    - RuleDeclaration

```

## Properties

`Result`

The action to be done when the condition is met. For example, rolling back an
execution for a failure condition.

_Required_: No

_Type_: String

_Allowed values_: `ROLLBACK | FAIL | RETRY | SKIP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

The rules that make up the condition.

_Required_: No

_Type_: Array of [RuleDeclaration](aws-properties-codepipeline-pipeline-ruledeclaration.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlockerDeclaration

EncryptionKey

All content copied from https://docs.aws.amazon.com/.
