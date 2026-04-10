# StackResourceDrift

Contains the drift information for a resource that has been checked for drift. This includes
actual and expected property values for resources in which CloudFormation has detected drift. Only
resource properties explicitly defined in the stack template are checked for drift. For more
information, see [Detect unmanaged\
configuration changes to stacks and resources with drift detection](../../../../services/cloudformation/latest/userguide/using-cfn-stack-drift.md).

Resources that don't currently support drift detection can't be checked. For a list of
resources that support drift detection, see [Resource type\
support for imports and drift detection](../../../../services/cloudformation/latest/userguide/resource-import-supported-resources.md).

Use [DetectStackResourceDrift](api-detectstackresourcedrift.md) to detect drift on individual resources, or
[DetectStackDrift](api-detectstackdrift.md) to detect drift on all resources in a given stack that
support drift detection.

## Contents

**LogicalResourceId**

The logical name of the resource specified in the template.

Type: String

Required: Yes

**ResourceType**

The type of the resource.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**StackId**

The ID of the stack.

Type: String

Required: Yes

**StackResourceDriftStatus**

Status of the resource's actual configuration compared to its expected configuration.

- `DELETED`: The resource differs from its expected template configuration
because the resource has been deleted.

- `MODIFIED`: One or more resource properties differ from their expected values
(as defined in the stack template and any values specified as template parameters).

- `IN_SYNC`: The resource's actual configuration matches its expected template
configuration.

- `NOT_CHECKED`: CloudFormation does not currently return this value.

- `UNKNOWN`: CloudFormation could not run drift detection for the resource. See the
`DriftStatusReason` for details.

Type: String

Valid Values: `IN_SYNC | MODIFIED | DELETED | NOT_CHECKED | UNKNOWN | UNSUPPORTED`

Required: Yes

**Timestamp**

Time at which CloudFormation performed drift detection on the stack resource.

Type: Timestamp

Required: Yes

**ActualProperties**

A JSON structure that contains the actual property values of the stack resource.

For resources whose `StackResourceDriftStatus` is `DELETED`, this
structure will not be present.

Type: String

Required: No

**DriftStatusReason**

The reason for the drift status.

Type: String

Required: No

**ExpectedProperties**

A JSON structure that contains the expected property values of the stack resource, as
defined in the stack template and any values specified as template parameters.

For resources whose `StackResourceDriftStatus` is `DELETED`, this
structure will not be present.

Type: String

Required: No

**ModuleInfo**

Contains information about the module from which the resource was created, if the resource
was created from a module included in the stack template.

Type: [ModuleInfo](api-moduleinfo.md) object

Required: No

**PhysicalResourceId**

The name or unique identifier that corresponds to a physical instance ID of a resource
supported by CloudFormation.

Type: String

Required: No

**PhysicalResourceIdContext.member.N**

Context information that enables CloudFormation to uniquely identify a resource. CloudFormation uses
context key-value pairs in cases where a resource's logical and physical IDs aren't enough to
uniquely identify that resource. Each context key-value pair specifies a unique resource that
contains the targeted resource.

Type: Array of [PhysicalResourceIdContextKeyValuePair](api-physicalresourceidcontextkeyvaluepair.md) objects

Array Members: Maximum number of 5 items.

Required: No

**PropertyDifferences.member.N**

A collection of the resource properties whose actual values differ from their expected
values. These will be present only for resources whose `StackResourceDriftStatus` is
`MODIFIED`.

Type: Array of [PropertyDifference](api-propertydifference.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/stackresourcedrift.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/stackresourcedrift.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/stackresourcedrift.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StackResourceDetail

StackResourceDriftInformation

All content copied from https://docs.aws.amazon.com/.
