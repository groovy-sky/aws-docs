# FindingSource

The source of the finding. This indicates how the access that generated the finding is
granted. It is populated for Amazon S3 bucket findings.

## Contents

**type**

Indicates the type of access that generated the finding.

Type: String

Valid Values: `POLICY | BUCKET_ACL | S3_ACCESS_POINT | S3_ACCESS_POINT_ACCOUNT`

Required: Yes

**detail**

Includes details about how the access that generated the finding is granted. This is
populated for Amazon S3 bucket findings.

Type: [FindingSourceDetail](api-findingsourcedetail.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/FindingSource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/FindingSource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/FindingSource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FindingDetails

FindingSourceDetail
