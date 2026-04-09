# MultiRegionProperties

Defines the structure for multi-Region cluster configurations, containing the witness region and linked cluster settings.

## Contents

**clusters**

The set of peered clusters that form the multi-Region cluster configuration. Each peered cluster represents a database instance in a different Region.

Type: Array of strings

Pattern: `arn:aws(-[^:]+)?:dsql:[a-z0-9-]{1,20}:[0-9]{12}:cluster/[a-z0-9]{26}`

Required: No

**witnessRegion**

The Region that serves as the witness region for a multi-Region cluster. The witness Region helps maintain cluster consistency and quorum.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 50.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dsql-2018-05-10/multiregionproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dsql-2018-05-10/multiregionproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dsql-2018-05-10/multiregionproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionDetails

ValidationExceptionField

All content copied from https://docs.aws.amazon.com/.
