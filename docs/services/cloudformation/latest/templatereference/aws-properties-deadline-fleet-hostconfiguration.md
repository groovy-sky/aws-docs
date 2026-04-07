This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet HostConfiguration

Provides a script that runs as a worker is starting up that you can use to provide
additional configuration for workers in your fleet.

To remove a script from a fleet, use the [UpdateFleet](https://docs.aws.amazon.com/deadline-cloud/latest/APIReference/API_UpdateFleet.html)
operation with the `hostConfiguration` `scriptBody` parameter set to an empty string ("").

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScriptBody" : String,
  "ScriptTimeoutSeconds" : Integer
}

```

### YAML

```yaml

  ScriptBody: String
  ScriptTimeoutSeconds: Integer

```

## Properties

`ScriptBody`

The text of the script that runs as a worker is starting up that you can use to provide
additional configuration for workers in your fleet. The script runs after a worker enters
the `STARTING` state and before the worker processes tasks.

For more information about using the script, see [Run scripts as an\
administrator to configure workers](https://docs.aws.amazon.com/deadline-cloud/latest/developerguide/smf-admin.html) in the _Deadline Cloud Developer_
_Guide_.

###### Important

The script runs as an administrative user ( `sudo root` on Linux, as an
Administrator on Windows).

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `15000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScriptTimeoutSeconds`

The maximum time that the host configuration can run. If the timeout expires, the worker
enters the `NOT RESPONDING` state and shuts down. You are charged for the time
that the worker is running the host configuration script.

###### Note

You should configure your fleet for a maximum of one worker while testing your host
configuration script to avoid starting additional workers.

The default is 300 seconds (5 minutes).

_Required_: No

_Type_: Integer

_Minimum_: `300`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FleetConfiguration

MemoryMiBRange
