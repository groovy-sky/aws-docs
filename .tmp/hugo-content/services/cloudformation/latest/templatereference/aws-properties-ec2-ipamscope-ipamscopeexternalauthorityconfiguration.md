This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMScope IpamScopeExternalAuthorityConfiguration

The configuration that links an Amazon VPC IPAM scope to an external authority system. It specifies the type of external system and the external resource identifier that identifies your account or instance in that system.

In IPAM, an external authority is a third-party IP address management system that provides CIDR blocks when you provision address space for top-level IPAM pools. This allows you to use your existing IP management system to control which address ranges are allocated to AWS while using Amazon VPC IPAM to manage subnets within those ranges.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExternalResourceIdentifier" : String,
  "IpamScopeExternalAuthorityType" : String
}

```

### YAML

```yaml

  ExternalResourceIdentifier: String
  IpamScopeExternalAuthorityType: String

```

## Properties

`ExternalResourceIdentifier`

The identifier for the external resource managing this scope. For Infoblox integrations, this is the Infoblox resource identifier in the format `<version>.identity.account.<entity_realm>.<entity_id>`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamScopeExternalAuthorityType`

The type of external authority managing this scope. Currently supports `Infoblox` for integration with Infoblox Universal DDI.

_Required_: Yes

_Type_: String

_Allowed values_: `infoblox`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::IPAMScope

Tag

All content copied from https://docs.aws.amazon.com/.
