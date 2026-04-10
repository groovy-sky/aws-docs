# Export

The `Export` structure describes the exported output values for a stack.

For more information, see [Get exported outputs from\
a deployed CloudFormation stack](../../../../services/cloudformation/latest/userguide/using-cfn-stack-exports.md).

## Contents

**ExportingStackId**

The stack that contains the exported output name and value.

Type: String

Required: No

**Name**

The name of exported output value. Use this name and the `Fn::ImportValue`
function to import the associated value into other stacks. The name is defined in the
`Export` field in the associated stack's `Outputs` section.

Type: String

Required: No

**Value**

The value of the exported output, such as a resource physical ID. This value is defined in
the `Export` field in the associated stack's `Outputs` section.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/export.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/export.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/export.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventFilter

HookResultSummary

All content copied from https://docs.aws.amazon.com/.
