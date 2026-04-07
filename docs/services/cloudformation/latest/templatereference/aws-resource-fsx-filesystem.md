This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem

The `AWS::FSx::FileSystem` resource is an Amazon FSx resource type
that specifies an Amazon FSx file system. You can create any of the following
supported file system types:

- Amazon FSx for Lustre

- Amazon FSx for NetApp ONTAP

- FSx for OpenZFS

- Amazon FSx for Windows File Server

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FSx::FileSystem",
  "Properties" : {
      "BackupId" : String,
      "FileSystemType" : String,
      "FileSystemTypeVersion" : String,
      "KmsKeyId" : String,
      "LustreConfiguration" : LustreConfiguration,
      "NetworkType" : String,
      "OntapConfiguration" : OntapConfiguration,
      "OpenZFSConfiguration" : OpenZFSConfiguration,
      "SecurityGroupIds" : [ String, ... ],
      "StorageCapacity" : Integer,
      "StorageType" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "WindowsConfiguration" : WindowsConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::FSx::FileSystem
Properties:
  BackupId: String
  FileSystemType: String
  FileSystemTypeVersion: String
  KmsKeyId: String
  LustreConfiguration:
    LustreConfiguration
  NetworkType: String
  OntapConfiguration:
    OntapConfiguration
  OpenZFSConfiguration:
    OpenZFSConfiguration
  SecurityGroupIds:
    - String
  StorageCapacity: Integer
  StorageType: String
  SubnetIds:
    - String
  Tags:
    - Tag
  WindowsConfiguration:
    WindowsConfiguration

```

## Properties

`BackupId`

The ID of the file system backup that you are using to create a file system. For more information,
see [CreateFileSystemFromBackup](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemFromBackup.html).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemType`

The type of Amazon FSx file system, which can be `LUSTRE`,
`WINDOWS`, `ONTAP`, or `OPENZFS`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemTypeVersion`

For FSx for Lustre file systems, sets the Lustre version for the file system
that you're creating. Valid values are `2.10`, `2.12`, and
`2.15`:

- `2.10` is supported by the Scratch and Persistent\_1 Lustre
deployment types.

- `2.12` is supported by all Lustre deployment types, except
for `PERSISTENT_2` with a metadata configuration mode.

- `2.15` is supported by all Lustre deployment types and is
recommended for all new file systems.

Default value is `2.10`, except for the following deployments:

- Default value is `2.12` when `DeploymentType` is set to
`PERSISTENT_2` without a metadata configuration mode.

- Default value is `2.15` when `DeploymentType` is set to
`PERSISTENT_2` with a metadata configuration mode.

_Required_: No

_Type_: String

_Pattern_: `^[0-9](.[0-9]*)*$`

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the AWS Key Management Service (AWS KMS) key used to encrypt Amazon FSx file
system data. Used as follows with Amazon FSx file system types:

- Amazon FSx for Lustre `PERSISTENT_1`
and `PERSISTENT_2` deployment types only.

`SCRATCH_1` and `SCRATCH_2` types are encrypted using
the Amazon FSx service AWS KMS key for your account.

- Amazon FSx for NetApp ONTAP

- Amazon FSx for OpenZFS

- Amazon FSx for Windows File Server

If this ID isn't specified, the Amazon FSx-managed key for your account is used. For more information,
see [Encrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Encrypt.html) in the
_AWS Key Management Service API Reference_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LustreConfiguration`

The Lustre configuration for the file system being created. This configuration is required if the `FileSystemType` is
set to `LUSTRE`.

###### Note

The following parameters are not supported when creating Lustre file systems with a data repository association.

- `AutoImportPolicy`

- `ExportPath`

- `ImportedChunkSize`

- `ImportPath`

_Required_: Conditional

_Type_: [LustreConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-filesystem-lustreconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkType`

The network type of the file system.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OntapConfiguration`

The ONTAP configuration properties of the FSx for ONTAP file system that you
are creating. This configuration is required if the `FileSystemType` is set to
`ONTAP`.

_Required_: Conditional

_Type_: [OntapConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-filesystem-ontapconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenZFSConfiguration`

The Amazon FSx for OpenZFS configuration properties for the file system that you are creating.
This configuration is required if the `FileSystemType` is set to `OPENZFS`.

_Required_: Conditional

_Type_: [OpenZFSConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-filesystem-openzfsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

A list of IDs specifying the security groups to apply to all network interfaces
created for file system access. This list isn't returned in later requests to
describe the file system.

###### Important

You must specify a security group if you are creating a Multi-AZ
FSx for ONTAP file system in a VPC subnet that has been shared with you.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageCapacity`

Sets the storage capacity of the file system that you're creating.

`StorageCapacity` is required if you are creating a new file system. It is not required if you are creating a
file system by restoring a backup.

**FSx for Lustre file systems** \- The amount of
storage capacity that you can configure depends on the value that you set for
`StorageType` and the Lustre `DeploymentType`, as
follows:

- For `SCRATCH_2`, `PERSISTENT_2` and
`PERSISTENT_1` deployment types using SSD storage type, the valid
values are 1200 GiB, 2400 GiB, and increments of 2400 GiB.

- For `PERSISTENT_1` HDD file systems, valid values are increments of
6000 GiB for 12 MB/s/TiB file systems and increments of 1800 GiB for 40 MB/s/TiB
file systems.

- For `SCRATCH_1` deployment type, valid values are 1200 GiB, 2400
GiB, and increments of 3600 GiB.

**FSx for ONTAP file systems** \- The amount of SSD storage capacity that you
can configure depends on the value of the `HAPairs` property. The minimum value is calculated as 1,024 GiB \* HAPairs and
the maximum is calculated as 524,288 GiB \* HAPairs, up to a maximum amount of SSD storage capacity of 1,048,576 GiB
(1 pebibyte).

**FSx for OpenZFS file systems** \- The amount of storage
capacity that you can configure is from 64 GiB up to 524,288 GiB (512 TiB). If you are creating a
file system from a backup, you can specify a storage capacity equal to or greater than the original
file system's storage capacity.

**FSx for Windows File Server file systems** \- The amount
of storage capacity that you can configure depends on the value that you set for
`StorageType` as follows:

- For SSD storage, valid values are 32 GiB-65,536 GiB (64 TiB).

- For HDD storage, valid values are 2000 GiB-65,536 GiB (64 TiB).

_Required_: Conditional

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageType`

Sets the storage class for the file system that you're creating. Valid values are
`SSD`, `HDD`, and `INTELLIGENT_TIERING`.

- Set to `SSD` to use solid state drive storage. SSD is supported on all Windows,
Lustre, ONTAP, and OpenZFS deployment types.

- Set to `HDD` to use hard disk drive storage, which is supported on
`SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types,
and on `PERSISTENT_1` Lustre file system deployment types.

- Set to `INTELLIGENT_TIERING` to use fully elastic, intelligently-tiered storage.
Intelligent-Tiering is only available for OpenZFS file systems with the Multi-AZ deployment type
and for Lustre file systems with the Persistent\_2 deployment type.

Default value is `SSD`. For more information, see [Storage\
type options](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/optimize-fsx-costs.html#storage-type-options) in the _FSx for Windows File Server User_
_Guide_, [FSx for Lustre storage classes](https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-fsx-lustre.html#lustre-storage-classes)
in the _FSx for Lustre User Guide_, and [Working with Intelligent-Tiering](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance-intelligent-tiering)
in the _Amazon FSx for OpenZFS User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `SSD | HDD | INTELLIGENT_TIERING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

Specifies the IDs of the subnets that the file system will be accessible from. For
Windows and ONTAP `MULTI_AZ_1` deployment types,provide exactly two subnet
IDs, one for the preferred file server and one for the standby file server. You specify
one of these subnets as the preferred subnet using the `WindowsConfiguration >
                PreferredSubnetID` or `OntapConfiguration > PreferredSubnetID`
properties. For more information about Multi-AZ file system configuration, see [Availability and durability: Single-AZ and Multi-AZ file systems](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html) in the
_Amazon FSx for Windows User Guide_ and [Availability and durability](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html) in the _Amazon FSx for ONTAP User_
_Guide_.

For Windows `SINGLE_AZ_1` and `SINGLE_AZ_2` and all Lustre
deployment types, provide exactly one subnet ID.
The file server is launched in that subnet's Availability Zone.

_Required_: Yes

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to associate with the file system. For more information, see [Tagging your\
Amazon FSx resources](https://docs.aws.amazon.com/fsx/latest/LustreGuide/tag-resources.html) in the _Amazon FSx for Lustre User_
_Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-filesystem-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WindowsConfiguration`

The configuration object for the Microsoft Windows file system you are creating. This configuration is required if `FileSystemType`
is set to `WINDOWS`.

_Required_: Conditional

_Type_: [WindowsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-filesystem-windowsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the file system resource ID. For example:

`{"Ref":"file_system_logical_id"}`

Returns `fs-0123456789abcdef6`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DNSName`

Returns the FSx for Windows file system's DNSName.

Example: `amznfsxp1honlek.corp.example.com`

`LustreMountName`

Returns the Lustre file system's `LustreMountName`.

Example for `SCRATCH_1` deployment types: This value is always `fsx`.

Example for `SCRATCH_2` and `PERSISTENT` deployment types: `2p3fhbmv`

`ResourceARN`

Returns the Amazon Resource Name (ARN) for the Amazon FSx file system.

Example: `arn:aws:fsx:us-east-2:111122223333:file-system/fs-0123abcd56789ef0a`

`RootVolumeId`

Returns the root volume ID of the FSx for OpenZFS file system.

Example: `fsvol-0123456789abcdefa`

## Examples

- [Create an Amazon FSx for Lustre File System](#aws-resource-fsx-filesystem--examples--Create_an_Amazon_FSx_for_Lustre_File_System)

- [Create an Amazon FSx for Windows File Server File System in a Self-managed Active Directory](#aws-resource-fsx-filesystem--examples--Create_an_Amazon_FSx_for_Windows_File_Server_File_System_in_a_Self-managed_Active_Directory)

- [Create an Amazon FSx for Windows File Server File System in an Amazon Managed Active Directory](#aws-resource-fsx-filesystem--examples--Create_an_Amazon_FSx_for_Windows_File_Server_File_System_in_an_Amazon_Managed_Active_Directory)

- [Create an Amazon FSx for Windows File Server File System with file access audit event logging enabled](#aws-resource-fsx-filesystem--examples--Create_an_Amazon_FSx_for_Windows_File_Server_File_System_with_file_access_audit_event_logging_enabled)

- [Create an Amazon FSx for NetApp ONTAP File System](#aws-resource-fsx-filesystem--examples--Create_an_Amazon_FSx_for_NetApp_ONTAP_File_System)

### Create an Amazon FSx for Lustre File System

The following examples create a 1.2 TiB persistent Amazon FSx for Lustre file
system, with a `PerUnitStorageThroughput` of 200 MB/s/TiB.

#### JSON

```json

{
    "Resources": {
        "BasicS3LinkedLustreFileSystem": {
            "Type": "AWS::FSx::FileSystem",
            "Properties": {
                "FileSystemType": "LUSTRE",
                "FileSystemTypeVersion": "2.12",
                "StorageCapacity": 1200,
                "SubnetIds": [
                    {
                        "Fn::ImportValue": "MySubnet01"
                    }
                ],
                "SecurityGroupIds": [
                    {
                        "Fn::ImportValue": "LustreIngressSecurityGroupId"
                    }
                ],
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": "CFNs3linkedLustre"
                    }
                ],
                "LustreConfiguration": {
                    "AutoImportPolicy" : "NEW",
                    "CopyTagsToBackups" : true,
                    "DeploymentType": "PERSISTENT_1",
                    "PerUnitStorageThroughput": 200,
                    "DataCompressionType": "LZ4",
                    "ImportPath": {
                        "Fn::Join": [
                            "",
                            [
                                "s3://",
                                {
                                    "Fn::ImportValue": "LustreCFNS3ImportBucketName"
                                }
                            ]
                        ]
                    },
                    "ExportPath": {
                        "Fn::Join": [
                            "",
                            [
                                "s3://",
                                {
                                    "Fn::ImportValue": "LustreCFNS3ExportPath"
                                }
                            ]
                        ]
                    },
                    "WeeklyMaintenanceStartTime": "2:20:30"
                }
            }
        }
    },
    "Outputs": {
        "FileSystemId": {
            "Value": {
                "Ref": "BasicS3LinkedLustreFileSystem"
            }
        }
    }
}

```

#### YAML

```yaml

Resources:
  BasicS3LinkedLustreFileSystem:
    Type: AWS::FSx::FileSystem
    Properties:
      FileSystemType: "LUSTRE"
      FileSystemTypeVersion: "2.12"
      StorageCapacity: 1200
      SubnetIds: [!ImportValue MySubnet01]
      SecurityGroupIds: [!ImportValue LustreIngressSecurityGroupId]
      Tags:
        - Key: "Name"
          Value: "CFNs3linkedLustre"
      LustreConfiguration:
        AutoImportPolicy: "NEW"
        CopyTagsToBackups: true
        DeploymentType: "PERSISTENT_1"
        PerUnitStorageThroughput: 200
        DataCompressionType: "LZ4"
        ImportPath: !Join ["", ["s3://", !ImportValue LustreCFNS3ImportBucketName]]
        ExportPath: !Join ["", ["s3://", !ImportValue LustreCFNS3ExportPath]]
        WeeklyMaintenanceStartTime: "2:20:30"
Outputs:
  FileSystemId:
    Value: !Ref BasicS3LinkedLustreFileSystem
```

### Create an Amazon FSx for Windows File Server File System in a Self-managed Active Directory

The following examples create a Multi-AZ Amazon FSx for Windows File Server
file system joined to a self-managed active directory.

#### JSON

```json

{
    "Resources": {
        "WindowsSelfManagedADFileSystemWithAllConfigs": {
            "Type": "AWS::FSx::FileSystem",
            "Properties": {
                "FileSystemType": "WINDOWS",
                "StorageCapacity": 32,
                "StorageType": "SSD",
                "SubnetIds": [
                    {
                        "Fn::ImportValue": "MySubnet01"
                    },
                    {
                        "Fn::ImportValue": "MySubnet02"
                    }
                ],
                "SecurityGroupIds": [
                    {
                        "Fn::ImportValue": "WindowsIngressSecurityGroupId"
                    }
                ],
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": "windows"
                    }
                ],
                "WindowsConfiguration": {
                    "ThroughputCapacity": 8,
                    "Aliases": [
                        "financials.corp.example.com"
                    ],
                    "WeeklyMaintenanceStartTime": "4:16:30",
                    "DailyAutomaticBackupStartTime": "01:00",
                    "AutomaticBackupRetentionDays": 30,
                    "CopyTagsToBackups": false,
                    "DeploymentType": "MULTI_AZ_1",
                    "PreferredSubnetId": {
                        "Fn::ImportValue": "MySubnet01"
                    },
                    "SelfManagedActiveDirectoryConfiguration": {
                        "DnsIps": [
                            {
                                "Fn::Select": [
                                    0,
                                    {
                                        "Fn::Split": [
                                            ",",
                                            {
                                                "Fn::ImportValue": "MySelfManagedADDnsIpAddresses"
                                            }
                                        ]
                                    }
                                ]
                            }
                        ],
                        "DomainName": {
                            "Fn::ImportValue": "SelfManagedADDomainName"
                        },
                        "FileSystemAdministratorsGroup": "MyDomainAdminGroup",
                        "OrganizationalUnitDistinguishedName": "OU=FileSystems,DC=corp,DC=example,DC=com",
                        "UserName": "Admin",
                        "Password": {
                            "Fn::Join": [
                                ":",
                                [
                                    "{{resolve:secretsmanager",
                                    {
                                        "Fn::ImportValue": "MySelfManagedADCredentialName"
                                    },
                                    "SecretString}}"
                                ]
                            ]
                        }
                    }
                }
            }
        }
    },
    "Outputs": {
        "FileSystemId": {
            "Value": {
                "Ref": "WindowsSelfManagedADFileSystemWithAllConfigs"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  WindowsSelfManagedADFileSystemWithAllConfigs:
    Type: 'AWS::FSx::FileSystem'
    Properties:
      FileSystemType: WINDOWS
      StorageCapacity: 32
      StorageType: SSD
      SubnetIds:
        - !ImportValue MySubnet01
        - !ImportValue MySubnet02
      SecurityGroupIds:
        - !ImportValue WindowsIngressSecurityGroupId
      Tags:
        - Key: Name
          Value: windows
      WindowsConfiguration:
        ThroughputCapacity: 8
        Aliases:
            - financials.corp.example.com
        WeeklyMaintenanceStartTime: '4:16:30'
        DailyAutomaticBackupStartTime: '01:00'
        AutomaticBackupRetentionDays: 30
        CopyTagsToBackups: false
        DeploymentType: MULTI_AZ_1
        PreferredSubnetId: !ImportValue MySubnet01
        SelfManagedActiveDirectoryConfiguration:
          DnsIps:
            - !Select
              - 0
              - !Split
                - ','
                - !ImportValue MySelfManagedADDnsIpAddresses
          DomainName:
            'Fn::ImportValue': SelfManagedADDomainName
          FileSystemAdministratorsGroup: MyDomainAdminGroup
          OrganizationalUnitDistinguishedName: 'OU=FileSystems,DC=corp,DC=example,DC=com'
          UserName: Admin
          Password: !Join
            - ':'
            - - '{{resolve:secretsmanager'
              - !ImportValue MySelfManagedADCredentialName
              - 'SecretString}}'
Outputs:
  FileSystemId:
    Value: !Ref WindowsSelfManagedADFileSystemWithAllConfigs

```

### Create an Amazon FSx for Windows File Server File System in an Amazon Managed Active Directory

The following examples create a Multi-AZ Amazon FSx for Windows File Server
file system using HDD storage that is joined to an AWS Managed Active
Directory.

#### JSON

```json

{
    "Resources": {
        "WindowsMadFileSystemWithAllConfigs": {
            "Type": "AWS::FSx::FileSystem",
            "Properties": {
                "FileSystemType": "WINDOWS",
                "StorageCapacity": 2000,
                "StorageType": "HDD",
                "SubnetIds": [
                    {
                        "Fn::ImportValue": "CfnFsxMadSubnet01"
                    },
                    {
                        "Fn::ImportValue": "CfnFsxMadSubnet02"
                    }
                ],
                "SecurityGroupIds": [
                    {
                        "Fn::ImportValue": "WindowsIngressSecurityGroupId"
                    }
                ],
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": "windows"
                    }
                ],
                "WindowsConfiguration": {
                    "ActiveDirectoryId": {
                        "Fn::ImportValue": "CfnFsxMadDirectoryServiceId"
                    },
                    "ThroughputCapacity": 8,
                    "Aliases": [
                        "financials.corp.example.com"
                    ],
                    "WeeklyMaintenanceStartTime": "4:16:30",
                    "DailyAutomaticBackupStartTime": "01:00",
                    "AutomaticBackupRetentionDays": 90,
                    "CopyTagsToBackups": false,
                    "DeploymentType": "MULTI_AZ_1",
                    "PreferredSubnetId": {
                        "Fn::ImportValue": "CfnFsxMadSubnet01"
                    }
                }
            }
        }
    },
    "Outputs": {
        "FileSystemId": {
            "Value": {
                "Ref": "WindowsMadFileSystemWithAllConfigs"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  WindowsMadFileSystemWithAllConfigs:
    Type: 'AWS::FSx::FileSystem'
    Properties:
      FileSystemType: WINDOWS
      StorageCapacity: 2000
      StorageType: SSD
      SubnetIds:
        - !ImportValue CfnFsxMadSubnet01
        - !ImportValue CfnFsxMadSubnet02
      SecurityGroupIds:
        - !ImportValue WindowsIngressSecurityGroupId
      Tags:
        - Key: Name
          Value: windows
      WindowsConfiguration:
        ActiveDirectoryId: !ImportValue CfnFsxMadDirectoryServiceId
        ThroughputCapacity: 8
        Aliases:
            - financials.corp.example.com
        WeeklyMaintenanceStartTime: '4:16:30'
        DailyAutomaticBackupStartTime: '01:00'
        AutomaticBackupRetentionDays: 90
        DeploymentType: MULTI_AZ_1
        PreferredSubnetId: !ImportValue CfnFsxMadSubnet01
        CopyTagsToBackups: false
Outputs:
  FileSystemId:
    Value: !Ref WindowsMadFileSystemWithAllConfigs

```

### Create an Amazon FSx for Windows File Server File System with file access audit event logging enabled

The following examples create a Multi-AZ Amazon FSx for Windows File Server
file system with file access auditing enabled. Audit event logs are emitted when
end users make successful or failed attempts to access files, folders, and file
shares on the file system.

#### JSON

```json

{
    "Resources": {
        "WindowsMadFileSystemWithAllConfigs": {
            "Type": "Dev::FSx::FileSystem",
            "Properties": {
                "FileSystemType": "WINDOWS",
                "StorageCapacity": 100,
                "SubnetIds": [
                    {
                        "Fn::ImportValue": "CfnFsxMadSubnet01"
                    },
                    {
                        "Fn::ImportValue": "CfnFsxMadSubnet02"
                    }
                ],
                "SecurityGroupIds": [
                    {
                        "Fn::ImportValue": "CfnWindowsIngressSecurityGroupId"
                    }
                ],
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": "windows"
                    }
                ],
                "WindowsConfiguration": {
                    "ActiveDirectoryId": {
                        "Fn::ImportValue": "CfnFsxMadDirectoryServiceId"
                    },
                    "ThroughputCapacity": 32,
                    "DeploymentType": "MULTI_AZ_1",
                    "PreferredSubnetId": {
                        "Fn::ImportValue": "CfnFsxMadSubnet01"
                    }
                    "AuditLogConfiguration": {
                        "FileAccessAuditLogLevel": "SUCCESS_AND_FAILURE",
                        "FileShareAccessAuditLogLevel": "SUCCESS_AND_FAILURE",
                        "AuditLogDestination":
                        {
                          "Fn::Select": [
                            0,
                            {
                              "Fn::Split": [
                                ":*",
                                {
                                  "Fn::ImportValue": "CfnWindowsFileAccessAuditingLogGroupArn"
                                }
                              ]
                            }
                          ]
                        }
                    }
                }
            }
        }
    },
    "Outputs": {
        "FileSystemId": {
             "Value": {
                "Ref": "WindowsMadFileSystemWithAllConfigs"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  WindowsMadFileSystemWithAllConfigs:
    Type: "AWS::FSx::FileSystem"
    Properties:
      FileSystemType: "WINDOWS"
      StorageCapacity: 100
      SubnetIds:
        - !ImportValue CfnFsxMadSubnet01
        - !ImportValue CfnFsxMadSubnet02
      SecurityGroupIds:
        - !ImportValue WindowsIngressSecurityGroupId
      Tags:
        - Key: Name
          Value: windows
      WindowsConfiguration:
        ActiveDirectoryId: !ImportValue CfnFsxMadDirectoryServiceId
        ThroughputCapacity: 32
        DeploymentType: MULTI_AZ_1
        PreferredSubnetId: !ImportValue CfnFsxMadSubnet01
        AuditLogConfiguration:
          FileAccessAuditLogLevel: "SUCCESS_AND_FAILURE"
          FileShareAccessAuditLogLevel: "SUCCESS_AND_FAILURE"
          AuditLogDestination: !Select
            - 0
            - !Split
              - ':*'
              - 'Fn::ImportValue': WindowsFileAccessAuditingLogGroupArn

Outputs:
    FileSystemId:
      Value: !Ref WindowsMadFileSystemWithAllConfigs
```

### Create an Amazon FSx for NetApp ONTAP File System

The following examples create a 1 TiB Amazon FSx for NetApp ONTAP file
system.

#### JSON

```json

{
   "OntapMultiAzFileSystemWithAllConfigs": {
      "Type": "AWS::FSx::FileSystem",
      "Condition": "OntapMazEnabled",
      "Properties": {
         "FileSystemType": "ONTAP",
         "StorageCapacity": 1024,
         "StorageType": "SSD",
         "SubnetIds": [
         "subnet-1234567890abcdef0",
         "subnet-abcdef01234567890"
         ],
         "SecurityGroupIds": [
         "sg-021345abcdef6789"
         ],
         "OntapConfiguration": {
            "AutomaticBackupRetentionDays": 3,
            "DailyAutomaticBackupStartTime": "07:00",
            "DeploymentType": "MULTI_AZ_1",
            "DiskIopsConfiguration": {
               "Iops": 10000,
               "Mode": "USER_PROVISIONED"
            },
            "PreferredSubnetId": "subnet-abcdef01234567890",
            "RouteTableIds": [
            "rtb-abcdef01234567890"
            ],
            "ThroughputCapacity": 512,
            "WeeklyMaintenanceStartTime": "4:16:30"
         },
         "Tags": [
            {
               "Key": "Name",
               "Value": "OntapFileSystem_MAZ"
            }
         ]
      }
   }
}
```

#### YAML

```yaml

OntapMultiAzFileSystemWithAllConfigs:
   Type: "AWS::FSx::FileSystem"
   Condition: OntapMazEnabled
   Properties:
      FileSystemType: "ONTAP"
      StorageCapacity: 1024
      StorageType: SSD
      SubnetIds: ["subnet-1234567890abcdef0", "subnet-abcdef01234567890"]
      SecurityGroupIds: ["sg-068ec2a396aa7c2c8"]
      OntapConfiguration:
         AutomaticBackupRetentionDays: 3
         DailyAutomaticBackupStartTime: "07:00"
         DeploymentType: "MULTI_AZ_1"
         DiskIopsConfiguration:
            Iops: 10000
            Mode: "USER_PROVISIONED"
         PreferredSubnetId: "subnet-abcdef01234567890"
         RouteTableIds: ["rtb-abcdef01234567890"]
         ThroughputCapacity: 512
         WeeklyMaintenanceStartTime: "4:16:30"
      Tags:
        - Key: "Name"
          Value: "OntapFileSystem_MAZ"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AuditLogConfiguration
