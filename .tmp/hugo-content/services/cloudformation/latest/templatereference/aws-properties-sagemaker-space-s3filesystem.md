This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space S3FileSystem

A custom file system in Amazon S3. This is only supported in Amazon SageMaker Unified Studio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Uri" : String
}

```

### YAML

```yaml

  S3Uri: String

```

## Properties

`S3Uri`

The Amazon S3 URI that specifies the location in S3 where files are stored, which is mounted
within the Studio environment. For example:
`s3://<bucket-name>/<prefix>/`.

_Required_: No

_Type_: String

_Pattern_: `(s3)://([^/]+)/?(.*)`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceSpec

SpaceAppLifecycleManagement

All content copied from https://docs.aws.amazon.com/.
