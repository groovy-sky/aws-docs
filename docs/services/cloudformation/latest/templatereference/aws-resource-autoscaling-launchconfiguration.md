This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::LaunchConfiguration

The `AWS::AutoScaling::LaunchConfiguration` resource specifies the launch
configuration that can be used by an Auto Scaling group to configure Amazon EC2 instances.

When you update the launch configuration for an Auto Scaling group, CloudFormation deletes
that resource and creates a new launch configuration with the updated properties and a new
name. Existing instances are not affected. To update existing instances when you update the
`AWS::AutoScaling::LaunchConfiguration` resource, you can specify an [UpdatePolicy\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) for the group. You can find sample update policies for rolling updates in
[Configure Amazon EC2\
Auto Scaling resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-ec2-auto-scaling.html).

###### Note

Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group
using either a [launch\
template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) or a launch configuration. We strongly recommend that you do not use
launch configurations. For more information, see [Launch configurations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html)
in the _Amazon EC2 Auto Scaling User Guide_.

For help migrating from launch configurations to launch templates, see [Migrate AWS CloudFormation stacks from launch configurations to launch\
templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.html) in the _Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AutoScaling::LaunchConfiguration",
  "Properties" : {
      "AssociatePublicIpAddress" : Boolean,
      "BlockDeviceMappings" : [ BlockDeviceMapping, ... ],
      "ClassicLinkVPCId" : String,
      "ClassicLinkVPCSecurityGroups" : [ String, ... ],
      "EbsOptimized" : Boolean,
      "IamInstanceProfile" : String,
      "ImageId" : String,
      "InstanceId" : String,
      "InstanceMonitoring" : Boolean,
      "InstanceType" : String,
      "KernelId" : String,
      "KeyName" : String,
      "LaunchConfigurationName" : String,
      "MetadataOptions" : MetadataOptions,
      "PlacementTenancy" : String,
      "RamDiskId" : String,
      "SecurityGroups" : [ String, ... ],
      "SpotPrice" : String,
      "UserData" : String
    }
}

```

### YAML

```yaml

Type: AWS::AutoScaling::LaunchConfiguration
Properties:
  AssociatePublicIpAddress: Boolean
  BlockDeviceMappings:
    - BlockDeviceMapping
  ClassicLinkVPCId: String
  ClassicLinkVPCSecurityGroups:
    - String
  EbsOptimized: Boolean
  IamInstanceProfile: String
  ImageId: String
  InstanceId: String
  InstanceMonitoring: Boolean
  InstanceType: String
  KernelId: String
  KeyName: String
  LaunchConfigurationName: String
  MetadataOptions:
    MetadataOptions
  PlacementTenancy: String
  RamDiskId: String
  SecurityGroups:
    - String
  SpotPrice: String
  UserData: String

```

## Properties

`AssociatePublicIpAddress`

Specifies whether to assign a public IPv4 address to the group's instances. If the
instance is launched into a default subnet, the default is to assign a public IPv4
address, unless you disabled the option to assign a public IPv4 address on the subnet.
If the instance is launched into a nondefault subnet, the default is not to assign a
public IPv4 address, unless you enabled the option to assign a public IPv4 address on
the subnet.

If you specify `true`, each instance in the Auto Scaling group receives a unique
public IPv4 address. For more information, see [Provide network connectivity for\
your Auto Scaling instances using Amazon VPC](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html) in the
_Amazon EC2 Auto Scaling User Guide_.

If you specify this property, you must specify at least one subnet for
`VPCZoneIdentifier` when you create your group.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BlockDeviceMappings`

The block device mapping entries that define the block devices to attach to the
instances at launch. By default, the block devices specified in the block device mapping
for the AMI are used. For more information, see [Block device\
mappings](../../../ec2/latest/userguide/block-device-mapping-concepts.md) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: Array of [BlockDeviceMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscaling-launchconfiguration-blockdevicemapping.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClassicLinkVPCId`

Available for backward compatibility.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClassicLinkVPCSecurityGroups`

Available for backward compatibility.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsOptimized`

Specifies whether the launch configuration is optimized for EBS I/O
( `true`) or not ( `false`). The optimization provides dedicated
throughput to Amazon EBS and an optimized configuration stack to provide optimal I/O
performance. This optimization is not available with all instance types. Additional fees
are incurred when you enable EBS optimization for an instance type that is not
EBS-optimized by default. For more information, see [Amazon EBS-optimized instances](../../../ec2/latest/userguide/ebs-optimized.md)
in the _Amazon EC2 User Guide_.

The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IamInstanceProfile`

The name or the Amazon Resource Name (ARN) of the instance profile associated with the
IAM role for the instance. The instance profile contains the IAM role. For more
information, see [IAM role for applications that run\
on Amazon EC2 instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/us-iam-role.html) in the _Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageId`

The ID of the Amazon Machine Image (AMI) that was assigned during registration. For
more information, see [Find a Linux AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html) in the
_Amazon EC2 User Guide_.

If you specify `InstanceId`, an `ImageId` is not
required.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceId`

The ID of the Amazon EC2 instance to use to create the launch configuration. When you use
an instance to create a launch configuration, all properties are derived from the instance
with the exception of `BlockDeviceMapping` and
`AssociatePublicIpAddress`. You can override any properties from the instance by
specifying them in the launch configuration.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceMonitoring`

Controls whether instances in this group are launched with detailed
( `true`) or basic ( `false`) monitoring.

The default value is `true` (enabled).

###### Important

When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and
your account is charged a fee. When you disable detailed monitoring, CloudWatch generates
metrics every 5 minutes. For more information, see [Configure\
monitoring for Auto Scaling instances](https://docs.aws.amazon.com/autoscaling/latest/userguide/enable-as-instance-metrics.html) in the
_Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceType`

Specifies the instance type of the EC2 instance. For information about available
instance types, see [Available\
instance types](../../../ec2/latest/userguide/instance-types.md#AvailableInstanceTypes) in the _Amazon EC2 User Guide_.

If you specify `InstanceId`, an `InstanceType` is not
required.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KernelId`

The ID of the kernel associated with the AMI.

###### Note

We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
information, see [User provided\
kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyName`

The name of the key pair. For more information, see [Amazon EC2 key pairs and Amazon EC2\
instances](../../../ec2/latest/userguide/ec2-key-pairs.md) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LaunchConfigurationName`

The name of the launch configuration. This name must be unique per Region per
account.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetadataOptions`

The metadata options for the instances. For more information, see [Configure the instance metadata options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds) in the
_Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: [MetadataOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-autoscaling-launchconfiguration-metadataoptions.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlacementTenancy`

The tenancy of the instance, either `default` or `dedicated`. An
instance with `dedicated` tenancy runs on isolated, single-tenant hardware
and can only be launched into a VPC. To launch dedicated instances into a shared tenancy
VPC (a VPC with the instance placement tenancy attribute set to `default`),
you must set the value of this property to `dedicated`.

If you specify `PlacementTenancy`, you must specify at least one subnet for
`VPCZoneIdentifier` when you create your group.

Valid values: `default` \| `dedicated`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RamDiskId`

The ID of the RAM disk to select.

###### Note

We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
information, see [User provided\
kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedKernels.html) in the _Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

A list that contains the security groups to assign to the instances in the Auto Scaling
group. The list can contain both the IDs of existing security groups and references to [SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.

For more information, see [Control traffic to resources using\
security groups](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) in the _Amazon Virtual Private Cloud User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpotPrice`

The maximum hourly price to be paid for any Spot Instance launched to fulfill the
request. Spot Instances are launched when the price you specify exceeds the current Spot
price. For more information, see [Request Spot\
Instances for fault-tolerant and flexible applications](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-template-spot-instances.html) in the
_Amazon EC2 Auto Scaling User Guide_.

Valid Range: Minimum value of 0.001

###### Note

When you change your maximum price by creating a new launch configuration, running
instances will continue to run as long as the maximum price for those running
instances is higher than the current Spot price.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserData`

The Base64-encoded user data to make available to the launched EC2 instances. For more
information, see [Instance metadata and user\
data](../../../ec2/latest/userguide/ec2-instance-metadata.md) in the _Amazon EC2 User Guide for Linux Instances_.

_Required_: No

_Type_: String

_Maximum_: `21847`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:
`mystack-mylaunchconfig-1DDYF1E3B3I`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Remarks

CloudFormation marks the Auto Scaling group as successful (by setting its status to
CREATE\_COMPLETE) when its desired capacity is reached. However, if `SpotPrice` is
set in the launch configuration, then desired capacity is not used as a criteria for
success. Whether your request is fulfilled depends on Spot Instance capacity and your
maximum price. If the current Spot price is less than your specified maximum price, Amazon
EC2 Auto Scaling uses the desired capacity as the target capacity for the group. If the
request for Spot Instances is unsuccessful, it keeps trying.

If your Auto Scaling instances receive a public IPv4 address and are launched in a VPC
that is defined in the same stack template, you must use the [DependsOn\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html).

## Examples

The following examples create launch configurations that can be used by an Auto Scaling
group to configure Amazon EC2 instances.

- [Amazon EBS-backed AMI and defined block device mappings](#aws-resource-autoscaling-launchconfiguration--examples--Amazon_EBS-backed_AMI_and_defined_block_device_mappings)

- [Instance store-backed AMI with Spot price and IAM role](#aws-resource-autoscaling-launchconfiguration--examples--Instance_store-backed_AMI_with_Spot_price_and_IAM_role)

- [Provisioned IOPS EBS-optimized volume with key-pair name and user data](#aws-resource-autoscaling-launchconfiguration--examples--Provisioned_IOPS_EBS-optimized_volume_with_key-pair_name_and_user_data)

### Amazon EBS-backed AMI and defined block device mappings

This example shows a launch configuration with a `BlockDeviceMappings`
property that lists two devices: a 30 gigabyte EBS root volume mapped to /dev/sda1 and a
100 gigabyte EBS volume mapped to /dev/sdm. The /dev/sdm volume uses the default EBS
volume type based on the region and is not deleted when terminating the instance it is
attached to.

CloudFormation supports parameters from the AWS Systems Manager
Parameter Store. In this example, the `ImageId` property references the latest
Amazon Linux 2 AMI (EBS-backed image) from the Parameter Store. For more information, see
[AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) in the _AWS Systems Manager User Guide_ and the blog post [Query for the latest Amazon Linux AMI IDs using AWS Systems Manager\
Parameter Store](https://aws.amazon.com/blogs/compute/query-for-the-latest-amazon-linux-ami-ids-using-aws-systems-manager-parameter-store) on the AWS Compute Blog.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters": {
    "LatestAmiId": {
      "Description": "Region specific image from the Parameter Store",
      "Type": "AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>",
      "Default": "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
    },
    "InstanceType": {
      "Description": "Amazon EC2 instance type for the instances",
      "Type": "String",
      "AllowedValues": [
        "t3.micro",
        "t3.small",
        "t3.medium"
      ],
      "Default": "t3.micro"
    }
  },
  "Resources":{
    "myLaunchConfig":{
      "Type":"AWS::AutoScaling::LaunchConfiguration",
      "Properties":{
        "ImageId":{ "Ref":"LatestAmiId" },
        "SecurityGroups":[ { "Ref":"myEC2SecurityGroup" } ],
        "InstanceType":{ "Ref":"InstanceType" },
        "BlockDeviceMappings":[
          {
            "DeviceName":"/dev/sda1",
            "Ebs":{
              "VolumeSize":"30",
              "VolumeType":"gp3"
            }
          },
          {
            "DeviceName":"/dev/sdm",
            "Ebs":{
              "VolumeSize":"100",
              "DeleteOnTermination":"false"
            }
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: 2010-09-09
Parameters:
  LatestAmiId:
    Description: Region specific image from the Parameter Store
    Type: 'AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>'
    Default: '/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2'
  InstanceType:
    Description: Amazon EC2 instance type for the instances
    Type: String
    AllowedValues:
      - t3.micro
      - t3.small
      - t3.medium
    Default: t3.micro
Resources:
  myLaunchConfig:
    Type: AWS::AutoScaling::LaunchConfiguration
    Properties:
      ImageId: !Ref LatestAmiId
      SecurityGroups:
        - !Ref myEC2SecurityGroup
      InstanceType:
        !Ref InstanceType
      BlockDeviceMappings:
        - DeviceName: /dev/sda1
          Ebs:
            VolumeSize: '30'
            VolumeType: gp3
        - DeviceName: /dev/sdm
          Ebs:
            VolumeSize: '100'
            DeleteOnTermination: false
```

### Instance store-backed AMI with Spot price and IAM role

This example shows a launch configuration that launches Spot Instances in the Auto
Scaling group. This launch configuration will only be active if the current Spot price is
less than the price in the template specification (0.045). It also demonstrates a launch
configuration that uses the `IamInstanceProfile` property. For an example of a
full template, including the definition of, and further references from the [InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html) object referenced here as `RootInstanceProfile`, see
[auto\_scaling\_with\_instance\_profile.template](https://s3.amazonaws.com/cloudformation-templates-us-east-1/auto_scaling_with_instance_profile.template).

In this example, the `ImageId` property references the latest Amazon Linux
AMI (instance store/S3-backed image) from the Parameter Store. The BlockDeviceMappings
property lists a virtual device `ephemeral0` mapped to /dev/sdc.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters": {
    "LatestAmiId": {
      "Description": "Region specific image from the Parameter Store",
      "Type": "AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>",
      "Default": "/aws/service/ami-amazon-linux-latest/amzn-ami-hvm-x86_64-s3"
    }
  },
  "Resources":{
    "myLaunchConfig":{
      "Type":"AWS::AutoScaling::LaunchConfiguration",
      "Properties":{
        "ImageId":{ "Ref":"LatestAmiId" },
        "SecurityGroups":[ { "Ref":"myEC2SecurityGroup" } ],
        "InstanceType":"m3.medium",
        "SpotPrice":"0.045",
        "IamInstanceProfile":{ "Ref":"RootInstanceProfile" },
        "BlockDeviceMappings": [
          {
            "DeviceName": "/dev/sdc",
            "VirtualName": "ephemeral0"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: 2010-09-09
Parameters:
  LatestAmiId:
    Description: Region specific image from the Parameter Store
    Type: 'AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>'
    Default: '/aws/service/ami-amazon-linux-latest/amzn-ami-hvm-x86_64-s3'
Resources:
  myLaunchConfig:
    Type: AWS::AutoScaling::LaunchConfiguration
    Properties:
      ImageId: !Ref LatestAmiId
      SecurityGroups:
        - !Ref myEC2SecurityGroup
      InstanceType: m3.medium
      SpotPrice: '0.045'
      IamInstanceProfile: !Ref RootInstanceProfile
      BlockDeviceMappings:
      - DeviceName: /dev/sdc
        VirtualName: ephemeral0
```

### Provisioned IOPS EBS-optimized volume with key-pair name and user data

This example demonstrates a launch configuration that configures the
`EbsOptmized` property to launch instances with a provisioned IOPS
EBS-optimized volume. This can increase the performance of your EBS-backed
instances.

###### Note

For instances that are not EBS–optimized by default, you must enable EBS
optimization to achieve the level of performance described in the [Amazon\
EBS-optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) documentation in the _Amazon Elastic Compute_
_Cloud User Guide_. For current generation instance types, EBS-optimization
is enabled by default at no additional cost. Enabling EBS optimization for a previous
generation instance type that is not EBS-optimized by default incurs additional
fees.

When you use a launch configuration such as this one, your `m1.large`
instances will contain optimized EBS root volumes with the provisioned IOPS settings that
you specified in the AMI. Because you cannot specify the IOPS settings in a launch
configuration, the AMI must be configured with a block device mapping that specifies the
desired number of IOPS. The following are key attributes of this EBS-optimized instance
configuration:

- An instance type of `m1.large` or greater. This is required for EBS
optimization. This optimization is only available for certain instance types and
sizes.

- An EBS-backed AMI with a volume type of `io1` and the number of IOPS
you want to provision for the volume.

- The size of the EBS volume must accommodate the IOPS you need. There is a 50:1
ratio between IOPS and Gibibytes (GiB) of storage.

For more information about IOPS performance with provisioned IOPS volumes, see [Provisioned IOPS\
SSD ( `io1` and `io2`) volumes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html#EBSVolumeTypes_piops) in the _Amazon_
_Elastic Compute Cloud User Guide_.

For more performance tips, see [Amazon EBS volume performance on Linux\
instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSPerformance.html) in the _Amazon Elastic Compute Cloud User Guide_.

#### JSON

```json

{
  "Resources":{
    "myLaunchConfig":{
      "Type":"AWS::AutoScaling::LaunchConfiguration",
      "Properties":{
        "ImageId":"ami-02354e95b3example",
        "SecurityGroups":[ { "Ref":"myEC2SecurityGroup" }, "myExistingEC2SecurityGroup" ],
        "InstanceType":"m1.large",
        "UserData": {"Fn::Base64": {"Fn::Join": ["", [
          "#!/bin/bash -x\n",
          "/opt/aws/bin/cfn-signal -e $? --stack ", {"Ref": "AWS::StackName"}, " --resource myASG ", " --region ", {"Ref": "AWS::Region"}, "\n"
        ]]}},
        "EbsOptimized":"true"
      }
    }
  }
}
```

#### YAML

```yaml

---
Resources:
  myLaunchConfig:
    Type: AWS::AutoScaling::LaunchConfiguration
    Properties:
      ImageId: ami-02354e95b3example
      SecurityGroups:
        - !Ref myEC2SecurityGroup
        - myExistingEC2SecurityGroup
      InstanceType: m1.large
      UserData: !Base64 |
        #!/bin/bash -x
        yum install -y aws-cfn-bootstrap
        /opt/aws/bin/cfn-signal -e $? --stack !Ref 'AWS::StackName' --resource myASG --region !Ref 'AWS::Region'
      EbsOptimized: true
```

## See also

- You can find additional useful snippets in the following sections of the _AWS CloudFormation User Guide_:

- For examples of Auto Scaling groups, see [Configure\
Amazon EC2 Auto Scaling resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-ec2-auto-scaling.html).

- For examples of launch templates, see [Create\
launch templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-ec2-launch-templates.html).

- [Migrate AWS CloudFormation stacks from launch configurations to\
launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.html) in the _Amazon EC2 Auto Scaling User_
_Guide_

- [CreateLaunchConfiguration](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateLaunchConfiguration.html) in the _Amazon EC2 Auto Scaling API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VCpuCountRequest

BlockDevice
