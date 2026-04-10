This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigurationRecorder RecordingMode

Specifies the default recording frequency that AWS Config uses to record configuration changes.

AWS Config supports _Continuous recording_ and _Daily recording_.

- Continuous recording allows you to record configuration changes continuously whenever a change occurs.

- Daily recording allows you to receive a configuration item (CI) representing the most recent state of your resources over the last 24-hour period, only if it’s different from the previous CI recorded.

###### Note

AWS Firewall Manager depends on continuous recording to monitor your resources. If you are using Firewall Manager,
it is recommended that you set the recording frequency to Continuous.

You can also override the recording frequency for specific resource types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordingFrequency" : String,
  "RecordingModeOverrides" : [ RecordingModeOverride, ... ]
}

```

### YAML

```yaml

  RecordingFrequency: String
  RecordingModeOverrides:
    - RecordingModeOverride

```

## Properties

`RecordingFrequency`

The default recording frequency that AWS Config uses to record configuration changes.

###### Important

Daily recording cannot be specified for the following resource types:

- `AWS::Config::ResourceCompliance`

- `AWS::Config::ConformancePackCompliance`

- `AWS::Config::ConfigurationRecorder`

For the **allSupported** ( `ALL_SUPPORTED_RESOURCE_TYPES`) recording strategy, these resource types will be set to Continuous recording.

_Required_: Yes

_Type_: String

_Allowed values_: `CONTINUOUS | DAILY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordingModeOverrides`

An array of `recordingModeOverride` objects for you to specify your overrides for the recording mode.
The `recordingModeOverride` object in the `recordingModeOverrides` array consists of three fields: a `description`, the new `recordingFrequency`, and an array of `resourceTypes` to override.

_Required_: No

_Type_: Array of [RecordingModeOverride](aws-properties-config-configurationrecorder-recordingmodeoverride.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordingGroup

RecordingModeOverride

All content copied from https://docs.aws.amazon.com/.
