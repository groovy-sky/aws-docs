This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile HiddenSageMakerImage

The SageMaker images that are hidden from the Studio user interface. You must specify the SageMaker
image name and version aliases.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SageMakerImageName" : String,
  "VersionAliases" : [ String, ... ]
}

```

### YAML

```yaml

  SageMakerImageName: String
  VersionAliases:
    - String

```

## Properties

`SageMakerImageName`

The SageMaker image name that you are hiding from the Studio user interface.

_Required_: No

_Type_: String

_Allowed values_: `sagemaker_distribution`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionAliases`

The version aliases you are hiding from the Studio user interface.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FSxLustreFileSystemConfig

IdleSettings

All content copied from https://docs.aws.amazon.com/.
