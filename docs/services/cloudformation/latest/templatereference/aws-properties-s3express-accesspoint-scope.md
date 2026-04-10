This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::AccessPoint Scope

You can use the access point scope to restrict access to specific prefixes, API operations, or a combination of both.

For more information, see [Manage the scope of your access points for directory buckets.](../../../s3/latest/userguide/access-points-directory-buckets-manage-scope.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Permissions" : [ String, ... ],
  "Prefixes" : [ String, ... ]
}

```

### YAML

```yaml

  Permissions:
    - String
  Prefixes:
    - String

```

## Properties

`Permissions`

You can include one or more API operations as permissions.

_Required_: No

_Type_: Array of String

_Allowed values_: `GetObject | GetObjectAttributes | ListMultipartUploadParts | ListBucket | ListBucketMultipartUploads | PutObject | DeleteObject | AbortMultipartUpload`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefixes`

You can specify any amount of prefixes, but the total length of characters of all prefixes must be less than 256 bytes in size.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublicAccessBlockConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
