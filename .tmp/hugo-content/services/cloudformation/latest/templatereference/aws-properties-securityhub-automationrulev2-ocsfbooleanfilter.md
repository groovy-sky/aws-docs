This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 OcsfBooleanFilter

Enables filtering of security findings based on boolean field values in OCSF.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "Filter" : BooleanFilter
}

```

### YAML

```yaml

  FieldName: String
  Filter:
    BooleanFilter

```

## Properties

`FieldName`

The name of the field.

_Required_: Yes

_Type_: String

_Allowed values_: `compliance.assessments.meets_criteria | vulnerabilities.is_exploit_available | vulnerabilities.is_fix_available`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Enables filtering of security findings based on boolean field values in OCSF.

_Required_: Yes

_Type_: [BooleanFilter](aws-properties-securityhub-automationrulev2-booleanfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumberFilter

OcsfDateFilter

All content copied from https://docs.aws.amazon.com/.
