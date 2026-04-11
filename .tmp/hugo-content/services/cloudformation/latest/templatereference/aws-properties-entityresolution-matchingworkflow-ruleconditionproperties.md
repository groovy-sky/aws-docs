This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow RuleConditionProperties

The properties of a rule condition that provides the ability to use more complex
syntax.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ RuleCondition, ... ]
}

```

### YAML

```yaml

  Rules:
    - RuleCondition

```

## Properties

`Rules`

A list of rule objects, each of which have fields `ruleName` and `condition`.

_Required_: Yes

_Type_: Array of [RuleCondition](aws-properties-entityresolution-matchingworkflow-rulecondition.md)

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleCondition

Tag

All content copied from https://docs.aws.amazon.com/.
