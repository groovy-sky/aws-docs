This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Bucket AccessRules

`AccessRules` is a property of the [AWS::Lightsail::Bucket](../userguide/aws-resource-lightsail-bucket.md) resource. It describes access rules for a bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowPublicOverrides" : Boolean,
  "GetObject" : String
}

```

### YAML

```yaml

  AllowPublicOverrides: Boolean
  GetObject: String

```

## Properties

`AllowPublicOverrides`

A Boolean value indicating whether the access control list (ACL) permissions that are
applied to individual objects override the `GetObject` option that is currently
specified.

When this is true, you can use the [PutObjectAcl](../../../s3/latest/api/api-putobjectacl.md) Amazon S3 API
operation to set individual objects to public (read-only) or private, using either the `public-read`
ACL or the `private` ACL.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GetObject`

Specifies the anonymous access to all objects in a bucket.

The following options can be specified:

- `public` \- Sets all objects in the bucket to public (read-only), making
them readable by everyone on the internet.

If the `GetObject` value is set to `public`, then all
objects in the bucket default to public regardless of the
`allowPublicOverrides` value.

- `private` \- Sets all objects in the bucket to private, making them
readable only by you and anyone that you grant access to.

If the `GetObject` value is set to `private`, and the
`allowPublicOverrides` value is set to `true`, then all
objects in the bucket default to private unless they are configured with a
`public-read` ACL. Individual objects with a `public-read`
ACL are readable by everyone on the internet.

_Required_: No

_Type_: String

_Allowed values_: `public | private`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lightsail::Bucket

Tag

All content copied from https://docs.aws.amazon.com/.
