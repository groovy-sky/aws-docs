This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain EBSOptions

The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to
data nodes in the OpenSearch Service domain. For more information, see [EBS volume size limits](../../../opensearch-service/latest/developerguide/limits.md#ebsresource) in the _Amazon OpenSearch Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EBSEnabled" : Boolean,
  "Iops" : Integer,
  "Throughput" : Integer,
  "VolumeSize" : Integer,
  "VolumeType" : String
}

```

### YAML

```yaml

  EBSEnabled: Boolean
  Iops: Integer
  Throughput: Integer
  VolumeSize: Integer
  VolumeType: String

```

## Properties

`EBSEnabled`

Specifies whether Amazon EBS volumes are attached to data nodes in the OpenSearch Service
domain.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Iops`

The number of I/O operations per second (IOPS) that the volume supports. This property
applies only to the `gp3` and provisioned IOPS EBS volume types.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Throughput`

The throughput (in MiB/s) of the EBS volumes attached to data nodes. Applies only to the
`gp3` volume type.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeSize`

The size (in GiB) of the EBS volume for each data node. The minimum and maximum size of an
EBS volume depends on the EBS volume type and the instance type to which it is attached. For
more information, see [EBS volume size\
limits](../../../opensearch-service/latest/developerguide/limits.md#ebsresource) in the _Amazon OpenSearch Service Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeType`

The EBS volume type to use with the OpenSearch Service domain. If you choose
`gp3`, you must also specify values for `Iops` and
`Throughput`. For more information about each type, see [Amazon EBS volume\
types](../../../ec2/latest/userguide/ebsvolumetypes.md) in the _Amazon EC2 User Guide for Linux Instances_.

_Required_: No

_Type_: String

_Allowed values_: `standard | gp2 | io1 | gp3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainEndpointOptions

EncryptionAtRestOptions

All content copied from https://docs.aws.amazon.com/.
