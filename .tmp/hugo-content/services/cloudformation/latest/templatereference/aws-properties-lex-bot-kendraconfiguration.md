This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot KendraConfiguration

Provides configuration information for the `AMAZON.KendraSearchIntent`
intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra
index and returns documents from the index that match the user's
utterance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KendraIndex" : String,
  "QueryFilterString" : String,
  "QueryFilterStringEnabled" : Boolean
}

```

### YAML

```yaml

  KendraIndex: String
  QueryFilterString:
    String
  QueryFilterStringEnabled: Boolean

```

## Properties

`KendraIndex`

The Amazon Resource Name (ARN) of the Amazon Kendra index that you want the
`AMAZON.KendraSearchIntent` intent to search. The index must be in the
same account and Region as the Amazon Lex bot.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:kendra:[a-z]+-(?:[a-z]+-)*[0-9]:[0-9]{12}:index/[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `32`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryFilterString`

A query filter that Amazon Lex sends to Amazon Kendra to filter the response from
a query. The filter is in the format defined by Amazon Kendra. For more
information, see [Filtering\
queries](../../../kendra/latest/dg/filtering.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryFilterStringEnabled`

Determines whether the `AMAZON.KendraSearchIntent` intent uses a
custom query string to query the Amazon Kendra index.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntentOverride

LambdaCodeHook

All content copied from https://docs.aws.amazon.com/.
