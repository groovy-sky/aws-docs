This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 CompositeFilter

Enables the creation of filtering criteria for security findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanFilters" : [ OcsfBooleanFilter, ... ],
  "DateFilters" : [ OcsfDateFilter, ... ],
  "MapFilters" : [ OcsfMapFilter, ... ],
  "NumberFilters" : [ OcsfNumberFilter, ... ],
  "Operator" : String,
  "StringFilters" : [ OcsfStringFilter, ... ]
}

```

### YAML

```yaml

  BooleanFilters:
    - OcsfBooleanFilter
  DateFilters:
    - OcsfDateFilter
  MapFilters:
    - OcsfMapFilter
  NumberFilters:
    - OcsfNumberFilter
  Operator: String
  StringFilters:
    - OcsfStringFilter

```

## Properties

`BooleanFilters`

Enables filtering based on boolean field values.

_Required_: No

_Type_: Array of [OcsfBooleanFilter](aws-properties-securityhub-automationrulev2-ocsfbooleanfilter.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateFilters`

Enables filtering based on date and timestamp fields.

_Required_: No

_Type_: Array of [OcsfDateFilter](aws-properties-securityhub-automationrulev2-ocsfdatefilter.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapFilters`

Enables the creation of filtering criteria for security findings.

_Required_: No

_Type_: Array of [OcsfMapFilter](aws-properties-securityhub-automationrulev2-ocsfmapfilter.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberFilters`

Enables filtering based on numerical field values.

_Required_: No

_Type_: Array of [OcsfNumberFilter](aws-properties-securityhub-automationrulev2-ocsfnumberfilter.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The logical operator used to combine multiple filter conditions.

_Required_: No

_Type_: String

_Allowed values_: `AND | OR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringFilters`

Enables filtering based on string field values.

_Required_: No

_Type_: Array of [OcsfStringFilter](aws-properties-securityhub-automationrulev2-ocsfstringfilter.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BooleanFilter

Criteria

All content copied from https://docs.aws.amazon.com/.
