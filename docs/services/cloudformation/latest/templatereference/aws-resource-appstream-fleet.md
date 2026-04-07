This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Fleet

The `AWS::AppStream::Fleet` resource creates a fleet for Amazon WorkSpaces Applications. A fleet consists of streaming instances that run a specified image when using Always-On or On-Demand.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::Fleet",
  "Properties" : {
      "ComputeCapacity" : ComputeCapacity,
      "Description" : String,
      "DisableIMDSV1" : Boolean,
      "DisconnectTimeoutInSeconds" : Integer,
      "DisplayName" : String,
      "DomainJoinInfo" : DomainJoinInfo,
      "EnableDefaultInternetAccess" : Boolean,
      "FleetType" : String,
      "IamRoleArn" : String,
      "IdleDisconnectTimeoutInSeconds" : Integer,
      "ImageArn" : String,
      "ImageName" : String,
      "InstanceType" : String,
      "MaxConcurrentSessions" : Integer,
      "MaxSessionsPerInstance" : Integer,
      "MaxUserDurationInSeconds" : Integer,
      "Name" : String,
      "Platform" : String,
      "RootVolumeConfig" : VolumeConfig,
      "SessionScriptS3Location" : S3Location,
      "StreamView" : String,
      "Tags" : [ Tag, ... ],
      "UsbDeviceFilterStrings" : [ String, ... ],
      "VpcConfig" : VpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::Fleet
Properties:
  ComputeCapacity:
    ComputeCapacity
  Description: String
  DisableIMDSV1: Boolean
  DisconnectTimeoutInSeconds: Integer
  DisplayName: String
  DomainJoinInfo:
    DomainJoinInfo
  EnableDefaultInternetAccess: Boolean
  FleetType: String
  IamRoleArn: String
  IdleDisconnectTimeoutInSeconds: Integer
  ImageArn: String
  ImageName: String
  InstanceType: String
  MaxConcurrentSessions: Integer
  MaxSessionsPerInstance: Integer
  MaxUserDurationInSeconds: Integer
  Name: String
  Platform: String
  RootVolumeConfig:
    VolumeConfig
  SessionScriptS3Location:
    S3Location
  StreamView: String
  Tags:
    - Tag
  UsbDeviceFilterStrings:
    - String
  VpcConfig:
    VpcConfig

```

## Properties

`ComputeCapacity`

The desired capacity for the fleet. This is not allowed for Elastic fleets.

_Required_: No

_Type_: [ComputeCapacity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-fleet-computecapacity.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description to display.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableIMDSV1`

Indicates whether Instance Metadata Service Version 1 (IMDSv1) is disabled for the fleet.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisconnectTimeoutInSeconds`

The amount of time that a streaming session remains active after users disconnect. If users try to reconnect to the streaming session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new streaming instance.

Specify a value between 60 and 36000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The fleet name to display.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainJoinInfo`

The name of the directory and organizational unit (OU) to use to join the fleet to a Microsoft Active Directory domain. This is not allowed for Elastic fleets.

_Required_: No

_Type_: [DomainJoinInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-fleet-domainjoininfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableDefaultInternetAccess`

Enables or disables default internet access for the fleet.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FleetType`

The fleet type.

ALWAYS\_ON

Provides users with instant-on access to their apps.
You are charged for all running instances in your fleet, even if no users are streaming apps.

ON\_DEMAND

Provide users with access to applications after they connect, which takes one to two minutes.
You are charged for instance streaming when users are connected and a
small hourly fee for instances that are not streaming apps.

ELASTIC

The pool of streaming instances is managed by Amazon WorkSpaces Applications. When a
user selects their application or desktop to launch, they will start streaming
after the app block has been downloaded and mounted to a streaming
instance.

_Allowed Values_: `ALWAYS_ON` \| `ELASTIC` \| `ON_DEMAND`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IamRoleArn`

The ARN of the IAM role that is applied to the fleet. To assume a role, the fleet instance calls the AWS Security Token Service `AssumeRole` API operation and passes the ARN of the role to use. The operation creates a new session with temporary credentials. WorkSpaces Applications retrieves the temporary credentials and creates the **appstream\_machine\_role** credential profile on the instance.

For more information, see [Using an IAM Role to Grant Permissions to Applications and Scripts Running on WorkSpaces Applications Streaming Instances](https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html) in the _Amazon WorkSpaces Applications Administration Guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdleDisconnectTimeoutInSeconds`

The amount of time that users can be idle (inactive) before they are disconnected
from their streaming session and the `DisconnectTimeoutInSeconds` time
interval begins. Users are notified before they are disconnected due to inactivity. If
they try to reconnect to the streaming session before the time interval specified in
`DisconnectTimeoutInSeconds` elapses, they are connected to their
previous session. Users are considered idle when they stop providing keyboard or mouse
input during their streaming session. File uploads and downloads, audio in, audio out,
and pixels changing do not qualify as user activity. If users continue to be idle after
the time interval in `IdleDisconnectTimeoutInSeconds` elapses, they are
disconnected.

To prevent users from being disconnected due to inactivity, specify a value of 0. Otherwise, specify a value between 60 and 36000.

If you enable this feature, we recommend that you specify a value that corresponds exactly to a whole number of minutes (for example, 60, 120, and 180). If you don't do this, the value is rounded to the nearest minute. For example, if you specify a value of 70, users are disconnected after 1 minute of inactivity. If you specify a value that is at the midpoint between two different minutes, the value is rounded up. For example, if you specify a value of 90, users are disconnected after 2 minutes of inactivity.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageArn`

The ARN of the public, private, or shared image to use.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageName`

The name of the image used to create the fleet.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The instance type to use when launching fleet instances. The following instance types are available for non-Elastic fleets:

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

The following instance types are available for Elastic fleets:

- stream.standard.small

- stream.standard.medium

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrentSessions`

The maximum number of concurrent sessions that can be run on an Elastic fleet. This setting is
required for Elastic fleets, but is not used for other fleet types.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSessionsPerInstance`

Max number of user sessions on an instance. This is applicable only for multi-session fleets.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxUserDurationInSeconds`

The maximum amount of time that a streaming session can remain active, in seconds. If users are still connected to a streaming instance five minutes before this limit is reached, they are prompted to save any open documents before being disconnected. After this time elapses, the instance is terminated and replaced by a new instance.

Specify a value between 600 and 432000.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A unique name for the fleet.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Platform`

The platform of the fleet. Platform is a required setting for Elastic fleets, and is not used for other fleet types.

_Required_: No

_Type_: String

_Allowed values_: `WINDOWS | WINDOWS_SERVER_2016 | WINDOWS_SERVER_2019 | WINDOWS_SERVER_2022 | WINDOWS_SERVER_2025 | AMAZON_LINUX2 | RHEL8 | ROCKY_LINUX8 | UBUNTU_PRO_2404`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RootVolumeConfig`

The current configuration of the root volume for fleet instances, including the storage size in GB.

_Required_: No

_Type_: [VolumeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-fleet-volumeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionScriptS3Location`

The S3 location of the session scripts configuration zip file. This only applies to Elastic fleets.

_Required_: No

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-fleet-s3location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamView`

The WorkSpaces Applications view that is displayed to your users when they stream from the fleet. When `APP` is specified, only the windows of applications opened by users display. When `DESKTOP` is specified, the standard desktop that is provided by the operating system displays.

The default value is `APP`.

_Required_: No

_Type_: String

_Allowed values_: `APP | DESKTOP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-fleet-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsbDeviceFilterStrings`

The USB device filter strings that specify which USB devices a user can redirect to the fleet streaming session, when using the Windows native client. This is allowed but not required for Elastic fleets.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

The VPC configuration for the fleet. This is required for Elastic fleets, but not required for other fleet types.

_Required_: No

_Type_: [VpcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-fleet-vpcconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CreateFleet](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_CreateFleet.html) in the _Amazon WorkSpaces Applications API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Attribute

ComputeCapacity
