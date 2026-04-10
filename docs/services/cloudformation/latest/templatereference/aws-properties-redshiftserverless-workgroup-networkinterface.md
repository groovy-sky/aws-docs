This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Workgroup NetworkInterface

Contains information about a network interface
in an Amazon Redshift Serverless managed VPC endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String,
  "NetworkInterfaceId" : String,
  "PrivateIpAddress" : String,
  "SubnetId" : String
}

```

### YAML

```yaml

  AvailabilityZone: String
  NetworkInterfaceId: String
  PrivateIpAddress: String
  SubnetId: String

```

## Properties

`AvailabilityZone`

The availability Zone.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceId`

The unique identifier of the network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateIpAddress`

The IPv4 address of the network interface within the subnet.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The unique identifier of the subnet.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Endpoint

PerformanceTarget

All content copied from https://docs.aws.amazon.com/.
