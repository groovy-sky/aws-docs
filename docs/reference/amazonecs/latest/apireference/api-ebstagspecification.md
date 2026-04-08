# EBSTagSpecification

The tag specifications of an Amazon EBS volume.

## Contents

**resourceType**

The type of volume resource.

Type: String

Valid Values: `volume`

Required: Yes

**propagateTags**

Determines whether to propagate the tags from the task definition to the
Amazon EBS volume. Tags can only propagate to a `SERVICE` specified in
`ServiceVolumeConfiguration`. If no value is specified, the tags
aren't propagated.

Type: String

Valid Values: `TASK_DEFINITION | SERVICE | NONE`

Required: No

**tags**

The tags applied to this Amazon EBS volume. `AmazonECSCreated` and
`AmazonECSManaged` are reserved tags that can't be used.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/ebstagspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/ebstagspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/ebstagspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DockerVolumeConfiguration

ECSExpressGatewayService
