This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotVersion

###### Note

Amazon Lex V2 is the only supported version in CloudFormation.

Specifies a new version of the bot based on the `DRAFT`
version. If the `DRAFT` version of this resource hasn't
changed since you created the last version, Amazon Lex doesn't
create a new version, it returns the last created version.

When you specify the first version of a bot, Amazon Lex sets
the version to 1. Subsequent versions increment by 1.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lex::BotVersion",
  "Properties" : {
      "BotId" : String,
      "BotVersionLocaleSpecification" : [ BotVersionLocaleSpecification, ... ],
      "Description" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lex::BotVersion
Properties:
  BotId: String
  BotVersionLocaleSpecification:
    - BotVersionLocaleSpecification
  Description: String

```

## Properties

`BotId`

The unique identifier of the bot.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z]+$`

_Minimum_: `10`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BotVersionLocaleSpecification`

Specifies the locales that Amazon Lex adds to this version.
You can choose the Draft version or any other previously published
version for each locale. When you specify a source version, the locale
data is copied from the source version to the new version.

_Required_: Yes

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lex-botversion-botversionlocalespecification.html) of [BotVersionLocaleSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lex-botversion-botversionlocalespecification.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the version.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BotVersion`

The version of the bot.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TextLogSetting

BotVersionLocaleDetails
