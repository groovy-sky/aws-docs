# QueryStage

Stage statistics such as input and output rows and bytes, execution time and stage
state. This information also includes substages and the query stage plan.

## Contents

**ExecutionTime**

Time taken to execute this stage.

Type: Long

Required: No

**InputBytes**

The number of bytes input into the stage for execution.

Type: Long

Required: No

**InputRows**

The number of rows input into the stage for execution.

Type: Long

Required: No

**OutputBytes**

The number of bytes output from the stage after execution.

Type: Long

Required: No

**OutputRows**

The number of rows output from the stage after execution.

Type: Long

Required: No

**QueryStagePlan**

Stage plan information such as name, identifier, sub plans, and source stages.

Type: [QueryStagePlanNode](api-querystageplannode.md) object

Required: No

**StageId**

The identifier for a stage.

Type: Long

Required: No

**State**

State of the stage after query execution.

Type: String

Required: No

**SubStages**

List of sub query stages that form this stage execution plan.

Type: Array of [QueryStage](api-querystage.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/querystage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/querystage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/querystage.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

QueryRuntimeStatisticsTimeline

QueryStagePlanNode
