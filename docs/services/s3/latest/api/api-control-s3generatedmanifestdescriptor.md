# S3GeneratedManifestDescriptor

Describes the specified job's generated manifest. Batch Operations jobs created with a
ManifestGenerator populate details of this descriptor after execution of the
ManifestGenerator.

## Contents

**Format**

The format of the generated manifest.

Type: String

Valid Values: `S3InventoryReport_CSV_20211130`

Required: No

**Location**

Contains the information required to locate a manifest object. Manifests can't be
imported from directory buckets. For more information, see [Directory\
buckets](../userguide/directory-buckets-overview.md).

Type: [JobManifestLocation](api-control-jobmanifestlocation.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/s3generatedmanifestdescriptor.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/s3generatedmanifestdescriptor.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/s3generatedmanifestdescriptor.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3DeleteObjectTaggingOperation

S3Grant

All content copied from https://docs.aws.amazon.com/.
