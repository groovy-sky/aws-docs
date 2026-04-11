This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Classifier

The `AWS::Glue::Classifier` resource creates an AWS Glue classifier that
categorizes data sources and specifies schemas. For more information, see [Adding Classifiers to\
a Crawler](../../../glue/latest/dg/add-classifier.md) and [Classifier Structure](../../../glue/latest/dg/aws-glue-api-crawler-classifiers.md#aws-glue-api-crawler-classifiers-Classifier) in the _AWS Glue Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Classifier",
  "Properties" : {
      "CsvClassifier" : CsvClassifier,
      "GrokClassifier" : GrokClassifier,
      "JsonClassifier" : JsonClassifier,
      "XMLClassifier" : XMLClassifier
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Classifier
Properties:
  CsvClassifier:
    CsvClassifier
  GrokClassifier:
    GrokClassifier
  JsonClassifier:
    JsonClassifier
  XMLClassifier:
    XMLClassifier

```

## Properties

`CsvClassifier`

A classifier for comma-separated values (CSV).

_Required_: Conditional

_Type_: [CsvClassifier](aws-properties-glue-classifier-csvclassifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrokClassifier`

A classifier that uses `grok`.

_Required_: Conditional

_Type_: [GrokClassifier](aws-properties-glue-classifier-grokclassifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JsonClassifier`

A classifier for JSON content.

_Required_: Conditional

_Type_: [JsonClassifier](aws-properties-glue-classifier-jsonclassifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XMLClassifier`

A classifier for XML content.

_Required_: Conditional

_Type_: [XMLClassifier](aws-properties-glue-classifier-xmlclassifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the classifier name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetRedshiftCatalog

CsvClassifier

All content copied from https://docs.aws.amazon.com/.
