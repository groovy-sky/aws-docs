This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain UserSettings

A collection of settings that apply to users of Amazon SageMaker Studio. These settings are specified when the
[CreateUserProfile](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html)
API is called, and as `DefaultUserSettings` when the [CreateDomain](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html) API is called.

`SecurityGroups` is aggregated when specified in both calls. For all other settings in
`UserSettings`, the values specified in `CreateUserProfile` take precedence over those
specified in `CreateDomain`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoMountHomeEFS" : String,
  "CodeEditorAppSettings" : CodeEditorAppSettings,
  "CustomFileSystemConfigs" : [ CustomFileSystemConfig, ... ],
  "CustomPosixUserConfig" : CustomPosixUserConfig,
  "DefaultLandingUri" : String,
  "ExecutionRole" : String,
  "JupyterLabAppSettings" : JupyterLabAppSettings,
  "JupyterServerAppSettings" : JupyterServerAppSettings,
  "KernelGatewayAppSettings" : KernelGatewayAppSettings,
  "RSessionAppSettings" : RSessionAppSettings,
  "RStudioServerProAppSettings" : RStudioServerProAppSettings,
  "SecurityGroups" : [ String, ... ],
  "SharingSettings" : SharingSettings,
  "SpaceStorageSettings" : DefaultSpaceStorageSettings,
  "StudioWebPortal" : String,
  "StudioWebPortalSettings" : StudioWebPortalSettings
}

```

### YAML

```yaml

  AutoMountHomeEFS: String
  CodeEditorAppSettings:
    CodeEditorAppSettings
  CustomFileSystemConfigs:
    - CustomFileSystemConfig
  CustomPosixUserConfig:
    CustomPosixUserConfig
  DefaultLandingUri: String
  ExecutionRole: String
  JupyterLabAppSettings:
    JupyterLabAppSettings
  JupyterServerAppSettings:
    JupyterServerAppSettings
  KernelGatewayAppSettings:
    KernelGatewayAppSettings
  RSessionAppSettings:
    RSessionAppSettings
  RStudioServerProAppSettings:
    RStudioServerProAppSettings
  SecurityGroups:
    - String
  SharingSettings:
    SharingSettings
  SpaceStorageSettings:
    DefaultSpaceStorageSettings
  StudioWebPortal: String
  StudioWebPortalSettings:
    StudioWebPortalSettings

```

## Properties

`AutoMountHomeEFS`

Indicates whether auto-mounting of an EFS volume is supported for the user profile. The
`DefaultAsDomain` value is only supported for user profiles. Do not use the
`DefaultAsDomain` value when setting this parameter for a domain.

SageMaker applies this setting only to private spaces that the user creates in the domain. SageMaker doesn't apply this setting to shared spaces.

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled | DefaultAsDomain`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeEditorAppSettings`

The Code Editor application settings.

SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.

_Required_: No

_Type_: [CodeEditorAppSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-codeeditorappsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomFileSystemConfigs`

The settings for assigning a custom file system to a user profile. Permitted users can
access this file system in Amazon SageMaker AI Studio.

SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.

_Required_: No

_Type_: Array of [CustomFileSystemConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-customfilesystemconfig.html)

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomPosixUserConfig`

Details about the POSIX identity that is used for file system operations.

SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.

_Required_: No

_Type_: [CustomPosixUserConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-customposixuserconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultLandingUri`

The default experience that the user is directed to when accessing the domain. The
supported values are:

- `studio::`: Indicates that Studio is the default experience. This value can
only be passed if `StudioWebPortal` is set to `ENABLED`.

- `app:JupyterServer:`: Indicates that Studio Classic is the default
experience.

_Required_: No

_Type_: String

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

The execution role for the user.

SageMaker applies this setting only to private spaces that the user creates in the domain. SageMaker doesn't apply this setting to shared spaces.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JupyterLabAppSettings`

The settings for the JupyterLab application.

SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.

_Required_: No

_Type_: [JupyterLabAppSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-jupyterlabappsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JupyterServerAppSettings`

The Jupyter server's app settings.

_Required_: No

_Type_: [JupyterServerAppSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-jupyterserverappsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KernelGatewayAppSettings`

The kernel gateway app settings.

_Required_: No

_Type_: [KernelGatewayAppSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-kernelgatewayappsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RSessionAppSettings`

A collection of settings that configure the `RSessionGateway` app.

_Required_: No

_Type_: [RSessionAppSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-rsessionappsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RStudioServerProAppSettings`

A collection of settings that configure user interaction with the
`RStudioServerPro` app.

_Required_: No

_Type_: [RStudioServerProAppSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-rstudioserverproappsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroups`

The security groups for the Amazon Virtual Private Cloud (VPC) that the domain uses for
communication.

Optional when the `CreateDomain.AppNetworkAccessType` parameter is set to
`PublicInternetOnly`.

Required when the `CreateDomain.AppNetworkAccessType` parameter is set to
`VpcOnly`, unless specified as part of the `DefaultUserSettings` for
the domain.

Amazon SageMaker AI adds a security group to allow NFS traffic from Amazon SageMaker AI Studio. Therefore, the number of security groups that you can specify is one less than the
maximum number shown.

SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `32 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharingSettings`

Specifies options for sharing Amazon SageMaker AI Studio notebooks.

_Required_: No

_Type_: [SharingSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-sharingsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpaceStorageSettings`

The storage settings for a space.

SageMaker applies these settings only to private spaces that the user creates in the domain. SageMaker doesn't apply these settings to shared spaces.

_Required_: No

_Type_: [DefaultSpaceStorageSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-defaultspacestoragesettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StudioWebPortal`

Whether the user can access Studio. If this value is set to `DISABLED`, the
user cannot access Studio, even if that is the default experience for the domain.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StudioWebPortalSettings`

Studio settings. If these settings are applied on a user level, they take priority over
the settings applied on a domain level.

_Required_: No

_Type_: [StudioWebPortalSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-studiowebportalsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UnifiedStudioSettings

AWS::SageMaker::Endpoint
