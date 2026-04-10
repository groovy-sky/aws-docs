This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Capability SsoIdentity

An IAM Identity CenterIAM; Identity Center identity (user or group) that can be assigned permissions in a capability.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Type" : String
}

```

### YAML

```yaml

  Id: String
  Type: String

```

## Properties

`Id`

The unique identifier of the IAM Identity CenterIAM; Identity Center user or group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of identity. Valid values are `SSO_USER` or `SSO_GROUP`.

_Required_: Yes

_Type_: String

_Allowed values_: `SSO_USER | SSO_GROUP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkAccess

Tag

All content copied from https://docs.aws.amazon.com/.
