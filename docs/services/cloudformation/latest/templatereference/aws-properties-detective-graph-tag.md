This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Detective::Graph Tag

The metadata that you apply to a resource to help you categorize and organize it.
Each tag consists of a key and an optional value, both of which you define. For more information about tags,
see [What is Tag Editor](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html)
in the _Tagging AWS Resources and Tag Editor User Guide_.

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

One part of a key-value pair that makes up a tag. A key is a general label that acts like a category for more specific tag values.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional part of a key-value pair that makes up a tag. A value acts as a descriptor in a tag category (key).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Detective::Graph

AWS::Detective::MemberInvitation
