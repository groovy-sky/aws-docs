This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan AssociatedAlarm

An Amazon CloudWatch alarm associated with a Region switch plan. These alarms can be used to
trigger automatic execution of the plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmType" : String,
  "CrossAccountRole" : String,
  "ExternalId" : String,
  "ResourceIdentifier" : String
}

```

### YAML

```yaml

  AlarmType: String
  CrossAccountRole: String
  ExternalId: String
  ResourceIdentifier: String

```

## Properties

`AlarmType`

The alarm type for an associated alarm. An associated CloudWatch alarm can be an application health alarm or a trigger alarm.

_Required_: Yes

_Type_: String

_Allowed values_: `applicationHealth | trigger`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossAccountRole`

The cross account role for the configuration.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z0-9-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID (secret key) for the configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

The resource identifier for alarms that you associate with a plan.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Asg

CustomActionLambdaConfiguration
