This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::AccessPoint PublicAccessBlockConfiguration

The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can
enable the configuration options in any combination. Bucket-level settings work alongside
account-level settings (which may inherit from organization-level policies). For more
information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockPublicAcls" : Boolean,
  "BlockPublicPolicy" : Boolean,
  "IgnorePublicAcls" : Boolean,
  "RestrictPublicBuckets" : Boolean
}

```

### YAML

```yaml

  BlockPublicAcls: Boolean
  BlockPublicPolicy: Boolean
  IgnorePublicAcls: Boolean
  RestrictPublicBuckets: Boolean

```

## Properties

`BlockPublicAcls`

Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects
in this bucket. Setting this element to `TRUE` causes the following behavior:

- PUT Bucket ACL and PUT Object ACL calls fail if the specified ACL is public.

- PUT Object calls fail if the request includes a public ACL.

- PUT Bucket calls fail if the request includes a public ACL.

Enabling this setting doesn't affect existing policies or ACLs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockPublicPolicy`

Specifies whether Amazon S3 should block public bucket policies for this bucket. Setting this element to
`TRUE` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy
allows public access.

Enabling this setting doesn't affect existing bucket policies.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IgnorePublicAcls`

Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket. Setting
this element to `TRUE` causes Amazon S3 to ignore all public ACLs on this bucket and objects in
this bucket.

Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new
public ACLs from being set.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestrictPublicBuckets`

Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element
to `TRUE` restricts access to this bucket to only AWS service principals and
authorized users within this account if the bucket has a public policy.

Enabling this setting doesn't affect previously stored bucket policies, except that public and
cross-account access within any public bucket policy, including non-public delegation to specific
accounts, is blocked.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::S3::AccessPoint

Tag
