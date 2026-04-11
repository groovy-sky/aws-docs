This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link OpenRtbAttributeModuleParameters

Describes the parameters of an open RTB attribute module.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : Action,
  "FilterConfiguration" : [ Filter, ... ],
  "FilterType" : String,
  "HoldbackPercentage" : Number
}

```

### YAML

```yaml

  Action:
    Action
  FilterConfiguration:
    - Filter
  FilterType: String
  HoldbackPercentage: Number

```

## Properties

`Action`

Describes a bid action.

_Required_: Yes

_Type_: [Action](aws-properties-rtbfabric-link-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterConfiguration`

Describes the configuration of a filter.

_Required_: Yes

_Type_: Array of [Filter](aws-properties-rtbfabric-link-filter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterType`

The filter type.

_Required_: Yes

_Type_: String

_Allowed values_: `INCLUDE | EXCLUDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HoldbackPercentage`

The hold back percentage.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NoBidModuleParameters

ResponderErrorMaskingForHttpCode

All content copied from https://docs.aws.amazon.com/.
