This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::StorageVirtualMachine Tag

Specifies a key-value pair for a resource tag.

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

A value that specifies the `TagKey`, the name of the tag. Tag keys must
be unique for the resource to which they are attached.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A value that specifies the `TagValue`, the value assigned to the
corresponding tag key. Tag values can be null and don't have to be unique in a tag set.
For example, you can have a key-value pair in a tag set of `finances : April`
and also of `payroll : April`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SelfManagedActiveDirectoryConfiguration

AWS::FSx::Volume

All content copied from https://docs.aws.amazon.com/.
