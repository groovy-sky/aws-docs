This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudVmCluster

The `AWS::ODB::CloudVmCluster` resource creates a VM cluster on the specified Exadata infrastructure in the Oracle Database. A VM cluster provides the compute resources for Oracle Database workloads.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ODB::CloudVmCluster",
  "Properties" : {
      "CloudExadataInfrastructureId" : String,
      "ClusterName" : String,
      "CpuCoreCount" : Integer,
      "DataCollectionOptions" : DataCollectionOptions,
      "DataStorageSizeInTBs" : Number,
      "DbNodes" : [ DbNode, ... ],
      "DbNodeStorageSizeInGBs" : Integer,
      "DbServers" : [ String, ... ],
      "DisplayName" : String,
      "GiVersion" : String,
      "Hostname" : String,
      "IamRoles" : [ IamRole, ... ],
      "IsLocalBackupEnabled" : Boolean,
      "IsSparseDiskgroupEnabled" : Boolean,
      "LicenseModel" : String,
      "MemorySizeInGBs" : Integer,
      "OdbNetworkId" : String,
      "ScanListenerPortTcp" : Integer,
      "SshPublicKeys" : [ String, ... ],
      "SystemVersion" : String,
      "Tags" : [ Tag, ... ],
      "TimeZone" : String
    }
}

```

### YAML

```yaml

Type: AWS::ODB::CloudVmCluster
Properties:
  CloudExadataInfrastructureId: String
  ClusterName: String
  CpuCoreCount: Integer
  DataCollectionOptions:
    DataCollectionOptions
  DataStorageSizeInTBs: Number
  DbNodes:
    - DbNode
  DbNodeStorageSizeInGBs: Integer
  DbServers:
    - String
  DisplayName: String
  GiVersion: String
  Hostname: String
  IamRoles:
    - IamRole
  IsLocalBackupEnabled: Boolean
  IsSparseDiskgroupEnabled: Boolean
  LicenseModel: String
  MemorySizeInGBs: Integer
  OdbNetworkId: String
  ScanListenerPortTcp: Integer
  SshPublicKeys:
    - String
  SystemVersion: String
  Tags:
    - Tag
  TimeZone: String

```

## Properties

`CloudExadataInfrastructureId`

The unique identifier of the Exadata infrastructure that this VM cluster belongs to.

Required when creating a VM cluster.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterName`

The name of the Grid Infrastructure (GI) cluster.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9-]*$`

_Minimum_: `1`

_Maximum_: `11`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CpuCoreCount`

The number of CPU cores enabled on the VM cluster.

Required when creating a VM cluster.

_Required_: Conditional

_Type_: Integer

_Minimum_: `0`

_Maximum_: `368`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataCollectionOptions`

The set of diagnostic collection options enabled for the VM cluster.

_Required_: No

_Type_: [DataCollectionOptions](aws-properties-odb-cloudvmcluster-datacollectionoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataStorageSizeInTBs`

The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DbNodes`

Property description not available.

_Required_: No

_Type_: Array of [DbNode](aws-properties-odb-cloudvmcluster-dbnode.md)

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DbNodeStorageSizeInGBs`

The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DbServers`

The list of database servers for the VM cluster.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisplayName`

The user-friendly name for the VM cluster.

Required when creating a VM cluster.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z_](?!.*--)[a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GiVersion`

The software version of the Oracle Grid Infrastructure (GI) for the VM cluster.

Required when creating a VM cluster.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Hostname`

The host name for the VM cluster.

Required when creating a VM cluster.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9-]*[a-zA-Z0-9]$`

_Minimum_: `1`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IamRoles`

The AWS Identity and Access Management (IAM) service roles associated with the VM cluster.

_Required_: No

_Type_: Array of [IamRole](aws-properties-odb-cloudvmcluster-iamrole.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsLocalBackupEnabled`

Specifies whether database backups to local Exadata storage are enabled for the VM cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IsSparseDiskgroupEnabled`

Specifies whether the VM cluster is configured with a sparse disk group.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LicenseModel`

The Oracle license model applied to the VM cluster.

_Required_: No

_Type_: String

_Allowed values_: `BRING_YOUR_OWN_LICENSE | LICENSE_INCLUDED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemorySizeInGBs`

The amount of memory, in gigabytes (GB), that's allocated for the VM cluster.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OdbNetworkId`

The unique identifier of the ODB network for the VM cluster.

Required when creating a VM cluster.

_Required_: Conditional

_Type_: String

_Pattern_: `(arn:(?:aws|aws-cn|aws-us-gov|aws-iso-{0,1}[a-z]{0,1}):[a-z0-9-]+:[a-z0-9-]*:[0-9]+:[a-z0-9-]+/[a-zA-Z0-9_~.-]{6,64}|[a-zA-Z0-9_~.-]{6,64})`

_Minimum_: `6`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScanListenerPortTcp`

The port number for TCP connections to the single client access name (SCAN) listener.

Valid values: `1024–8999` with the following exceptions: `2484`,
`6100`, `6200`, `7060`, `7070`,
`7085`, and `7879`

Default: `1521`

_Required_: No

_Type_: Integer

_Minimum_: `1024`

_Maximum_: `8999`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SshPublicKeys`

The public key portion of one or more key pairs used for SSH access to the VM cluster.

Required when creating a VM cluster.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SystemVersion`

The operating system version of the image chosen for the VM cluster.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to assign to the Vm Cluster.

_Required_: No

_Type_: Array of [Tag](aws-properties-odb-cloudvmcluster-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The time zone of the VM cluster.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier of the VM cluster. For example:

`{ "Ref": "myVmCluster" }`

For the VM cluster `myVmCluster`, `Ref` returns the unique identifier of the VM cluster.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`CloudVmClusterArn`

The Amazon Resource Name (ARN) of the VM cluster.

`CloudVmClusterId`

The unique identifier of the VM cluster.

`ComputeModel`

The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An
ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores
elastically allocated from a pool of compute and storage servers. An OCPU is a legacy
physical measure of compute resources. OCPUs are based on the physical core of a processor
with hyper-threading enabled.

`DiskRedundancy`

The type of redundancy configured for the VM cluster. `NORMAL` is 2-way redundancy. `HIGH` is 3-way redundancy.

`Domain`

The domain of the VM cluster.

`ListenerPort`

The port number configured for the listener on the VM cluster.

`NodeCount`

The number of nodes in the VM cluster.

`Ocid`

The OCID of the VM cluster.

`OciResourceAnchorName`

The name of the OCI resource anchor for the VM cluster.

`OciUrl`

The HTTPS link to the VM cluster in OCI.

`ScanDnsName`

The FQDN of the DNS record for the Single Client Access Name (SCAN) IP addresses that are associated with the VM cluster.

`ScanIpIds`

The OCID of the SCAN IP addresses that are associated with the VM cluster.

`Shape`

The hardware model name of the Exadata infrastructure that's running the VM cluster.

`StorageSizeInGBs`

The amount of local node storage, in gigabytes (GB), that's allocated to the VM cluster.

`VipIds`

The virtual IP (VIP) addresses that are associated with the VM cluster. Oracle's Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the VM cluster to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DataCollectionOptions

All content copied from https://docs.aws.amazon.com/.
