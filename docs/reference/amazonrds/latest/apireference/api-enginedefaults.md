# EngineDefaults

Contains the result of a successful invocation of the `DescribeEngineDefaultParameters` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**DBParameterGroupFamily**

Specifies the name of the DB parameter group family that the engine default parameters apply to.

Type: String

Required: No

**Marker**

An optional pagination token provided by a previous
EngineDefaults request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords` .

Type: String

Required: No

**Parameters.Parameter.N**

Contains a list of engine default parameters.

Type: Array of [Parameter](api-parameter.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/enginedefaults.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/enginedefaults.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/enginedefaults.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Endpoint

Event

All content copied from https://docs.aws.amazon.com/.
