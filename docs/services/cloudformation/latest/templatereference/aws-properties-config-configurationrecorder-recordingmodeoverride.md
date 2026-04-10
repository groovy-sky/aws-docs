This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigurationRecorder RecordingModeOverride

An object for you to specify your overrides for the recording mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "RecordingFrequency" : String,
  "ResourceTypes" : [ String, ... ]
}

```

### YAML

```yaml

  Description: String
  RecordingFrequency: String
  ResourceTypes:
    - String

```

## Properties

`Description`

A description that you provide for the override.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordingFrequency`

The recording frequency that will be applied to all the resource types specified in the override.

- Continuous recording allows you to record configuration changes continuously whenever a change occurs.

- Daily recording allows you to receive a configuration item (CI) representing the most recent state of your resources over the last 24-hour period, only if it’s different from the previous CI recorded.

###### Note

AWS Firewall Manager depends on continuous recording to monitor your resources. If you are using Firewall Manager,
it is recommended that you set the recording frequency to Continuous.

_Required_: Yes

_Type_: String

_Allowed values_: `CONTINUOUS | DAILY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTypes`

A comma-separated list that specifies which resource types AWS Config
includes in the override.

###### Important

Daily recording cannot be specified for the following resource types:

- `AWS::Config::ResourceCompliance`

- `AWS::Config::ConformancePackCompliance`

- `AWS::Config::ConfigurationRecorder`

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordingMode

RecordingStrategy

All content copied from https://docs.aws.amazon.com/.
