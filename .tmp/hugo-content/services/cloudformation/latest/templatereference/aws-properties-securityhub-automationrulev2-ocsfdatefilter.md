This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 OcsfDateFilter

Enables filtering of security findings based on date and timestamp fields in OCSF.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "Filter" : DateFilter
}

```

### YAML

```yaml

  FieldName: String
  Filter:
    DateFilter

```

## Properties

`FieldName`

The name of the field.

_Required_: Yes

_Type_: String

_Allowed values_: `finding_info.created_time_dt | finding_info.first_seen_time_dt | finding_info.last_seen_time_dt | finding_info.modified_time_dt`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Enables filtering of security findings based on date and timestamp fields in OCSF.

_Required_: Yes

_Type_: [DateFilter](aws-properties-securityhub-automationrulev2-datefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OcsfBooleanFilter

OcsfFindingFilters

All content copied from https://docs.aws.amazon.com/.
