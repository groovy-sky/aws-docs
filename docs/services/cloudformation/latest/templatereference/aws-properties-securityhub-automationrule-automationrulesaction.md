This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRule AutomationRulesAction

One or more actions that AWS Security Hub CSPM takes when a finding matches the defined criteria
of a rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FindingFieldsUpdate" : AutomationRulesFindingFieldsUpdate,
  "Type" : String
}

```

### YAML

```yaml

  FindingFieldsUpdate:
    AutomationRulesFindingFieldsUpdate
  Type: String

```

## Properties

`FindingFieldsUpdate`

Specifies that the automation rule action is an update to a finding field.

_Required_: Yes

_Type_: [AutomationRulesFindingFieldsUpdate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-automationrulesfindingfieldsupdate.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of action that Security Hub CSPM takes when a finding matches the defined criteria of a rule.

_Required_: Yes

_Type_: String

_Allowed values_: `FINDING_FIELDS_UPDATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecurityHub::AutomationRule

AutomationRulesFindingFieldsUpdate
