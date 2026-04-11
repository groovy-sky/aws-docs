This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Pipeline S3Location

The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline
definition from this location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "ETag" : String,
  "Key" : String,
  "Version" : String
}

```

### YAML

```yaml

  Bucket: String
  ETag: String
  Key: String
  Version: String

```

## Properties

`Bucket`

The name of the S3 bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ETag`

A file checksum of the pipeline definition file.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The object key (or key name) which uniquely identifies the object in an S3 bucket.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version ID of the pipeline definition file. If not specified, Amazon SageMaker will retrieve the latest
version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipelineDefinition

Tag

All content copied from https://docs.aws.amazon.com/.
