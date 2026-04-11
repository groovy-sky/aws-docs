This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 AutomationRulesFindingFieldsUpdateV2

Allows you to define the structure for modifying specific fields in security findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String,
  "SeverityId" : Integer,
  "StatusId" : Integer
}

```

### YAML

```yaml

  Comment: String
  SeverityId: Integer
  StatusId: Integer

```

## Properties

`Comment`

Notes or contextual information for findings that are modified by the automation rule.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeverityId`

The severity level to be assigned to findings that match the automation rule criteria.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusId`

The status to be applied to findings that match automation rule criteria.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutomationRulesActionV2

BooleanFilter

All content copied from https://docs.aws.amazon.com/.
