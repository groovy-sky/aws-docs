This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Field

Creates a field in the Cases domain. This field is used to define the case object
model (that is, defines what data can be captured on cases) in a Cases domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cases::Field",
  "Properties" : {
      "Attributes" : FieldAttributes,
      "Description" : String,
      "DomainId" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cases::Field
Properties:
  Attributes:
    FieldAttributes
  Description: String
  DomainId: String
  Name: String
  Tags:
    - Tag
  Type: String

```

## Properties

`Attributes`

Union of field attributes.

_Required_: No

_Type_: [FieldAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-field-fieldattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Description of the field.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainId`

The unique identifier of the Cases domain.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: Updates are not supported.

`Name`

Name of the field.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[\S]$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-field-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Type of the field.

_Required_: Yes

_Type_: String

_Allowed values_: `Text | Number | Boolean | DateTime | SingleSelect | Url | User`

_Update requires_: Updates are not supported.

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the field. For example:

`arn:aws:cases:us-west-2:123456789012:domain/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111/field/a1b2c3d4-5678-90ab-cdef-EXAMPLE22222`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTime`

Timestamp at which the resource was created.

`FieldArn`

The Amazon Resource Name (ARN) of the field.

`FieldId`

Unique identifier of the field.

`LastModifiedTime`

Timestamp at which the resource was created or last modified.

`Namespace`

Namespace of the field.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

FieldAttributes
