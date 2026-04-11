This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space SpaceJupyterLabAppSettings

The settings for the JupyterLab application within a space.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppLifecycleManagement" : SpaceAppLifecycleManagement,
  "CodeRepositories" : [ CodeRepository, ... ],
  "DefaultResourceSpec" : ResourceSpec
}

```

### YAML

```yaml

  AppLifecycleManagement:
    SpaceAppLifecycleManagement
  CodeRepositories:
    - CodeRepository
  DefaultResourceSpec:
    ResourceSpec

```

## Properties

`AppLifecycleManagement`

Settings that are used to configure and manage the lifecycle of JupyterLab applications in
a space.

_Required_: No

_Type_: [SpaceAppLifecycleManagement](aws-properties-sagemaker-space-spaceapplifecyclemanagement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeRepositories`

A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterLab application.

_Required_: No

_Type_: Array of [CodeRepository](aws-properties-sagemaker-space-coderepository.md)

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultResourceSpec`

Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs
on.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-space-resourcespec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpaceIdleSettings

SpaceSettings

All content copied from https://docs.aws.amazon.com/.
