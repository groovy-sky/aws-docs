---
title: "EngineDefaults"
---

# EngineDefaults

Represents the output of a `DescribeEngineDefaultParameters`
operation.

## Contents

###### Note

In the following list, the required parameters are described first.

**CacheNodeTypeSpecificParameters.CacheNodeTypeSpecificParameter.N**

A list of parameters specific to a particular cache node type. Each element in the
list contains detailed information about one parameter.

Type: Array of [CacheNodeTypeSpecificParameter](api-cachenodetypespecificparameter.md) objects

Required: No

**CacheParameterGroupFamily**

Specifies the name of the cache parameter group family to which the engine default
parameters apply.

Valid values are: `memcached1.4` \| `memcached1.5` \|
`memcached1.6` \| `redis2.6` \| `redis2.8` \|
`redis3.2` \| `redis4.0` \| `redis5.0` \|
`redis6.0` \| `redis6.x` \| `redis7`

Type: String

Required: No

**Marker**

Provides an identifier to allow retrieval of paginated results.

Type: String

Required: No

**Parameters.Parameter.N**

Contains a list of engine default parameters.

Type: Array of [Parameter](api-parameter.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/EngineDefaults)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/EngineDefaults)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/EngineDefaults)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Endpoint

Event

All content copied from https://docs.aws.amazon.com/.
