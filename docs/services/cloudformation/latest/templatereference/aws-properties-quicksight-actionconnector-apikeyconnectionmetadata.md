This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector APIKeyConnectionMetadata

Configuration for API key-based authentication to external services.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKey" : String,
  "BaseEndpoint" : String,
  "Email" : String
}

```

### YAML

```yaml

  ApiKey: String
  BaseEndpoint: String
  Email: String

```

## Properties

`ApiKey`

The API key used for authentication.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseEndpoint`

The base URL endpoint for the external service.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Email`

The email address associated with the API key, if required.

_Required_: No

_Type_: String

_Pattern_: `^[\w.%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QuickSight::ActionConnector

AuthConfig

All content copied from https://docs.aws.amazon.com/.
