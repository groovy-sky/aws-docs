This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::Application Tag

A set of key-value pairs that are used to manage the resource. Tags can only be
applied to permission sets and cannot be applied to corresponding roles that IAM Identity Center
creates in AWS accounts.

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

The key for the tag.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w+=,.@-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the tag.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w+=,.@-]+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SignInOptions

AWS::SSO::ApplicationAssignment

All content copied from https://docs.aws.amazon.com/.
