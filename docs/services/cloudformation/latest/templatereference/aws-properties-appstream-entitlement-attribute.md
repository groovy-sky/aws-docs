This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Entitlement Attribute

An attribute that belongs to an entitlement. Application entitlements work by matching a
supported SAML 2.0 attribute name to a value when a user identity federates to an AppStream
2.0 SAML application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

A supported AWS IAM SAML PrincipalTag attribute that is matched to a value when a user
identity federates to an WorkSpaces Applications SAML application.

The following are supported values:

- roles

- department

- organization

- groups

- title

- costCenter

- userType

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A value that is matched to a supported SAML attribute name when a user identity federates to an WorkSpaces Applications SAML application.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppStream::Entitlement

AWS::AppStream::Fleet
