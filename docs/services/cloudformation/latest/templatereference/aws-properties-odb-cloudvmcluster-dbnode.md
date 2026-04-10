This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudVmCluster DbNode

Information about a DB node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackupIpId" : String,
  "BackupVnic2Id" : String,
  "CpuCoreCount" : Integer,
  "DbNodeArn" : String,
  "DbNodeId" : String,
  "DbNodeStorageSizeInGBs" : Integer,
  "DbServerId" : String,
  "DbSystemId" : String,
  "HostIpId" : String,
  "Hostname" : String,
  "MemorySizeInGBs" : Integer,
  "Ocid" : String,
  "Status" : String,
  "Tags" : [ Tag, ... ],
  "Vnic2Id" : String,
  "VnicId" : String
}

```

### YAML

```yaml

  BackupIpId: String
  BackupVnic2Id: String
  CpuCoreCount: Integer
  DbNodeArn: String
  DbNodeId: String
  DbNodeStorageSizeInGBs: Integer
  DbServerId: String
  DbSystemId: String
  HostIpId: String
  Hostname: String
  MemorySizeInGBs: Integer
  Ocid: String
  Status: String
  Tags:
    - Tag
  Vnic2Id: String
  VnicId: String

```

## Properties

`BackupIpId`

The Oracle Cloud ID (OCID) of the backup IP address that's associated with the DB node.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`BackupVnic2Id`

The OCID of the second backup VNIC.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`CpuCoreCount`

Number of CPU cores enabled on the DB node.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DbNodeArn`

The Amazon Resource Name (ARN) of the DB node.

_Required_: No

_Type_: String

_Pattern_: `arn:(?:aws|aws-cn|aws-us-gov|aws-iso-{0,1}[a-z]{0,1}):[a-z0-9-]+:[a-z0-9-]*:[0-9]+:[a-z0-9-]+/[a-z0-9-_]{6,64}`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DbNodeId`

The unique identifier of the DB node.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_~.-]+`

_Minimum_: `6`

_Maximum_: `64`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DbNodeStorageSizeInGBs`

The amount of local node storage, in gigabytes (GBs), that's allocated on the DB node.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DbServerId`

The unique identifier of the Db server that is associated with the DB node.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_~.-]+`

_Minimum_: `6`

_Maximum_: `64`

`DbSystemId`

The OCID of the DB system.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`HostIpId`

The OCID of the host IP address that's associated with the DB node.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Hostname`

The host name for the DB node.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MemorySizeInGBs`

The allocated memory in GBs on the DB node.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Ocid`

The OCID of the DB node.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Status`

The current status of the DB node.

_Required_: No

_Type_: String

_Allowed values_: `AVAILABLE | FAILED | PROVISIONING | TERMINATED | TERMINATING | UPDATING | STOPPING | STOPPED | STARTING`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-odb-cloudvmcluster-tag.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Vnic2Id`

The OCID of the second VNIC.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VnicId`

The OCID of the VNIC.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataCollectionOptions

IamRole

All content copied from https://docs.aws.amazon.com/.
