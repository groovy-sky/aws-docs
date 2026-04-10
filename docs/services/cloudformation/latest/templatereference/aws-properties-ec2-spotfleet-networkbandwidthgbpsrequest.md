This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet NetworkBandwidthGbpsRequest

The minimum and maximum amount of baseline network bandwidth, in gigabits per second
(Gbps). For more information, see [Amazon EC2 instance network bandwidth](../../../ec2/latest/userguide/ec2-instance-network-bandwidth.md) in the _Amazon EC2 User Guide_.

Default: No minimum or maximum limits

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

The maximum amount of network bandwidth, in Gbps. To specify no maximum limit, omit this
parameter.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Min`

The minimum amount of network bandwidth, in Gbps. To specify no minimum limit, omit this
parameter.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemoryMiBRequest

NetworkInterfaceCountRequest

All content copied from https://docs.aws.amazon.com/.
