# TemplateConfiguration

The configuration details of a generated template.

## Contents

**DeletionPolicy**

The `DeletionPolicy` assigned to resources in the generated template. Supported
values are:

- `DELETE` \- delete all resources when the stack is deleted.

- `RETAIN` \- retain all resources when the stack is deleted.

For more information, see [DeletionPolicy\
attribute](../../../../services/cloudformation/latest/userguide/aws-attribute-deletionpolicy.md) in the _AWS CloudFormation User Guide_.

Type: String

Valid Values: `DELETE | RETAIN`

Required: No

**UpdateReplacePolicy**

The `UpdateReplacePolicy` assigned to resources in the generated template.
Supported values are:

- `DELETE` \- delete all resources when the resource is replaced during an update
operation.

- `RETAIN` \- retain all resources when the resource is replaced during an update
operation.

For more information, see [UpdateReplacePolicy attribute](../../../../services/cloudformation/latest/userguide/aws-attribute-updatereplacepolicy.md) in the _AWS CloudFormation User Guide_.

Type: String

Valid Values: `DELETE | RETAIN`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/templateconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/templateconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/templateconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TemplateParameter

All content copied from https://docs.aws.amazon.com/.
