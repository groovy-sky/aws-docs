This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::AppImageConfig ContainerConfig

The configuration used to run the application image container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerArguments" : [ String, ... ],
  "ContainerEntrypoint" : [ String, ... ],
  "ContainerEnvironmentVariables" : [ CustomImageContainerEnvironmentVariable, ... ]
}

```

### YAML

```yaml

  ContainerArguments:
    - String
  ContainerEntrypoint:
    - String
  ContainerEnvironmentVariables:
    - CustomImageContainerEnvironmentVariable

```

## Properties

`ContainerArguments`

The arguments for the container when you're running the application.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerEntrypoint`

The entrypoint used to run the application in the container.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerEnvironmentVariables`

The environment variables to set in the container

_Required_: No

_Type_: Array of [CustomImageContainerEnvironmentVariable](aws-properties-sagemaker-appimageconfig-customimagecontainerenvironmentvariable.md)

_Minimum_: `0`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeEditorAppImageConfig

CustomImageContainerEnvironmentVariable

All content copied from https://docs.aws.amazon.com/.
