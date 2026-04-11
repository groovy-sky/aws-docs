This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem OpenZFSConfiguration

The OpenZFS configuration for the file system that's being created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutomaticBackupRetentionDays" : Integer,
  "CopyTagsToBackups" : Boolean,
  "CopyTagsToVolumes" : Boolean,
  "DailyAutomaticBackupStartTime" : String,
  "DeploymentType" : String,
  "DiskIopsConfiguration" : DiskIopsConfiguration,
  "EndpointIpAddressRange" : String,
  "EndpointIpv6AddressRange" : String,
  "Options" : [ String, ... ],
  "PreferredSubnetId" : String,
  "ReadCacheConfiguration" : ReadCacheConfiguration,
  "RootVolumeConfiguration" : RootVolumeConfiguration,
  "RouteTableIds" : [ String, ... ],
  "ThroughputCapacity" : Integer,
  "WeeklyMaintenanceStartTime" : String
}

```

### YAML

```yaml

  AutomaticBackupRetentionDays: Integer
  CopyTagsToBackups: Boolean
  CopyTagsToVolumes: Boolean
  DailyAutomaticBackupStartTime: String
  DeploymentType: String
  DiskIopsConfiguration:
    DiskIopsConfiguration
  EndpointIpAddressRange: String
  EndpointIpv6AddressRange: String
  Options:
    - String
  PreferredSubnetId: String
  ReadCacheConfiguration:
    ReadCacheConfiguration
  RootVolumeConfiguration:
    RootVolumeConfiguration
  RouteTableIds:
    - String
  ThroughputCapacity: Integer
  WeeklyMaintenanceStartTime: String

```

## Properties

`AutomaticBackupRetentionDays`

The number of days to retain automatic backups. Setting this property to
`0` disables automatic backups. You can retain automatic backups for a
maximum of 90 days. The default is `30`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTagsToBackups`

A Boolean value indicating whether tags for the file system should be copied to
backups. This value defaults to `false`. If it's set to `true`,
all tags for the file system are copied to all automatic and user-initiated backups
where the user doesn't specify tags. If this value is `true`, and you specify
one or more tags, only the specified tags are copied to backups. If you specify one or
more tags when creating a user-initiated backup, no tags are copied from the file
system, regardless of this value.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTagsToVolumes`

A Boolean value indicating whether tags for the file system should be copied to volumes.
This value defaults to `false`. If it's set to `true`, all tags
for the file system are copied to volumes where the user doesn't specify tags. If this
value is `true`, and you specify one or more tags, only the specified tags
are copied to volumes. If you specify one or more tags when creating the volume, no
tags are copied from the file system, regardless of this value.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DailyAutomaticBackupStartTime`

A recurring daily time, in the format `HH:MM`. `HH` is the
zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the
hour. For example, `05:00` specifies 5 AM daily.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentType`

Specifies the file system deployment type. Valid values are the following:

- `MULTI_AZ_1`\- Creates file systems with high availability and durability by replicating your data and supporting failover across multiple Availability Zones in the same AWS Region.

- `SINGLE_AZ_HA_2`\- Creates file systems with high availability and throughput capacities of 160 - 10,240 MB/s using an NVMe L2ARC cache by deploying a primary and standby file system within the same Availability Zone.

- `SINGLE_AZ_HA_1`\- Creates file systems with high availability and throughput capacities of 64 - 4,096 MB/s by deploying a primary and standby file system within the same Availability Zone.

- `SINGLE_AZ_2`\- Creates file systems with throughput capacities of 160 - 10,240 MB/s
using an NVMe L2ARC cache that automatically recover within a single Availability Zone.

- `SINGLE_AZ_1`\- Creates file systems with throughput capacities of 64 - 4,096 MBs that automatically recover within a single Availability Zone.

For a list of which AWS Regions each deployment type is available in, see [Deployment type availability](../../../fsx/latest/openzfsguide/availability-durability.md#available-aws-regions).
For more information on the differences in performance between deployment types, see [File system performance](../../../fsx/latest/openzfsguide/performance.md#zfs-fs-performance)
in the _Amazon FSx for OpenZFS User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `SINGLE_AZ_1 | SINGLE_AZ_2 | SINGLE_AZ_HA_1 | SINGLE_AZ_HA_2 | MULTI_AZ_1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DiskIopsConfiguration`

The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP, Amazon FSx for Windows File Server, or FSx for OpenZFS file system. By default, Amazon FSx
automatically provisions 3 IOPS per GB of storage capacity. You can provision additional IOPS per
GB of storage. The configuration consists of the total number of provisioned SSD IOPS
and how it is was provisioned, or the mode (by the customer or by Amazon FSx).

_Required_: No

_Type_: [DiskIopsConfiguration](aws-properties-fsx-filesystem-diskiopsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointIpAddressRange`

(Multi-AZ only) Specifies the IPv4 address range in which the endpoints to access your
file system will be created. By default in the Amazon FSx API and Amazon FSx console, Amazon FSx
selects an available /28 IP address range for you from one of the VPC's CIDR ranges.
You can have overlapping endpoint IP addresses for file systems deployed in the
same VPC/route tables, as long as they don't overlap with any subnet.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{9,17}$`

_Minimum_: `9`

_Maximum_: `17`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointIpv6AddressRange`

(Multi-AZ only) Specifies the IP address range in which the endpoints to access
your file system will be created. By default in the Amazon FSx API and Amazon FSx console,
Amazon FSx selects an available /118 IP address
range for you from one of the VPC's CIDR ranges. You can have overlapping endpoint
IP addresses for file systems deployed in the same VPC/route tables, as long as they
don't overlap with any subnet.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{4,43}$`

_Minimum_: `4`

_Maximum_: `43`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Options`

To delete a file system if there are child volumes present below the root volume,
use the string `DELETE_CHILD_VOLUMES_AND_SNAPSHOTS`. If your file system
has child volumes and you don't use this option, the delete request will fail.

_Required_: No

_Type_: Array of String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredSubnetId`

Required when `DeploymentType` is set to `MULTI_AZ_1`. This specifies the subnet in which you want the preferred file server to be located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReadCacheConfiguration`

Specifies the optional provisioned SSD read cache on file systems that use the Intelligent-Tiering storage class.

_Required_: No

_Type_: [ReadCacheConfiguration](aws-properties-fsx-filesystem-readcacheconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RootVolumeConfiguration`

The configuration Amazon FSx uses when creating the root value of the Amazon FSx for OpenZFS
file system. All volumes are children of the root volume.

_Required_: No

_Type_: [RootVolumeConfiguration](aws-properties-fsx-filesystem-rootvolumeconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteTableIds`

(Multi-AZ only) Specifies the route tables in which Amazon FSx creates the rules
for routing traffic to the correct file server. You should specify all virtual private cloud
(VPC) route tables associated with the subnets in which your clients are located. By default,
Amazon FSx selects your VPC's default route table.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThroughputCapacity`

Specifies the throughput of an Amazon FSx for OpenZFS file system, measured in megabytes per second (MBps). Required if
you are creating a new file system.

Valid values depend on the `DeploymentType` that you choose, as follows:

- For `MULTI_AZ_1` and `SINGLE_AZ_2`, valid values are 160, 320, 640,
1280, 2560, 3840, 5120, 7680, or 10240 MBps.

- For `SINGLE_AZ_1`, valid values are 64, 128, 256, 512, 1024, 2048, 3072, or 4096 MBps.

You pay for additional throughput capacity that you provision.

_Required_: No

_Type_: Integer

_Minimum_: `8`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeeklyMaintenanceStartTime`

The preferred start time to perform weekly maintenance, formatted d:HH:MM in the UTC
time zone, where d is the weekday number, from 1 through 7, beginning with Monday and ending with Sunday.

For example, `1:05:00` specifies maintenance at 5 AM Monday.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OntapConfiguration

ReadCacheConfiguration

All content copied from https://docs.aws.amazon.com/.
