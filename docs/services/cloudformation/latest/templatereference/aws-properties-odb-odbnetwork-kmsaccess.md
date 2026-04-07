This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::OdbNetwork KmsAccess

Configuration for AWS Key Management Service (KMS) access from the ODB network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "Ipv4Addresses" : [ String, ... ],
  "KmsPolicyDocument" : String,
  "Status" : String
}

```

### YAML

```yaml

  DomainName: String
  Ipv4Addresses:
    - String
  KmsPolicyDocument: String
  Status: String

```

## Properties

`DomainName`

The domain name for AWS Key Management Service (KMS) access configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4Addresses`

The IPv4 addresses allowed for AWS Key Management Service (KMS) access.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsPolicyDocument`

The AWS Key Management Service (KMS) policy document that defines permissions for key usage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status of the AWS Key Management Service (KMS) access configuration.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | ENABLING | DISABLED | DISABLING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CrossRegionS3RestoreSourcesAccess

ManagedS3BackupAccess
