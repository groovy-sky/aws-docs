This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource AuroraPostgreSqlParameters

Parameters for Amazon Aurora PostgreSQL-Compatible Edition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Database" : String,
  "Host" : String,
  "Port" : Number
}

```

### YAML

```yaml

  Database: String
  Host: String
  Port: Number

```

## Properties

`Database`

The Amazon Aurora PostgreSQL database to connect to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

The Amazon Aurora PostgreSQL-Compatible host to connect to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port that Amazon Aurora PostgreSQL is listening on.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuroraParameters

CredentialPair

All content copied from https://docs.aws.amazon.com/.
