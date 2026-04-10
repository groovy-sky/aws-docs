This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsPath PathFilter

Describes a set of filters for a path analysis. Use path filters to scope the analysis when
there can be multiple resulting paths.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationAddress" : String,
  "DestinationPortRange" : FilterPortRange,
  "SourceAddress" : String,
  "SourcePortRange" : FilterPortRange
}

```

### YAML

```yaml

  DestinationAddress: String
  DestinationPortRange:
    FilterPortRange
  SourceAddress: String
  SourcePortRange:
    FilterPortRange

```

## Properties

`DestinationAddress`

The destination IPv4 address.

_Required_: No

_Type_: String

_Pattern_: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPortRange`

The destination port range.

_Required_: No

_Type_: [FilterPortRange](aws-properties-ec2-networkinsightspath-filterportrange.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceAddress`

The source IPv4 address.

_Required_: No

_Type_: String

_Pattern_: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePortRange`

The source port range.

_Required_: No

_Type_: [FilterPortRange](aws-properties-ec2-networkinsightspath-filterportrange.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterPortRange

Tag

All content copied from https://docs.aws.amazon.com/.
