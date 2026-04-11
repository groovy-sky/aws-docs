This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 OcsfNumberFilter

Enables filtering of security findings based on numerical field values in OCSF.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "Filter" : NumberFilter
}

```

### YAML

```yaml

  FieldName: String
  Filter:
    NumberFilter

```

## Properties

`FieldName`

The name of the field.

_Required_: Yes

_Type_: String

_Allowed values_: `activity_id | compliance.status_id | confidence_score | finding_info.related_events_count | vendor_attributes.severity_id`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Enables filtering of security findings based on numerical field values in OCSF.

_Required_: Yes

_Type_: [NumberFilter](aws-properties-securityhub-automationrulev2-numberfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OcsfMapFilter

OcsfStringFilter

All content copied from https://docs.aws.amazon.com/.
