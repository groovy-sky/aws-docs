This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SalesforceKnowledgeArticleConfiguration

Provides the configuration information for the knowledge article types that Amazon Kendra indexes. Amazon Kendra indexes standard knowledge articles and the
standard fields of knowledge articles, or the custom fields of custom knowledge
articles, but not both

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomKnowledgeArticleTypeConfigurations" : [ SalesforceCustomKnowledgeArticleTypeConfiguration, ... ],
  "IncludedStates" : [ String, ... ],
  "StandardKnowledgeArticleTypeConfiguration" : SalesforceStandardKnowledgeArticleTypeConfiguration
}

```

### YAML

```yaml

  CustomKnowledgeArticleTypeConfigurations:
    - SalesforceCustomKnowledgeArticleTypeConfiguration
  IncludedStates:
    - String
  StandardKnowledgeArticleTypeConfiguration:
    SalesforceStandardKnowledgeArticleTypeConfiguration

```

## Properties

`CustomKnowledgeArticleTypeConfigurations`

Configuration information for custom Salesforce knowledge articles.

_Required_: No

_Type_: Array of [SalesforceCustomKnowledgeArticleTypeConfiguration](aws-properties-kendra-datasource-salesforcecustomknowledgearticletypeconfiguration.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedStates`

Specifies the document states that should be included when Amazon Kendra indexes
knowledge articles. You must specify at least one state.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandardKnowledgeArticleTypeConfiguration`

Configuration information for standard Salesforce knowledge articles.

_Required_: No

_Type_: [SalesforceStandardKnowledgeArticleTypeConfiguration](aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceCustomKnowledgeArticleTypeConfiguration

SalesforceStandardKnowledgeArticleTypeConfiguration

All content copied from https://docs.aws.amazon.com/.
