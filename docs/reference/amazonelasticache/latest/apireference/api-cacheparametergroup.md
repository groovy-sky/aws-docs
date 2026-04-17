---
title: "CacheParameterGroup"
---

# CacheParameterGroup

Represents the output of a `CreateCacheParameterGroup` operation.

## Contents

###### Note

In the following list, the required parameters are described first.

**ARN**

The ARN (Amazon Resource Name) of the cache parameter group.

Type: String

Required: No

**CacheParameterGroupFamily**

The name of the cache parameter group family that this cache parameter group is
compatible with.

Valid values are: `memcached1.4` \| `memcached1.5` \|
`memcached1.6` \| `redis2.6` \| `redis2.8` \|
`redis3.2` \| `redis4.0` \| `redis5.0` \|
`redis6.x` \| `redis7`

Type: String

Required: No

**CacheParameterGroupName**

The name of the cache parameter group.

Type: String

Required: No

**Description**

The description for this cache parameter group.

Type: String

Required: No

**IsGlobal**

Indicates whether the parameter group is associated with a Global datastore

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/CacheParameterGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/CacheParameterGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/CacheParameterGroup)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CacheNodeUpdateStatus

CacheParameterGroupStatus

All content copied from https://docs.aws.amazon.com/.
