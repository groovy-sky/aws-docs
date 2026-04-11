This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration AdConditioningConfiguration

The setting that indicates what conditioning MediaTailor will perform on ads that the ad decision server (ADS) returns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StreamingMediaFileConditioning" : String
}

```

### YAML

```yaml

  StreamingMediaFileConditioning: String

```

## Properties

`StreamingMediaFileConditioning`

For ads that have media files with streaming delivery and supported file extensions, indicates what transcoding action MediaTailor takes when it first receives these ads from the ADS.
`TRANSCODE` indicates that MediaTailor must transcode the ads.
`NONE` indicates that you have already transcoded the ads outside of MediaTailor and don't need them transcoded as part of the ad insertion workflow.
For more information about ad conditioning see [Using preconditioned ads](../../../mediatailor/latest/ug/precondition-ads.md) in the AWS Elemental MediaTailor user guide.

_Required_: Yes

_Type_: String

_Allowed values_: `TRANSCODE | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaTailor::PlaybackConfiguration

AdDecisionServerConfiguration

All content copied from https://docs.aws.amazon.com/.
