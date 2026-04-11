This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FrameCaptureS3Settings

Sets up Amazon S3 as the destination for this Frame Capture output.

The parent of this entity is FrameCaptureCdnSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CannedAcl" : String
}

```

### YAML

```yaml

  CannedAcl: String

```

## Properties

`CannedAcl`

Specify the canned ACL to apply to each S3 request. Defaults to none.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameCaptureOutputSettings

FrameCaptureSettings

All content copied from https://docs.aws.amazon.com/.
