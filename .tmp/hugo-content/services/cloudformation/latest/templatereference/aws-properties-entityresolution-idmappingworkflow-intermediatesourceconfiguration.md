This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdMappingWorkflow IntermediateSourceConfiguration

The Amazon S3 location that temporarily stores your data while it processes.
Your information won't be saved permanently.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IntermediateS3Path" : String
}

```

### YAML

```yaml

  IntermediateS3Path: String

```

## Properties

`IntermediateS3Path`

The Amazon S3 location (bucket and prefix). For example:
`s3://provider_bucket/DOC-EXAMPLE-BUCKET`

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdMappingWorkflowOutputSource

ProviderProperties

All content copied from https://docs.aws.amazon.com/.
