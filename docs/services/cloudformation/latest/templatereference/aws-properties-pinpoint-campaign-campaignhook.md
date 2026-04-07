This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign CampaignHook

Specifies settings for invoking an Lambda function that customizes a segment for a
campaign.

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
invokes to customize a segment for a campaign.

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CampaignEventFilter

CampaignInAppMessage
