# DBProxyTargetGroup

Represents a set of RDS DB instances, Aurora DB clusters, or both that a proxy can connect to. Currently, each target group
is associated with exactly one RDS DB instance or Aurora DB cluster.

This data type is used as a response element in the `DescribeDBProxyTargetGroups` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**ConnectionPoolConfig**

The settings that determine the size and behavior of the connection pool for the target group.

Type: [ConnectionPoolConfigurationInfo](api-connectionpoolconfigurationinfo.md) object

Required: No

**CreatedDate**

The date and time when the target group was first created.

Type: Timestamp

Required: No

**DBProxyName**

The identifier for the RDS proxy associated with this target group.

Type: String

Required: No

**IsDefault**

Indicates whether this target group is the first one used for connection requests by the associated proxy.
Because each proxy is currently associated with a single target group, currently this setting
is always `true`.

Type: Boolean

Required: No

**Status**

The current status of this target group. A status of `available` means the
target group is correctly associated with a database. Other values indicate that you must wait for
the target group to be ready, or take some action to resolve an issue.

Type: String

Required: No

**TargetGroupArn**

The Amazon Resource Name (ARN) representing the target group.

Type: String

Required: No

**TargetGroupName**

The identifier for the target group. This name must be unique for all target groups owned by your AWS account in the specified AWS Region.

Type: String

Required: No

**UpdatedDate**

The date and time when the target group was last updated.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbproxytargetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbproxytargetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbproxytargetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBProxyTarget

DBRecommendation

All content copied from https://docs.aws.amazon.com/.
