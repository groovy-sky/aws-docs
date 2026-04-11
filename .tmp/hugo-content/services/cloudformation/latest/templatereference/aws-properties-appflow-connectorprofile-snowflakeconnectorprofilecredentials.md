This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile SnowflakeConnectorProfileCredentials

The connector-specific profile credentials required when using Snowflake.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Password" : String,
  "Username" : String
}

```

### YAML

```yaml

  Password: String
  Username: String

```

## Properties

`Password`

The password that corresponds to the user name.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The name of the user.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SnowflakeConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-snowflakeconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlackConnectorProfileProperties

SnowflakeConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
