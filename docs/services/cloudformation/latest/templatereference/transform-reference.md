This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# Transform reference

Transforms are macros hosted by CloudFormation. Unlike custom macros, a transform doesn't
require any special permissions to use it because it is hosted by CloudFormation. Transforms can
be used in templates in any account within CloudFormation. Also, there is no charge incurred
when using transforms. CloudFormation treats a transform the same as any other macro in terms of
evaluation order and scope.

For more information about how macros work, see [Using CloudFormation macros to perform custom\
processing on templates](../userguide/template-macros.md) in the _AWS CloudFormation User Guide_.

You can use the following transforms in your CloudFormation templates.

###### Topics

- [AWS::CodeDeployBlueGreen transform](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-codedeploybluegreen.html)

- [AWS::Include transform](transform-aws-include.md)

- [AWS::LanguageExtensions transform](transform-aws-languageextensions.md)

- [AWS::SecretsManager transform](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-secretsmanager.html)

- [AWS::Serverless transform](transform-aws-serverless.md)

- [AWS::ServiceCatalog transform](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-servicecatalog.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Rule functions

AWS::CodeDeployBlueGreen
