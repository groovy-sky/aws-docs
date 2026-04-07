This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VolumeAttachment

Attaches an Amazon EBS volume to a running instance and exposes it to the instance with
the specified device name.

Before this resource can be deleted (and therefore the volume detached), you must first
unmount the volume in the instance. Failure to do so results in the volume being stuck in
the busy state while it is trying to detach, which could possibly damage the file system or
the data it contains.

If an Amazon EBS volume is the root device of an instance, it cannot be detached while
the instance is in the "running" state. To detach the root volume, stop the instance
first.

If the root volume is detached from an instance with an AWS Marketplace product
code, then the product codes from that volume are no longer associated with the
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VolumeAttachment",
  "Properties" : {
      "Device" : String,
      "EbsCardIndex" : Integer,
      "InstanceId" : String,
      "VolumeId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VolumeAttachment
Properties:
  Device: String
  EbsCardIndex: Integer
  InstanceId: String
  VolumeId: String

```

## Properties

`Device`

The device name (for example, `/dev/sdh` or `xvdh`).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsCardIndex`

The index of the EBS card. Some instance types support multiple EBS cards. The default EBS card index is 0.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceId`

The ID of the instance to which the volume attaches. This value can be a reference to an
[`AWS::EC2::Instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource, or it can be the physical ID of an
existing EC2 instance.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeId`

The ID of the Amazon EBS volume. The volume and instance must be within the same
Availability Zone. This value can be a reference to an [`AWS::EC2::Volume`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html) resource, or it can be the volume ID of an
existing Amazon EBS volume.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Attach an EBS volume to a running instance

This example attaches an EC2 EBS volume to the EC2 instance with the logical name
"Ec2Instance".

#### JSON

```json

"NewVolume" : {
   "Type" : "AWS::EC2::Volume",
   "Properties" : {
      "Size" : "100",
      "AvailabilityZone" : { "Fn::GetAtt" : [ "Ec2Instance", "AvailabilityZone" ] },
      "Tags" : [ {
         "Key" : "MyTag",
          "Value" : "TagValue"
      } ]
   }
},

"MountPoint" : {
   "Type" : "AWS::EC2::VolumeAttachment",
   "Properties" : {
      "InstanceId" : { "Ref" : "Ec2Instance" },
      "VolumeId"  : { "Ref" : "NewVolume" },
      "Device" : "/dev/sdh"
   }
}
```

#### YAML

```yaml

NewVolume:
  Type: AWS::EC2::Volume
  Properties:
    Size: 100
    AvailabilityZone: !GetAtt Ec2Instance.AvailabilityZone
    Tags:
      - Key: MyTag
        Value: TagValue
  DeletionPolicy: Snapshot

MountPoint:
  Type: AWS::EC2::VolumeAttachment
  Properties:
    InstanceId: !Ref Ec2Instance
    VolumeId: !Ref NewVolume
    Device: /dev/sdh
```

## See also

- [Attach an EBS volume to\
an instance](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-attaching-volume.html)

- [Detach an EBS volume\
from an instance](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-detaching-volume.html)

- [AttachVolume](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-AttachVolume.html) in the _Amazon EC2 API_
_Reference_

- [DetachVolume](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DetachVolume.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EC2::VPC
