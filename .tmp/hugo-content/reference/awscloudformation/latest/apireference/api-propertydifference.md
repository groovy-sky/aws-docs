# PropertyDifference

Information about a resource property whose actual value differs from its expected value, as
defined in the stack template and any values specified as template parameters. These will be
present only for resources whose `StackResourceDriftStatus` is `MODIFIED`.
For more information, see [Detect unmanaged\
configuration changes to stacks and resources with drift detection](../../../../services/cloudformation/latest/userguide/using-cfn-stack-drift.md).

## Contents

**ActualValue**

The actual property value of the resource property.

Type: String

Required: Yes

**DifferenceType**

The type of property difference.

- `ADD`: A value has been added to a resource property that's an array or list
data type.

- `REMOVE`: The property has been removed from the current resource
configuration.

- `NOT_EQUAL`: The current property value differs from its expected value (as
defined in the stack template and any values specified as template parameters).

Type: String

Valid Values: `ADD | REMOVE | NOT_EQUAL`

Required: Yes

**ExpectedValue**

The expected property value of the resource property, as defined in the stack template and
any values specified as template parameters.

Type: String

Required: Yes

**PropertyPath**

The fully-qualified path to the resource property.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/propertydifference.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/propertydifference.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/propertydifference.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PhysicalResourceIdContextKeyValuePair

RequiredActivatedType

All content copied from https://docs.aws.amazon.com/.
