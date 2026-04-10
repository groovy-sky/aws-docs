This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource CredentialPair

The combination of user name and password that are used as credentials.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlternateDataSourceParameters" : [ DataSourceParameters, ... ],
  "Password" : String,
  "Username" : String
}

```

### YAML

```yaml

  AlternateDataSourceParameters:
    - DataSourceParameters
  Password: String
  Username: String

```

## Properties

`AlternateDataSourceParameters`

A set of alternate data source parameters that you want to share for these
credentials. The credentials are applied in tandem with the data source parameters when
you copy a data source by using a create or update request. The API operation compares
the `DataSourceParameters` structure that's in the request with the
structures in the `AlternateDataSourceParameters` allow list. If the
structures are an exact match, the request is allowed to use the new data source with
the existing credentials. If the `AlternateDataSourceParameters` list is
null, the `DataSourceParameters` originally used with these
`Credentials` is automatically allowed.

_Required_: No

_Type_: Array of [DataSourceParameters](aws-properties-quicksight-datasource-datasourceparameters.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

Password.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

User name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuroraPostgreSqlParameters

DatabricksParameters

All content copied from https://docs.aws.amazon.com/.
