This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdMappingWorkflow IdMappingWorkflowOutputSource

A list of `IdMappingWorkflowOutputSource` objects, each of which contains
fields `outputS3Path` and `KMSArn`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KMSArn" : String,
  "OutputS3Path" : String
}

```

### YAML

```yaml

  KMSArn: String
  OutputS3Path: String

```

## Properties

`KMSArn`

Customer AWS KMS ARN for encryption at rest. If not provided, system will use
an AWS Entity Resolution managed KMS key.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):kms:.*:[0-9]+:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputS3Path`

The S3 path to which AWS Entity Resolution will write the output table.

_Required_: Yes

_Type_: String

_Pattern_: `^s3://([^/]+)/?(.*?([^/]+)/?)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdMappingWorkflowInputSource

IntermediateSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
