This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain DefaultSpaceSettings

The default settings for shared spaces that users create in the domain.

SageMaker applies these settings only to shared spaces. It doesn't apply them to private
spaces.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomFileSystemConfigs" : [ CustomFileSystemConfig, ... ],
  "CustomPosixUserConfig" : CustomPosixUserConfig,
  "ExecutionRole" : String,
  "JupyterLabAppSettings" : JupyterLabAppSettings,
  "JupyterServerAppSettings" : JupyterServerAppSettings,
  "KernelGatewayAppSettings" : KernelGatewayAppSettings,
  "SecurityGroups" : [ String, ... ],
  "SpaceStorageSettings" : DefaultSpaceStorageSettings
}

```

### YAML

```yaml

  CustomFileSystemConfigs:
    - CustomFileSystemConfig
  CustomPosixUserConfig:
    CustomPosixUserConfig
  ExecutionRole: String
  JupyterLabAppSettings:
    JupyterLabAppSettings
  JupyterServerAppSettings:
    JupyterServerAppSettings
  KernelGatewayAppSettings:
    KernelGatewayAppSettings
  SecurityGroups:
    - String
  SpaceStorageSettings:
    DefaultSpaceStorageSettings

```

## Properties

`CustomFileSystemConfigs`

The settings for assigning a custom file system to a domain. Permitted users can access
this file system in Amazon SageMaker AI Studio.

_Required_: No

_Type_: Array of [CustomFileSystemConfig](aws-properties-sagemaker-domain-customfilesystemconfig.md)

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomPosixUserConfig`

Property description not available.

_Required_: No

_Type_: [CustomPosixUserConfig](aws-properties-sagemaker-domain-customposixuserconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

The ARN of the execution role for the space.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JupyterLabAppSettings`

Property description not available.

_Required_: No

_Type_: [JupyterLabAppSettings](aws-properties-sagemaker-domain-jupyterlabappsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JupyterServerAppSettings`

The JupyterServer app settings.

_Required_: No

_Type_: [JupyterServerAppSettings](aws-properties-sagemaker-domain-jupyterserverappsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KernelGatewayAppSettings`

The KernelGateway app settings.

_Required_: No

_Type_: [KernelGatewayAppSettings](aws-properties-sagemaker-domain-kernelgatewayappsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroups`

The security group IDs for the Amazon VPC that the space uses for
communication.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `32 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceStorageSettings`

Property description not available.

_Required_: No

_Type_: [DefaultSpaceStorageSettings](aws-properties-sagemaker-domain-defaultspacestoragesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultEbsStorageSettings

DefaultSpaceStorageSettings

All content copied from https://docs.aws.amazon.com/.
