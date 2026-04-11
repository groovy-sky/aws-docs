This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy IngressTlsProtocolExpression

The structure for a TLS related condition matching on the incoming mail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : IngressTlsProtocolToEvaluate,
  "Operator" : String,
  "Value" : String
}

```

### YAML

```yaml

  Evaluate:
    IngressTlsProtocolToEvaluate
  Operator: String
  Value: String

```

## Properties

`Evaluate`

The left hand side argument of a TLS condition expression.

_Required_: Yes

_Type_: [IngressTlsProtocolToEvaluate](aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocoltoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The matching operator for a TLS condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `MINIMUM_TLS_VERSION | IS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The right hand side argument of a TLS condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `TLS1_2 | TLS1_3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressStringToEvaluate

IngressTlsProtocolToEvaluate

All content copied from https://docs.aws.amazon.com/.
