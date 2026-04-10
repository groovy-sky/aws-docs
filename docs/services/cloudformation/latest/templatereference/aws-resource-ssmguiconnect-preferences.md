This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMGuiConnect::Preferences

Specify new or changed connection recording preferences for
your AWS Systems Manager GUI Connect connections.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMGuiConnect::Preferences",
  "Properties" : {
      "ConnectionRecordingPreferences" : ConnectionRecordingPreferences
    }
}

```

### YAML

```yaml

Type: AWS::SSMGuiConnect::Preferences
Properties:
  ConnectionRecordingPreferences:
    ConnectionRecordingPreferences

```

## Properties

`ConnectionRecordingPreferences`

The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region. This includes details such as which S3 bucket recordings are stored in.

_Required_: No

_Type_: [ConnectionRecordingPreferences](aws-properties-ssmguiconnect-preferences-connectionrecordingpreferences.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The primary identifier for the AWS CloudFormation resource.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Systems Manager GUI Connect

ConnectionRecordingPreferences

All content copied from https://docs.aws.amazon.com/.
