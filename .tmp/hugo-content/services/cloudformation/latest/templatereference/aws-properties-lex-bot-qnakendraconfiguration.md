This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot QnAKendraConfiguration

Contains details about the configuration of the Amazon Kendra index used for the `AMAZON.QnAIntent`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExactResponse" : Boolean,
  "KendraIndex" : String,
  "QueryFilterString" : String,
  "QueryFilterStringEnabled" : Boolean
}

```

### YAML

```yaml

  ExactResponse: Boolean
  KendraIndex: String
  QueryFilterString:
    String
  QueryFilterStringEnabled: Boolean

```

## Properties

`ExactResponse`

Specifies whether to return an exact response from the Amazon Kendra index or to let the Amazon Bedrock model you select generate a response based on the results. To use this feature, you must first add FAQ questions to your index by following the steps at [Adding frequently asked questions (FAQs) to an index](../../../kendra/latest/dg/in-creating-faq.md).

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KendraIndex`

The ARN of the Amazon Kendra index to use.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryFilterString`

Contains the Amazon Kendra filter string to use if enabled. For more information on the Amazon Kendra search filter JSON format, see [Using document attributes to filter search results](../../../kendra/latest/dg/filtering.md#search-filtering).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryFilterStringEnabled`

Specifies whether to enable an Amazon Kendra filter string or not.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QnAIntentConfiguration

Replication

All content copied from https://docs.aws.amazon.com/.
