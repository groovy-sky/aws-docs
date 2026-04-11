# ParameterDeclaration

The `ParameterDeclaration` data type.

## Contents

**DefaultValue**

The default value of the parameter.

Type: String

Required: No

**Description**

The description that's associate with the parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**NoEcho**

Flag that indicates whether the parameter value is shown as plain text in logs and in the
AWS Management Console.

Type: Boolean

Required: No

**ParameterConstraints**

The criteria that CloudFormation uses to validate parameter values.

Type: [ParameterConstraints](api-parameterconstraints.md) object

Required: No

**ParameterKey**

The name that's associated with the parameter.

Type: String

Required: No

**ParameterType**

The type of parameter.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/parameterdeclaration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/parameterdeclaration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/parameterdeclaration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterConstraints

PhysicalResourceIdContextKeyValuePair

All content copied from https://docs.aws.amazon.com/.
