# EngineConfiguration

The engine configuration for the workgroup, which includes the minimum/maximum number of Data Processing Units (DPU) that queries should use when
running in provisioned capacity. If not specified, Athena uses default values (Default value for min is 4 and for max is Minimum of 124 and allocated DPUs).

To specify DPU values for PC queries the WG containing EngineConfiguration should have the following values:
The name of the Classifications should be `athena-query-engine-properties`, with the only allowed properties as `max-dpu-count` and `min-dpu-count`.

## Contents

**AdditionalConfigs**

Contains additional notebook engine `MAP<string, string>` parameter
mappings in the form of key-value pairs. To specify an Athena notebook that
the Jupyter server will download and serve, specify a value for the [StartSession:NotebookVersion](api-startsession.md#athena-StartSession-request-NotebookVersion) field, and then add a key named
`NotebookId` to `AdditionalConfigs` that has the value of the
Athena notebook ID.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 255.

Key Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Value Length Constraints: Maximum length of 51200.

Required: No

**Classifications**

The configuration classifications that can be specified for the engine.

Type: Array of [Classification](https://docs.aws.amazon.com/athena/latest/APIReference/API_Classification.html) objects

Required: No

**CoordinatorDpuSize**

The number of DPUs to use for the coordinator. A coordinator is a special executor
that orchestrates processing work and manages other executors in a notebook session. The
default is 1.

Type: Integer

Valid Range: Fixed value of 1.

Required: No

**DefaultExecutorDpuSize**

The default number of DPUs to use for executors. An executor is the smallest unit of
compute that a notebook session can request from Athena. The default is
1.

Type: Integer

Valid Range: Fixed value of 1.

Required: No

**MaxConcurrentDpus**

The maximum number of DPUs that can run concurrently.

Type: Integer

Valid Range: Minimum value of 2. Maximum value of 5000.

Required: No

**SparkProperties**

Specifies custom jar files and Spark properties for use cases like cluster encryption,
table formats, and general Spark tuning.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 255.

Key Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Value Length Constraints: Maximum length of 51200.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/EngineConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/EngineConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/EngineConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EncryptionConfiguration

EngineVersion
