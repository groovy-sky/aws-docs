# QueryStagePlanNode

Stage plan information such as name, identifier, sub plans, and remote sources.

## Contents

**Children**

Stage plan information such as name, identifier, sub plans, and remote sources of
child plan nodes/

Type: Array of [QueryStagePlanNode](api-querystageplannode.md) objects

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

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/querystageplannode.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/querystageplannode.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/querystageplannode.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

QueryStage

ResultConfiguration
