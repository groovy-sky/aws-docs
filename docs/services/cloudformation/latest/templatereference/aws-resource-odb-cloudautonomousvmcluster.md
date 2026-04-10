This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudAutonomousVmCluster

The `AWS::ODB::CloudAutonomousVmCluster` resource creates an Autonomous VM cluster. An Autonomous VM cluster provides the infrastructure for running Autonomous Databases.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ODB::CloudAutonomousVmCluster",
  "Properties" : {
      "AutonomousDataStorageSizeInTBs" : Number,
      "CloudExadataInfrastructureId" : String,
      "CpuCoreCountPerNode" : Integer,
      "DbServers" : [ String, ... ],
      "Description" : String,
      "DisplayName" : String,
      "IamRoles" : [ IamRole, ... ],
      "IsMtlsEnabledVmCluster" : Boolean,
      "LicenseModel" : String,
      "MaintenanceWindow" : MaintenanceWindow,
      "MemoryPerOracleComputeUnitInGBs" : Integer,
      "OdbNetworkId" : String,
      "ScanListenerPortNonTls" : Integer,
      "ScanListenerPortTls" : Integer,
      "Tags" : [ Tag, ... ],
      "TimeZone" : String,
      "TotalContainerDatabases" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::ODB::CloudAutonomousVmCluster
Properties:
  AutonomousDataStorageSizeInTBs: Number
  CloudExadataInfrastructureId: String
  CpuCoreCountPerNode: Integer
  DbServers:
    - String
  Description: String
  DisplayName: String
  IamRoles:
    - IamRole
  IsMtlsEnabledVmCluster: Boolean
  LicenseModel: String
  MaintenanceWindow:
    MaintenanceWindow
  MemoryPerOracleComputeUnitInGBs: Integer
  OdbNetworkId: String
  ScanListenerPortNonTls: Integer
  ScanListenerPortTls: Integer
  Tags:
    - Tag
  TimeZone: String
  TotalContainerDatabases: Integer

```

## Properties

`AutonomousDataStorageSizeInTBs`

The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB.

Required when creating an Autonomous VM cluster.

_Required_: Conditional

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CloudExadataInfrastructureId`

The unique identifier of the Cloud Exadata Infrastructure containing this Autonomous VM cluster.

Required when creating an Autonomous VM cluster.

_Required_: Conditional

_Type_: String

_Pattern_: `(arn:(?:aws|aws-cn|aws-us-gov|aws-iso-{0,1}[a-z]{0,1}):[a-z0-9-]+:[a-z0-9-]*:[0-9]+:[a-z0-9-]+/[a-zA-Z0-9_~.-]{6,64}|[a-zA-Z0-9_~.-]{6,64})`

_Minimum_: `6`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CpuCoreCountPerNode`

The number of CPU cores enabled per node in the Autonomous VM cluster.

Required when creating an Autonomous VM cluster.

_Required_: Conditional

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DbServers`

The list of database servers associated with the Autonomous VM cluster.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The user-provided description of the Autonomous VM cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisplayName`

The display name of the Autonomous VM cluster.

Required when creating an Autonomous VM cluster.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z_](?!.*--)[a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IamRoles`

The AWS Identity and Access Management (IAM) service roles associated with the Autonomous VM cluster.

_Required_: No

_Type_: Array of [IamRole](aws-properties-odb-cloudautonomousvmcluster-iamrole.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsMtlsEnabledVmCluster`

Specifies whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LicenseModel`

The Oracle license model that applies to the Autonomous VM cluster. Valid values are `LICENSE_INCLUDED` or `BRING_YOUR_OWN_LICENSE`.

_Required_: No

_Type_: String

_Allowed values_: `BRING_YOUR_OWN_LICENSE | LICENSE_INCLUDED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaintenanceWindow`

The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window.

_Required_: No

_Type_: [MaintenanceWindow](aws-properties-odb-cloudautonomousvmcluster-maintenancewindow.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemoryPerOracleComputeUnitInGBs`

The amount of memory allocated per Oracle Compute Unit, in GB.

Required when creating an Autonomous VM cluster.

_Required_: Conditional

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OdbNetworkId`

The unique identifier of the ODB network associated with this Autonomous VM cluster.

Required when creating an Autonomous VM cluster.

_Required_: Conditional

_Type_: String

_Pattern_: `(arn:(?:aws|aws-cn|aws-us-gov|aws-iso-{0,1}[a-z]{0,1}):[a-z0-9-]+:[a-z0-9-]*:[0-9]+:[a-z0-9-]+/[a-zA-Z0-9_~.-]{6,64}|[a-zA-Z0-9_~.-]{6,64})`

_Minimum_: `6`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScanListenerPortNonTls`

The SCAN listener port for non-TLS (TCP) protocol. The default is 1521.

_Required_: No

_Type_: Integer

_Minimum_: `1024`

_Maximum_: `8999`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScanListenerPortTls`

The SCAN listener port for TLS (TCP) protocol. The default is 2484.

_Required_: No

_Type_: Integer

_Minimum_: `1024`

_Maximum_: `8999`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to assign to the Autonomous Vm Cluster.

_Required_: No

_Type_: Array of [Tag](aws-properties-odb-cloudautonomousvmcluster-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The time zone of the Autonomous VM cluster.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TotalContainerDatabases`

The total number of Autonomous Container Databases that can be created with the allocated local storage.

Required when creating an Autonomous VM cluster.

_Required_: Conditional

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier of the Autonomous VM cluster. For example:

`{ "Ref": "myAutonomousVmCluster" }`

For the Autonomous VM cluster `myAutonomousVmCluster`, `Ref` returns the unique identifier of the Autonomous VM cluster.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`AutonomousDataStoragePercentage`

The percentage of data storage currently in use for Autonomous Databases in the Autonomous VM cluster.

`AvailableAutonomousDataStorageSizeInTBs`

The available data storage space for Autonomous Databases in the Autonomous VM cluster, in TB.

`AvailableContainerDatabases`

The number of Autonomous CDBs that you can create with the currently available storage.

`AvailableCpus`

The number of CPU cores available for allocation to Autonomous Databases.

`CloudAutonomousVmClusterArn`

The Amazon Resource Name (ARN) for the Autonomous VM cluster.

`CloudAutonomousVmClusterId`

The unique identifier of the Autonomous VM cluster.

`ComputeModel`

The compute model of the Autonomous VM cluster: ECPU or OCPU.

`CpuCoreCount`

The total number of CPU cores in the Autonomous VM cluster.

`CpuPercentage`

The percentage of total CPU cores currently in use in the Autonomous VM cluster.

`DataStorageSizeInGBs`

The total data storage allocated to the Autonomous VM cluster, in GB.

`DataStorageSizeInTBs`

The total data storage allocated to the Autonomous VM cluster, in TB.

`DbNodeStorageSizeInGBs`

The local node storage allocated to the Autonomous VM cluster, in gigabytes (GB).

`Domain`

The domain name for the Autonomous VM cluster.

`ExadataStorageInTBsLowestScaledValue`

The minimum value to which you can scale down the Exadata storage, in TB.

`Hostname`

The hostname for the Autonomous VM cluster.

`MaxAcdsLowestScaledValue`

The minimum value to which you can scale down the maximum number of Autonomous CDBs.

`MemorySizeInGBs`

The total amount of memory allocated to the Autonomous VM cluster, in gigabytes (GB).

`NodeCount`

The number of database server nodes in the Autonomous VM cluster.

`NonProvisionableAutonomousContainerDatabases`

The number of Autonomous CDBs that can't be provisioned because of resource constraints.

`Ocid`

The Oracle Cloud Identifier (OCID) of the Autonomous VM cluster.

`OciResourceAnchorName`

The name of the OCI resource anchor associated with this Autonomous VM cluster.

`OciUrl`

The URL for accessing the OCI console page for this Autonomous VM cluster.

`ProvisionableAutonomousContainerDatabases`

The number of Autonomous CDBs that can be provisioned in the Autonomous VM cluster.

`ProvisionedAutonomousContainerDatabases`

The number of Autonomous CDBs currently provisioned in the Autonomous VM cluster.

`ProvisionedCpus`

The number of CPU cores currently provisioned in the Autonomous VM cluster.

`ReclaimableCpus`

The number of CPU cores that can be reclaimed from terminated or scaled-down Autonomous Databases.

`ReservedCpus`

The number of CPU cores reserved for system operations and redundancy.

`Shape`

The shape of the Exadata infrastructure for the Autonomous VM cluster.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Oracle Database@AWS

IamRole

All content copied from https://docs.aws.amazon.com/.
