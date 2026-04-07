This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::UserGroup Tag

A tag that can be added to an ElastiCache cluster or replication group. Tags are
composed of a Key/Value pair. You can use tags to categorize and track all your
ElastiCache resources, with the exception of global replication group. When you add or
remove tags on replication groups, those actions will be replicated to all nodes in the
replication group. A tag with a null Value is permitted.

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

The key for the tag. May not be null.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!aws:)[a-zA-Z0-9 _\.\/=+:\-@]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag's value. May be null.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9 _\.\/=+:\-@]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElastiCache::UserGroup

Next
