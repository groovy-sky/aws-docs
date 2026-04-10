This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet RowLevelPermissionTagConfiguration

The element you can use to define tags for row-level security.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String,
  "TagRuleConfigurations" : [ [ , ... ], ... ],
  "TagRules" : [ RowLevelPermissionTagRule, ... ]
}

```

### YAML

```yaml

  Status: String
  TagRuleConfigurations:
    -
    -
  TagRules:
    - RowLevelPermissionTagRule

```

## Properties

`Status`

The status of row-level security tags. If enabled, the status is `ENABLED`. If disabled, the status is `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagRuleConfigurations`

The configuration of tags on a dataset to set row-level security.

_Required_: No

_Type_: Array of Array

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagRules`

A set of rules associated with row-level security, such as the tag names and columns that they are assigned to.

_Required_: Yes

_Type_: Array of [RowLevelPermissionTagRule](aws-properties-quicksight-dataset-rowlevelpermissiontagrule.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RowLevelPermissionDataSet

RowLevelPermissionTagRule

All content copied from https://docs.aws.amazon.com/.
