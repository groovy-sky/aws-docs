This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::Subscriber AwsLogSource

Adds a natively supported AWS service as an Amazon Security Lake source. Enables source types for member accounts in required AWS Regions, based on the parameters you specify. You can choose any source type in any Region for either accounts that are part of a trusted organization or standalone accounts. Once you add an AWS service as a source, Security Lake starts collecting logs and events from it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceName" : String,
  "SourceVersion" : String
}

```

### YAML

```yaml

  SourceName: String
  SourceVersion: String

```

## Properties

`SourceName`

Source name of the natively supported AWS service that is supported as an Amazon Security Lake source. For the list of sources supported by Amazon Security Lake see [Collecting data from AWS services](../../../security-lake/latest/userguide/internal-sources.md) in the Amazon Security Lake User Guide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceVersion`

Source version of the natively supported AWS service that is supported as an Amazon Security Lake source. For more details about source versions supported by Amazon Security Lake see [OCSF source identification](../../../security-lake/latest/userguide/open-cybersecurity-schema-framework.md#ocsf-source-identification) in the Amazon Security Lake User Guide.

_Required_: No

_Type_: String

_Pattern_: `^(latest|[0-9]\.[0-9])$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityLake::Subscriber

CustomLogSource

All content copied from https://docs.aws.amazon.com/.
