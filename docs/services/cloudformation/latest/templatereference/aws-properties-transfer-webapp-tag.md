This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::WebApp Tag

Creates a key-value pair for a specific resource. Tags are metadata that you can use
to search for and group a resource for various purposes. You can apply tags to servers,
users, and roles. A tag key can take more than one value. For example, to group servers
for accounting purposes, you might create a tag called `Group` and assign the
values `Research` and `Accounting` to that group.

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

The name assigned to the tag that you create.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Contains one or more values that you assigned to the key name you create.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IdentityProviderDetails

Vpc
