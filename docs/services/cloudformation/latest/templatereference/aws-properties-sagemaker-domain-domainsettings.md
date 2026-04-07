This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain DomainSettings

A collection of settings that apply to the `SageMaker Domain`. These settings
are specified through the `CreateDomain` API call.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DockerSettings" : DockerSettings,
  "ExecutionRoleIdentityConfig" : String,
  "IpAddressType" : String,
  "RStudioServerProDomainSettings" : RStudioServerProDomainSettings,
  "SecurityGroupIds" : [ String, ... ],
  "UnifiedStudioSettings" : UnifiedStudioSettings
}

```

### YAML

```yaml

  DockerSettings:
    DockerSettings
  ExecutionRoleIdentityConfig: String
  IpAddressType: String
  RStudioServerProDomainSettings:
    RStudioServerProDomainSettings
  SecurityGroupIds:
    - String
  UnifiedStudioSettings:
    UnifiedStudioSettings

```

## Properties

`DockerSettings`

A collection of settings that configure the domain's Docker interaction.

_Required_: No

_Type_: [DockerSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-dockersettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleIdentityConfig`

The configuration for attaching a SageMaker AI user profile name to the execution
role as a [sts:SourceIdentity key](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_monitor.html).

_Required_: No

_Type_: String

_Allowed values_: `USER_PROFILE_NAME | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address type for the domain. Specify `ipv4` for IPv4-only
connectivity or `dualstack` for both IPv4 and IPv6 connectivity. When you
specify `dualstack`, the subnet must support IPv6 CIDR blocks. If not specified,
defaults to `ipv4`.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUALSTACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RStudioServerProDomainSettings`

A collection of settings that configure the `RStudioServerPro` Domain-level
app.

_Required_: No

_Type_: [RStudioServerProDomainSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-rstudioserverprodomainsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The security groups for the Amazon Virtual Private Cloud that the `Domain` uses for
communication between Domain-level apps and user apps.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `32 | 3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnifiedStudioSettings`

The settings that apply to an SageMaker AI domain when you use it in Amazon SageMaker Unified Studio.

_Required_: No

_Type_: [UnifiedStudioSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-domain-unifiedstudiosettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DockerSettings

EFSFileSystemConfig
