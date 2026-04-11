This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource DocumentsMetadataConfiguration

Document metadata files that contain information such as the document access control
information, source URI, document author, and custom attributes. Each metadata file
contains metadata about a single document.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Prefix" : String
}

```

### YAML

```yaml

  S3Prefix: String

```

## Properties

`S3Prefix`

A prefix used to filter metadata configuration files in the AWS S3
bucket. The S3 bucket might contain multiple metadata files. Use `S3Prefix`
to include only the desired metadata files.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeValue

GoogleDriveConfiguration

All content copied from https://docs.aws.amazon.com/.
