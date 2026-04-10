# ModuleInfo

Contains information about the module from which the resource was created, if the resource
was created from a module included in the stack template.

For more information about modules, see [Create reusable resource configurations\
that can be included across templates with CloudFormation modules](../../../../services/cloudformation/latest/userguide/modules.md) in the
_AWS CloudFormation User Guide_.

## Contents

**LogicalIdHierarchy**

A concatenated list of the logical IDs of the module or modules that contains the resource.
Modules are listed starting with the inner-most nested module, and separated by
`/`.

In the following example, the resource was created from a module, `moduleA`,
that's nested inside a parent module, `moduleB`.

`moduleA/moduleB`

For more information, see [Reference module resources in\
CloudFormation templates](../../../../services/cloudformation/latest/userguide/module-ref-resources.md) in the _AWS CloudFormation User Guide_.

Type: String

Required: No

**TypeHierarchy**

A concatenated list of the module type or types that contains the resource. Module types are
listed starting with the inner-most nested module, and separated by `/`.

In the following example, the resource was created from a module of type
`AWS::First::Example::MODULE`, that's nested inside a parent module of type
`AWS::Second::Example::MODULE`.

`AWS::First::Example::MODULE/AWS::Second::Example::MODULE`

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/moduleinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/moduleinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/moduleinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedExecution

OperationEntry

All content copied from https://docs.aws.amazon.com/.
