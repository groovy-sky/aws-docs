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

Type: [QueryStagePlanNode](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryStagePlanNode.html) object

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

Type: Array of [QueryStage](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryStage.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/QueryStage)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/QueryStage)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/QueryStage)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryRuntimeStatisticsTimeline

QueryStagePlanNode
