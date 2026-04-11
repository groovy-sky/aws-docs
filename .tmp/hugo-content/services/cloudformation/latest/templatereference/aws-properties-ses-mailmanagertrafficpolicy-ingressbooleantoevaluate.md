This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy IngressBooleanToEvaluate

The union type representing the allowed types of operands for a boolean
condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Analysis" : IngressAnalysis,
  "IsInAddressList" : IngressIsInAddressList
}

```

### YAML

```yaml

  Analysis:
    IngressAnalysis
  IsInAddressList:
    IngressIsInAddressList

```

## Properties

`Analysis`

The structure type for a boolean condition stating the Add On ARN and its returned
value.

_Required_: No

_Type_: [IngressAnalysis](aws-properties-ses-mailmanagertrafficpolicy-ingressanalysis.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsInAddressList`

The structure type for a boolean condition that provides the address lists to evaluate incoming traffic on.

_Required_: No

_Type_: [IngressIsInAddressList](aws-properties-ses-mailmanagertrafficpolicy-ingressisinaddresslist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressBooleanExpression

IngressIpToEvaluate

All content copied from https://docs.aws.amazon.com/.
