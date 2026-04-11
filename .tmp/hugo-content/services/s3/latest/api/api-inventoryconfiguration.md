# InventoryConfiguration

Specifies the S3 Inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket\
inventory](restbucketgetinventoryconfig.md) in the _Amazon S3 API Reference_.

## Contents

**Destination**

Contains information about where to publish the inventory results.

Type: [InventoryDestination](api-inventorydestination.md) data type

Required: Yes

**Id**

The ID used to identify the inventory configuration.

Type: String

Required: Yes

**IncludedObjectVersions**

Object versions to include in the inventory list. If set to `All`, the list includes all
the object versions, which adds the version-related fields `VersionId`,
`IsLatest`, and `DeleteMarker` to the list. If set to `Current`, the
list does not contain these version-related fields.

Type: String

Valid Values: `All | Current`

Required: Yes

**IsEnabled**

Specifies whether the inventory is enabled or disabled. If set to `True`, an inventory
list is generated. If set to `False`, no inventory list is generated.

Type: Boolean

Required: Yes

**Schedule**

Specifies the schedule for generating inventory results.

Type: [InventorySchedule](api-inventoryschedule.md) data type

Required: Yes

**Filter**

Specifies an inventory filter. The inventory only includes objects that meet the filter's
criteria.

Type: [InventoryFilter](api-inventoryfilter.md) data type

Required: No

**OptionalFields**

Contains the optional fields that are included in the inventory results.

Type: Array of strings

Valid Values: `Size | LastModifiedDate | StorageClass | ETag | IsMultipartUploaded | ReplicationStatus | EncryptionStatus | ObjectLockRetainUntilDate | ObjectLockMode | ObjectLockLegalHoldStatus | IntelligentTieringAccessTier | BucketKeyStatus | ChecksumAlgorithm | ObjectAccessControlList | ObjectOwner | LifecycleExpirationDate`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/inventoryconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/inventoryconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/inventoryconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntelligentTieringFilter

InventoryDestination

All content copied from https://docs.aws.amazon.com/.
