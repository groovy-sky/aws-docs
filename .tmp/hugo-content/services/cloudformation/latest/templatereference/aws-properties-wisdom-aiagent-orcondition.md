This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent OrCondition

A list of conditions which would be applied together with an `OR` condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AndConditions" : [ TagCondition, ... ],
  "TagCondition" : TagCondition
}

```

### YAML

```yaml

  AndConditions:
    - TagCondition
  TagCondition:
    TagCondition

```

## Properties

`AndConditions`

A list of conditions which would be applied together with an `AND` condition.

_Required_: No

_Type_: Array of [TagCondition](aws-properties-wisdom-aiagent-tagcondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagCondition`

A leaf node condition which can be used to specify a tag condition.

_Required_: No

_Type_: [TagCondition](aws-properties-wisdom-aiagent-tagcondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrchestrationAIAgentConfiguration

SelfServiceAIAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
