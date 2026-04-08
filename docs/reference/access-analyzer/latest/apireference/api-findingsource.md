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

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/findingsource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/findingsource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/findingsource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FindingDetails

FindingSourceDetail
