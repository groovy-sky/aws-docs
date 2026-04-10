This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::WebExperience SamlProviderConfiguration

Information about the SAML 2.0-compliant identity provider (IdP) used to authenticate
end users of an Amazon Q Business web experience.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationUrl" : String
}

```

### YAML

```yaml

  AuthenticationUrl: String

```

## Properties

`AuthenticationUrl`

The URL where Amazon Q Business end users will be redirected for authentication.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*$`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenIDConnectProviderConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
