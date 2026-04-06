# AclGrantee

You specify each grantee as a type-value pair using one of these types. You can specify
only one type of grantee. For more information, see [PutBucketAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAcl.html).

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**id**

The value specified is the canonical user ID of an AWS account.

Type: String

Required: No

**uri**

Used for granting permissions to a predefined group.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/AclGrantee)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/AclGrantee)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/AclGrantee)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AccessPreviewSummary

AnalysisRule
