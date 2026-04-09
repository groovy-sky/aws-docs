# DeleteMarkerReplication

Specifies whether S3 on Outposts replicates delete markers. If you specify a
`Filter` element in your replication configuration, you must also include a
`DeleteMarkerReplication` element. If your `Filter` includes a
`Tag` element, the `DeleteMarkerReplication` element's
`Status` child element must be set to `Disabled`, because
S3 on Outposts does not support replicating delete markers for tag-based rules.

For more information about delete marker replication, see [How delete operations affect replication](../userguide/s3outpostsreplication.md#outposts-replication-what-is-replicated) in the
_Amazon S3 User Guide_.

## Contents

**Status**

Indicates whether to replicate delete markers.

Type: String

Valid Values: `Enabled | Disabled`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/deletemarkerreplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/deletemarkerreplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/deletemarkerreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Credentials

DeleteMultiRegionAccessPointInput

All content copied from https://docs.aws.amazon.com/.
