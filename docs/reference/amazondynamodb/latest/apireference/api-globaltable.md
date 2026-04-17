---
title: "GlobalTable"
---

# GlobalTable

Represents the properties of a global table.

## Contents

###### Note

In the following list, the required parameters are described first.

**GlobalTableName**

The global table name.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**ReplicationGroup**

The Regions where the global table has replicas.

Type: Array of [Replica](api-replica.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/GlobalTable)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/GlobalTable)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/GlobalTable)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalSecondaryIndexWarmThroughputDescription

GlobalTableDescription

All content copied from https://docs.aws.amazon.com/.
