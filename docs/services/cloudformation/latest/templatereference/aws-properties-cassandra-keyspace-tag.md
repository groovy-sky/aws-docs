This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Keyspace Tag

Describes a tag. A tag is a key-value pair. You can add up to 50 tags to a single Amazon Keyspaces resource.

AWS-assigned tag names and values are automatically assigned the `aws:` prefix, which the user cannot assign.
AWS-assigned tag names do not count towards the tag limit of 50. User-assigned tag names have the
prefix `user:` in the Cost Allocation Report. You cannot backdate the application of a tag.

For more information, see [Adding tags and labels to Amazon Keyspaces resources](https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html) in the _Amazon Keyspaces Developer_
_Guide_.

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

The key of the tag. Tag keys are case sensitive. Each Amazon Keyspaces resource can only have up to one tag with the same key. If you try to add an
existing tag (same key), the existing tag value will be updated to the new value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the tag. Tag values are case-sensitive and can be null.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplicationSpecification

AWS::Cassandra::Table
