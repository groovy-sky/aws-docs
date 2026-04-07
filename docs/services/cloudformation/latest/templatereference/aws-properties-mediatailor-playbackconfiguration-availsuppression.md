This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration AvailSuppression

The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FillPolicy" : String,
  "Mode" : String,
  "Value" : String
}

```

### YAML

```yaml

  FillPolicy: String
  Mode: String
  Value: String

```

## Properties

`FillPolicy`

Defines the policy to apply to the avail suppression mode. `BEHIND_LIVE_EDGE` will always use the full avail suppression policy. `AFTER_LIVE_EDGE` mode can be used to invoke partial ad break fills when a session starts mid-break.

_Required_: No

_Type_: String

_Allowed values_: `PARTIAL_AVAIL | FULL_AVAIL_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Sets the ad suppression mode. By default, ad suppression is off and all ad breaks are filled with ads or slate. When Mode is set to `BEHIND_LIVE_EDGE`, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window. When Mode is set to `AFTER_LIVE_EDGE`, ad suppression is active and MediaTailor won't fill ad breaks that are within the live edge plus the avail suppression value.

_Required_: No

_Type_: String

_Allowed values_: `OFF | BEHIND_LIVE_EDGE | AFTER_LIVE_EDGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A live edge offset time in HH:MM:SS. MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AdsInteractionLog

Bumper
