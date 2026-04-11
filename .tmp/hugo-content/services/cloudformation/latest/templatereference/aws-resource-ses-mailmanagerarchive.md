This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerArchive

Creates a new email archive resource for storing and retaining emails.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MailManagerArchive",
  "Properties" : {
      "ArchiveName" : String,
      "KmsKeyArn" : String,
      "Retention" : ArchiveRetention,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::MailManagerArchive
Properties:
  ArchiveName: String
  KmsKeyArn: String
  Retention:
    ArchiveRetention
  Tags:
    - Tag

```

## Properties

`ArchiveName`

A unique name for the new archive.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*[a-zA-Z0-9]$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The Amazon Resource Name (ARN) of the KMS key for encrypting emails in the archive.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov|-eusc):kms:[a-z0-9-]{1,20}:[0-9]{12}:(key|alias)/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Retention`

The period for retaining emails in the archive before automatic deletion.

_Required_: No

_Type_: [ArchiveRetention](aws-properties-ses-mailmanagerarchive-archiveretention.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-mailmanagerarchive-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ArchiveArn`

The Amazon Resource Name (ARN) of the archive.

`ArchiveId`

The unique identifier of the archive.

`ArchiveState`

The current state of the archive:

- `ACTIVE` – The archive is ready and available for use.

- `PENDING_DELETION` – The archive has been marked for deletion
and will be permanently deleted in 30 days. No further modifications can be made
in this state.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ArchiveRetention

All content copied from https://docs.aws.amazon.com/.
