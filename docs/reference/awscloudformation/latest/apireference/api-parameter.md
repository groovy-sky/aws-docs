# Parameter

The `Parameter` data type.

## Contents

**ParameterKey**

The key associated with the parameter. If you don't specify a key and value for a particular
parameter, CloudFormation uses the default value that's specified in your template.

Type: String

Required: No

**ParameterValue**

The input value associated with the parameter.

Type: String

Required: No

**ResolvedValue**

Read-only. The value that corresponds to a Systems Manager parameter key. This field is returned only
for Systems Manager parameter types in the template. For more information, see [Specify\
existing resources at runtime with CloudFormation-supplied parameter types](../../../../services/cloudformation/latest/userguide/cloudformation-supplied-parameter-types.md) in the
_AWS CloudFormation User Guide_.

Type: String

Required: No

**UsePreviousValue**

During a stack update, use the existing parameter value that the stack is using for a given
parameter key. If you specify `true`, do not specify a parameter value.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/parameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/parameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/parameter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Output

ParameterConstraints

All content copied from https://docs.aws.amazon.com/.
