This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioTrack

Information about one audio track to extract. You can select multiple tracks.

The parent of this entity is AudioTrackSelection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Track" : Integer
}

```

### YAML

```yaml

  Track: Integer

```

## Properties

`Track`

1-based integer value that maps to a specific audio track

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioSilenceFailoverSettings

AudioTrackSelection

All content copied from https://docs.aws.amazon.com/.
