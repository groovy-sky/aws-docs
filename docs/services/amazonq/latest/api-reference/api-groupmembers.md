# GroupMembers

A list of users or sub groups that belong to a group. This is for generating
Amazon Q Business chat results only from document a user has access to.

## Contents

**memberGroups**

A list of sub groups that belong to a group. For example, the sub groups "Research",
"Engineering", and "Sales and Marketing" all belong to the group "Company".

Type: Array of [MemberGroup](api-membergroup.md) objects

Required: No

**memberUsers**

A list of users that belong to a group. For example, a list of interns all belong to
the "Interns" group.

Type: Array of [MemberUser](api-memberuser.md) objects

Required: No

**s3PathForGroupMembers**

Information required for Amazon Q Business to find a specific file in an Amazon S3
bucket.

Type: [S3](api-s3.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/groupmembers.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/groupmembers.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/groupmembers.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FailedDocument

GroupStatusDetail

All content copied from https://docs.aws.amazon.com/.
