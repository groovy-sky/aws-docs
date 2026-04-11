This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCZonalShift::ZonalAutoshiftConfiguration ControlCondition

A control condition is an alarm that you specify for a practice run. When you configure practice runs
with zonal autoshift for a resource, you specify Amazon CloudWatch alarms, which you create in CloudWatch
to use with the practice run. The alarms that you specify are an
_outcome alarm_, to monitor application health during practice runs and,
optionally, a _blocking alarm_, to block practice runs from starting or to interrupt
a practice run in progress.

Control condition alarms do not apply for autoshifts.

For more information, see
[Considerations when you configure zonal autoshift](../../../r53recovery/latest/dg/arc-zonal-autoshift-considerations.md) in the ARC Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmIdentifier" : String,
  "Type" : String
}

```

### YAML

```yaml

  AlarmIdentifier: String
  Type: String

```

## Properties

`AlarmIdentifier`

The Amazon Resource Name (ARN) for an Amazon CloudWatch alarm that you specify as a control condition for a practice run.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `8`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of alarm specified for a practice run. You can only specify Amazon CloudWatch alarms for practice runs, so the
only valid value is `CLOUDWATCH`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]*$`

_Minimum_: `8`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ARCZonalShift::ZonalAutoshiftConfiguration

PracticeRunConfiguration

All content copied from https://docs.aws.amazon.com/.
