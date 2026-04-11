# JobManifestSpec

Describes the format of a manifest. If the manifest is in CSV format, also describes the
columns contained within the manifest.

## Contents

**Format**

Indicates which of the available formats the specified manifest uses.

Type: String

Valid Values: `S3BatchOperations_CSV_20180820 | S3InventoryReport_CSV_20161130`

Required: Yes

**Fields**

If the specified manifest object is in the `S3BatchOperations_CSV_20180820`
format, this element describes which columns contain the required data.

Type: Array of strings

Valid Values: `Ignore | Bucket | Key | VersionId`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/jobmanifestspec.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/jobmanifestspec.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/jobmanifestspec.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobManifestLocation

JobOperation

All content copied from https://docs.aws.amazon.com/.
