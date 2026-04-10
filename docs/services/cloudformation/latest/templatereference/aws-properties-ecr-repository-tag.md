This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::Repository Tag

The metadata to apply to a resource to help you categorize and organize them. Each tag
consists of a key and a value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

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

One part of a key-value pair that make up a tag. A `key` is a general label
that acts like a category for more specific tag values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A `value` acts as a descriptor within a tag category (key).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecyclePolicy

AWS::ECR::RepositoryCreationTemplate

All content copied from https://docs.aws.amazon.com/.
