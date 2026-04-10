This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Classifier JsonClassifier

A classifier for `JSON` content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JsonPath" : String,
  "Name" : String
}

```

### YAML

```yaml

  JsonPath: String
  Name: String

```

## Properties

`JsonPath`

A `JsonPath` string defining the JSON data for the classifier to classify.
AWS Glue supports a subset of `JsonPath`, as described in [Writing JsonPath\
Custom Classifiers](../../../glue/latest/dg/custom-classifier.md#custom-classifier-json).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the classifier.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [JsonClassifier Structure](../../../glue/latest/dg/aws-glue-api-crawler-classifiers.md#aws-glue-api-crawler-classifiers-JsonClassifier) in the _AWS Glue Developer_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrokClassifier

XMLClassifier

All content copied from https://docs.aws.amazon.com/.
