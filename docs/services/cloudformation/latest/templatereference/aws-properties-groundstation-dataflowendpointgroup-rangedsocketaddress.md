This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup RangedSocketAddress

A socket address with a port range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "PortRange" : IntegerRange
}

```

### YAML

```yaml

  Name: String
  PortRange:
    IntegerRange

```

## Properties

`Name`

IPv4 socket address.

_Required_: No

_Type_: String

_Pattern_: `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`

_Minimum_: `7`

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortRange`

Port range of a socket address.

_Required_: No

_Type_: [IntegerRange](aws-properties-groundstation-dataflowendpointgroup-integerrange.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RangedConnectionDetails

SecurityDetails

All content copied from https://docs.aws.amazon.com/.
