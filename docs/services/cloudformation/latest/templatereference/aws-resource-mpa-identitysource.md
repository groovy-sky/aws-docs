This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MPA::IdentitySource

Creates a new identity source. For more information, see [Identity Source](https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html) in the _Multi-party approval User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MPA::IdentitySource",
  "Properties" : {
      "IdentitySourceParameters" : IdentitySourceParameters,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MPA::IdentitySource
Properties:
  IdentitySourceParameters:
    IdentitySourceParameters
  Tags:
    - Tag

```

## Properties

`IdentitySourceParameters`

A ` IdentitySourceParameters` object. Contains details for the resource that provides identities to the identity source. For example, an IAM Identity Center instance.

_Required_: Yes

_Type_: [IdentitySourceParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mpa-identitysource-identitysourceparameters.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags that you have added to the specified resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mpa-identitysource-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreationTime`

Timestamp when the identity source was created.

`IdentitySourceArn`

Amazon Resource Name (ARN) for the identity source.

`IdentitySourceParameters.IamIdentityCenter.ApprovalPortalUrl`

URL for the approval portal associated with the IAM Identity Center instance.

`IdentitySourceType`

The type of resource that provided identities to the identity source. For example, an IAM Identity Center instance.

`Status`

Status for the identity source. For example, if the identity source is `ACTIVE`.

`StatusCode`

Status code of the identity source.

`StatusMessage`

Message describing the status for the identity source.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

IamIdentityCenter
