---
title: "CacheParameterGroupStatus"
---

# CacheParameterGroupStatus

Status of the cache parameter group.

## Contents

###### Note

In the following list, the required parameters are described first.

**CacheNodeIdsToReboot.CacheNodeId.N**

A list of the cache node IDs which need to be rebooted for parameter changes to be
applied. A node ID is a numeric identifier (0001, 0002, etc.).

Type: Array of strings

Required: No

**CacheParameterGroupName**

The name of the cache parameter group.

Type: String

Required: No

**ParameterApplyStatus**

The status of parameter updates.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/CacheParameterGroupStatus)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/CacheParameterGroupStatus)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/CacheParameterGroupStatus)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CacheParameterGroup

CacheSecurityGroup

All content copied from https://docs.aws.amazon.com/.
