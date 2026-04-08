# CreateLaunchConfiguration

Creates a launch configuration.

If you exceed your maximum limit of launch configurations, the call fails. To query
this limit, call the [DescribeAccountLimits](api-describeaccountlimits.md) API.
For information about updating this limit, see [Quotas for \
Amazon EC2 Auto Scaling](../../../../services/autoscaling/ec2/userguide/ec2-auto-scaling-quotas.md) in the _Amazon EC2 Auto Scaling User Guide_.

For more information, see [Launch\
configurations](../../../../services/autoscaling/ec2/userguide/launch-configurations.md) in the _Amazon EC2 Auto Scaling User Guide_.

###### Note

Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group using either a
launch template or a launch configuration. We strongly recommend that you do not use
launch configurations. They do not provide full functionality for Amazon EC2 Auto Scaling or Amazon EC2.
For information about using launch templates, see [Launch templates](../../../../services/autoscaling/ec2/userguide/launch-templates.md) in the _Amazon EC2 Auto Scaling User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**AssociatePublicIpAddress**

Specifies whether to assign a public IPv4 address to the group's instances. If the
instance is launched into a default subnet, the default is to assign a public IPv4
address, unless you disabled the option to assign a public IPv4 address on the subnet.
If the instance is launched into a nondefault subnet, the default is not to assign a
public IPv4 address, unless you enabled the option to assign a public IPv4 address on
the subnet.

If you specify `true`, each instance in the Auto Scaling group receives a unique
public IPv4 address. For more information, see [Provide network connectivity for\
your Auto Scaling instances using Amazon VPC](../../../../services/autoscaling/ec2/userguide/asg-in-vpc.md) in the
_Amazon EC2 Auto Scaling User Guide_.

If you specify this property, you must specify at least one subnet for
`VPCZoneIdentifier` when you create your group.

Type: Boolean

Required: No

**BlockDeviceMappings.member.N**

The block device mapping entries that define the block devices to attach to the
instances at launch. By default, the block devices specified in the block device mapping
for the AMI are used. For more information, see [Block device\
mappings](../../../../services/ec2/latest/userguide/block-device-mapping-concepts.md) in the _Amazon EC2 User Guide_.

Type: Array of [BlockDeviceMapping](api-blockdevicemapping.md) objects

Required: No

**ClassicLinkVPCId**

Available for backward compatibility.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**ClassicLinkVPCSecurityGroups.member.N**

Available for backward compatibility.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**EbsOptimized**

Specifies whether the launch configuration is optimized for EBS I/O
( `true`) or not ( `false`). The optimization provides dedicated
throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O
performance. This optimization is not available with all instance types. Additional fees
are incurred when you enable EBS optimization for an instance type that is not
EBS-optimized by default. For more information, see [Amazon EBS-optimized instances](../../../../services/ec2/latest/userguide/ebs-optimized.md)
in the _Amazon EC2 User Guide_.

The default value is `false`.

Type: Boolean

Required: No

**IamInstanceProfile**

The name or the Amazon Resource Name (ARN) of the instance profile associated with the
IAM role for the instance. The instance profile contains the IAM role. For more
information, see [IAM role for applications that run\
on Amazon EC2 instances](../../../../services/autoscaling/ec2/userguide/us-iam-role.md) in the _Amazon EC2 Auto Scaling User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**ImageId**

The ID of the Amazon Machine Image (AMI) that was assigned during registration. For
more information, see [Find a Linux AMI](../../../../services/ec2/latest/userguide/finding-an-ami.md) in the
_Amazon EC2 User Guide_.

If you specify `InstanceId`, an `ImageId` is not
required.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**InstanceId**

The ID of the instance to use to create the launch configuration. The new launch
configuration derives attributes from the instance, except for the block device
mapping.

To create a launch configuration with a block device mapping or override any other
instance attributes, specify them as part of the same request.

For more information, see [Create a launch\
configuration](../../../../services/autoscaling/ec2/userguide/create-launch-config.md) in the _Amazon EC2 Auto Scaling User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 19.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**InstanceMonitoring**

Controls whether instances in this group are launched with detailed
( `true`) or basic ( `false`) monitoring.

The default value is `true` (enabled).

###### Important

When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and
your account is charged a fee. When you disable detailed monitoring, CloudWatch generates
metrics every 5 minutes. For more information, see [Configure\
monitoring for Auto Scaling instances](../../../../services/autoscaling/latest/userguide/enable-as-instance-metrics.md) in the
_Amazon EC2 Auto Scaling User Guide_.

Type: [InstanceMonitoring](api-instancemonitoring.md) object

Required: No

**InstanceType**

Specifies the instance type of the EC2 instance. For information about available
instance types, see [Available\
instance types](../../../../services/ec2/latest/userguide/instance-types.md#AvailableInstanceTypes) in the _Amazon EC2 User Guide_.

If you specify `InstanceId`, an `InstanceType` is not
required.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**KernelId**

The ID of the kernel associated with the AMI.

###### Note

We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
information, see [User provided\
kernels](../../../../services/ec2/latest/userguide/userprovidedkernels.md) in the _Amazon EC2 User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**KeyName**

The name of the key pair. For more information, see [Amazon EC2 key pairs and Amazon EC2\
instances](../../../../services/ec2/latest/userguide/ec2-key-pairs.md) in the _Amazon EC2 User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**LaunchConfigurationName**

The name of the launch configuration. This name must be unique per Region per
account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: Yes

**MetadataOptions**

The metadata options for the instances. For more information, see [Configure the instance metadata options](../../../../services/autoscaling/ec2/userguide/create-launch-config.md#launch-configurations-imds) in the
_Amazon EC2 Auto Scaling User Guide_.

Type: [InstanceMetadataOptions](api-instancemetadataoptions.md) object

Required: No

**PlacementTenancy**

The tenancy of the instance, either `default` or `dedicated`. An
instance with `dedicated` tenancy runs on isolated, single-tenant hardware
and can only be launched into a VPC. To launch dedicated instances into a shared tenancy
VPC (a VPC with the instance placement tenancy attribute set to `default`),
you must set the value of this property to `dedicated`.

If you specify `PlacementTenancy`, you must specify at least one subnet for
`VPCZoneIdentifier` when you create your group.

Valid values: `default` \| `dedicated`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**RamdiskId**

The ID of the RAM disk to select.

###### Note

We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
information, see [User provided\
kernels](../../../../services/ec2/latest/userguide/userprovidedkernels.md) in the _Amazon EC2 User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**SecurityGroups.member.N**

A list that contains the security group IDs to assign to the instances in the Auto Scaling
group. For more information, see [Control traffic to your AWS\
resources using security groups](../../../../services/vpc/latest/userguide/vpc-security-groups.md) in the _Amazon Virtual Private_
_Cloud User Guide_.

Type: Array of strings

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**SpotPrice**

The maximum hourly price to be paid for any Spot Instance launched to fulfill the
request. Spot Instances are launched when the price you specify exceeds the current Spot
price. For more information, see [Request Spot\
Instances for fault-tolerant and flexible applications](../../../../services/autoscaling/ec2/userguide/launch-template-spot-instances.md) in the
_Amazon EC2 Auto Scaling User Guide_.

Valid Range: Minimum value of 0.001

###### Note

When you change your maximum price by creating a new launch configuration, running
instances will continue to run as long as the maximum price for those running
instances is higher than the current Spot price.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**UserData**

The user data to make available to the launched EC2 instances. For more information,
see [Instance metadata and user data](../../../../services/ec2/latest/userguide/ec2-instance-metadata.md) (Linux) and [Instance metadata and\
user data](../../../../services/ec2/latest/windowsguide/ec2-instance-metadata.md) (Windows). If you are using a command line tool, base64-encoding
is performed for you, and you can load the text from a file. Otherwise, you must provide
base64-encoded text. User data is limited to 16 KB.

Type: String

Length Constraints: Maximum length of 21847.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AlreadyExists**

You already have an Auto Scaling group or launch configuration with this name.

**message**

HTTP Status Code: 400

**LimitExceeded**

You have already reached a limit for your Amazon EC2 Auto Scaling resources
(for example, Auto Scaling groups, launch configurations, or lifecycle hooks). For more
information, see [DescribeAccountLimits](api-describeaccountlimits.md).

**message**

HTTP Status Code: 400

**ResourceContention**

You already have a pending update to an Amazon EC2 Auto Scaling resource (for example, an Auto Scaling group,
instance, or load balancer).

**message**

HTTP Status Code: 500

## Examples

### Example

This example illustrates one usage of CreateLaunchConfiguration.

#### Sample Request

```

https://autoscaling.amazonaws.com/?Action=CreateLaunchConfiguration
&LaunchConfigurationName=my-lc
&ImageId=ami-0abcdef1234567890
&InstanceType=t2.micro
&SecurityGroups.member.1=sg-1a2b3c4d
&Version=2011-01-01
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/autoscaling-2011-01-01/createlaunchconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/autoscaling-2011-01-01/createlaunchconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateAutoScalingGroup

CreateOrUpdateTags
