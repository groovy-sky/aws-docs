# PlacementStrategy

The task placement strategy for a task or service. For more information, see [Task placement strategies](../../../../services/amazonecs/latest/developerguide/task-placement-strategies.md) in the _Amazon_
_Elastic Container Service Developer Guide_.

## Contents

**field**

The field to apply the placement strategy against. For the `spread`
placement strategy, valid values are `instanceId` (or `host`,
which has the same effect), or any platform or custom attribute that's applied to a
container instance, such as `attribute:ecs.availability-zone`. For the
`binpack` placement strategy, valid values are `cpu` and
`memory`. For the `random` placement strategy, this field is
not used.

Type: String

Required: No

**type**

The type of placement strategy. The `random` placement strategy randomly
places tasks on available candidates. The `spread` placement strategy spreads
placement across available candidates evenly based on the `field` parameter.
The `binpack` strategy places tasks on available candidates that have the
least available amount of the resource that's specified with the `field`
parameter. For example, if you binpack on memory, a task is placed on the instance with
the least amount of remaining memory but still enough to run the task.

Type: String

Valid Values: `random | spread | binpack`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/placementstrategy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/placementstrategy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/placementstrategy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PlacementConstraint

PlatformDevice
