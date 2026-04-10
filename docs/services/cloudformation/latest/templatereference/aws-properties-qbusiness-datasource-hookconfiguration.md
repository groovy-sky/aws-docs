This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource HookConfiguration

Provides the configuration information for invoking a Lambda function in
AWS Lambda to alter document metadata and content when ingesting
documents into Amazon Q Business.

You can configure your Lambda function using the
`PreExtractionHookConfiguration` parameter if you want to apply advanced
alterations on the original or raw documents.

If you want to apply advanced alterations on the Amazon Q Business structured documents,
you must configure your Lambda function using
`PostExtractionHookConfiguration`.

You can only invoke one Lambda function. However, this function can invoke
other functions it requires.

For more information, see [Custom document enrichment](../../../amazonq/latest/business-use-dg/custom-document-enrichment.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InvocationCondition" : DocumentAttributeCondition,
  "LambdaArn" : String,
  "RoleArn" : String,
  "S3BucketName" : String
}

```

### YAML

```yaml

  InvocationCondition:
    DocumentAttributeCondition
  LambdaArn: String
  RoleArn: String
  S3BucketName: String

```

## Properties

`InvocationCondition`

The condition used for when a Lambda function should be invoked.

For example, you can specify a condition that if there are empty date-time values,
then Amazon Q Business should invoke a function that inserts the current date-time.

_Required_: No

_Type_: [DocumentAttributeCondition](aws-properties-qbusiness-datasource-documentattributecondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaArn`

The Amazon Resource Name (ARN) of the Lambda function during ingestion.
For more information, see [Using Lambda functions for Amazon Q Business document enrichment](../../../amazonq/latest/qbusiness-ug/cde-lambda-operations.md).

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:lambda:[a-z-]*-[0-9]:[0-9]{12}:function:[a-zA-Z0-9-_]+(/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})?(:[a-zA-Z0-9-_]+)?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of a role with permission to run
`PreExtractionHookConfiguration` and
`PostExtractionHookConfiguration` for altering document metadata and
content during the document ingestion process.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

Stores the original, raw documents or the structured, parsed documents before and
after altering them. For more information, see [Data contracts for Lambda functions](../../../amazonq/latest/business-use-dg/cde-lambda-operations.md#cde-lambda-operations-data-contracts).

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentEnrichmentConfiguration

ImageExtractionConfiguration

All content copied from https://docs.aws.amazon.com/.
