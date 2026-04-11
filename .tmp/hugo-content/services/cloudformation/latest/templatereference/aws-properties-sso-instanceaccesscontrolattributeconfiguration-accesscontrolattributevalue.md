This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::InstanceAccessControlAttributeConfiguration AccessControlAttributeValue

The value used for mapping a specified attribute to an identity source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Source" : [ String, ... ]
}

```

### YAML

```yaml

  Source:
    - String

```

## Properties

`Source`

The identity source to use when mapping a specified attribute to IAM Identity Center.

_Required_: Yes

_Type_: Array of String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessControlAttribute

AWS::SSO::PermissionSet

All content copied from https://docs.aws.amazon.com/.
