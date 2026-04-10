This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::Project

The `AWS::Rekognition::Project` type creates an Amazon Rekognition Custom Labels
project. A project is a group of resources needed to create and manage versions of an
Amazon Rekognition Custom Labels model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Rekognition::Project",
  "Properties" : {
      "ProjectName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Rekognition::Project
Properties:
  ProjectName: String
  Tags:
    - Tag

```

## Properties

`ProjectName`

The name of the project to create.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9][a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-rekognition-project-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns
the project name. For example:

`{ "Ref": "CircuitBoardProject" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Fn::GetAtt`
returns a value for a specified attribute of this type. The
following are the available attributes and sample return values. For more information about
using `Fn::GetAtt`, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name of the project.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
