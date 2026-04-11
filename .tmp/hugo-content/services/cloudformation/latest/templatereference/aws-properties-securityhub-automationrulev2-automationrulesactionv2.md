This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 AutomationRulesActionV2

Allows you to configure automated responses.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExternalIntegrationConfiguration" : ExternalIntegrationConfiguration,
  "FindingFieldsUpdate" : AutomationRulesFindingFieldsUpdateV2,
  "Type" : String
}

```

### YAML

```yaml

  ExternalIntegrationConfiguration:
    ExternalIntegrationConfiguration
  FindingFieldsUpdate:
    AutomationRulesFindingFieldsUpdateV2
  Type: String

```

## Properties

`ExternalIntegrationConfiguration`

The settings for integrating automation rule actions with external systems or service.

_Required_: No

_Type_: [ExternalIntegrationConfiguration](aws-properties-securityhub-automationrulev2-externalintegrationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingFieldsUpdate`

Specifies that the automation rule action is an update to a finding field.

_Required_: No

_Type_: [AutomationRulesFindingFieldsUpdateV2](aws-properties-securityhub-automationrulev2-automationrulesfindingfieldsupdatev2.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of action that Security Hub CSPM takes when a finding matches the defined criteria of a rule.

_Required_: Yes

_Type_: String

_Allowed values_: `FINDING_FIELDS_UPDATE | EXTERNAL_INTEGRATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityHub::AutomationRuleV2

AutomationRulesFindingFieldsUpdateV2

All content copied from https://docs.aws.amazon.com/.
