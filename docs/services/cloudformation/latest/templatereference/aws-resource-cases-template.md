This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Template

Creates a template in the Cases domain. This template is used to define the case object
model (that is, to define what data can be captured on cases) in a Cases domain. A template
must have a unique name within a domain, and it must reference existing field IDs and layout
IDs. Additionally, multiple fields with same IDs are not allowed within the same Template. A
template can be either Active or Inactive, as indicated by its status. Inactive templates
cannot be used to create cases.

Other template APIs are:

- [DeleteTemplate](https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_DeleteTemplate.html)

- [GetTemplate](https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_GetTemplate.html)

- [ListTemplates](https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_ListTemplates.html)

- [UpdateTemplate](https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_UpdateTemplate.html)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cases::Template",
  "Properties" : {
      "Description" : String,
      "DomainId" : String,
      "LayoutConfiguration" : LayoutConfiguration,
      "Name" : String,
      "RequiredFields" : [ RequiredField, ... ],
      "Rules" : [ TemplateRule, ... ],
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cases::Template
Properties:
  Description: String
  DomainId: String
  LayoutConfiguration:
    LayoutConfiguration
  Name: String
  RequiredFields:
    - RequiredField
  Rules:
    - TemplateRule
  Status: String
  Tags:
    - Tag

```

## Properties

`Description`

A brief description of the template.

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

`LayoutConfiguration`

Object to store configuration of layouts associated to the template.

_Required_: No

_Type_: [LayoutConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-template-layoutconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The template name.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[\S]$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequiredFields`

A list of fields that must contain a value for a case to be successfully created with this
template.

_Required_: No

_Type_: Array of [RequiredField](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-template-requiredfield.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

A list of case rules (also known as [case field conditions](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html)) on a template.

_Required_: No

_Type_: Array of [TemplateRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-template-templaterule.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the template.

_Required_: No

_Type_: String

_Allowed values_: `Active | Inactive`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-template-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the template. For example:

`arn:aws:cases:us-west-2:123456789012:domain/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111/template/a1b2c3d4-5678-90ab-cdef-EXAMPLE55555`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTime`

Timestamp at which the resource was created.

`LastModifiedTime`

Timestamp at which the resource was created or last modified.

`TemplateArn`

The Amazon Resource Name (ARN) of the template.

`TemplateId`

A unique identifier of a template.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

LayoutConfiguration
