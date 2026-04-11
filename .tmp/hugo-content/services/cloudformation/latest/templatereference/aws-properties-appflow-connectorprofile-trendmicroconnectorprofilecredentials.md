This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile TrendmicroConnectorProfileCredentials

The connector-specific profile credentials required when using Trend Micro.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiSecretKey" : String
}

```

### YAML

```yaml

  ApiSecretKey: String

```

## Properties

`ApiSecretKey`

The Secret Access Key portion of the credentials.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [TrendmicroConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-trendmicroconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeConnectorProfileProperties

VeevaConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
