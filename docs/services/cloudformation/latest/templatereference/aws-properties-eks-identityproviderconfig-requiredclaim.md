This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::IdentityProviderConfig RequiredClaim

A key-value pair that describes a required claim in the identity token. If set, each
claim is verified to be present in the token with a matching value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key to match from the token.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The value for the key from the token.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OidcIdentityProviderConfig

Tag

All content copied from https://docs.aws.amazon.com/.
