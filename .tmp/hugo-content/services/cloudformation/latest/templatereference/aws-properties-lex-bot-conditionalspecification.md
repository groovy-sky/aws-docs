This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ConditionalSpecification

Provides a list of conditional branches. Branches are evaluated in
the order that they are entered in the list. The first branch with a
condition that evaluates to true is executed. The last branch in the
list is the default branch. The default branch should not have any condition
expression. The default branch is executed if no other branch has a
matching condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionalBranches" : [ ConditionalBranch, ... ],
  "DefaultBranch" : DefaultConditionalBranch,
  "IsActive" : Boolean
}

```

### YAML

```yaml

  ConditionalBranches:
    - ConditionalBranch
  DefaultBranch:
    DefaultConditionalBranch
  IsActive: Boolean

```

## Properties

`ConditionalBranches`

A list of conditional branches. A conditional branch is made up of a
condition, a response and a next step. The response and next step are
executed when the condition is true.

_Required_: Yes

_Type_: Array of [ConditionalBranch](aws-properties-lex-bot-conditionalbranch.md)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultBranch`

The conditional branch that should be followed when the conditions
for other branches are not satisfied. A conditional branch is made up
of a condition, a response and a next step.

_Required_: Yes

_Type_: [DefaultConditionalBranch](aws-properties-lex-bot-defaultconditionalbranch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsActive`

Determines whether a conditional branch is active. When
`IsActive` is false, the conditions are not
evaluated.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalBranch

ConversationLogSettings

All content copied from https://docs.aws.amazon.com/.
