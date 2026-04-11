This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::Application PortalOptionsConfiguration

A structure that describes the options for the portal associated with an
application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SignInOptions" : SignInOptions,
  "Visibility" : String
}

```

### YAML

```yaml

  SignInOptions:
    SignInOptions
  Visibility: String

```

## Properties

`SignInOptions`

A structure that describes the sign-in options for the access portal.

_Required_: No

_Type_: [SignInOptions](aws-properties-sso-application-signinoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Indicates whether this application is visible in the access portal.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSO::Application

SignInOptions

All content copied from https://docs.aws.amazon.com/.
