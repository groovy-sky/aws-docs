This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ApiKeyCredentials

The API key credentials required for API key authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKey" : String,
  "ApiSecretKey" : String
}

```

### YAML

```yaml

  ApiKey: String
  ApiSecretKey: String

```

## Properties

`ApiKey`

The API key required for API key authentication.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiSecretKey`

The API secret key required for API key authentication.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmplitudeConnectorProfileCredentials

BasicAuthCredentials

All content copied from https://docs.aws.amazon.com/.
