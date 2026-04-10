This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster EbsBlockDeviceConfig

`EbsBlockDeviceConfig` is a subproperty of the `EbsConfiguration` property type. `EbsBlockDeviceConfig` defines the number and type of EBS volumes to associate with all EC2 instances in an EMR cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VolumeSpecification" : VolumeSpecification,
  "VolumesPerInstance" : Integer
}

```

### YAML

```yaml

  VolumeSpecification:
    VolumeSpecification
  VolumesPerInstance: Integer

```

## Properties

`VolumeSpecification`

EBS volume specifications such as volume type, IOPS, size (GiB) and throughput (MiB/s)
that are requested for the EBS volume attached to an Amazon EC2 instance in the
cluster.

_Required_: Yes

_Type_: [VolumeSpecification](aws-properties-emr-cluster-volumespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumesPerInstance`

Number of EBS volumes with a specific volume configuration that are associated with
every instance in the instance group

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configuration

EbsConfiguration

All content copied from https://docs.aws.amazon.com/.
