This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem WindowsConfiguration

The Microsoft Windows configuration for the file system that's being created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActiveDirectoryId" : String,
  "Aliases" : [ String, ... ],
  "AuditLogConfiguration" : AuditLogConfiguration,
  "AutomaticBackupRetentionDays" : Integer,
  "CopyTagsToBackups" : Boolean,
  "DailyAutomaticBackupStartTime" : String,
  "DeploymentType" : String,
  "DiskIopsConfiguration" : DiskIopsConfiguration,
  "FsrmConfiguration" : FsrmConfiguration,
  "PreferredSubnetId" : String,
  "SelfManagedActiveDirectoryConfiguration" : SelfManagedActiveDirectoryConfiguration,
  "ThroughputCapacity" : Integer,
  "WeeklyMaintenanceStartTime" : String
}

```

### YAML

```yaml

  ActiveDirectoryId: String
  Aliases:
    - String
  AuditLogConfiguration:
    AuditLogConfiguration
  AutomaticBackupRetentionDays: Integer
  CopyTagsToBackups: Boolean
  DailyAutomaticBackupStartTime: String
  DeploymentType: String
  DiskIopsConfiguration:
    DiskIopsConfiguration
  FsrmConfiguration:
    FsrmConfiguration
  PreferredSubnetId: String
  SelfManagedActiveDirectoryConfiguration:
    SelfManagedActiveDirectoryConfiguration
  ThroughputCapacity: Integer
  WeeklyMaintenanceStartTime: String

```

## Properties

`ActiveDirectoryId`

The ID for an existing AWS Managed Microsoft Active Directory (AD)
instance that the file system should join when it's created. Required if you are joining the file system to an existing
AWS Managed Microsoft AD.

_Required_: Conditional

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Aliases`

An array of one or more DNS alias names that you want to associate with the Amazon FSx
file system. Aliases allow you to use existing DNS names to access the data in your
Amazon FSx file system. You can associate up to 50 aliases with a file system at any
time.

For more information, see [Working with DNS\
Aliases](../../../fsx/latest/windowsguide/managing-dns-aliases.md) and [Walkthrough 5: Using DNS aliases to access your file system](../../../fsx/latest/windowsguide/walkthrough05-file-system-custom-cname.md), including
additional steps you must take to be able to access your file system using a DNS
alias.

An alias name has to meet the following requirements:

- Formatted as a fully-qualified domain name (FQDN),
`hostname.domain`, for example,
`accounting.example.com`.

- Can contain alphanumeric characters, the underscore (\_), and the hyphen
(-).

- Cannot start or end with a hyphen.

- Can start with a numeric.

For DNS alias names, Amazon FSx stores alphabetical characters as lowercase letters
(a-z), regardless of how you specify them: as uppercase letters, lowercase letters, or
the corresponding letters in escape codes.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuditLogConfiguration`

The configuration that Amazon FSx for Windows File Server uses to audit and log
user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server
file system.

_Required_: No

_Type_: [AuditLogConfiguration](aws-properties-fsx-filesystem-auditlogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticBackupRetentionDays`

The number of days to retain automatic backups. Setting this property to
`0` disables automatic backups. You can retain automatic backups for a
maximum of 90 days. The default is `30`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTagsToBackups`

A boolean flag indicating whether tags for the file system should be copied to
backups. This value defaults to false. If it's set to true, all tags for the file
system are copied to all automatic and user-initiated backups where the user
doesn't specify tags. If this value is true, and you specify one or more tags, only
the specified tags are copied to backups. If you specify one or more tags when creating a
user-initiated backup, no tags are copied from the file system, regardless of this value.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DailyAutomaticBackupStartTime`

A recurring daily time, in the format `HH:MM`. `HH` is the
zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the
hour. For example, `05:00` specifies 5 AM daily.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentType`

Specifies the file system deployment type, valid values are the following:

- `MULTI_AZ_1` \- Deploys a high availability file system that is configured
for Multi-AZ redundancy to tolerate temporary Availability Zone (AZ) unavailability. You
can only deploy a Multi-AZ file system in AWS Regions that have a minimum of three Availability Zones. Also
supports HDD storage type

- `SINGLE_AZ_1` \- (Default) Choose to deploy a file system that is configured for single AZ redundancy.

- `SINGLE_AZ_2` \- The latest generation Single AZ file system.
Specifies a file system that is configured for single AZ redundancy and supports HDD storage type.

For more information, see
[Availability and Durability: Single-AZ and Multi-AZ File Systems](../../../fsx/latest/windowsguide/high-availability-multiaz.md).

_Required_: No

_Type_: String

_Allowed values_: `MULTI_AZ_1 | SINGLE_AZ_1 | SINGLE_AZ_2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DiskIopsConfiguration`

The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for Windows file system.
By default, Amazon FSx automatically provisions 3 IOPS per GiB of storage capacity. You can provision additional
IOPS per GiB of storage, up to the maximum limit associated with your chosen throughput capacity.

_Required_: No

_Type_: [DiskIopsConfiguration](aws-properties-fsx-filesystem-diskiopsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FsrmConfiguration`

Property description not available.

_Required_: No

_Type_: [FsrmConfiguration](aws-properties-fsx-filesystem-fsrmconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredSubnetId`

Required when `DeploymentType` is set to `MULTI_AZ_1`. This
specifies the subnet in which you want the preferred file server to be located. For
in-AWS applications, we recommend that you launch your clients in the
same availability zone as your preferred file server to reduce cross-availability zone data
transfer costs and minimize latency.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedActiveDirectoryConfiguration`

The configuration that Amazon FSx uses to join a FSx for Windows File Server file system or an FSx for ONTAP storage virtual machine (SVM) to
a self-managed (including on-premises) Microsoft Active Directory (AD)
directory. For more information, see
[Using Amazon FSx for Windows with your self-managed Microsoft Active Directory](../../../fsx/latest/windowsguide/self-managed-ad.md) or
[Managing FSx for ONTAP SVMs](../../../fsx/latest/ontapguide/managing-svms.md).

_Required_: No

_Type_: [SelfManagedActiveDirectoryConfiguration](aws-properties-fsx-filesystem-selfmanagedactivedirectoryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThroughputCapacity`

Sets the throughput capacity of an Amazon FSx file system, measured in megabytes per
second (MB/s), in 2 to the _n_ th increments, between 2^3 (8) and 2^11
(2048).

###### Note

To increase storage capacity, a file system must have a minimum throughput capacity
of 16 MB/s.

_Required_: Conditional

_Type_: Integer

_Minimum_: `8`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeeklyMaintenanceStartTime`

The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC
time zone, where d is the weekday number, from 1 through 7, beginning with Monday and ending with Sunday.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserAndGroupQuotas

AWS::FSx::S3AccessPointAttachment

All content copied from https://docs.aws.amazon.com/.
