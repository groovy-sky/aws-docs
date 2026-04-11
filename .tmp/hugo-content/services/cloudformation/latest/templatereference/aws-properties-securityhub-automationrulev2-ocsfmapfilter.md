This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 OcsfMapFilter

Enables filtering of security findings based on map field values in OCSF.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "Filter" : MapFilter
}

```

### YAML

```yaml

  FieldName: String
  Filter:
    MapFilter

```

## Properties

`FieldName`

The name of the field.

_Required_: Yes

_Type_: String

_Allowed values_: `resources.tags`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Enables filtering of security findings based on map field values in OCSF.

_Required_: Yes

_Type_: [MapFilter](aws-properties-securityhub-automationrulev2-mapfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OcsfFindingFilters

OcsfNumberFilter

All content copied from https://docs.aws.amazon.com/.
