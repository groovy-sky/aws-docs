This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Instance

The `AWS::Lightsail::Instance` resource specifies an Amazon Lightsail instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Instance",
  "Properties" : {
      "AddOns" : [ AddOn, ... ],
      "AvailabilityZone" : String,
      "BlueprintId" : String,
      "BundleId" : String,
      "Hardware" : Hardware,
      "InstanceName" : String,
      "KeyPairName" : String,
      "Location" : Location,
      "Networking" : Networking,
      "State" : State,
      "Tags" : [ Tag, ... ],
      "UserData" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Instance
Properties:
  AddOns:
    - AddOn
  AvailabilityZone: String
  BlueprintId: String
  BundleId: String
  Hardware:
    Hardware
  InstanceName: String
  KeyPairName: String
  Location:
    Location
  Networking:
    Networking
  State:
    State
  Tags:
    - Tag
  UserData: String

```

## Properties

`AddOns`

An array of add-ons for the instance.

###### Note

If the instance has an add-on enabled when performing a delete instance request, the
add-on is automatically disabled before the instance is deleted.

_Required_: No

_Type_: Array of [AddOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-instance-addon.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The Availability Zone for the instance.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`BlueprintId`

The blueprint ID for the instance (for example, `os_amlinux_2016_03`).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`BundleId`

The bundle ID for the instance (for example, `micro_1_0`).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`Hardware`

The hardware properties for the instance, such as the vCPU count, attached disks, and
amount of RAM.

###### Important

The instance restarts when performing an attach disk or detach disk request. This
resets the public IP address of your instance if a static IP isn't attached to
it.

_Required_: No

_Type_: [Hardware](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-instance-hardware.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InstanceName`

The name of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][\w\-.]*[a-zA-Z0-9]$`

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyPairName`

The name of the key pair to use for the instance.

If no key pair name is specified, the Regional Lightsail default key
pair is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

The location for the instance, such as the AWS Region and Availability
Zone.

###### Note

The `Location` property is read-only and should not be specified in a
create instance or update instance request.

_Required_: No

_Type_: [Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-instance-location.html)

_Update requires_: Updates are not supported.

`Networking`

The public ports and the monthly amount of data transfer allocated for the
instance.

_Required_: No

_Type_: [Networking](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-instance-networking.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The status code and the state (for example, `running`) of the
instance.

###### Note

The `State` property is read-only and should not be specified in a create
instance or update instance request.

_Required_: No

_Type_: [State](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-instance-state.html)

_Update requires_: Updates are not supported.

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-instance-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserData`

The optional launch script for the instance.

Specify a launch script to configure an instance with additional user data. For example,
you might want to specify `apt-get -y update` as a launch script.

###### Note

Depending on the blueprint of your instance, the command to get software on your
instance varies. Amazon Linux and CentOS use `yum`, Debian and Ubuntu use
`apt-get`, and FreeBSD uses `pkg`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Hardware.CpuCount`

The number of vCPUs the instance has.

`Hardware.RamSizeInGb`

The amount of RAM in GB on the instance (for example, `1.0`).

`InstanceArn`

The Amazon Resource Name (ARN) of the instance (for example,
`arn:aws:lightsail:us-east-2:123456789101:Instance/244ad76f-8aad-4741-809f-12345EXAMPLE`).

`Ipv6Addresses`

The IPv6 addresses of the instance.

`IsStaticIp`

A Boolean value indicating whether the instance has a static IP assigned to it.

`Location.AvailabilityZone`

The AWS Region and Availability Zone where the instance is
located.

`Location.RegionName`

The AWS Region of the instance.

`Networking.MonthlyTransfer.GbPerMonthAllocated`

The amount of allocated monthly data transfer (in GB) for an instance.

`PrivateIpAddress`

The private IP address of the instance.

`PublicIpAddress`

The public IP address of the instance.

`ResourceType`

The resource type of the instance (for example, `Instance`).

`SshKeyName`

The name of the SSH key pair used by the instance.

`State.Code`

The status code of the instance.

`State.Name`

The state of the instance (for example, `running` or
`pending`).

`SupportCode`

The support code of the instance.

Include this code in your email to support when you have questions about an instance or
another resource in Lightsail. This code helps our support team to look up
your Lightsail information.

`UserName`

The user name for connecting to the instance (for example,
`ec2-user`).

## Remarks

_Attaching a static IP to an instance_

You cannot attach a static IP to an instance using the instance resource. Instead,
you must use the static IP resource to attach a static IP to an instance. To attach a
static IP to an instance, the instance must be in a `running` state.

_Network ports_

If no network ports are specified when performing a create instance request, the
default network ports are opened when the instance is created.

To open ports on your instance when performing a create instance request, you must
specify all the ports that you want to open, including the default ports. The default
ports are not automatically opened when you specify the ports you want to open.

_Disk attach and detach_

The instance restarts when performing an attach disk or detach disk request. This
resets the public IP address of your instance if a static IP isn't attached to
it.

If you detach a disk (for eample, `DiskA`) and attach a different disk
(for example, `DiskB`) in the same request, and the attach disk request
fails, CloudFormation will attempt to roll back the changes so that
`DiskA` is re-attached to the instance. However, if you delete
`DiskA` before CloudFormation attemps the roll-back, then the
roll-back will fail and the instance will not have either disk attached.

_Read-only properties_

The `State`, `Location`, `CpuCount`, and
`RamSizeInGb` properties are read-only and should not be specified in a
create instance or update instance request.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AddOn
