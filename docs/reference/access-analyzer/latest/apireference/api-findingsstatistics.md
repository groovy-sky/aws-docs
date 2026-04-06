# FindingsStatistics

Contains information about the aggregate statistics for an external or unused access
analyzer. Only one parameter can be used in a `FindingsStatistics`
object.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**externalAccessFindingsStatistics**

The aggregate statistics for an external access analyzer.

Type: [ExternalAccessFindingsStatistics](api-externalaccessfindingsstatistics.md) object

Required: No

**internalAccessFindingsStatistics**

The aggregate statistics for an internal access analyzer. This includes information
about active, archived, and resolved findings related to internal access within your AWS
organization or account.

Type: [InternalAccessFindingsStatistics](api-internalaccessfindingsstatistics.md) object

Required: No

**unusedAccessFindingsStatistics**

The aggregate statistics for an unused access analyzer.

Type: [UnusedAccessFindingsStatistics](api-unusedaccessfindingsstatistics.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/FindingsStatistics)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/FindingsStatistics)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/FindingsStatistics)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FindingSourceDetail

FindingSummary
