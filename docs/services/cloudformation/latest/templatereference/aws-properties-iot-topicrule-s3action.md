This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule S3Action

Describes an action to write data to an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "CannedAcl" : String,
  "Key" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  BucketName: String
  CannedAcl: String
  Key: String
  RoleArn: String

```

## Properties

`BucketName`

The Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CannedAcl`

The Amazon S3 canned ACL that controls access to the object identified by the object
key. For more information, see [S3 canned ACLs](../../../s3/latest/dev/acl-overview.md#canned-acl).

_Required_: No

_Type_: String

_Allowed values_: `private | public-read | public-read-write | aws-exec-read | authenticated-read | bucket-owner-read | bucket-owner-full-control | log-delivery-write`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The object key. For more information, see [Actions, resources, and condition keys for Amazon S3](../../../s3/latest/dev/list-amazons3.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants access.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RepublishActionHeaders

SigV4Authorization

All content copied from https://docs.aws.amazon.com/.
