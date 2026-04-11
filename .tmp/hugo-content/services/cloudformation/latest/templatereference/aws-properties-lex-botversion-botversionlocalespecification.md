This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotVersion BotVersionLocaleSpecification

Specifies the locale that Amazon Lex adds to this version.
You can choose the Draft version or any other previously published
version for each locale. When you specify a source version, the locale
data is copied from the source version to the new version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BotVersionLocaleDetails" : BotVersionLocaleDetails,
  "LocaleId" : String
}

```

### YAML

```yaml

  BotVersionLocaleDetails:
    BotVersionLocaleDetails
  LocaleId: String

```

## Properties

`BotVersionLocaleDetails`

The version of a bot used for a bot locale.

_Required_: Yes

_Type_: [BotVersionLocaleDetails](aws-properties-lex-botversion-botversionlocaledetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocaleId`

The identifier of the locale to add to the version.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BotVersionLocaleDetails

AWS::Lex::ResourcePolicy

All content copied from https://docs.aws.amazon.com/.
