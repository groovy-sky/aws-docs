This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent KnowledgeBaseAssociationConfigurationData

The data of the configuration for a `KNOWLEDGE_BASE` type Amazon Q in
Connect Assistant Association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentTagFilter" : TagFilter,
  "MaxResults" : Number,
  "OverrideKnowledgeBaseSearchType" : String
}

```

### YAML

```yaml

  ContentTagFilter:
    TagFilter
  MaxResults: Number
  OverrideKnowledgeBaseSearchType: String

```

## Properties

`ContentTagFilter`

An object that can be used to specify Tag conditions.

_Required_: No

_Type_: [TagFilter](aws-properties-wisdom-aiagent-tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxResults`

The maximum number of results to return per page.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideKnowledgeBaseSearchType`

The search type to be used against the Knowledge Base for this request. The values can be `SEMANTIC`
which uses vector embeddings or `HYBRID` which use vector embeddings and raw text

_Required_: No

_Type_: String

_Allowed values_: `HYBRID | SEMANTIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EmailResponseAIAgentConfiguration

ManualSearchAIAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
