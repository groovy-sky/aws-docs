# GroupMembers

A list of users or sub groups that belong to a group. This is for generating
Amazon Q Business chat results only from document a user has access to.

## Contents

**memberGroups**

A list of sub groups that belong to a group. For example, the sub groups "Research",
"Engineering", and "Sales and Marketing" all belong to the group "Company".

Type: Array of [MemberGroup](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_MemberGroup.html) objects

Required: No

**memberUsers**

A list of users that belong to a group. For example, a list of interns all belong to
the "Interns" group.

Type: Array of [MemberUser](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_MemberUser.html) objects

Required: No

**s3PathForGroupMembers**

Information required for Amazon Q Business to find a specific file in an Amazon S3
bucket.

Type: [S3](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_S3.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/GroupMembers)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/GroupMembers)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/GroupMembers)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FailedDocument

GroupStatusDetail
