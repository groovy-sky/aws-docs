# QueryStagePlanNode

Stage plan information such as name, identifier, sub plans, and remote sources.

## Contents

**Children**

Stage plan information such as name, identifier, sub plans, and remote sources of
child plan nodes/

Type: Array of [QueryStagePlanNode](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryStagePlanNode.html) objects

Required: No

**Identifier**

Information about the operation this query stage plan node is performing.

Type: String

Required: No

**Name**

Name of the query stage plan that describes the operation this stage is performing as
part of query execution.

Type: String

Required: No

**RemoteSources**

Source plan node IDs.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/QueryStagePlanNode)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/QueryStagePlanNode)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/QueryStagePlanNode)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryStage

ResultConfiguration
