This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPC Tag

Specifies a tag. For more information, see [Resource tags](../userguide/aws-properties-resource-tags.md).

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

The tag key.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

This example specifies two tags for the VPC.

#### JSON

```json

"Tags" : [
   {
      "Key" : "key1",
      "Value" : "value1"
   },
   {
      "Key" : "key2",
      "Value" : "value2"
   }
]
```

#### YAML

```yaml

Tags:
  - Key: "key1"
    Value: "value1"
  - Key: "key2"
    Value: "value2"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPC

AWS::EC2::VPCBlockPublicAccessExclusion

All content copied from https://docs.aws.amazon.com/.
