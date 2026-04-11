This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy IngressIpv4Expression

The union type representing the allowed types for the left hand side of an IP
condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : IngressIpToEvaluate,
  "Operator" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Evaluate:
    IngressIpToEvaluate
  Operator: String
  Values:
    - String

```

## Properties

`Evaluate`

The left hand side argument of an IP condition expression.

_Required_: Yes

_Type_: [IngressIpToEvaluate](aws-properties-ses-mailmanagertrafficpolicy-ingressiptoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The matching operator for an IP condition expression.

_Required_: Yes

_Type_: String

_Allowed values_: `CIDR_MATCHES | NOT_CIDR_MATCHES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The right hand side argument of an IP condition expression.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressIpToEvaluate

IngressIpv6Expression

All content copied from https://docs.aws.amazon.com/.
