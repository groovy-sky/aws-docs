This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration ManifestServiceInteractionLog

Settings for customizing what events are included in logs for interactions with the origin server.

For more information about manifest service logs, including descriptions of the event types, see [MediaTailor manifest logs description and event types](../../../mediatailor/latest/ug/log-types.md)
in AWS Elemental MediaTailor User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludeEventTypes" : [ String, ... ]
}

```

### YAML

```yaml

  ExcludeEventTypes:
    - String

```

## Properties

`ExcludeEventTypes`

Indicates that MediaTailor won't emit the selected events in the logs for playback sessions that are initialized with this configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManifestProcessingRules

Tag

All content copied from https://docs.aws.amazon.com/.
