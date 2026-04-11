This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet Ec2EbsVolume

Specifies the EBS volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iops" : Integer,
  "SizeGiB" : Integer,
  "ThroughputMiB" : Integer
}

```

### YAML

```yaml

  Iops: Integer
  SizeGiB: Integer
  ThroughputMiB: Integer

```

## Properties

`Iops`

The IOPS per volume.

_Required_: No

_Type_: Integer

_Minimum_: `3000`

_Maximum_: `16000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeGiB`

The EBS volume size in GiB.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThroughputMiB`

The throughput per volume in MiB.

_Required_: No

_Type_: Integer

_Minimum_: `125`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomerManagedWorkerCapabilities

FleetAmountCapability

All content copied from https://docs.aws.amazon.com/.
