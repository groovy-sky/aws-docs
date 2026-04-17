---
title: "CacheSecurityGroup"
---

# CacheSecurityGroup

Represents the output of one of the following operations:

- `AuthorizeCacheSecurityGroupIngress`

- `CreateCacheSecurityGroup`

- `RevokeCacheSecurityGroupIngress`

## Contents

###### Note

In the following list, the required parameters are described first.

**ARN**

The ARN of the cache security group,

Type: String

Required: No

**CacheSecurityGroupName**

The name of the cache security group.

Type: String

Required: No

**Description**

The description of the cache security group.

Type: String

Required: No

**EC2SecurityGroups.EC2SecurityGroup.N**

A list of Amazon EC2 security groups that are associated with this cache security
group.

Type: Array of [EC2SecurityGroup](api-ec2securitygroup.md) objects

Required: No

**OwnerId**

The Amazon account ID of the cache security group owner.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/CacheSecurityGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/CacheSecurityGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/CacheSecurityGroup)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CacheParameterGroupStatus

CacheSecurityGroupMembership

All content copied from https://docs.aws.amazon.com/.
