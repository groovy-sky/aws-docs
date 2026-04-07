This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application AutoStopConfiguration

The configuration for an application to automatically stop after a certain amount of
time being idle.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "IdleTimeoutMinutes" : Integer
}

```

### YAML

```yaml

  Enabled: Boolean
  IdleTimeoutMinutes: Integer

```

## Properties

`Enabled`

Enables the application to automatically stop after a certain amount of time being idle.
Defaults to true.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`IdleTimeoutMinutes`

The amount of idle time in minutes after which your application will automatically stop.
Defaults to 15 minutes.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10080`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AutoStartConfiguration

CloudWatchLoggingConfiguration
