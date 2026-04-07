This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::Activity

An activity is a task that you write in any programming language and host on any machine
that has access to AWS Step Functions. Activities must poll Step Functions using the
`GetActivityTask` API action and respond using `SendTask*` API
actions. This function makes Step Functions aware of your activity and returns
an identifier for use in a state machine and when polling from the activity.

For information about creating an activity, see [Creating an\
Activity State Machine](https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-creating-activity-state-machine.html) in the _AWS Step Functions Developer_
_Guide_ and [CreateActivity](https://docs.aws.amazon.com/step-functions/latest/apireference/API_CreateActivity.html)
in the _AWS Step Functions API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::StepFunctions::Activity",
  "Properties" : {
      "EncryptionConfiguration" : EncryptionConfiguration,
      "Name" : String,
      "Tags" : [ TagsEntry, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::StepFunctions::Activity
Properties:
  EncryptionConfiguration:
    EncryptionConfiguration
  Name: String
  Tags:
    - TagsEntry

```

## Properties

`EncryptionConfiguration`

Encryption configuration for the activity.

Activity configuration is immutable, and resource names must be unique. To set customer managed keys for encryption, you must create a **new Activity**. If you attempt to change the configuration in your CFN template for an existing activity, you will receive an `ActivityAlreadyExists` exception.

To update your activity to include customer managed keys, set a new activity name within your CloudFormation template.

_Required_: No

_Type_: [EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-activity-encryptionconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the activity.

A name must _not_ contain:

- white space

- brackets `< > { } [ ]`

- wildcard characters `? *`

- special characters ``" # % \ ^ | ~ ` $ & , ; : /``

- control characters ( `U+0000-001F`, `U+007F-009F`, `U+FFFE-FFFF`)

- surrogates ( `U+D800-DFFF`)

- invalid characters ( ` U+10FFFF`)

To enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and \_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The list of tags to add to a resource.

Tags may only contain Unicode letters, digits, white space, or these symbols: `_ . : / = + - @`.

_Required_: No

_Type_: Array of [TagsEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-activity-tagsentry.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic
function, `Ref` returns the ARN of the created activity. For example:

`{ "Ref": "MyActivity" }`

Returns a value similar to the following:

`arn:aws:states:us-east-1:111122223333:activity:myActivity`

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

`Arn`

Returns the ARN of the resource.

`Name`

Returns the name of the activity. For example:

`{ "Fn::GetAtt": ["MyActivity", "Name"] }`

Returns a value similar to the following:

`myActivity`

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

## Examples

The following examples create Step Functions activities.

Note that Activity configuration is immutable, and resource names must be unique. To set customer managed keys for encryption, you must create a **new Activity**. If you attempt to change the configuration in your CFN template for an existing activity, you will receive an `ActivityAlreadyExists` exception.

To update your activity to include customer managed keys, set a new activity name within your CFN template. The following shows an example that creates a new activity with a customer managed key configuration:

- [Create an activity](#aws-resource-stepfunctions-activity--examples--Create_an_activity)

- [Create an activity with a customer managed AWS KMS key](#aws-resource-stepfunctions-activity--examples--Create_an_activity_with_a_customer_managed_key)

### Create an activity

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: An example template for a new Activity.
Resources:
    Activity:
    Type: AWS::StepFunctions::Activity
    Properties:
        Name: myActivity
        EncryptionConfiguration:
            Type: AWS_OWNED_KEY
        Tags:
        -
          Key: "keyname1"
          Value: "value1"
        -
          Key: "keyname2"
          Value: "value2"

```

#### JSON

```json

{
"AWSTemplateFormatVersion" : "2010-09-09",
"Description" : "An example template for a Step Functions activity.",
"Resources" : {
    "MyActivity" : {
        "Type" : "AWS::StepFunctions::Activity",
        "Properties" : {
            "Name" : "myActivity",
            "EncryptionConfiguration" : {
                "Type": "AWS_OWNED_KEY",
            }
            "Tags": [
                    {
                        "Key": "keyname1",
                        "Value": "value1"
                    },
                    {
                        "Key": "keyname2",
                        "Value": "value2"
                    }
                ]
            }
        }
    }
}
```

### Create an activity with a customer managed AWS KMS key

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: An example template for an Activity with customer managed key encryption.
Resources:
    Activity:
    Type: AWS::StepFunctions::Activity
    Properties:
        Name: ActivityWithKmsEncryption
        EncryptionConfiguration:
            Type: CUSTOMER_MANAGED_KMS_KEY
            KmsKeyId: !Ref MyKmsKey
            KmsDataKeyReusePeriodSeconds: 100
        Tags:
          -
            Key: "keyname1"
            Value: "value1"
          -
            Key: "keyname2"
            Value: "value2"

    MyKmsKey:
        Type: AWS::KMS::Key
        Properties:
            Description: Symmetric KMS key used for encryption/decryption
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Step Functions

EncryptionConfiguration
