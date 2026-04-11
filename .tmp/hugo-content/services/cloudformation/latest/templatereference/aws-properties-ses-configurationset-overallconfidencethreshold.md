This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet OverallConfidenceThreshold

Defines the validation threshold settings. This object determines the minimum score required for SES to allow an email to be sent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfidenceVerdictThreshold" : String
}

```

### YAML

```yaml

  ConfidenceVerdictThreshold: String

```

## Properties

`ConfidenceVerdictThreshold`

The validation threshold that determines the minimum score required for SES to allow an email to be sent.

Valid Values:

- `HIGH` – Allows emails to be sent only to addresses with high delivery likelihood. This provides maximum protection for your sender reputation but may suppress some legitimate addresses with medium delivery confidence.

- `MEDIUM` – Allows emails to be sent to addresses with medium or high delivery likelihood. This balances reputation protection with delivery reach by allowing addresses with medium and high delivery confidence. This suppresses delivery to email addresses with low delivery confidence.

- `MANAGED` – Amazon SES automatically manages the threshold to suppress invalid addresses. This option allows Amazon SES to optimize the validation threshold based on your sending patterns and reputation.

_Required_: Yes

_Type_: String

_Pattern_: `MEDIUM|HIGH|MANAGED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GuardianOptions

ReputationOptions

All content copied from https://docs.aws.amazon.com/.
