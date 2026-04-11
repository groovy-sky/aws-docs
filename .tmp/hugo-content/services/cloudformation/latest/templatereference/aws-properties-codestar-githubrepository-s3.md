This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStar::GitHubRepository S3

The `S3` property type specifies information about the Amazon S3 bucket
that contains the code to be committed to the new repository.

`S3` is a property of the `AWS::CodeStar::GitHubRepository`
resource.

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

The name of the Amazon S3 bucket that contains the ZIP file with the content to be
committed to the new repository.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`Key`

The S3 object key or file name for the ZIP file.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`ObjectVersion`

The object version of the ZIP file, if versioning is enabled for the Amazon S3
bucket.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Code

Next

All content copied from https://docs.aws.amazon.com/.
