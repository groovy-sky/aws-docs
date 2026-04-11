This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::MLTransform FindMatchesParameters

The parameters to configure the find matches transform.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccuracyCostTradeoff" : Number,
  "EnforceProvidedLabels" : Boolean,
  "PrecisionRecallTradeoff" : Number,
  "PrimaryKeyColumnName" : String
}

```

### YAML

```yaml

  AccuracyCostTradeoff: Number
  EnforceProvidedLabels: Boolean
  PrecisionRecallTradeoff: Number
  PrimaryKeyColumnName: String

```

## Properties

`AccuracyCostTradeoff`

The value that is selected when tuning your transform for a balance between accuracy and
cost. A value of 0.5 means that the system balances accuracy and cost concerns. A value of 1.0
means a bias purely for accuracy, which typically results in a higher cost, sometimes
substantially higher. A value of 0.0 means a bias purely for cost, which results in a less
accurate `FindMatches` transform, sometimes with unacceptable accuracy.

Accuracy measures how well the transform finds true positives and true negatives. Increasing accuracy requires more machine resources and cost. But it also results in increased recall.

Cost measures how many compute resources, and thus money, are consumed to run the
transform.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnforceProvidedLabels`

The value to switch on or off to force the output to match the provided labels from users. If the value is `True`, the `find matches` transform forces the output to match the provided labels. The results override the normal conflation results. If the value is `False`, the `find matches` transform does not ensure all the labels provided are respected, and the results rely on the trained model.

Note that setting this value to true may increase the conflation execution time.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrecisionRecallTradeoff`

The value selected when tuning your transform for a balance between precision and recall.
A value of 0.5 means no preference; a value of 1.0 means a bias purely for precision, and a
value of 0.0 means a bias for recall. Because this is a tradeoff, choosing values close to 1.0
means very low recall, and choosing values close to 0.0 results in very low precision.

The precision metric indicates how often your model is correct when it predicts a match.

The recall metric indicates that for an actual match, how often your model predicts the
match.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryKeyColumnName`

The name of a column that uniquely identifies rows in the source table. Used to help identify matching records.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::MLTransform

GlueTables

All content copied from https://docs.aws.amazon.com/.
