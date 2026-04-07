# TableRecordExpirationJobMetrics

Provides metrics for the record expiration job that most recently ran for a table. The metrics provide insight into the amount of data that was removed when the job ran.

## Contents

**deletedDataFiles**

The total number of data files that were removed when the job ran.

Type: Long

Required: No

**deletedRecords**

The total number of records that were removed when the job ran.

Type: Long

Required: No

**removedFilesSize**

The total size (in bytes) of the data files that were removed when the job ran.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/TableRecordExpirationJobMetrics)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/TableRecordExpirationJobMetrics)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/TableRecordExpirationJobMetrics)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TableRecordExpirationConfigurationValue

TableRecordExpirationSettings
