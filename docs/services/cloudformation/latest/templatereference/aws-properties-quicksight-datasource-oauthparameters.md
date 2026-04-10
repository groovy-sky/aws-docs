This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource OAuthParameters

An object that contains information needed to create a data source connection that uses OAuth client credentials. This option is available for data source connections that are made with Snowflake and Starburst.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdentityProviderResourceUri" : String,
  "IdentityProviderVpcConnectionProperties" : VpcConnectionProperties,
  "OAuthScope" : String,
  "TokenProviderUrl" : String
}

```

### YAML

```yaml

  IdentityProviderResourceUri: String
  IdentityProviderVpcConnectionProperties:
    VpcConnectionProperties
  OAuthScope: String
  TokenProviderUrl: String

```

## Properties

`IdentityProviderResourceUri`

The resource uri of the identity provider.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProviderVpcConnectionProperties`

Property description not available.

_Required_: No

_Type_: [VpcConnectionProperties](aws-properties-quicksight-datasource-vpcconnectionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthScope`

The OAuth scope.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenProviderUrl`

The token endpoint URL of the identity provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MySqlParameters

OracleParameters

All content copied from https://docs.aws.amazon.com/.
