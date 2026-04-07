This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBInstance AdditionalStorageVolume

Contains details about an additional storage volume for a DB instance. RDS support
additional storage volumes for RDS for Oracle and RDS for SQL Server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocatedStorage" : String,
  "Iops" : Integer,
  "MaxAllocatedStorage" : Integer,
  "StorageThroughput" : Integer,
  "StorageType" : String,
  "VolumeName" : String
}

```

### YAML

```yaml

  AllocatedStorage: String
  Iops: Integer
  MaxAllocatedStorage: Integer
  StorageThroughput: Integer
  StorageType: String
  VolumeName: String

```

## Properties

`AllocatedStorage`

The amount of storage allocated for the additional storage volume, in gibibytes (GiB).
The minimum is 20 GiB. The maximum is 65,536 GiB (64 TiB).

_Required_: No

_Type_: String

_Pattern_: `^[0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Iops`

The number of I/O operations per second (IOPS) provisioned for the additional storage
volume.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAllocatedStorage`

The upper limit in gibibytes (GiB) to which RDS can automatically scale the storage of
the additional storage volume.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageThroughput`

The storage throughput value for the additional storage volume, in mebibytes per second (MiBps). This setting applies only to the General Purpose SSD ( `gp3`) storage type.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageType`

The storage type for the additional storage volume.

Valid Values: `GP3 | IO2`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeName`

The name of the additional storage volume.

Valid Values: `RDSDBDATA2 | RDSDBDATA3 | RDSDBDATA4`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RDS::DBInstance

CertificateDetails
