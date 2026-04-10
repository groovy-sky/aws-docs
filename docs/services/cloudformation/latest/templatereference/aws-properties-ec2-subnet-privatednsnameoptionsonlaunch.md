This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Subnet PrivateDnsNameOptionsOnLaunch

Specifies the options for instance hostnames.

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

Indicates whether to respond to DNS queries for instance hostname with DNS AAAA
records.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableResourceNameDnsARecord`

Indicates whether to respond to DNS queries for instance hostnames with DNS A
records.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostnameType`

The type of hostname for EC2 instances. For IPv4 only subnets, an instance DNS name
must be based on the instance IPv4 address. For IPv6 only subnets, an instance DNS name
must be based on the instance ID. For dual-stack subnets, you can specify whether DNS
names use the instance IPv4 address or the instance ID.

_Required_: No

_Type_: String

_Allowed values_: `ip-name | resource-name`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlockPublicAccessStates

Tag

All content copied from https://docs.aws.amazon.com/.
