This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource SnowflakeParameters

The parameters for Snowflake.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationType" : String,
  "Database" : String,
  "DatabaseAccessControlRole" : String,
  "Host" : String,
  "OAuthParameters" : OAuthParameters,
  "Warehouse" : String
}

```

### YAML

```yaml

  AuthenticationType: String
  Database: String
  DatabaseAccessControlRole: String
  Host: String
  OAuthParameters:
    OAuthParameters
  Warehouse: String

```

## Properties

`AuthenticationType`

The authentication type that you want to use for your connection. This parameter accepts OAuth and non-OAuth authentication types.

_Required_: No

_Type_: String

_Allowed values_: `PASSWORD | TOKEN | X509 | KEYPAIR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Database`

Database.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseAccessControlRole`

The database access control role.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

Host.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthParameters`

An object that contains information needed to create a data source connection between an Quick Sight account and Snowflake.

_Required_: No

_Type_: [OAuthParameters](aws-properties-quicksight-datasource-oauthparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Warehouse`

Warehouse.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3TablesParameters

SparkParameters

All content copied from https://docs.aws.amazon.com/.
