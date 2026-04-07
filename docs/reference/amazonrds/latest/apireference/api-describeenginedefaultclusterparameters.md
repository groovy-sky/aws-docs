# DescribeEngineDefaultClusterParameters

Returns the default engine and system parameter information for the cluster database engine.

For more information on Amazon Aurora, see
[What is Amazon Aurora?](../../../../services/amazonrds/latest/aurorauserguide/chap-auroraoverview.md) in the _Amazon Aurora User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBParameterGroupFamily**

The name of the DB cluster parameter group family to return engine parameter information for.

Type: String

Required: Yes

**Filters.Filter.N**

This parameter isn't currently supported.

Type: Array of [Filter](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Filter.html) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeEngineDefaultClusterParameters` request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Required: No

## Response Elements

The following element is returned by the service.

**EngineDefaults**

Contains the result of a successful invocation of the `DescribeEngineDefaultParameters` action.

Type: [EngineDefaults](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_EngineDefaults.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/rds-2014-10-31/DescribeEngineDefaultClusterParameters)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeDBSubnetGroups

DescribeEngineDefaultParameters
