This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space

Creates a private space or a space used for real time collaboration in a domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Space",
  "Properties" : {
      "DomainId" : String,
      "OwnershipSettings" : OwnershipSettings,
      "SpaceDisplayName" : String,
      "SpaceName" : String,
      "SpaceSettings" : SpaceSettings,
      "SpaceSharingSettings" : SpaceSharingSettings,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Space
Properties:
  DomainId: String
  OwnershipSettings:
    OwnershipSettings
  SpaceDisplayName: String
  SpaceName: String
  SpaceSettings:
    SpaceSettings
  SpaceSharingSettings:
    SpaceSharingSettings
  Tags:
    - Tag

```

## Properties

`DomainId`

The ID of the associated domain.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OwnershipSettings`

The collection of ownership settings for a space.

_Required_: No

_Type_: [OwnershipSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-space-ownershipsettings.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpaceDisplayName`

The name of the space that appears in the Studio UI.

_Required_: No

_Type_: String

_Pattern_: `^(?!\s*$).+`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceName`

The name of the space.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpaceSettings`

A collection of space settings.

_Required_: No

_Type_: [SpaceSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-space-spacesettings.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceSharingSettings`

A collection of space sharing settings.

_Required_: No

_Type_: [SpaceSharingSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-space-spacesharingsettings.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-space-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the domain that the space is associated with and the name of the
space.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`SpaceArn`

The space's Amazon Resource Name (ARN).

`Url`

Returns the URL of the space. If the space is created with AWS IAM Identity Center (Successor to
AWS Single Sign-On) authentication, users can navigate to the URL after appending the respective
redirect parameter for the application type to be federated through AWS IAM Identity Center.

The following application types are supported:

- Studio Classic: `&redirect=JupyterServer`

- JupyterLab: `&redirect=JupyterLab`

- Code Editor, based on Code-OSS, Visual Studio Code - Open Source: `&redirect=CodeEditor`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TemplateProviderDetail

CodeRepository
