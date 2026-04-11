This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::ImageBuilder

The `AWS::AppStream::ImageBuilder` resource creates an image builder for Amazon WorkSpaces Applications. An image builder is a virtual machine that is used to create an image.

The initial state of the image builder is `PENDING`. When it is ready, the state is `RUNNING`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::ImageBuilder",
  "Properties" : {
      "AccessEndpoints" : [ AccessEndpoint, ... ],
      "AppstreamAgentVersion" : String,
      "Description" : String,
      "DisplayName" : String,
      "DomainJoinInfo" : DomainJoinInfo,
      "EnableDefaultInternetAccess" : Boolean,
      "IamRoleArn" : String,
      "ImageArn" : String,
      "ImageName" : String,
      "InstanceType" : String,
      "Name" : String,
      "RootVolumeConfig" : VolumeConfig,
      "SoftwaresToInstall" : [ String, ... ],
      "SoftwaresToUninstall" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VpcConfig" : VpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::ImageBuilder
Properties:
  AccessEndpoints:
    - AccessEndpoint
  AppstreamAgentVersion: String
  Description: String
  DisplayName: String
  DomainJoinInfo:
    DomainJoinInfo
  EnableDefaultInternetAccess: Boolean
  IamRoleArn: String
  ImageArn: String
  ImageName: String
  InstanceType: String
  Name: String
  RootVolumeConfig:
    VolumeConfig
  SoftwaresToInstall:
    - String
  SoftwaresToUninstall:
    - String
  Tags:
    - Tag
  VpcConfig:
    VpcConfig

```

## Properties

`AccessEndpoints`

The list of virtual private cloud (VPC) interface endpoint objects. Administrators can connect to the image builder only through the specified endpoints.

_Required_: No

_Type_: Array of [AccessEndpoint](aws-properties-appstream-imagebuilder-accessendpoint.md)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppstreamAgentVersion`

The version of the WorkSpaces Applications agent to use for this image builder. To use the latest version of the WorkSpaces Applications agent, specify \[LATEST\].

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description to display.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The image builder name to display.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainJoinInfo`

The name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain.

_Required_: No

_Type_: [DomainJoinInfo](aws-properties-appstream-imagebuilder-domainjoininfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableDefaultInternetAccess`

Enables or disables default internet access for the image builder.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleArn`

The ARN of the IAM role that is applied to the image builder. To assume a role, the image builder calls the AWS Security Token Service `AssumeRole` API operation and passes the ARN of the role to use. The operation creates a new session with temporary credentials. WorkSpaces Applications retrieves the temporary credentials and creates the **appstream\_machine\_role** credential profile on the instance.

For more information, see [Using an IAM Role to Grant Permissions to Applications and Scripts Running on WorkSpaces Applications Streaming Instances](../../../appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.md) in the _Amazon WorkSpaces Applications Administration Guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageArn`

The ARN of the public, private, or shared image to use.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageName`

The name of the image used to create the image builder.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The instance type to use when launching the image builder. The following instance types are available:

- stream.standard.small

- stream.standard.medium

- stream.standard.large

- stream.compute.large

- stream.compute.xlarge

- stream.compute.2xlarge

- stream.compute.4xlarge

- stream.compute.8xlarge

- stream.memory.large

- stream.memory.xlarge

- stream.memory.2xlarge

- stream.memory.4xlarge

- stream.memory.8xlarge

- stream.memory.z1d.large

- stream.memory.z1d.xlarge

- stream.memory.z1d.2xlarge

- stream.memory.z1d.3xlarge

- stream.memory.z1d.6xlarge

- stream.memory.z1d.12xlarge

- stream.graphics-design.large

- stream.graphics-design.xlarge

- stream.graphics-design.2xlarge

- stream.graphics-design.4xlarge

- stream.graphics.g4dn.xlarge

- stream.graphics.g4dn.2xlarge

- stream.graphics.g4dn.4xlarge

- stream.graphics.g4dn.8xlarge

- stream.graphics.g4dn.12xlarge

- stream.graphics.g4dn.16xlarge

- stream.graphics.g5.xlarge

- stream.graphics.g5.2xlarge

- stream.graphics.g5.4xlarge

- stream.graphics.g5.8xlarge

- stream.graphics.g5.16xlarge

- stream.graphics.g5.12xlarge

- stream.graphics.g5.24xlarge

- stream.graphics.g6.xlarge

- stream.graphics.g6.2xlarge

- stream.graphics.g6.4xlarge

- stream.graphics.g6.8xlarge

- stream.graphics.g6.16xlarge

- stream.graphics.g6.12xlarge

- stream.graphics.g6.24xlarge

- stream.graphics.gr6.4xlarge

- stream.graphics.gr6.8xlarge

- stream.graphics.g6f.large

- stream.graphics.g6f.xlarge

- stream.graphics.g6f.2xlarge

- stream.graphics.g6f.4xlarge

- stream.graphics.gr6f.4xlarge

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A unique name for the image builder.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RootVolumeConfig`

The current configuration of the root volume for the image builder, including the storage size in GB.

_Required_: No

_Type_: [VolumeConfig](aws-properties-appstream-imagebuilder-volumeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SoftwaresToInstall`

The list of license included applications to install on the image builder during creation.

Possible values include the following:

- Microsoft\_Office\_2021\_LTSC\_Professional\_Plus\_32Bit

- Microsoft\_Office\_2021\_LTSC\_Professional\_Plus\_64Bit

- Microsoft\_Office\_2024\_LTSC\_Professional\_Plus\_32Bit

- Microsoft\_Office\_2024\_LTSC\_Professional\_Plus\_64Bit

- Microsoft\_Visio\_2021\_LTSC\_Professional\_32Bit

- Microsoft\_Visio\_2021\_LTSC\_Professional\_64Bit

- Microsoft\_Visio\_2024\_LTSC\_Professional\_32Bit

- Microsoft\_Visio\_2024\_LTSC\_Professional\_64Bit

- Microsoft\_Project\_2021\_Professional\_32Bit

- Microsoft\_Project\_2021\_Professional\_64Bit

- Microsoft\_Project\_2024\_Professional\_32Bit

- Microsoft\_Project\_2024\_Professional\_64Bit

- Microsoft\_Office\_2021\_LTSC\_Standard\_32Bit

- Microsoft\_Office\_2021\_LTSC\_Standard\_64Bit

- Microsoft\_Office\_2024\_LTSC\_Standard\_32Bit

- Microsoft\_Office\_2024\_LTSC\_Standard\_64Bit

- Microsoft\_Visio\_2021\_LTSC\_Standard\_32Bit

- Microsoft\_Visio\_2021\_LTSC\_Standard\_64Bit

- Microsoft\_Visio\_2024\_LTSC\_Standard\_32Bit

- Microsoft\_Visio\_2024\_LTSC\_Standard\_64Bit

- Microsoft\_Project\_2021\_Standard\_32Bit

- Microsoft\_Project\_2021\_Standard\_64Bit

- Microsoft\_Project\_2024\_Standard\_32Bit

- Microsoft\_Project\_2024\_Standard\_64Bit

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SoftwaresToUninstall`

The list of license included applications to uninstall from the image builder during creation.

Possible values include the following:

- Microsoft\_Office\_2021\_LTSC\_Professional\_Plus\_32Bit

- Microsoft\_Office\_2021\_LTSC\_Professional\_Plus\_64Bit

- Microsoft\_Office\_2024\_LTSC\_Professional\_Plus\_32Bit

- Microsoft\_Office\_2024\_LTSC\_Professional\_Plus\_64Bit

- Microsoft\_Visio\_2021\_LTSC\_Professional\_32Bit

- Microsoft\_Visio\_2021\_LTSC\_Professional\_64Bit

- Microsoft\_Visio\_2024\_LTSC\_Professional\_32Bit

- Microsoft\_Visio\_2024\_LTSC\_Professional\_64Bit

- Microsoft\_Project\_2021\_Professional\_32Bit

- Microsoft\_Project\_2021\_Professional\_64Bit

- Microsoft\_Project\_2024\_Professional\_32Bit

- Microsoft\_Project\_2024\_Professional\_64Bit

- Microsoft\_Office\_2021\_LTSC\_Standard\_32Bit

- Microsoft\_Office\_2021\_LTSC\_Standard\_64Bit

- Microsoft\_Office\_2024\_LTSC\_Standard\_32Bit

- Microsoft\_Office\_2024\_LTSC\_Standard\_64Bit

- Microsoft\_Visio\_2021\_LTSC\_Standard\_32Bit

- Microsoft\_Visio\_2021\_LTSC\_Standard\_64Bit

- Microsoft\_Visio\_2024\_LTSC\_Standard\_32Bit

- Microsoft\_Visio\_2024\_LTSC\_Standard\_64Bit

- Microsoft\_Project\_2021\_Standard\_32Bit

- Microsoft\_Project\_2021\_Standard\_64Bit

- Microsoft\_Project\_2024\_Standard\_32Bit

- Microsoft\_Project\_2024\_Standard\_64Bit

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs.

_Required_: No

_Type_: Array of [Tag](aws-properties-appstream-imagebuilder-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

The VPC configuration for the image builder. You can specify only one subnet.

_Required_: No

_Type_: [VpcConfig](aws-properties-appstream-imagebuilder-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`StreamingUrl`

The URL to start an image builder streaming session, returned as a string.

## See also

- [CreateImageBuilder](../../../../reference/appstream2/latest/apireference/api-createimagebuilder.md) in the _Amazon WorkSpaces Applications API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

AccessEndpoint

All content copied from https://docs.aws.amazon.com/.
