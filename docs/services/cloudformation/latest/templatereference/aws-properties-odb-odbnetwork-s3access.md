This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::OdbNetwork S3Access

The configuration for Amazon S3 access from the ODB network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "Ipv4Addresses" : [ String, ... ],
  "S3PolicyDocument" : String,
  "Status" : String
}

```

### YAML

```yaml

  DomainName: String
  Ipv4Addresses:
    - String
  S3PolicyDocument: String
  Status: String

```

## Properties

`DomainName`

The domain name for the Amazon S3 access.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4Addresses`

The IPv4 addresses for the Amazon S3 access.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3PolicyDocument`

The endpoint policy for the Amazon S3 access.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the Amazon S3 access.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | ENABLING | DISABLED | DISABLING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedServices

ServiceNetworkEndpoint

All content copied from https://docs.aws.amazon.com/.
