This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance PrivateDnsNameOptions

The type of hostnames to assign to instances in the subnet at launch. For IPv4 only subnets, an
instance DNS name must be based on the instance IPv4 address. For IPv6 only subnets, an instance
DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS
names use the instance IPv4 address or the instance ID. For more information, see [Amazon EC2 instance hostname types](../../../ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon Elastic Compute Cloud User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableResourceNameDnsAAAARecord" : Boolean,
  "EnableResourceNameDnsARecord" : Boolean,
  "HostnameType" : String
}

```

### YAML

```yaml

  EnableResourceNameDnsAAAARecord: Boolean
  EnableResourceNameDnsARecord: Boolean
  HostnameType: String

```

## Properties

`EnableResourceNameDnsAAAARecord`

Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records. For more information, see [Amazon EC2 instance hostname types](../../../ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon Elastic Compute Cloud User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EnableResourceNameDnsARecord`

Indicates whether to respond to DNS queries for instance hostnames with DNS A records. For more information, see [Amazon EC2 instance hostname types](../../../ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon Elastic Compute Cloud User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`HostnameType`

The type of hostnames to assign to instances in the subnet at launch. For IPv4 only subnets, an
instance DNS name must be based on the instance IPv4 address. For IPv6 only subnets, an instance
DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS
names use the instance IPv4 address or the instance ID. For more information, see [Amazon EC2 instance hostname types](../../../ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon Elastic Compute Cloud User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `ip-name | resource-name`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkInterface

PrivateIpAddressSpecification

All content copied from https://docs.aws.amazon.com/.
