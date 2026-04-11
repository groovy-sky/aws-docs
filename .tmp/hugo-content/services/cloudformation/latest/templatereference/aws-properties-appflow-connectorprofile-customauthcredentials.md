This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile CustomAuthCredentials

The custom credentials required for custom authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialsMap" : {Key: Value, ...},
  "CustomAuthenticationType" : String
}

```

### YAML

```yaml

  CredentialsMap:
    Key: Value
  CustomAuthenticationType: String

```

## Properties

`CredentialsMap`

A map that holds custom authentication credentials.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\w]{1,128}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAuthenticationType`

The custom authentication type that the connector uses.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorProfileProperties

CustomConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
