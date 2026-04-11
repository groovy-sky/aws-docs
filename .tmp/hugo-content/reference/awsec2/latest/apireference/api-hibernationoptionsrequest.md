# HibernationOptionsRequest

Indicates whether your instance is configured for hibernation. This parameter is valid
only if the instance meets the [hibernation\
prerequisites](../../../../services/ec2/latest/userguide/hibernating-prerequisites.md). For more information, see [Hibernate your Amazon EC2\
instance](../../../../services/ec2/latest/userguide/hibernate.md) in the _Amazon EC2 User Guide_.

## Contents

**Configured**

Set to `true` to enable your instance for hibernation.

For Spot Instances, if you set `Configured` to `true`, either
omit the `InstanceInterruptionBehavior` parameter (for [`SpotMarketOptions`](api-spotmarketoptions.md)), or set it to
`hibernate`. When `Configured` is true:

- If you omit `InstanceInterruptionBehavior`, it defaults to
`hibernate`.

- If you set `InstanceInterruptionBehavior` to a value other than
`hibernate`, you'll get an error.

Default: `false`

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/hibernationoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/hibernationoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/hibernationoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HibernationOptions

HistoryRecord
