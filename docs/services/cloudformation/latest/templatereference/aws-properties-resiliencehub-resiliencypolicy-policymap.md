This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResilienceHub::ResiliencyPolicy PolicyMap

The type of resiliency policy to be created, including the recovery time objective (RTO)
and recovery point objective (RPO) in seconds.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AZ" : FailurePolicy,
  "Hardware" : FailurePolicy,
  "Region" : FailurePolicy,
  "Software" : FailurePolicy
}

```

### YAML

```yaml

  AZ:
    FailurePolicy
  Hardware:
    FailurePolicy
  Region:
    FailurePolicy
  Software:
    FailurePolicy

```

## Properties

`AZ`

Defines the RTO and RPO targets for Availability Zone disruption.

_Required_: Yes

_Type_: [FailurePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Hardware`

Defines the RTO and RPO targets for hardware disruption.

_Required_: Yes

_Type_: [FailurePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

Defines the RTO and RPO targets for Regional disruption.

_Required_: No

_Type_: [FailurePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Software`

Defines the RTO and RPO targets for software disruption.

_Required_: Yes

_Type_: [FailurePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resiliencehub-resiliencypolicy-failurepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FailurePolicy

Next
