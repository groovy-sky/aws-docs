This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Classifier

The `AWS::Glue::Classifier` resource creates an AWS Glue classifier that
categorizes data sources and specifies schemas. For more information, see [Adding Classifiers to\
a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-classifier.html) and [Classifier Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-crawler-classifiers.html#aws-glue-api-crawler-classifiers-Classifier) in the _AWS Glue Developer Guide_.

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

_Type_: [CsvClassifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-classifier-csvclassifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrokClassifier`

A classifier that uses `grok`.

_Required_: Conditional

_Type_: [GrokClassifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-classifier-grokclassifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JsonClassifier`

A classifier for JSON content.

_Required_: Conditional

_Type_: [JsonClassifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-classifier-jsonclassifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XMLClassifier`

A classifier for XML content.

_Required_: Conditional

_Type_: [XMLClassifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-classifier-xmlclassifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the classifier name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetRedshiftCatalog

CsvClassifier
