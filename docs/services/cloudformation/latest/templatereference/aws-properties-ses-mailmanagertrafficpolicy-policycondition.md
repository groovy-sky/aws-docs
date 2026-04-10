This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy PolicyCondition

The email traffic filtering conditions which are contained in a traffic policy
resource.

###### Important

This data type is a UNION, so only one of the following members can be specified
when used or returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanExpression" : IngressBooleanExpression,
  "IpExpression" : IngressIpv4Expression,
  "Ipv6Expression" : IngressIpv6Expression,
  "StringExpression" : IngressStringExpression,
  "TlsExpression" : IngressTlsProtocolExpression
}

```

### YAML

```yaml

  BooleanExpression:
    IngressBooleanExpression
  IpExpression:
    IngressIpv4Expression
  Ipv6Expression:
    IngressIpv6Expression
  StringExpression:
    IngressStringExpression
  TlsExpression:
    IngressTlsProtocolExpression

```

## Properties

`BooleanExpression`

This represents a boolean type condition matching on the incoming mail. It performs
the boolean operation configured in 'Operator' and evaluates the 'Protocol' object
against the 'Value'.

_Required_: No

_Type_: [IngressBooleanExpression](aws-properties-ses-mailmanagertrafficpolicy-ingressbooleanexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpExpression`

This represents an IP based condition matching on the incoming mail. It performs the
operation configured in 'Operator' and evaluates the 'Protocol' object against the
'Value'.

_Required_: No

_Type_: [IngressIpv4Expression](aws-properties-ses-mailmanagertrafficpolicy-ingressipv4expression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Expression`

This represents an IPv6 based condition matching on the incoming mail. It performs the
operation configured in 'Operator' and evaluates the 'Protocol' object against the
'Value'.

_Required_: No

_Type_: [IngressIpv6Expression](aws-properties-ses-mailmanagertrafficpolicy-ingressipv6expression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringExpression`

This represents a string based condition matching on the incoming mail. It performs
the string operation configured in 'Operator' and evaluates the 'Protocol' object
against the 'Value'.

_Required_: No

_Type_: [IngressStringExpression](aws-properties-ses-mailmanagertrafficpolicy-ingressstringexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TlsExpression`

This represents a TLS based condition matching on the incoming mail. It performs the
operation configured in 'Operator' and evaluates the 'Protocol' object against the
'Value'.

_Required_: No

_Type_: [IngressTlsProtocolExpression](aws-properties-ses-mailmanagertrafficpolicy-ingresstlsprotocolexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressTlsProtocolToEvaluate

PolicyStatement

All content copied from https://docs.aws.amazon.com/.
