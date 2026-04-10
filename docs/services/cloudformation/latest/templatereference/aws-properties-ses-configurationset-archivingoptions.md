This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet ArchivingOptions

Used to associate a configuration set with a MailManager archive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArchiveArn" : String
}

```

### YAML

```yaml

  ArchiveArn: String

```

## Properties

`ArchiveArn`

The Amazon Resource Name (ARN) of the MailManager archive where the Amazon SES API v2 will archive sent
emails.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::ConfigurationSet

ConditionThreshold

All content copied from https://docs.aws.amazon.com/.
