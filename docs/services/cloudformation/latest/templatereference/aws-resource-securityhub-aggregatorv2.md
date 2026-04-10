This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AggregatorV2

Enables aggregation across AWS Regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::AggregatorV2",
  "Properties" : {
      "LinkedRegions" : [ String, ... ],
      "RegionLinkingMode" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::AggregatorV2
Properties:
  LinkedRegions:
    - String
  RegionLinkingMode: String
  Tags:
    Key: Value

```

## Properties

`LinkedRegions`

The list of Regions that are linked to the aggregation Region.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionLinkingMode`

Determines how Regions are linked to an Aggregator V2.

_Required_: Yes

_Type_: String

_Allowed values_: `SPECIFIED_REGIONS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to be applied to the AggregatorV2.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AggregatorV2ARN` for the `AggregatorV2` created: `arn:aws:securityhub:region:123456789012:aggregatorv2/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AggregationRegion`

The AWS Region where data is aggregated.

`AggregatorV2Arn`

The ARN of the AggregatorV2.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Security Hub CSPM

AWS::SecurityHub::AutomationRule

All content copied from https://docs.aws.amazon.com/.
