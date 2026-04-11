This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy IngressStringToEvaluate

The union type representing the allowed types for the left hand side of a string
condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Analysis" : IngressAnalysis,
  "Attribute" : String
}

```

### YAML

```yaml

  Analysis:
    IngressAnalysis
  Attribute: String

```

## Properties

`Analysis`

The structure type for a string condition stating the Add On ARN and its returned value.

_Required_: No

_Type_: [IngressAnalysis](aws-properties-ses-mailmanagertrafficpolicy-ingressanalysis.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attribute`

The enum type representing the allowed attribute types for a string condition.

_Required_: No

_Type_: String

_Allowed values_: `RECIPIENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressStringExpression

IngressTlsProtocolExpression

All content copied from https://docs.aws.amazon.com/.
