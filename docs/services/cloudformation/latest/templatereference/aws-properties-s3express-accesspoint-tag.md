This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::AccessPoint Tag

A key-value pair that you use to label your access points. You can add tags to new access points when you create them, or you can add tags to existing access points. Tags can help you organize and control access to access points. For more information, see [Using tags for attribute-based access control (ABAC)](../../../s3/latest/userguide/tagging.md#using-tags-for-abac).

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

The key of the tag. Tags are key-value pairs that you use to label your access points. Tags can help you organize and control access to access points. For more information, see [Tagging S3 resources for cost allocation or attribute-based access control (ABAC)](../../../s3/latest/userguide/tagging.md).

_Required_: Yes

_Type_: String

_Pattern_: `^(?!aws:.*)([\p{L}\p{Z}\p{N}_.:=+\/\-@%]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the tag. Tags are key-value pairs that you use to label your access points. Tags can help you organize and control access to access points. For more information, see [Tagging S3 resources for cost allocation or attribute-based access control (ABAC)](../../../s3/latest/userguide/tagging.md).

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:=+\/\-@%]*)$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Scope

VpcConfiguration
