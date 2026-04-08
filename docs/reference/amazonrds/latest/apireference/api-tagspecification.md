# TagSpecification

The tags to apply to resources when creating or modifying a DB instance or DB cluster. When you specify a tag, you must specify the resource type to tag, otherwise the request will fail.

## Contents

###### Note

In the following list, the required parameters are described first.

**ResourceType**

The type of resource to tag on creation.

Valid Values:

- `auto-backup` \- The DB instance's automated backup.

- `cluster-auto-backup` \- The DB cluster's automated backup.

Type: String

Required: No

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/tagspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/tagspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/tagspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Tag

TargetHealth
