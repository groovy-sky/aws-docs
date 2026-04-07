This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::OdbNetwork StsAccess

Configuration for AWS Security Token Service (STS) access from the ODB network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "Ipv4Addresses" : [ String, ... ],
  "Status" : String,
  "StsPolicyDocument" : String
}

```

### YAML

```yaml

  DomainName: String
  Ipv4Addresses:
    - String
  Status: String
  StsPolicyDocument: String

```

## Properties

`DomainName`

The domain name for AWS Security Token Service (STS) access configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4Addresses`

The IPv4 addresses allowed for AWS Security Token Service (STS) access.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status of the AWS Security Token Service (STS) access configuration.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | ENABLING | DISABLED | DISABLING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StsPolicyDocument`

The AWS Security Token Service (STS) policy document that defines permissions for token service usage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceNetworkEndpoint

Tag
