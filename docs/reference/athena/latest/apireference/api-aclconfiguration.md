# AclConfiguration

Indicates that an Amazon S3 canned ACL should be set to control ownership of
stored query results, including data files inserted by Athena as the result
of statements like CTAS or INSERT INTO. When Athena stores query results in
Amazon S3, the canned ACL is set with the `x-amz-acl` request
header. For more information about S3 Object Ownership, see [Object Ownership settings](../../../../services/s3/latest/userguide/about-object-ownership.md#object-ownership-overview) in the _Amazon S3 User_
_Guide_.

## Contents

**S3AclOption**

The Amazon S3 canned ACL that Athena should specify when storing
query results, including data files inserted by Athena as the result
of statements like CTAS or INSERT INTO. Currently the only supported canned ACL is
`BUCKET_OWNER_FULL_CONTROL`. If a query runs in a workgroup and the
workgroup overrides client-side settings, then the Amazon S3 canned ACL
specified in the workgroup's settings is used for all queries that run in the workgroup.
For more information about Amazon S3 canned ACLs, see [Canned ACL](../../../../services/s3/latest/userguide/acl-overview.md#canned-acl) in the _Amazon S3 User Guide_.

Type: String

Valid Values: `BUCKET_OWNER_FULL_CONTROL`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/AclConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/AclConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/AclConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Data Types

ApplicationDPUSizes
