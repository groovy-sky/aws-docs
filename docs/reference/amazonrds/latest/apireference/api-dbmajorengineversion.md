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

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbmajorengineversion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbmajorengineversion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbmajorengineversion.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DBInstanceStatusInfo

DBParameterGroup
