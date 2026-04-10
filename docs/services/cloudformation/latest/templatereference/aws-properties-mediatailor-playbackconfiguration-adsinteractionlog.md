This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration AdsInteractionLog

Settings for customizing what events are included in logs for interactions with the ad decision server (ADS).

For more information about ADS logs, inlcuding descriptions of the event types, see [MediaTailor ADS logs description and event types](../../../mediatailor/latest/ug/ads-log-format.md)
in AWS Elemental MediaTailor User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludeEventTypes" : [ String, ... ],
  "PublishOptInEventTypes" : [ String, ... ]
}

```

### YAML

```yaml

  ExcludeEventTypes:
    - String
  PublishOptInEventTypes:
    - String

```

## Properties

`ExcludeEventTypes`

Indicates that MediaTailor won't emit the selected events in the logs for playback sessions that are initialized with this configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublishOptInEventTypes`

Indicates that MediaTailor emits `RAW_ADS_RESPONSE` logs for playback sessions that are initialized with this configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdMarkerPassthrough

AvailSuppression

All content copied from https://docs.aws.amazon.com/.
