This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudExadataInfrastructure

The `AWS::ODB::CloudExadataInfrastructure` resource creates an Exadata infrastructure. An Exadata infrastructure provides the underlying compute and storage resources for Oracle Database workloads.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ODB::CloudExadataInfrastructure",
  "Properties" : {
      "AvailabilityZone" : String,
      "AvailabilityZoneId" : String,
      "ComputeCount" : Integer,
      "CustomerContactsToSendToOCI" : [ CustomerContact, ... ],
      "DatabaseServerType" : String,
      "DisplayName" : String,
      "MaintenanceWindow" : MaintenanceWindow,
      "Shape" : String,
      "StorageCount" : Integer,
      "StorageServerType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ODB::CloudExadataInfrastructure
Properties:
  AvailabilityZone: String
  AvailabilityZoneId: String
  ComputeCount: Integer
  CustomerContactsToSendToOCI:
    - CustomerContact
  DatabaseServerType: String
  DisplayName: String
  MaintenanceWindow:
    MaintenanceWindow
  Shape: String
  StorageCount: Integer
  StorageServerType: String
  Tags:
    - Tag

```

## Properties

`AvailabilityZone`

The name of the Availability Zone (AZ) where the Exadata infrastructure is located.

Required when creating an Exadata infrastructure. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the infrastructure.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZoneId`

The AZ ID of the AZ where the Exadata infrastructure is located.

Required when creating an Exadata infrastructure. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the infrastructure.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComputeCount`

The number of database servers for the Exadata infrastructure.

Required when creating an Exadata infrastructure.

_Required_: Conditional

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomerContactsToSendToOCI`

The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.

_Required_: No

_Type_: Array of [CustomerContact](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-odb-cloudexadatainfrastructure-customercontact.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseServerType`

The database server model type of the Exadata infrastructure. For the list of valid model names, use the `ListDbSystemShapes` operation.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\/.=-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisplayName`

The user-friendly name for the Exadata infrastructure.

Required when creating an Exadata infrastructure.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z_](?!.*--)[a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaintenanceWindow`

The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window.

_Required_: No

_Type_: [MaintenanceWindow](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-odb-cloudexadatainfrastructure-maintenancewindow.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Shape`

The model name of the Exadata infrastructure.

Required when creating an Exadata infrastructure.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\/.=-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageCount`

The number of storage servers that are activated for the Exadata infrastructure.

Required when creating an Exadata infrastructure.

_Required_: Conditional

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageServerType`

The storage server model type of the Exadata infrastructure. For the list of valid model names, use the `ListDbSystemShapes` operation.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\/.=-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to assign to the Exadata Infrastructure.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-odb-cloudexadatainfrastructure-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier of the Exadata infrastructure. For example:

`{ "Ref": "myExadataInfrastructure" }`

For the Exadata infrastructure `myExadataInfrastructure`, `Ref` returns the unique identifier of the Exadata infrastructure.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`ActivatedStorageCount`

The number of storage servers requested for the Exadata infrastructure.

`AdditionalStorageCount`

The number of storage servers requested for the Exadata infrastructure.

`AvailableStorageSizeInGBs`

The amount of available storage, in gigabytes (GB), for the Exadata infrastructure.

`CloudExadataInfrastructureArn`

The Amazon Resource Name (ARN) for the Exadata infrastructure.

`CloudExadataInfrastructureId`

The unique identifier for the Exadata infrastructure.

`ComputeModel`

The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An
ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores
elastically allocated from a pool of compute and storage servers. An OCPU is a legacy
physical measure of compute resources. OCPUs are based on the physical core of a processor
with hyper-threading enabled.

`CpuCount`

The total number of CPU cores that are allocated to the Exadata infrastructure.

`DataStorageSizeInTBs`

The size of the Exadata infrastructure's data disk group, in terabytes (TB).

`DbNodeStorageSizeInGBs`

The size of the Exadata infrastructure's local node storage, in gigabytes (GB).

`DbServerIds`

The list of database server identifiers for the Exadata infrastructure.

`DbServerVersion`

The software version of the database servers (dom0) in the Exadata infrastructure.

`MaxCpuCount`

The total number of CPU cores available on the Exadata infrastructure.

`MaxDataStorageInTBs`

The total amount of data disk group storage, in terabytes (TB), that's available on the Exadata infrastructure.

`MaxDbNodeStorageSizeInGBs`

The total amount of local node storage, in gigabytes (GB), that's available on the Exadata infrastructure.

`MaxMemoryInGBs`

The total amount of memory, in gigabytes (GB), that's available on the Exadata infrastructure.

`MemorySizeInGBs`

The amount of memory, in gigabytes (GB), that's allocated on the Exadata infrastructure.

`Ocid`

The OCID of the Exadata infrastructure.

`OciResourceAnchorName`

The name of the OCI resource anchor for the Exadata infrastructure.

`OciUrl`

The HTTPS link to the Exadata infrastructure in OCI.

`StorageServerVersion`

The software version of the storage servers on the Exadata infrastructure.

`TotalStorageSizeInGBs`

The total amount of storage, in gigabytes (GB), on the the Exadata infrastructure.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CustomerContact
