This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::TaskTemplate

Specifies a task template for a Amazon Connect instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::TaskTemplate",
  "Properties" : {
      "ClientToken" : String,
      "Constraints" : Constraints,
      "ContactFlowArn" : String,
      "Defaults" : [ DefaultFieldValue, ... ],
      "Description" : String,
      "Fields" : [ Field, ... ],
      "InstanceArn" : String,
      "Name" : String,
      "SelfAssignContactFlowArn" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::TaskTemplate
Properties:
  ClientToken: String
  Constraints:
    Constraints
  ContactFlowArn: String
  Defaults:
    - DefaultFieldValue
  Description: String
  Fields:
    - Field
  InstanceArn: String
  Name: String
  SelfAssignContactFlowArn: String
  Status: String
  Tags:
    - Tag

```

## Properties

`ClientToken`

A unique, case-sensitive identifier that you provide to ensure the idempotency of the
request.

_Required_: No

_Type_: String

_Pattern_: `^$|[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Constraints`

Constraints that are applicable to the fields listed.

The values can be represented in either JSON or YAML format. For an example of the
JSON configuration, see _Examples_ at the bottom of this page.

_Required_: No

_Type_: [Constraints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-constraints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContactFlowArn`

The Amazon Resource Name (ARN) of the flow that runs by default when a task is created
by referencing this template. `ContactFlowArn` is not required when there is
a field with `fieldType` = `QUICK_CONNECT`.

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Defaults`

The default values for fields when a task is created by referencing this
template.

_Required_: No

_Type_: Array of [DefaultFieldValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-defaultfieldvalue.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the task template.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fields`

Fields that are part of the template. A template requires at least one field that has
type `Name`.

_Required_: No

_Type_: Array of [Field](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-field.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the task template.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfAssignContactFlowArn`

The Amazon Resource Name (ARN) of the flow.

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the task template.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | INACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-tasktemplate-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the task template. For example:

`{ "Ref": "myTaskTemplate" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the task template.

## Examples

### JSON format for Constraints

Following is an example of the JSON format for the `Constraints`
property.

#### JSON

```json

{
  "Type" : "AWS::Connect::TaskTemplate",
  "Properties" : {
      "ClientToken" : String,
      "Constraints" : Constraints,
      "ContactFlowArn" : String,
      "Defaults" : [ DefaultFieldValue, ... ],
      "Description" : String,
      "Fields" : [ Field, ... ],
      "InstanceArn" : String,
      "Name" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Constraints
