This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::VpcAttachment VpcOptions

Describes the VPC options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplianceModeSupport" : Boolean,
  "DnsSupport" : Boolean,
  "Ipv6Support" : Boolean,
  "SecurityGroupReferencingSupport" : Boolean
}

```

### YAML

```yaml

  ApplianceModeSupport: Boolean
  DnsSupport: Boolean
  Ipv6Support: Boolean
  SecurityGroupReferencingSupport: Boolean

```

## Properties

`ApplianceModeSupport`

Indicates whether appliance mode is supported. If enabled, traffic flow between a source and destination use the same Availability Zone for the VPC attachment for the lifetime of that flow. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsSupport`

Indicates whether DNS is supported.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Support`

Indicates whether IPv6 is supported.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupReferencingSupport`

Indicates whether security group referencing is enabled for this VPC attachment. The default is `true`. However, at the core network policy-level the default is set to `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
