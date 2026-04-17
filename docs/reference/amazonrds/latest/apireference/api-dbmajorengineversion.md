---
title: "DBMajorEngineVersion"
---

# DBMajorEngineVersion

This data type is used as a response element in the operation
`DescribeDBMajorEngineVersions`.

## Contents

###### Note

In the following list, the required parameters are described first.

**Engine**

The name of the database engine.

Type: String

Required: No

**MajorEngineVersion**

The major version number of the database engine.

Type: String

Required: No

**SupportedEngineLifecycles.SupportedEngineLifecycle.N**

A list of the lifecycles supported by this engine for the
`DescribeDBMajorEngineVersions` operation.

Type: Array of [SupportedEngineLifecycle](api-supportedenginelifecycle.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/rds-2014-10-31/DBMajorEngineVersion)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/rds-2014-10-31/DBMajorEngineVersion)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/rds-2014-10-31/DBMajorEngineVersion)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBInstanceStatusInfo

DBParameterGroup

All content copied from https://docs.aws.amazon.com/.
