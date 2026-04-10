This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResilienceHub::ResiliencyPolicy FailurePolicy

Defines a failure policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RpoInSecs" : Integer,
  "RtoInSecs" : Integer
}

```

### YAML

```yaml

  RpoInSecs: Integer
  RtoInSecs: Integer

```

## Properties

`RpoInSecs`

Recovery Point Objective (RPO) in seconds.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RtoInSecs`

Recovery Time Objective (RTO) in seconds.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ResilienceHub::ResiliencyPolicy

PolicyMap

All content copied from https://docs.aws.amazon.com/.
