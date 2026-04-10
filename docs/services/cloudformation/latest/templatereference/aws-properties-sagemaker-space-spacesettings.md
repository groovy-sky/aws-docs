This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space SpaceSettings

A collection of space settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppType" : String,
  "CodeEditorAppSettings" : SpaceCodeEditorAppSettings,
  "CustomFileSystems" : [ CustomFileSystem, ... ],
  "JupyterLabAppSettings" : SpaceJupyterLabAppSettings,
  "JupyterServerAppSettings" : JupyterServerAppSettings,
  "KernelGatewayAppSettings" : KernelGatewayAppSettings,
  "RemoteAccess" : String,
  "SpaceManagedResources" : String,
  "SpaceStorageSettings" : SpaceStorageSettings
}

```

### YAML

```yaml

  AppType: String
  CodeEditorAppSettings:
    SpaceCodeEditorAppSettings
  CustomFileSystems:
    - CustomFileSystem
  JupyterLabAppSettings:
    SpaceJupyterLabAppSettings
  JupyterServerAppSettings:
    JupyterServerAppSettings
  KernelGatewayAppSettings:
    KernelGatewayAppSettings
  RemoteAccess: String
  SpaceManagedResources: String
  SpaceStorageSettings:
    SpaceStorageSettings

```

## Properties

`AppType`

The type of app created within the space.

If using the [UpdateSpace](../../../../reference/sagemaker/latest/apireference/api-updatespace.md) API, you can't change the app type of your
space by specifying a different value for this field.

_Required_: No

_Type_: String

_Allowed values_: `JupyterServer | KernelGateway | TensorBoard | RStudioServerPro | RSessionGateway | JupyterLab | CodeEditor`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeEditorAppSettings`

The Code Editor application settings.

_Required_: No

_Type_: [SpaceCodeEditorAppSettings](aws-properties-sagemaker-space-spacecodeeditorappsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomFileSystems`

A file system, created by you, that you assign to a space for an Amazon SageMaker AI
Domain. Permitted users can access this file system in Amazon SageMaker AI Studio.

_Required_: No

_Type_: Array of [CustomFileSystem](aws-properties-sagemaker-space-customfilesystem.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JupyterLabAppSettings`

The settings for the JupyterLab application.

_Required_: No

_Type_: [SpaceJupyterLabAppSettings](aws-properties-sagemaker-space-spacejupyterlabappsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JupyterServerAppSettings`

The JupyterServer app settings.

_Required_: No

_Type_: [JupyterServerAppSettings](aws-properties-sagemaker-space-jupyterserverappsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KernelGatewayAppSettings`

The KernelGateway app settings.

_Required_: No

_Type_: [KernelGatewayAppSettings](aws-properties-sagemaker-space-kernelgatewayappsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoteAccess`

A setting that enables or disables remote access for a SageMaker space. When enabled, this
allows you to connect to the remote space from your local IDE.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceManagedResources`

If you enable this option, SageMaker AI creates the following resources on your
behalf when you create the space:

- The user profile that possesses the space.

- The app that the space contains.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceStorageSettings`

The storage settings for a space.

_Required_: No

_Type_: [SpaceStorageSettings](aws-properties-sagemaker-space-spacestoragesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpaceJupyterLabAppSettings

SpaceSharingSettings

All content copied from https://docs.aws.amazon.com/.
