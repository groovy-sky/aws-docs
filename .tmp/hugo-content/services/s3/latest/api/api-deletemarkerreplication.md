# DeleteMarkerReplication

Specifies whether Amazon S3 replicates delete markers. If you specify a `Filter` in your
replication configuration, you must also include a `DeleteMarkerReplication` element. If your
`Filter` includes a `Tag` element, the `DeleteMarkerReplication` `Status` must be set to Disabled, because Amazon S3 does not support replicating delete markers
for tag-based rules. For an example configuration, see [Basic Rule\
Configuration](../dev/replication-add-config.md#replication-config-min-rule-config).

For more information about delete marker replication, see [Basic Rule Configuration](../dev/delete-marker-replication.md).

###### Note

If you are using an earlier version of the replication configuration, Amazon S3 handles replication of
delete markers differently. For more information, see [Backward Compatibility](../dev/replication-add-config.md#replication-backward-compat-considerations).

## Contents

**Status**

Indicates whether to replicate delete markers.

###### Note

Indicates whether to replicate delete markers.

Type: String

Valid Values: `Enabled | Disabled`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deletemarkerreplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deletemarkerreplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deletemarkerreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteMarkerEntry

Destination

All content copied from https://docs.aws.amazon.com/.
