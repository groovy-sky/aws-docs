This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection AuthorizationCodeProperties

The authorization code properties of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationCode" : String,
  "RedirectUri" : String
}

```

### YAML

```yaml

  AuthorizationCode: String
  RedirectUri: String

```

## Properties

`AuthorizationCode`

The authorization code of a connection.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedirectUri`

The redirect URI of a connection.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthenticationConfigurationInput

AwsLocation

All content copied from https://docs.aws.amazon.com/.
