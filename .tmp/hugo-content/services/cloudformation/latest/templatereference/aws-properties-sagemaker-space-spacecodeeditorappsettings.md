This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space SpaceCodeEditorAppSettings

The application settings for a Code Editor space.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppLifecycleManagement" : SpaceAppLifecycleManagement,
  "DefaultResourceSpec" : ResourceSpec
}

```

### YAML

```yaml

  AppLifecycleManagement:
    SpaceAppLifecycleManagement
  DefaultResourceSpec:
    ResourceSpec

```

## Properties

`AppLifecycleManagement`

Settings that are used to configure and manage the lifecycle of CodeEditor applications in
a space.

_Required_: No

_Type_: [SpaceAppLifecycleManagement](aws-properties-sagemaker-space-spaceapplifecyclemanagement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultResourceSpec`

Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs
on.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-space-resourcespec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpaceAppLifecycleManagement

SpaceIdleSettings

All content copied from https://docs.aws.amazon.com/.
