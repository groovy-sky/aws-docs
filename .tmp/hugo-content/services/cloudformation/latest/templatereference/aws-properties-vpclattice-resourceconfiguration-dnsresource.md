This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ResourceConfiguration DnsResource

The domain name of the resource configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "IpAddressType" : String
}

```

### YAML

```yaml

  DomainName: String
  IpAddressType: String

```

## Properties

`DomainName`

The domain name of the resource configuration.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address type for the resource configuration. Dualstack is not currently supported.

_Required_: Yes

_Type_: String

_Allowed values_: `IPV4 | IPV6 | DUALSTACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VpcLattice::ResourceConfiguration

ResourceConfigurationDefinition

All content copied from https://docs.aws.amazon.com/.
