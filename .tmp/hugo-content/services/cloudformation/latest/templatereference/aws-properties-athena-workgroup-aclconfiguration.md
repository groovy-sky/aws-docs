This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup AclConfiguration

Indicates that an Amazon S3 canned ACL should be set to control ownership of
stored query results, including data files inserted by Athena as the result
of statements like CTAS or INSERT INTO. When Athena stores query results in
Amazon S3, the canned ACL is set with the `x-amz-acl` request
header. For more information about S3 Object Ownership, see [Object Ownership settings](../../../s3/latest/userguide/about-object-ownership.md#object-ownership-overview) in the _Amazon S3 User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3AclOption" : String
}

```

### YAML

```yaml

  S3AclOption: String

```

## Properties

`S3AclOption`

The Amazon S3 canned ACL that Athena should specify when storing
query results, including data files inserted by Athena as the result
of statements like CTAS or INSERT INTO. Currently the only supported canned ACL is
`BUCKET_OWNER_FULL_CONTROL`. If a query runs in a workgroup and the
workgroup overrides client-side settings, then the Amazon S3 canned ACL
specified in the workgroup's settings is used for all queries that run in the workgroup.
For more information about Amazon S3 canned ACLs, see [Canned ACL](../../../s3/latest/userguide/acl-overview.md#canned-acl) in the _Amazon S3 User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `BUCKET_OWNER_FULL_CONTROL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Athena::WorkGroup

Classification

All content copied from https://docs.aws.amazon.com/.
