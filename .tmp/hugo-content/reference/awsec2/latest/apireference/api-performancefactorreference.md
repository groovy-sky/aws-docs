# PerformanceFactorReference

Specify an instance family to use as the baseline reference for CPU performance. All
instance types that match your specified attributes will be compared against the CPU
performance of the referenced instance family, regardless of CPU manufacturer or
architecture.

###### Note

Currently, only one instance family can be specified in the list.

## Contents

**InstanceFamily** (request), **instanceFamily** (response)

The instance family to use as a baseline reference.

###### Note

Ensure that you specify the correct value for the instance family. The instance
family is everything before the period ( `.`) in the instance type name. For
example, in the instance type `c6i.large`, the instance family is
`c6i`, not `c6`. For more information, see [Amazon EC2\
instance type naming conventions](../../../../services/ec2/latest/instancetypes/instance-type-names.md) in _Amazon EC2 Instance_
_Types_.

The following instance families are _not supported_ for performance
protection:

- `c1`

- `g3` \| `g3s`

- `hpc7g`

- `m1` \| `m2`

- `mac1` \| `mac2` \| `mac2-m1ultra` \|
`mac2-m2` \| `mac2-m2pro`

- `p3dn` \| `p4d` \| `p5`

- `t1`

- `u-12tb1` \| `u-18tb1` \| `u-24tb1` \|
`u-3tb1` \| `u-6tb1` \| `u-9tb1` \|
`u7i-12tb` \| `u7in-16tb` \| `u7in-24tb` \|
`u7in-32tb`

If you enable performance protection by specifying a supported instance family, the
returned instance types will exclude the above unsupported instance families.

If you specify an unsupported instance family as a value for baseline performance, the
API returns an empty response for [GetInstanceTypesFromInstanceRequirements](api-getinstancetypesfrominstancerequirements.md) and an exception for [CreateFleet](api-createfleet.md), [RequestSpotFleet](api-requestspotfleet.md), [ModifyFleet](api-modifyfleet.md), and [ModifySpotFleetRequest](api-modifyspotfleetrequest.md).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/performancefactorreference.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/performancefactorreference.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/performancefactorreference.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PeeringTgwInfo

PerformanceFactorReferenceRequest
