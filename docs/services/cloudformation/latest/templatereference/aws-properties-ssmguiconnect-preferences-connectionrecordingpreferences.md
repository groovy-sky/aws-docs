This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMGuiConnect::Preferences ConnectionRecordingPreferences

The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region. This includes details such as which S3 bucket recordings are stored in.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KMSKeyArn" : String,
  "RecordingDestinations" : RecordingDestinations
}

```

### YAML

```yaml

  KMSKeyArn: String
  RecordingDestinations:
    RecordingDestinations

```

## Properties

`KMSKeyArn`

The ARN of a AWS KMS key that is used to encrypt data while it is being processed by the service. This key must exist in the same AWS Region as the node you start an RDP connection to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordingDestinations`

Determines where recordings of RDP connections are stored.

_Required_: Yes

_Type_: [RecordingDestinations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmguiconnect-preferences-recordingdestinations.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SSMGuiConnect::Preferences

RecordingDestinations
