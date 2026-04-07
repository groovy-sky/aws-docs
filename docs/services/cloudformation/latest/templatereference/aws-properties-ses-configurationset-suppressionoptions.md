This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet SuppressionOptions

An object that contains information about the suppression list preferences for your
account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SuppressedReasons" : [ String, ... ],
  "ValidationOptions" : ValidationOptions
}

```

### YAML

```yaml

  SuppressedReasons:
    - String
  ValidationOptions:
    ValidationOptions

```

## Properties

`SuppressedReasons`

A list that contains the reasons that email addresses are automatically added to the
suppression list for your account. This list can contain any or all of the
following:

- `COMPLAINT` – Amazon SES adds an email address to the suppression
list for your account when a message sent to that address results in a
complaint.

- `BOUNCE` – Amazon SES adds an email address to the suppression list
for your account when a message sent to that address results in a hard
bounce.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationOptions`

The configuration settings for email automatic validation.

_Required_: No

_Type_: [ValidationOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-configurationset-validationoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SendingOptions

Tag
