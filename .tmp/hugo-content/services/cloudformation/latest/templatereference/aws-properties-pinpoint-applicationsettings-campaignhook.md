This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::ApplicationSettings CampaignHook

Specifies the Lambda function to use by default as a code hook for campaigns in the
application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LambdaFunctionName" : String,
  "Mode" : String,
  "WebUrl" : String
}

```

### YAML

```yaml

  LambdaFunctionName: String
  Mode: String
  WebUrl: String

```

## Properties

`LambdaFunctionName`

The name or Amazon Resource Name (ARN) of the Lambda function that Amazon Pinpoint
invokes to send messages for campaigns in the application.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The mode that Amazon Pinpoint uses to invoke the Lambda function. Possible values
are:

- `FILTER` \- Invoke the function to customize the segment that's used
by a campaign.

- `DELIVERY` \- (Deprecated) Previously, invoked the function to send a
campaign through a custom channel. This functionality is not supported anymore.
To send a campaign through a custom channel, use the
`CustomDeliveryConfiguration` and
`CampaignCustomMessage` objects of the campaign.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebUrl`

The web URL that Amazon Pinpoint calls to invoke the Lambda function over
HTTPS.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Pinpoint::ApplicationSettings

Limits

All content copied from https://docs.aws.amazon.com/.
