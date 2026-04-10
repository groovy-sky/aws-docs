# Perform custom processing on CloudFormation templates with template macros

With macros, you can perform custom processing on templates, from simple actions like
find-and-replace operations to extensive transformations of entire templates.

To get an idea of the breadth of possibilities, consider the `AWS::Include` and
`AWS::Serverless` transforms, which are macros hosted by CloudFormation:

- [AWS::Include transform](../templatereference/transform-aws-include.md) enables you to insert boilerplate template snippets into your templates.

- [AWS::Serverless transform](../templatereference/transform-aws-serverless.md) takes an entire template written in the AWS Serverless Application Model (AWS SAM) syntax and transforms
and expands it into a compliant CloudFormation template. For more information about
serverless applications and AWS SAM, see [AWS Serverless Application Model Developer Guide](../../../serverless-application-model/latest/developerguide/what-is-sam.md).

###### Topics

- [Billing](#template-macros-billing)

- [Macro examples](#template-macros-examples-list)

- [Related resources](#template-macros-related-resources)

- [Overview of CloudFormation macros](template-macros-overview.md)

- [Create a CloudFormation macro definition](template-macros-author.md)

- [Example simple string replacement macro](macros-example.md)

- [Troubleshoot the processed template](template-macros-troubleshoot-processed-template.md)

## Billing

When a macro runs, the owner of the Lambda function is billed for any charges related
to the execution of that function.

The `AWS::Include` and `AWS::Serverless` transforms are macros hosted by CloudFormation. There is no charge for using them.

## Macro examples

In addition to the examples in this section, you can find example macros, including
source code and templates, in our [GitHub repository](https://github.com/aws-cloudformation/aws-cloudformation-templates/tree/main/CloudFormation/MacrosExamples). These examples are provided 'as-is' for instructional
purposes.

## Related resources

- [AWS::CloudFormation::Macro](../templatereference/aws-resource-cloudformation-macro.md)

- [CloudFormation template Transform section](transform-section-structure.md)

- [Fn::Transform](../templatereference/intrinsic-function-reference-transform.md)

- [AWS::Serverless transform](../templatereference/transform-aws-serverless.md)

- [AWS::Include transform](../templatereference/transform-aws-include.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

cfn-response module

Macros overview

All content copied from https://docs.aws.amazon.com/.
