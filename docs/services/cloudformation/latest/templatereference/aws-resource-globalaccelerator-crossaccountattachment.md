This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GlobalAccelerator::CrossAccountAttachment

Create a cross-account attachment in AWS Global Accelerator. You create a cross-account attachment to
specify the _principals_ who have permission to work with _resources_
in accelerators in their own account. You specify, in the same attachment, the resources that are shared.

A principal can be an AWS account number or the Amazon Resource Name (ARN) for an
accelerator. For account numbers that are listed as principals, to work with a resource listed in the attachment,
you must sign in to an account specified as a principal. Then, you can work with resources that are listed,
with any of your accelerators. If an accelerator ARN is listed in the cross-account attachment as a principal,
anyone with permission to make updates to the accelerator can work with resources that are listed in the
attachment.

Specify each principal and resource separately. To specify two CIDR address pools, list
them individually under `Resources`, and so on. For a command line operation, for example,
you might use a statement like the following:

` "Resources": [{"Cidr": "169.254.60.0/24"},{"Cidr": "169.254.59.0/24"}]`

For more information, see [Working with cross-account attachments and resources in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/cross-account-resources.html) in the _AWS Global Accelerator Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GlobalAccelerator::CrossAccountAttachment",
  "Properties" : {
      "Name" : String,
      "Principals" : [ String, ... ],
      "Resources" : [ Resource, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GlobalAccelerator::CrossAccountAttachment
Properties:
  Name: String
  Principals:
    - String
  Resources:
    - Resource
  Tags:
    - Tag

```

## Properties

`Name`

The name of the cross-account attachment.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{0,64}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principals`

The principals included in the cross-account attachment.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resources`

The resources included in the cross-account attachment.

_Required_: No

_Type_: Array of [Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-globalaccelerator-crossaccountattachment-resource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Add tags for a cross-account attachment.

For more information, see [Tagging\
in AWS Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the _AWS Global Accelerator Developer Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-globalaccelerator-crossaccountattachment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the cross-account attachment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachmentArn`

The Amazon Resource Name (ARN) of the cross-account attachment.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Resource
