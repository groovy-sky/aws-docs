This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::ResourceDataSync

The `AWS::SSM::ResourceDataSync` resource creates, updates, or deletes a
resource data sync for AWS Systems Manager. A resource data sync helps you view
data from multiple sources in a single location. Systems Manager offers two types
of resource data sync: `SyncToDestination` and
`SyncFromSource`.

You can configure Systems Manager Inventory to use the
`SyncToDestination` type to synchronize Inventory data from multiple
AWS Regions to a single Amazon S3 bucket.

You can configure Systems Manager Explorer to use the `SyncFromSource` type to
synchronize operational work items (OpsItems) and operational data (OpsData) from
multiple AWS Regions. This type can synchronize OpsItems and OpsData
from multiple AWS accounts and Regions or from an
`EntireOrganization` by using AWS Organizations.

A resource data sync is an asynchronous operation that returns immediately. After a
successful initial sync is completed, the system continuously syncs data.

By default, data is not encrypted in Amazon S3. We strongly recommend that
you enable encryption in Amazon S3 to ensure secure data storage. We also
recommend that you secure access to the Amazon S3 bucket by creating a
restrictive bucket policy.

For more information, see [Configuring Inventory Collection](../../../systems-manager/latest/userguide/sysman-inventory-configuring.md#sysman-inventory-datasync) and [Setting Up\
Systems Manager Explorer to Display Data from Multiple Accounts and Regions](../../../systems-manager/latest/userguide/explorer-resource-data-sync.md)
in the _AWS Systems Manager User Guide_.

###### Important

The following _Syntax_ section shows all fields that are
supported for a resource data sync. The _Examples_ section below
shows the recommended way to specify configurations for each sync type. Refer to the
_Examples_ section when you create your resource data
sync.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::ResourceDataSync",
  "Properties" : {
      "BucketName" : String,
      "BucketPrefix" : String,
      "BucketRegion" : String,
      "KMSKeyArn" : String,
      "S3Destination" : S3Destination,
      "SyncFormat" : String,
      "SyncName" : String,
      "SyncSource" : SyncSource,
      "SyncType" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSM::ResourceDataSync
Properties:
  BucketName: String
  BucketPrefix: String
  BucketRegion: String
  KMSKeyArn: String
  S3Destination:
    S3Destination
  SyncFormat: String
  SyncName: String
  SyncSource:
    SyncSource
  SyncType: String

```

## Properties

`BucketName`

The name of the S3 bucket where the aggregated data is stored.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketPrefix`

An Amazon S3 prefix for the bucket.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketRegion`

The AWS Region with the S3 bucket targeted by the resource data sync.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KMSKeyArn`

The Amazon Resource Name (ARN) of an encryption key for a destination in Amazon S3. You can use a KMS key to encrypt inventory data in Amazon S3.
You must specify a key that exist in the same AWS Region as the
destination Amazon S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Destination`

Configuration information for the target S3 bucket.

_Required_: No

_Type_: [S3Destination](aws-properties-ssm-resourcedatasync-s3destination.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyncFormat`

A supported sync format. The following format is currently supported: JsonSerDe

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyncName`

A name for the resource data sync.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyncSource`

Information about the source where the data was synchronized.

_Required_: No

_Type_: [SyncSource](aws-properties-ssm-resourcedatasync-syncsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyncType`

The type of resource data sync. If `SyncType` is `SyncToDestination`,
then the resource data sync synchronizes data to an S3 bucket. If the `SyncType` is
`SyncFromSource` then the resource data sync synchronizes data from AWS Organizations or from
multiple AWS Regions.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the resource data sync, such as
`TestResourceDataSync`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Create a SyncToDestination resource data sync](#aws-resource-ssm-resourcedatasync--examples--Create_a_SyncToDestination_resource_data_sync)

- [Create a SyncFromSource resource data sync with type SingleAccountMultipleRegions](#aws-resource-ssm-resourcedatasync--examples--Create_a_SyncFromSource_resource_data_sync_with_type_SingleAccountMultipleRegions)

- [Create a SyncFromSource resource data sync with type EntireOrganization](#aws-resource-ssm-resourcedatasync--examples--Create_a_SyncFromSource_resource_data_sync_with_type_EntireOrganization)

- [Creating a SyncFromSource resource data sync with type OrganizationalUnits](#aws-resource-ssm-resourcedatasync--examples--Creating_a_SyncFromSource_resource_data_sync_with_type_OrganizationalUnits)

### Create a SyncToDestination resource data sync

The following example synchronizes Systems Manager Inventory metadata in
the US East (Ohio) Region (us-east-2) to a single Amazon S3 bucket.
Resource data sync then automatically updates the centralized data when new data
is collected.

#### JSON

```json

{
    "Description": "Create a resource data sync for Systems Manager",
    "Resources": {
        "BasicResourceDataSync": {
            "Type": "AWS::SSM::ResourceDataSync",
            "Properties": {
                "SyncName": "test-sync",
                "SyncType": "SyncToDestination",
                "S3Destination": {
                    "BucketName": "amzn-s3-demo-bucket",
                    "BucketRegion": "us-east-2",
                    "SyncFormat": "JsonSerDe",
                    "BucketPrefix": "cfn",
                    "KMSKeyArn": "kmsKeyARN"
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: Create a resource data sync for Systems Manager
Resources:
  BasicResourceDataSync:
    Type: AWS::SSM::ResourceDataSync
    Properties:
      SyncName: test-sync
      SyncType: SyncToDestination
      S3Destination:
        BucketName: amzn-s3-demo-bucket
        BucketRegion: us-east-2
        SyncFormat: JsonSerDe
        BucketPrefix: cfn
        KMSKeyArn: kmsKeyARN
```

### Create a SyncFromSource resource data sync with type SingleAccountMultipleRegions

The following example synchronizes Systems Manager Explorer OpsData and OpsItems
from multiple AWS Regions in a single AWS account.

#### JSON

```json

{
    "Description": "Create a resource data sync for Systems Manager Explorer",
    "Resources": {
        "BasicResourceDataSync": {
            "Type": "AWS::SSM::ResourceDataSync",
            "Properties": {
                "SyncName": "test-sync",
                "SyncType": "SyncFromSource",
                "SyncSource": {
                    "SourceType": "SingleAccountMultiRegions",
                    "SourceRegions": [
                        "us-east-1",
                        "us-west-1",
                        "us-west-2"
                    ],
                    "IncludeFutureRegions": false
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: Create a resource data sync for Systems Manager Explorer
Resources:
  BasicResourceDataSync:
    Type: AWS::SSM::ResourceDataSync
    Properties:
      SyncName: test-sync
      SyncType: SyncFromSource
      SyncSource:
        SourceType: SingleAccountMultiRegions
        SourceRegions:
        - us-east-1
        - us-west-1
        - us-west-2
        IncludeFutureRegions: false
```

### Create a SyncFromSource resource data sync with type EntireOrganization

The following example synchronizes Systems Manager Explorer OpsData and OpsItems
from your entire organization in AWS Organizationsin the us-west-1 Region.

#### JSON

```json

{
    "Description": "Create a resource data sync for Systems Manager Explorer",
    "Resources": {
        "BasicResourceDataSync": {
            "Type": "AWS::SSM::ResourceDataSync",
            "Properties": {
                "SyncName": "test-sync",
                "SyncType": "SyncFromSource",
                "SyncSource": {
                    "SourceType": "AwsOrganizations",
                    "AwsOrganizationsSource": {
                        "OrganizationSourceType": "EntireOrganization"
                    },
                    "SourceRegions": [
                        "us-west-1"
                    ],
                    "IncludeFutureRegions": false
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: Create a resource data sync for Systems Manager Explorer
Resources:
  BasicResourceDataSync:
    Type: AWS::SSM::ResourceDataSync
    Properties:
      SyncName: test-sync
      SyncType: SyncFromSource
      SyncSource:
        SourceType: AwsOrganizations
        AwsOrganizationsSource:
          OrganizationSourceType: EntireOrganization
        SourceRegions:
        - us-west-1
        IncludeFutureRegions: false
```

### Creating a SyncFromSource resource data sync with type OrganizationalUnits

The following example synchronizes Systems Manager Explorer OpsData and OpsItems
from organization unit 12345 in AWS Organizations in the us-west-1 Region.

#### JSON

```json

{
    "Description": "Create a resource data sync for Systems Manager Explorer",
    "Resources": {
        "BasicResourceDataSync": {
            "Type": "AWS::SSM::ResourceDataSync",
            "Properties": {
                "SyncName": "test-sync",
                "SyncType": "SyncFromSource",
                "SyncSource": {
                    "SourceType": "AwsOrganizations",
                    "AwsOrganizationsSource": {
                        "OrganizationSourceType": "OrganizationalUnits",
                        "OrganizationalUnits": [
                            "ou-12345"
                        ]
                    },
                    "SourceRegions": [
                        "us-west-1"
                    ],
                    "IncludeFutureRegions": false
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: Create a resource data sync for Systems Manager Explorer
Resources:
  BasicResourceDataSync:
    Type: AWS::SSM::ResourceDataSync
    Properties:
      SyncName: test-sync
      SyncType: SyncFromSource
      SyncSource:
        SourceType: AwsOrganizations
        AwsOrganizationsSource:
          OrganizationSourceType: OrganizationalUnits
          OrganizationalUnits:
          - ou-12345
        SourceRegions:
        - us-west-1
        IncludeFutureRegions: false
```

## See also

- [What\
is AWS Systems Manager?](../../../systems-manager/latest/userguide/what-is-systems-manager.md)

- [AWS Systems Manager Inventory](../../../systems-manager/latest/userguide/systems-manager-inventory.md)

- [Configuring inventory collection](../../../systems-manager/latest/userguide/sysman-inventory-configuring.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AwsOrganizationsSource

All content copied from https://docs.aws.amazon.com/.
