This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup EngineConfiguration

The engine configuration for the workgroup, which includes the minimum/maximum number of Data Processing Units (DPU) that queries should use when
running in provisioned capacity. If not specified, Athena uses default values (Default value for min is 4 and for max is Minimum of 124 and allocated DPUs).

To specify DPU values for PC queries the WG containing EngineConfiguration should have the following values:
The name of the Classifications should be `athena-query-engine-properties`, with the only allowed properties as `max-dpu-count` and `min-dpu-count`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalConfigs" : {Key: Value, ...},
  "Classifications" : [ Classification, ... ],
  "CoordinatorDpuSize" : Integer,
  "DefaultExecutorDpuSize" : Integer,
  "MaxConcurrentDpus" : Integer,
  "SparkProperties" : {Key: Value, ...}
}

```

### YAML

```yaml

  AdditionalConfigs:
    Key: Value
  Classifications:
    - Classification
  CoordinatorDpuSize: Integer
  DefaultExecutorDpuSize: Integer
  MaxConcurrentDpus: Integer
  SparkProperties:
    Key: Value

```

## Properties

`AdditionalConfigs`

Contains additional notebook engine `MAP<string, string>` parameter
mappings in the form of key-value pairs. To specify an Athena notebook that
the Jupyter server will download and serve, specify a value for the StartSessionRequest$NotebookVersion field, and then add a key named
`NotebookId` to `AdditionalConfigs` that has the value of the
Athena notebook ID.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Classifications`

The configuration classifications that can be specified for the engine.

_Required_: No

_Type_: Array of [Classification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-athena-workgroup-classification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CoordinatorDpuSize`

The number of DPUs to use for the coordinator. A coordinator is a special executor
that orchestrates processing work and manages other executors in a notebook session. The
default is 1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultExecutorDpuSize`

The default number of DPUs to use for executors. An executor is the smallest unit of
compute that a notebook session can request from Athena. The default is
1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrentDpus`

The maximum number of DPUs that can run concurrently.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparkProperties`

Specifies custom jar files and Spark properties for use cases like cluster encryption,
table formats, and general Spark tuning.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EncryptionConfiguration

EngineVersion
