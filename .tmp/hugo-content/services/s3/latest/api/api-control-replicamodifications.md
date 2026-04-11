# ReplicaModifications

A filter that you can use to specify whether replica modification sync is enabled.
S3 on Outposts replica modification sync can help you keep object metadata synchronized
between replicas and source objects. By default, S3 on Outposts replicates metadata from the
source objects to the replicas only. When replica modification sync is enabled,
S3 on Outposts replicates metadata changes made to the replica copies back to the source
object, making the replication bidirectional.

To replicate object metadata modifications on replicas, you can specify this element and
set the `Status` of this element to `Enabled`.

###### Note

You must enable replica modification sync on the source and destination buckets to
replicate replica metadata changes between the source and the replicas.

## Contents

**Status**

Specifies whether S3 on Outposts replicates modifications to object metadata on
replicas.

Type: String

Valid Values: `Enabled | Disabled`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/replicamodifications.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/replicamodifications.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/replicamodifications.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegionReport

ReplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
