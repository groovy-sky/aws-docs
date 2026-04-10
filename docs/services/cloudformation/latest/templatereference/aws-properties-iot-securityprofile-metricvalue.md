This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile MetricValue

The value to be compared with the `metric`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidrs" : [ String, ... ],
  "Count" : String,
  "Number" : Number,
  "Numbers" : [ Number, ... ],
  "Ports" : [ Integer, ... ],
  "Strings" : [ String, ... ]
}

```

### YAML

```yaml

  Cidrs:
    - String
  Count: String
  Number:
    Number
  Numbers:
    - Number
  Ports:
    - Integer
  Strings:
    - String

```

## Properties

`Cidrs`

If the `comparisonOperator` calls for a set of CIDRs, use this
to specify that set to be compared with the `metric`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Count`

If the `comparisonOperator` calls for a numeric value, use this
to specify that numeric value to be compared with the `metric`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Number`

The numeric values of a metric.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Numbers`

The numeric value of a metric.

_Required_: No

_Type_: Array of Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ports`

If the `comparisonOperator` calls for a set of ports, use this
to specify that set to be compared with the `metric`.

_Required_: No

_Type_: Array of Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Strings`

The string values of a metric.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricToRetain

StatisticalThreshold

All content copied from https://docs.aws.amazon.com/.
