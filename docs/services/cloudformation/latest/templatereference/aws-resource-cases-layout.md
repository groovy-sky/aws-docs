This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Layout

Creates a layout in the Cases domain. Layouts define the following configuration in
the top section and More Info tab of the Cases user interface:

- Fields to display to the users

- Field ordering

###### Note

Title and Status fields cannot be part of layouts since they are not
configurable.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cases::Layout",
  "Properties" : {
      "Content" : LayoutContent,
      "DomainId" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cases::Layout
Properties:
  Content:
    LayoutContent
  DomainId: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Content`

Object to store union of different versions of layout content.

_Required_: Yes

_Type_: [LayoutContent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-layout-layoutcontent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainId`

The unique identifier of the Cases domain.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: Updates are not supported.

`Name`

The name of the layout.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[\S]$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-layout-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the layout. For example:

`arn:aws:cases:us-west-2:123456789012:domain/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111/layout/a1b2c3d4-5678-90ab-cdef-EXAMPLE44444`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTime`

Timestamp at which the resource was created.

`LastModifiedTime`

Timestamp at which the resource was created or last modified.

`LayoutArn`

The Amazon Resource Name (ARN) of the newly created layout.

`LayoutId`

The unique identifier of the layout.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TextAttributes

BasicLayout
