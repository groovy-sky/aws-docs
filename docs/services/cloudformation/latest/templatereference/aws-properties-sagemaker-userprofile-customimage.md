This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile CustomImage

A custom SageMaker AI image. For more information, see
[Bring your own SageMaker AI image](../../../sagemaker/latest/dg/studio-byoi.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppImageConfigName" : String,
  "ImageName" : String,
  "ImageVersionNumber" : Integer
}

```

### YAML

```yaml

  AppImageConfigName: String
  ImageName: String
  ImageVersionNumber: Integer

```

## Properties

`AppImageConfigName`

The name of the AppImageConfig.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageName`

The name of the CustomImage. Must be unique to your account.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageVersionNumber`

The version number of the CustomImage.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomFileSystemConfig

CustomPosixUserConfig

All content copied from https://docs.aws.amazon.com/.
