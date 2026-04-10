This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# Alexa::ASK::Skill AuthenticationConfiguration

The `AuthenticationConfiguration` property type specifies the Login with
Amazon (LWA) configuration used to authenticate with the Alexa service. Only Login with
Amazon security profiles created through the [Build Skills with the Alexa Skills Kit developer documentation](https://developer.amazon.com/docs/ask-overviews/build-skills-with-the-alexa-skills-kit.html) are supported for authentication. A client ID, client secret, and
refresh token are required. You can generate a client ID and client secret by creating a
new [security profile](https://developer.amazon.com/lwa/sp/create-security-profile.html) on the Amazon Developer Portal or you can retrieve them
from an existing profile. You can then retrieve the refresh token using the Alexa Skills
Kit CLI. For instructions, see [util-command](https://developer.amazon.com/docs/smapi/ask-cli-command-reference.html)
in the [ASK CLI Command Reference](https://developer.amazon.com/docs/smapi/ask-cli-command-reference.html).

`AuthenticationConfiguration` is a property of the
`Alexa::ASK::Skill` resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ClientSecret" : String,
  "RefreshToken" : String
}

```

### YAML

```yaml

  ClientId: String
  ClientSecret: String
  RefreshToken: String

```

## Properties

`ClientId`

Client ID from Login with Amazon (LWA).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

Client secret from Login with Amazon (LWA).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshToken`

Refresh token from Login with Amazon (LWA). This token is secret.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Alexa::ASK::Skill

Overrides

All content copied from https://docs.aws.amazon.com/.
