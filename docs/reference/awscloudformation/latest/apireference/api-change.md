# Change

The `Change` structure describes the changes CloudFormation will perform if you
execute the change set.

## Contents

**HookInvocationCount**

Is either `null`, if no Hooks invoke for the resource, or contains the number
of Hooks that will invoke for the resource.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**ResourceChange**

A `ResourceChange` structure that describes the resource and action that
CloudFormation will perform.

Type: [ResourceChange](api-resourcechange.md) object

Required: No

**Type**

The type of entity that CloudFormation changes.

- `Resource` This change is for a resource.

Type: String

Valid Values: `Resource`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/change.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/change.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/change.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchDescribeTypeConfigurationsError

ChangeSetHook

All content copied from https://docs.aws.amazon.com/.
