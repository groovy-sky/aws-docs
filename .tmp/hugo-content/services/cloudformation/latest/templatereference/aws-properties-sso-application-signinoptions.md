This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::Application SignInOptions

A structure that describes the sign-in options for an application portal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationUrl" : String,
  "Origin" : String
}

```

### YAML

```yaml

  ApplicationUrl: String
  Origin: String

```

## Properties

`ApplicationUrl`

The URL that accepts authentication requests for an application. This is a required
parameter if the `Origin` parameter is `APPLICATION`.

_Required_: No

_Type_: String

_Pattern_: `^http(s)?:\/\/[-a-zA-Z0-9+&@#\/%?=~_|!:,.;]*[-a-zA-Z0-9+&bb@#\/%?=~_|]$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Origin`

This determines how IAM Identity Center navigates the user to the target application. It can be one
of the following values:

- `APPLICATION`: IAM Identity Center redirects the customer to the configured
`ApplicationUrl`.

- `IDENTITY_CENTER`: IAM Identity Center uses SAML identity-provider initiated
authentication to sign the customer directly into a SAML-based
application.

_Required_: Yes

_Type_: String

_Allowed values_: `IDENTITY_CENTER | APPLICATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortalOptionsConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
