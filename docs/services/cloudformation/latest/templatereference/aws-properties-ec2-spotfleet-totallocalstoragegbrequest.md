This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet TotalLocalStorageGBRequest

The minimum and maximum amount of total local storage, in GB.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Number,
  "Min" : Number
}

```

### YAML

```yaml

  Max: Number
  Min: Number

```

## Properties

`Max`

The maximum amount of total local storage, in GB. To specify no maximum limit, omit this
parameter.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Min`

The minimum amount of total local storage, in GB. To specify no minimum limit, omit this
parameter.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetGroupsConfig

VCpuCountRangeRequest

All content copied from https://docs.aws.amazon.com/.
