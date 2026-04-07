# StorageLensConfiguration

A container for the Amazon S3 Storage Lens configuration.

## Contents

**AccountLevel**

A container for all the account-level configurations of your S3 Storage Lens
configuration.

Type: [AccountLevel](api-control-accountlevel.md) data type

Required: Yes

**Id**

A container for the Amazon S3 Storage Lens configuration ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[a-zA-Z0-9\-\_\.]+`

Required: Yes

**IsEnabled**

A container for whether the S3 Storage Lens configuration is enabled.

Type: Boolean

Required: Yes

**AwsOrg**

A container for the AWS organization for this S3 Storage Lens configuration.

Type: [StorageLensAwsOrg](api-control-storagelensawsorg.md) data type

Required: No

**DataExport**

A container to specify the properties of your S3 Storage Lens metrics export including, the
destination, schema and format.

Type: [StorageLensDataExport](api-control-storagelensdataexport.md) data type

Required: No

**Exclude**

A container for what is excluded in this configuration. This container can only be valid
if there is no `Include` container submitted, and it's not empty.

Type: [Exclude](api-control-exclude.md) data type

Required: No

**ExpandedPrefixesDataExport**

A container that configures your S3 Storage Lens expanded prefixes metrics report.

Type: [StorageLensExpandedPrefixesDataExport](api-control-storagelensexpandedprefixesdataexport.md) data type

Required: No

**Include**

A container for what is included in this configuration. This container can only be valid
if there is no `Exclude` container submitted, and it's not empty.

Type: [Include](api-control-include.md) data type

Required: No

**PrefixDelimiter**

A container for all prefix delimiters that are used for object keys in this S3 Storage Lens
configuration. The prefix delimiters determine how S3 Storage Lens counts prefix depth, by
separating the hierarchical levels in object keys.

###### Note

- If either a prefix delimiter or existing delimiter is undefined, Amazon S3 uses the
delimiter that’s defined.

- If both the prefix delimiter and existing delimiter are undefined, S3 uses `/`
as the default delimiter.

- When custom delimiters are used, both the prefix delimiter and existing
delimiter must specify the same special character. Otherwise, your request results
in an error.

Type: String

Length Constraints: Maximum length of 1.

Required: No

**StorageLensArn**

The Amazon Resource Name (ARN) of the S3 Storage Lens configuration. This property is read-only
and follows the following format: `
               arn:aws:s3:us-east-1:example-account-id:storage-lens/your-dashboard-name
                  `

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `arn:[a-z\-]+:s3:[a-z0-9\-]+:\d{12}:storage\-lens\/.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/StorageLensConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/StorageLensConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/StorageLensConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StorageLensAwsOrg

StorageLensDataExport
