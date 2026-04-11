This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeCommit::Repository S3

Information about the Amazon S3 bucket that contains the code that will be
committed to the new repository. Changes to this property are ignored after initial
resource creation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Key" : String,
  "ObjectVersion" : String
}

```

### YAML

```yaml

  Bucket: String
  Key: String
  ObjectVersion: String

```

## Properties

`Bucket`

The name of the Amazon S3 bucket that contains the ZIP file with the content
that will be committed to the new repository. This can be specified using the name of
the bucket in the AWS account. Changes to this property are ignored after
initial resource creation.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key to use for accessing the Amazon S3 bucket. Changes to this property
are ignored after initial resource creation. For more information, see [Creating\
object key names](../../../s3/latest/userguide/object-keys.md) and [Uploading objects](../../../s3/latest/userguide/upload-objects.md) in the
Amazon S3 User Guide.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectVersion`

The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket. Changes to this property are ignored after initial resource creation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RepositoryTrigger

Tag

All content copied from https://docs.aws.amazon.com/.
