# ImageBuilder

Describes a virtual machine that is used to create an image.

## Contents

**Name**

The name of the image builder.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**AccessEndpoints**

The list of virtual private cloud (VPC) interface endpoint objects. Administrators can connect to the image builder only through the specified endpoints.

Type: Array of [AccessEndpoint](api-accessendpoint.md) objects

Array Members: Minimum number of 1 item. Maximum number of 4 items.

Required: No

**AppstreamAgentVersion**

The version of the WorkSpaces Applications agent that is currently being used by the image builder.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**Arn**

The ARN for the image builder.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**CreatedTime**

The time stamp when the image builder was created.

Type: Timestamp

Required: No

**Description**

The description to display.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**DisableIMDSV1**

Indicates whether Instance Metadata Service Version 1 (IMDSv1) is disabled for the image builder.

Type: Boolean

Required: No

**DisplayName**

The image builder name to display.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**DomainJoinInfo**

The name of the directory and organizational unit (OU) to use to join the image builder to a Microsoft Active Directory domain.

Type: [DomainJoinInfo](api-domainjoininfo.md) object

Required: No

**EnableDefaultInternetAccess**

Enables or disables default internet access for the image builder.

Type: Boolean

Required: No

**IamRoleArn**

The ARN of the IAM role that is applied to the image builder. To assume a role, the image builder calls the AWS Security Token Service (STS) `AssumeRole` API operation and passes the ARN of the role to use. The operation creates a new session with temporary credentials. WorkSpaces Applications retrieves the temporary credentials and creates the **appstream\_machine\_role** credential profile on the instance.

For more information, see [Using an IAM Role to Grant Permissions to Applications and Scripts Running on WorkSpaces Applications Streaming Instances](../../../../services/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.md) in the _Amazon WorkSpaces Applications Administration Guide_.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**ImageArn**

The ARN of the image from which this builder was created.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**ImageBuilderErrors**

The image builder errors.

Type: Array of [ResourceError](api-resourceerror.md) objects

Required: No

**InstanceType**

The instance type for the image builder. The following instance types are available:

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

Type: String

Length Constraints: Minimum length of 1.

Required: No

**LatestAppstreamAgentVersion**

Indicates whether the image builder is using the latest WorkSpaces Applications agent version or not.

Type: String

Valid Values: `TRUE | FALSE`

Required: No

**NetworkAccessConfiguration**

Describes the network details of the fleet or image builder instance.

Type: [NetworkAccessConfiguration](api-networkaccessconfiguration.md) object

Required: No

**Platform**

The operating system platform of the image builder.

Type: String

Valid Values: `WINDOWS | WINDOWS_SERVER_2016 | WINDOWS_SERVER_2019 | WINDOWS_SERVER_2022 | WINDOWS_SERVER_2025 | AMAZON_LINUX2 | RHEL8 | ROCKY_LINUX8 | UBUNTU_PRO_2404`

Required: No

**RootVolumeConfig**

The current configuration of the root volume for the image builder, including the storage size in GB.

Type: [VolumeConfig](api-volumeconfig.md) object

Required: No

**State**

The state of the image builder.

Type: String

Valid Values: `PENDING | UPDATING_AGENT | RUNNING | STOPPING | STOPPED | REBOOTING | SNAPSHOTTING | DELETING | FAILED | UPDATING | PENDING_QUALIFICATION | PENDING_SYNCING_APPS | SYNCING_APPS | PENDING_IMAGE_IMPORT`

Required: No

**StateChangeReason**

The reason why the last state change occurred.

Type: [ImageBuilderStateChangeReason](api-imagebuilderstatechangereason.md) object

Required: No

**VpcConfig**

The VPC configuration of the image builder.

Type: [VpcConfig](api-vpcconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/imagebuilder.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/imagebuilder.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/imagebuilder.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Image

ImageBuilderStateChangeReason

All content copied from https://docs.aws.amazon.com/.
