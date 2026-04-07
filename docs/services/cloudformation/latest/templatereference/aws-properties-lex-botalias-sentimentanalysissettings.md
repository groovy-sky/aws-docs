This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotAlias SentimentAnalysisSettings

Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of
user utterances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DetectSentiment" : Boolean
}

```

### YAML

```yaml

  DetectSentiment: Boolean

```

## Properties

`DetectSentiment`

Sets whether Amazon Lex uses Amazon Comprehend to detect the sentiment of user
utterances.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3BucketLogDestination

Tag
