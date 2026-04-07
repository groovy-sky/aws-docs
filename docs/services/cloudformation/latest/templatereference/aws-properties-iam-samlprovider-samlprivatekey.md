This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::SAMLProvider SAMLPrivateKey

Contains the private keys for the SAML provider.

This data type is used as a response element in the [GetSAMLProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetSAMLProvider.html)
operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyId" : String,
  "Timestamp" : String
}

```

### YAML

```yaml

  KeyId: String
  Timestamp: String

```

## Properties

`KeyId`

The unique identifier for the SAML private key.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Z0-9]+`

_Minimum_: `22`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timestamp`

The date and time, in [ISO 8601 date-time](http://www.iso.org/iso/iso8601) format, when the private key was uploaded.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IAM::SAMLProvider

Tag
