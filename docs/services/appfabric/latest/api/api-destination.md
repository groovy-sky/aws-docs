# Destination

Contains information about an audit log destination.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**firehoseStream**

Contains information about an Amazon Data Firehose delivery stream.

Type: [FirehoseStream](api-firehosestream.md) object

Required: No

**s3Bucket**

Contains information about an Amazon S3 bucket.

Type: [S3Bucket](api-s3bucket.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/destination.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/destination.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/destination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Credential

DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
