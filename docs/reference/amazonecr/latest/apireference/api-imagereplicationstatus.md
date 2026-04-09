# ImageReplicationStatus

The status of the replication process for an image.

## Contents

**failureCode**

The failure code for a replication that has failed.

Type: String

Required: No

**region**

The destination Region for the image replication.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 25.

Pattern: `[0-9a-z-]{2,25}`

Required: No

**registryId**

The AWS account ID associated with the registry to which the image belongs.

Type: String

Pattern: `[0-9]{12}`

Required: No

**status**

The image replication status.

Type: String

Valid Values: `IN_PROGRESS | COMPLETE | FAILED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/imagereplicationstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/imagereplicationstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/imagereplicationstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageReferrer

ImageScanFinding

All content copied from https://docs.aws.amazon.com/.
