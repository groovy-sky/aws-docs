This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource HookConfiguration

Provides the configuration information for invoking a Lambda function in AWS Lambda to alter document metadata and content when ingesting documents into
Amazon Kendra. You can configure your Lambda function using [PreExtractionHookConfiguration](../../../kendra/latest/dg/api-customdocumentenrichmentconfiguration.md) if you want to apply advanced alterations on
the original or raw documents. If you want to apply advanced alterations on the Amazon Kendra structured documents, you must configure your Lambda function using
[PostExtractionHookConfiguration](../../../kendra/latest/dg/api-customdocumentenrichmentconfiguration.md). You can only invoke one Lambda function.
However, this function can invoke other functions it requires.

For more information, see [Customizing document metadata\
during the ingestion process](../../../kendra/latest/dg/custom-document-enrichment.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InvocationCondition" : DocumentAttributeCondition,
  "LambdaArn" : String,
  "S3Bucket" : String
}

```

### YAML

```yaml

  InvocationCondition:
    DocumentAttributeCondition
  LambdaArn: String
  S3Bucket: String

```

## Properties

`InvocationCondition`

The condition used for when a Lambda function should be invoked.

For example, you can specify a condition that if there are empty date-time values,
then Amazon Kendra should invoke a function that inserts the current
date-time.

_Required_: No

_Type_: [DocumentAttributeCondition](aws-properties-kendra-datasource-documentattributecondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaArn`

The Amazon Resource Name (ARN) of an IAM role with permission to run a Lambda function
during ingestion. For more information, see [an IAM roles for Amazon Kendra](../../../kendra/latest/dg/iam-roles.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Bucket`

Stores the original, raw documents or the structured, parsed documents before and
after altering them. For more information, see [Data contracts for Lambda functions](../../../kendra/latest/dg/custom-document-enrichment.md#cde-data-contracts-lambda).

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GoogleDriveConfiguration

InlineCustomDocumentEnrichmentConfiguration

All content copied from https://docs.aws.amazon.com/.
