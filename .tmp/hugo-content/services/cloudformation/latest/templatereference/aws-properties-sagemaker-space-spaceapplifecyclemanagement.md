This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space SpaceAppLifecycleManagement

Settings that are used to configure and manage the lifecycle of Amazon SageMaker Studio
applications in a space.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleSettings" : SpaceIdleSettings
}

```

### YAML

```yaml

  IdleSettings:
    SpaceIdleSettings

```

## Properties

`IdleSettings`

Settings related to idle shutdown of Studio applications.

_Required_: No

_Type_: [SpaceIdleSettings](aws-properties-sagemaker-space-spaceidlesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3FileSystem

SpaceCodeEditorAppSettings

All content copied from https://docs.aws.amazon.com/.
