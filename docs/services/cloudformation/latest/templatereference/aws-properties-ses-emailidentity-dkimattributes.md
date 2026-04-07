This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::EmailIdentity DkimAttributes

Used to enable or disable DKIM authentication for an email identity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SigningEnabled" : Boolean
}

```

### YAML

```yaml

  SigningEnabled: Boolean

```

## Properties

`SigningEnabled`

Sets the DKIM signing configuration for the identity.

When you set this value `true`, then the messages that are sent from the
identity are signed using DKIM. If you set this value to `false`, your
messages are sent without DKIM signing.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConfigurationSetAttributes

DkimSigningAttributes
