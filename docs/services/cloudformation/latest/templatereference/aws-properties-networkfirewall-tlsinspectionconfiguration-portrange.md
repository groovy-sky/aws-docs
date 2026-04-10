This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::TLSInspectionConfiguration PortRange

A single port range specification. This is used for source and destination port ranges
in the stateless rule [MatchAttributes](../userguide/aws-properties-networkfirewall-rulegroup-matchattributes.md), `SourcePorts`, and
`DestinationPorts` settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FromPort" : Integer,
  "ToPort" : Integer
}

```

### YAML

```yaml

  FromPort: Integer
  ToPort: Integer

```

## Properties

`FromPort`

The lower limit of the port range. This must be less than or equal to the
`ToPort` specification.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToPort`

The upper limit of the port range. This must be greater than or equal to the
`FromPort` specification.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CheckCertificateRevocationStatus

ServerCertificate

All content copied from https://docs.aws.amazon.com/.
