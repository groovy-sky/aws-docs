This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance Volume

Specifies a volume to attach to an instance.

`Volume` is an embedded property of the [AWS::EC2::Instance](../userguide/aws-properties-ec2-instance.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Device" : String,
  "VolumeId" : String
}

```

### YAML

```yaml

  Device: String
  VolumeId: String

```

## Properties

`Device`

The device name (for example, `/dev/sdh` or `xvdh`).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeId`

The ID of the EBS volume. The volume and instance must be within the same Availability
Zone.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::InstanceConnectEndpoint

All content copied from https://docs.aws.amazon.com/.
