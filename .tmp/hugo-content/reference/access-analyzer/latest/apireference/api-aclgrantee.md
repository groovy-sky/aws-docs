# AclGrantee

You specify each grantee as a type-value pair using one of these types. You can specify
only one type of grantee. For more information, see [PutBucketAcl](../../../../services/s3/latest/api/api-putbucketacl.md).

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

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/aclgrantee.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/aclgrantee.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/aclgrantee.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AccessPreviewSummary

AnalysisRule
