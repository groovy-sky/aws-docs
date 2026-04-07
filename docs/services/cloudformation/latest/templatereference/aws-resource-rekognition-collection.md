This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::Collection

The `AWS::Rekognition::Collection` type creates a server-side container called a collection. You can use a collection
to store information about detected faces and search for known faces in images, stored videos, and streaming videos.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Rekognition::Collection",
  "Properties" : {
      "CollectionId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Rekognition::Collection
Properties:
  CollectionId: String
  Tags:
    - Tag

```

## Properties

`CollectionId`

ID for the collection that you are creating.

_Required_: Yes

_Type_: String

_Pattern_: `\A[a-zA-Z0-9_\.\-]+$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A set of tags (key-value pairs) that you want to attach to the collection.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rekognition-collection-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns
the collection ID. For example:

`{ "Ref": "MyCollection" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Fn::GetAtt`
returns a value for a specified attribute of this type. The
following are the available attributes and sample return values. For more information about
using `Fn::GetAtt`, see
[Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Arn`

Returns the Amazon Resource Name of the collection.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Rekognition

Tag
