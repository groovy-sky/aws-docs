This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Classifier GrokClassifier

A classifier that uses `grok` patterns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Classification" : String,
  "CustomPatterns" : String,
  "GrokPattern" : String,
  "Name" : String
}

```

### YAML

```yaml

  Classification: String
  CustomPatterns: String
  GrokPattern: String
  Name: String

```

## Properties

`Classification`

An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, and
so on.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomPatterns`

Optional custom grok patterns defined by this classifier. For more information, see
custom patterns in [Writing Custom\
Classifiers](../../../glue/latest/dg/custom-classifier.md).

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `16000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrokPattern`

The grok pattern applied to a data store by this classifier. For more information, see
built-in patterns in [Writing Custom\
Classifiers](../../../glue/latest/dg/custom-classifier.md).

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\t]*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the classifier.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CsvClassifier

JsonClassifier

All content copied from https://docs.aws.amazon.com/.
