This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate MessageTemplateAttachment

Information about the message template attachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachmentId" : String,
  "AttachmentName" : String,
  "S3PresignedUrl" : String
}

```

### YAML

```yaml

  AttachmentId: String
  AttachmentName: String
  S3PresignedUrl: String

```

## Properties

`AttachmentId`

The identifier of the attachment file.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttachmentName`

The name of the attachment file being uploaded. The name should include the file
extension.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3PresignedUrl`

The S3 Presigned URL for the attachment file. When generating the PreSignedUrl, please
ensure that the expires-in time is set to 30 minutes. The URL can be generated through
the AWS Console or through the AWS CLI. For more
information, see [Sharing objects with\
presigned URLs](../../../s3/latest/userguide/shareobjectpresignedurl.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GroupingConfiguration

MessageTemplateAttributes
