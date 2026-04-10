This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::Filter

Details about a filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::InspectorV2::Filter",
  "Properties" : {
      "Description" : String,
      "FilterAction" : String,
      "FilterCriteria" : FilterCriteria,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::InspectorV2::Filter
Properties:
  Description: String
  FilterAction: String
  FilterCriteria:
    FilterCriteria
  Name: String
  Tags:
    Key: Value

```

## Properties

`Description`

A description of the filter.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterAction`

The action that is to be applied to the findings that match the filter.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | SUPPRESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterCriteria`

Details on the filter criteria associated with this filter.

_Required_: Yes

_Type_: [FilterCriteria](aws-properties-inspectorv2-filter-filtercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the filter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags attached to the filter.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{2,127}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the filter. For example:

`arn:aws:inspector2:us-east-1:012345678901:owner/012345678901/filter/c1c0fe5d28e39baa`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Number (ARN) associated with this filter.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScopeSettings

DateFilter

All content copied from https://docs.aws.amazon.com/.
