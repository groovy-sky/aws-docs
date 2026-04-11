This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::EmailChannel

A _channel_ is a type of platform that you can deliver messages to.
You can use the email channel to send email to users. Before you can use Amazon Pinpoint
to send email, you must enable the email channel for an Amazon Pinpoint
application.

The EmailChannel resource represents the status, identity, and other settings of the
email channel for an application

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::EmailChannel",
  "Properties" : {
      "ApplicationId" : String,
      "ConfigurationSet" : String,
      "Enabled" : Boolean,
      "FromAddress" : String,
      "Identity" : String,
      "OrchestrationSendingRoleArn" : String,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::EmailChannel
Properties:
  ApplicationId: String
  ConfigurationSet: String
  Enabled: Boolean
  FromAddress: String
  Identity: String
  OrchestrationSendingRoleArn: String
  RoleArn: String

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that you're specifying the
email channel for.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationSet`

The [Amazon SES\
configuration set](../../../../reference/ses/latest/apireference/api-configurationset.md) that you want to apply to messages that you send
through the channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether to enable the email channel for the application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FromAddress`

The verified email address that you want to send email from when you send
email through the channel.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Identity`

The Amazon Resource Name (ARN) of the identity, verified with Amazon Simple
Email Service (Amazon SES), that you want to use when you send email through the
channel.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrchestrationSendingRoleArn`

The ARN of an IAM role for Amazon Pinpoint to use to send email from your campaigns or journeys through Amazon SES.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the AWS Identity and Access Management (IAM) role that you want Amazon Pinpoint to
use when it submits email-related event data for the channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier ( `ApplicationId`) for
the Amazon Pinpoint application that the channel is associated with.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Id`

(Deprecated) An identifier for the email channel. This property is retained
only for backward compatibility.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WriteTreatmentResource

AWS::Pinpoint::EmailTemplate

All content copied from https://docs.aws.amazon.com/.
