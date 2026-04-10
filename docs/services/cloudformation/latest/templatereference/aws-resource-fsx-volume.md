This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume

Creates an FSx for ONTAP or Amazon FSx for OpenZFS storage volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FSx::Volume",
  "Properties" : {
      "BackupId" : String,
      "Name" : String,
      "OntapConfiguration" : OntapConfiguration,
      "OpenZFSConfiguration" : OpenZFSConfiguration,
      "Tags" : [ Tag, ... ],
      "VolumeType" : String
    }
}

```

### YAML

```yaml

Type: AWS::FSx::Volume
Properties:
  BackupId: String
  Name: String
  OntapConfiguration:
    OntapConfiguration
  OpenZFSConfiguration:
    OpenZFSConfiguration
  Tags:
    - Tag
  VolumeType: String

```

## Properties

`BackupId`

Specifies the ID of the volume backup to use to create a new volume.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the volume.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,203}$`

_Minimum_: `1`

_Maximum_: `203`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OntapConfiguration`

The configuration of an Amazon FSx for NetApp ONTAP volume.

_Required_: No

_Type_: [OntapConfiguration](aws-properties-fsx-volume-ontapconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenZFSConfiguration`

The configuration of an Amazon FSx for OpenZFS volume.

_Required_: No

_Type_: [OpenZFSConfiguration](aws-properties-fsx-volume-openzfsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-fsx-volume-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeType`

The type of the volume.

_Required_: No

_Type_: String

_Allowed values_: `ONTAP | OPENZFS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID for the volume. For example:

`{"Ref":"vol_logical_id"}`

Returns `fsvol-0123456789abcdef6`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceARN`

Returns the volume's Amazon Resource Name (ARN).

Example: `arn:aws:fsx:us-east-2:111122223333:volume/fs-0123456789abcdef9/fsvol-01234567891112223`

`UUID`

Returns the volume's universally unique identifier (UUID).

Example: `abcd0123-cd45-ef67-11aa-1111aaaa23bc`

`VolumeId`

Returns the volume's ID.

Example: `fsvol-0123456789abcdefa`

## Examples

- [Create an ONTAP volume](#aws-resource-fsx-volume--examples--Create_an_ONTAP_volume)

- [Create an ONTAP volume from a backup](#aws-resource-fsx-volume--examples--Create_an_ONTAP_volume_from_a_backup)

### Create an ONTAP volume

#### JSON

```json

{
  "OntapVolumeWithAllConfigs": {
    "Type": "AWS::FSx::Volume",
    "Properties": {
      "Name": "volume1",
      "OntapConfiguration": {
        "JunctionPath": "/volume1",
        "SecurityStyle": "UNIX",
        "SizeInBytes": 419430400,
        "StorageEfficiencyEnabled": true,
        "StorageVirtualMachineId": {
          "Ref": "OntapStorageVirtualMachineWithAllConfigs"
        },
        "TieringPolicy": {
          "CoolingPeriod": 41,
          "Name": "AUTO"
        }
      },
      "Tags": [
        {
          "Key": "Name",
          "Value": "OntapVolume"
        }
      ],
      "VolumeType": "ONTAP"
    }
  }
}
```

#### YAML

```yaml

OntapVolumeWithAllConfigs:
  Type: 'AWS::FSx::Volume'
  Properties:
    Name: volume1
    OntapConfiguration:
      JunctionPath: /volume1
      SecurityStyle: UNIX
      SizeInBytes: 419430400
      StorageEfficiencyEnabled: true
      StorageVirtualMachineId: !Ref OntapStorageVirtualMachineWithAllConfigs
      TieringPolicy:
        CoolingPeriod: 41
        Name: AUTO
    Tags:
      - Key: Name
        Value: OntapVolume
    VolumeType: ONTAP
```

### Create an ONTAP volume from a backup

This example creates a volume from an existing backup:
`backup-0123abc456defghij`

#### JSON

```json

{
  "OntapVolumeFromBackupWithAllConfigs": {
    "Type": "AWS::FSx::Volume",
    "Properties": {
    "BackupId": "backup-0123abc456defghij",
      "Name": "volume11",
      "OntapConfiguration": {
        "JunctionPath": "/volume11",
        "SecurityStyle": "UNIX",
        "SizeInBytes": 419430400,
        "StorageEfficiencyEnabled": true,
        "StorageVirtualMachineId": {
          "Ref": "StorageVirtualMachineWithAllConfigs"
        },
        "TieringPolicy": {
          "CoolingPeriod": 42,
          "Name": "AUTO"
        }
      }
    }
  }
}
```

#### YAML

```yaml

OntapVolumeFromBackupWithAllConfigs:
  Type: "AWS::FSx::Volume"
  Properties:
  BackupId: "backup-0123abc456defghij"
    Name: "volume11"
    OntapConfiguration:
      JunctionPath: "/volume11"
      SecurityStyle: "UNIX"
      SizeInBytes: 419430400
      StorageEfficiencyEnabled: True
      StorageVirtualMachineId: !Ref StorageVirtualMachineWithAllConfigs
      TieringPolicy:
        CoolingPeriod: 42
        Name: "AUTO"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AggregateConfiguration

All content copied from https://docs.aws.amazon.com/.
