This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy IngressBooleanExpression

The structure for a boolean condition matching on the incoming mail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : IngressBooleanToEvaluate,
  "Operator" : String
}

```

### YAML

```yaml

  Evaluate:
    IngressBooleanToEvaluate
  Operator: String

```

## Properties

`Evaluate`

The operand on which to perform a boolean condition operation.

_Required_: Yes

_Type_: [IngressBooleanToEvaluate](aws-properties-ses-mailmanagertrafficpolicy-ingressbooleantoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The matching operator for a boolean condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `IS_TRUE | IS_FALSE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressAnalysis

IngressBooleanToEvaluate

All content copied from https://docs.aws.amazon.com/.
