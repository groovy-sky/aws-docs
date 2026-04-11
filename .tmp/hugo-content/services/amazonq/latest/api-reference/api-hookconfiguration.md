# HookConfiguration

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

For more information, see [Custom document enrichment](../business-use-dg/custom-document-enrichment.md).

## Contents

**invocationCondition**

The condition used for when a Lambda function should be invoked.

For example, you can specify a condition that if there are empty date-time values,
then Amazon Q Business should invoke a function that inserts the current date-time.

Type: [DocumentAttributeCondition](api-documentattributecondition.md) object

Required: No

**lambdaArn**

The Amazon Resource Name (ARN) of the Lambda function during ingestion.
For more information, see [Using Lambda functions for Amazon Q Business document enrichment](../qbusiness-ug/cde-lambda-operations.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:aws[a-zA-Z-]*:lambda:[a-z-]*-[0-9]:[0-9]{12}:function:[a-zA-Z0-9-_]+(/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})?(:[a-zA-Z0-9-_]+)?`

Required: No

**roleArn**

The Amazon Resource Name (ARN) of a role with permission to run
`PreExtractionHookConfiguration` and
`PostExtractionHookConfiguration` for altering document metadata and
content during the document ingestion process.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: No

**s3BucketName**

Stores the original, raw documents or the structured, parsed documents before and
after altering them. For more information, see [Data contracts for Lambda functions](../business-use-dg/cde-lambda-operations.md#cde-lambda-operations-data-contracts).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/hookconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/hookconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/hookconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HallucinationReductionConfiguration

IdcAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
