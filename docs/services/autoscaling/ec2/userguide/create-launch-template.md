# Create a launch template for an Auto Scaling group

Before you can create an Auto Scaling group using a launch template, you must create a launch
template that contains the configuration information to launch an instance, including
the ID of the Amazon Machine Image (AMI).

To create new launch templates, use the following procedures.

###### Contents

- [Create your launch template (console)](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html#create-launch-template-for-auto-scaling)

- [Change the default network interface settings (console)](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html#change-network-interface)

- [Modify the storage configuration (console)](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html#modify-storage-configuration)

- [Create a launch template from an existing instance (console)](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html#create-launch-template-from-instance)

- [Related resources](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html#create-launch-template-related-resources)

- [Limitations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html#create-launch-template-limitations)

###### Important

Launch template parameters are not fully validated when you create the launch
template. If you specify incorrect values for parameters, or if you do not use
supported parameter combinations, no instances can launch using this launch
template. Be sure to specify the correct values for the parameters and use supported
parameter combinations. For example, to launch instances with an Arm-based AWS
Graviton or Graviton2 AMI, you must specify an Arm-compatible instance type. For
more information, see [Launch template\
restrictions](../../../ec2/latest/userguide/launch-template-restrictions.md) in the _Amazon EC2 User Guide_.

## Create your launch template (console)

The following steps describe how to configure a basic launch template:

- Specify the Amazon Machine Image (AMI) from which to launch the
instances.

- Choose an instance type that is compatible with the AMI that you specify.

- Specify the key pair to use when connecting to instances, for example,
using SSH.

- Add one or more security groups to allow network access to the
instances.

- Specify whether to attach additional volumes to each instance.

- Add custom tags (key-value pairs) to the instances and volumes.

###### To create a launch template

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. On the navigation pane, under **Instances**, choose
    **Launch Templates**.

3. Choose **Create launch template**. Enter a name and
    provide a description for the initial version of the launch template.

4. (Optional) Under **Auto Scaling guidance**, select the
    check box to have Amazon EC2 provide guidance to help create a template to use
    with Amazon EC2 Auto Scaling.

5. Under **Launch template contents**, fill out each
    required field and any optional fields as needed.
1. **Application and OS Images (Amazon Machine**
      **Image)**: (Required) Choose the ID of the AMI for your
       instances. You can search through all available AMIs, or select an
       AMI from the **Recents** or **Quick**
      **Start** list. If you don't see the AMI that you need,
       choose **Browse more AMIs** to browse the full AMI
       catalog.

      To choose a custom AMI, you must first create your AMI from a
       customized instance. For more information, see [Create an Amazon EBS-backed\
       AMI](../../../ec2/latest/userguide/creating-an-ami-ebs.md) in the _Amazon EC2 User Guide_.

2. For **Instance type**, choose a single instance
       type that's compatible with the AMI that you specified.

      Alternatively, to use attribute-based instance type selection,
       choose **Advanced**, **Specify instance**
      **type attributes**, and then specify the following
       options:

- **Number of vCPUs**: Enter the minimum
and maximum number of vCPUs. To indicate no limits, enter a
minimum of 0, and keep the maximum blank.

- **Amount of memory (MiB)**: Enter the
minimum and maximum amount of memory, in MiB. To indicate no
limits, enter a minimum of 0, and keep the maximum blank.

- Expand **Optional instance type**
**attributes** and choose **Add**
**attribute** to further limit the types of
instances that can be used to fulfill your desired capacity.
For information about each attribute, see [InstanceRequirementsRequest](../../../../reference/awsec2/latest/apireference/api-instancerequirementsrequest.md) in the _Amazon EC2 API Reference_.

- **Resulting instance types**: You can
view the instance types that match the specified compute
requirements, such as vCPUs, memory, and storage.

- To exclude instance types, choose **Add**
**attribute**. From the
**Attribute** list, choose
**Excluded instance types**. From the
**Attribute value** list, select the
instance types to exclude.

3. **Key pair (login)**: For **Key pair**
      **name**, choose an existing key pair, or choose
       **Create new key pair** to create a new one.
       For more information, see [Amazon EC2 key pairs\
       and Linux instances](../../../ec2/latest/userguide/ec2-key-pairs.md) in the
       _Amazon EC2 User Guide_.

4. **Network settings**: For **Firewall**
      **(security groups)**, use one or more security groups,
       or keep this blank and configure one or more security groups as part
       of the network interface. For more information, see [Amazon EC2\
       security groups for Linux instances](../../../ec2/latest/userguide/ec2-security-groups.md) in the
       _Amazon EC2 User Guide_.

      If you don't specify any security groups in your launch template,
       Amazon EC2 uses the default security group for the VPC that your Auto Scaling
       group will launch instances into. By default, this security group
       doesn't allow inbound traffic from external networks. For more
       information, see [Default\
       security groups for your VPCs](https://docs.aws.amazon.com/vpc/latest/userguide/default-security-group.html) in the
       _Amazon VPC User Guide_.

5. Do one of the following:

- Change the default network interface settings. For
example, you can enable or disable the public IPv4
addressing feature, which overrides the auto-assign public
IPv4 addresses setting on the subnet. For more information,
see [Change the default network interface settings (console)](#change-network-interface).

- Skip this step to keep the default network interface
settings.

6. Do one of the following:

- Modify the storage configuration. For more information,
see [Modify the storage configuration (console)](#modify-storage-configuration).

- Skip this step to keep the default storage
configuration.

7. For **Resource tags**, specify tags by providing
       key and value combinations. If you specify instance tags in your
       launch template and then you choose to propagate your Auto Scaling group's
       tags to its instances, all the tags are merged. If the same tag key
       is specified for a tag in your launch template and a tag in your
       Auto Scaling group, then the tag value from the group takes precedence.
6. (Optional) Configure advanced settings. For example, you can choose an
    IAM role that your application can use when it accesses other AWS
    resources or specify the instance user data that can be used to perform
    common automated configuration tasks after an instance starts. For more
    information, see [Create a launch template using advanced settings](advanced-settings-for-your-launch-template.md).

7. When you are ready to create the launch template, choose **Create**
**launch template**.

8. To create an Auto Scaling group, choose **Create Auto Scaling**
**group** from the confirmation page.

## Change the default network interface settings (console)

Network interfaces provide connectivity to other resources in your VPC and the
internet. For more information, see [Provide network connectivity for your Auto Scaling instances using Amazon VPC](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html).

This section shows you how to change the default network interface settings. For
example, you can define whether you want to assign a public IPv4 address to each
instance instead of defaulting to the auto-assign public IPv4 addresses setting on
the subnet.

###### Considerations and limitations

When changing the default network interface settings, keep in mind the
following considerations and limitations:

- You must configure the security groups as part of the network interface,
not in the **Security groups** section of the template. You
cannot specify security groups in both places.

- If you specify an existing network interface ID, you can launch only one
instance. To do this, you must use the AWS CLI or an SDK to create the Auto Scaling
group. When you create the group, you must specify the Availability Zone,
but not the subnet ID. Also, you can specify an existing network interface
only if it has a device index of 0.

- You cannot auto-assign a public IPv4 address if you specify more than one
network interface. You also cannot specify duplicate device indexes across
network interfaces. Both the primary and secondary network interfaces reside
in the same subnet.

- When an instance launches, a private address is automatically allocated to
each network interface. The address comes from the CIDR range of the subnet
in which the instance is launched. For information on specifying CIDR blocks
(or IP address ranges) for your VPC or subnet, see the
[Amazon VPC User Guide](https://docs.aws.amazon.com/vpc/latest/userguide).

###### To change the default network interface settings

1. Under **Network settings**, expand **Advanced**
**network configuration**.

2. Choose **Add network interface** to configure the primary
    network interface, paying attention to the following fields:
01. **Device index**: Keep the default value, 0, to
        apply your changes to the primary network interface (eth0).

02. **Network interface**: Keep the default value,
        **New interface**, to have Amazon EC2 Auto Scaling
        automatically create a new network interface when an instance is
        launched. Alternatively, you can choose an existing, available
        network interface with a device index of 0, but this limits your
        Auto Scaling group to one instance.

03. **Description**: (Optional) Enter a descriptive
        name.

04. **Subnet**: Keep the default **Don't**
       **include in launch template** setting.

       If the AMI specifies a subnet for the network interface, this
        results in an error. We recommend turning off **Auto Scaling**
       **guidance** as a workaround. After you make this change,
        you will not receive an error message. However, regardless of where
        the subnet is specified, the subnet settings of the Auto Scaling group take
        precedence and cannot be overridden.

05. **Auto-assign public IP**: Change whether your
        network interface with a device index of 0 receives a public IPv4
        address. By default, instances in a default subnet receive a public
        IPv4 address, while instances in a nondefault subnet do not. Select
        **Enable** or **Disable** to
        override the subnet's default setting.

06. **Security groups**: Choose one or more security
        groups for the network interface. Each security group must be
        configured for the VPC that your Auto Scaling group will launch instances
        into. For more information, see [Amazon EC2\
        security groups for Linux instances](../../../ec2/latest/userguide/ec2-security-groups.md) in the
        _Amazon EC2 User Guide_.

07. **Delete on termination**: Choose
        **Yes** to delete the network interface when
        the instance is terminated, or choose **No** to
        keep the network interface.

08. **Elastic Fabric Adapter**: To support high
        performance computing and machine learning use cases, change the
        network interface into an Elastic Fabric Adapter network interface.
        For more information, see [Elastic Fabric Adapter](../../../ec2/latest/userguide/efa.md) in the
        _Amazon EC2 User Guide_.

09. **Network card index**: Choose
        **0** to attach the primary network interface
        to the network card with a device index of 0. If this option isn't
        available, keep the default value, **Don't include in launch**
       **template**. Attaching the network interface to a
        specific network card is available only for supported instance
        types. For more information, see [Network\
        cards](../../../ec2/latest/userguide/using-eni.md#network-cards) in the _Amazon EC2 User Guide_.

10. **ENA Express**: For instance types that support
        ENA Express, choose **Enable** to enable ENA
        Express or **Disable** to disable it. For more
        information, see [Improve\
        network performance with ENA Express on Linux instances](../../../ec2/latest/userguide/ena-express.md)
        in the _Amazon EC2 User Guide_.

11. **ENA Express UDP**: If you enable **ENA**
       **Express**, you can optionally use it for UDP traffic.
        Choose **Enable** to enable ENA Express UDP or
        **Disable** to disable it.
3. To add a secondary network interface, choose **Add network**
**interface**.

## Modify the storage configuration (console)

You can modify the storage configuration for instances launched from an
Amazon EBS-backed AMI or an instance store-backed AMI. You can also specify additional
EBS volumes to attach to the instances. The AMI includes one or more volumes of
storage, including the root volume ( **Volume 1 (AMI Root)**).

###### To modify the storage configuration

1. In **Configure storage**, modify the size or type of
    volume.

If the value you specify for volume size is outside the limits of the
    volume type, or smaller than the snapshot size, an error message is
    displayed. To help you address the issue, this message gives the minimum or
    maximum value that the field can accept.

Only volumes associated with an Amazon EBS-backed AMI appear. To display
    information about the storage configuration for an instance launched from an
    instance store-backed AMI, choose **Show details** from the
    **Instance store volumes** section.

To specify all EBS volume parameters, switch to the
    **Advanced** view in the top right corner.

2. For advanced options, expand the volume that you want to modify and
    configure the volume as follows:
1. **Storage type**: The type of volume (EBS or
       ephemeral) to associate with your instance. The instance store
       (ephemeral) volume type is only available if you select an instance
       type that supports it. For more information, see [Amazon EBS volumes](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volumes.html) in the _Amazon EBS User_
      _Guide_ and [Amazon EC2\
       instance store](../../../ec2/latest/userguide/instancestorage.md) in the
       _Amazon EC2 User Guide_.

2. **Device name**: Select from the list of
       available device names for the volume.

3. **Snapshot**: Select the snapshot from which to
       create the volume. You can search for available shared and public
       snapshots by entering text into the **Snapshot**
       field.

4. **Size (GiB)**: For EBS volumes, you can specify
       a storage size. If you have selected an AMI and instance that are
       eligible for the free tier, keep in mind that to stay within the
       free tier, you must stay under 30 GiB of total storage. For more
       information, see [Constraints\
       on the size and configuration of an EBS volume](https://docs.aws.amazon.com/ebs/latest/userguide/volume_constraints.html) in the
       _Amazon EBS User Guide_.

5. **Volume type**: For EBS volumes, choose the
       volume type. For more information, see [Amazon EBS volume\
       types](../../../ebs/latest/userguide/ebs-volume-types.md) in the _Amazon EBS User_
      _Guide_.

6. **IOPS**: If you have selected a Provisioned IOPS
       SSD ( `io1` and `io2`) or General Purpose SSD
       ( `gp3`) volume type, then you can enter the number of
       I/O operations per second (IOPS) that the volume can support. This
       is required for io1, io2, and gp3 volumes. It is not supported for
       gp2, st1, sc1, or standard volumes.

7. **Delete on termination**: For EBS volumes,
       choose **Yes** to delete the volume when the
       instance is terminated, or choose **No** to keep
       the volume.

8. **Encrypted**: If the instance type supports EBS
       encryption, you can choose **Yes** to enable
       encryption for the volume. If you have enabled encryption by default
       in this Region, encryption is enabled for you. For more information,
       see [Amazon EBS\
       encryption](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html) and [Enable Amazon EBS encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/encryption-by-default.html) in the _Amazon EBS_
      _User Guide_.

      The default effect of setting this parameter varies with the
       choice of volume source, as described in the following table. In all
       cases, you must have permission to use the specified AWS KMS key.

      Encryption outcomes If `Encrypted` parameter is set
      to...And if source of volume is...Then the default encryption state is...NotesNoNew (empty) volumeUnencrypted\*N/AUnencrypted snapshot that you ownUnencrypted\*Encrypted snapshot that you ownEncrypted by same keyUnencrypted snapshot that is shared with
      youUnencrypted\*Encrypted snapshot that is shared with
      youEncrypted by default KMS keyYesNew volumeEncrypted by default KMS keyTo use a non-default KMS key,
      specify a value for the **KMS key**
      parameter. Unencrypted snapshot that you ownEncrypted by default KMS keyEncrypted snapshot that you ownEncrypted by same keyUnencrypted snapshot that is shared with
      youEncrypted by default KMS keyEncrypted snapshot that is shared with
      youEncrypted by default KMS key

      \\* If encryption by default is enabled, all newly created volumes
       (whether or not the **Encrypted** parameter is set
       to **Yes**) are encrypted using the default
       KMS key. If you set both the **Encrypted** and
       **KMS key** parameters, then you can specify a
       non-default KMS key.

9. **KMS key**: If you chose
       **Yes** for **Encrypted**,
       then you must select a customer managed key to use to encrypt the volume. If
       you have enabled encryption by default in this Region, the default
       customer managed key is selected for you. You can select a different key or
       specify the ARN of any customer managed key that you previously created using
       the AWS Key Management Service.
3. To specify additional volumes to attach to the instances launched by this
    launch template, choose **Add new volume**.

## Create a launch template from an existing instance (console)

###### To create a launch template from an existing instance

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. On the navigation pane, under **Instances**, choose
    **Instances**.

3. Select the instance and choose **Actions**,
    **Image and templates**, **Create template from**
**instance**.

4. Provide a name and description.

5. Under **Auto Scaling guidance**, select the check box.

6. Adjust any settings as required, and choose **Create launch**
**template**.

7. To create an Auto Scaling group, choose **Create Auto Scaling**
**group** from the confirmation page.

## Related resources

We provide a few JSON and YAML template snippets that you can use to understand
how to declare launch templates in your CloudFormation stack templates. For more information,
see the [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) and [Create launch templates with CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-ec2-launch-templates.html) sections of the
_AWS CloudFormation User Guide_.

For more information about launch templates, see [Launching an instance from a\
launch template](../../../ec2/latest/userguide/ec2-launch-templates.md) in the _Amazon EC2 User Guide_.

## Limitations

- While you can specify a subnet in a launch template, doing so isn't
necessary if you only use the launch template to create Auto Scaling groups. You
can't specify the subnet for an Auto Scaling group by specifying the subnet in a
launch template. The subnets for the Auto Scaling group are taken from the Auto Scaling
group's own resource definition.

- For other limitations on user-defined network interfaces, see [Change the default network interface settings (console)](#change-network-interface).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Launch templates

Create a launch template using advanced settings
