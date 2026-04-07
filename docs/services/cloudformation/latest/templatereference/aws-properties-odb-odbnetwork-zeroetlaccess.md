This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::OdbNetwork ZeroEtlAccess

The configuration for Zero-ETL access from the ODB network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidr" : String,
  "Status" : String
}

```

### YAML

```yaml

  Cidr: String
  Status: String

```

## Properties

`Cidr`

The CIDR block for the Zero-ETL access.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the Zero-ETL access.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | ENABLING | DISABLED | DISABLING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::ODB::OdbPeeringConnection
