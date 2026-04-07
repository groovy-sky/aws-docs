This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Trigger Predicate

Defines the predicate of the trigger, which determines when it fires.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditions" : [ Condition, ... ],
  "Logical" : String
}

```

### YAML

```yaml

  Conditions:
    - Condition
  Logical: String

```

## Properties

`Conditions`

A list of the conditions that determine when the trigger will fire.

_Required_: No

_Type_: Array of [Condition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-trigger-condition.html)

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logical`

An optional field if only one condition is listed. If multiple conditions are listed,
then this field is required.

_Required_: No

_Type_: String

_Allowed values_: `AND | ANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Predicate Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-trigger.html#aws-glue-api-jobs-trigger-Predicate) in the _AWS Glue Developer_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotificationProperty

AWS::Glue::UsageProfile
