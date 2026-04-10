This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolClient AnalyticsConfiguration

The settings for Amazon Pinpoint analytics configuration. With an analytics
configuration, your application can collect user-activity metrics for user notifications
with a Amazon Pinpoint campaign.

Amazon Pinpoint isn't available in all AWS Regions. For a list of
available Regions, see [Amazon Cognito and Amazon Pinpoint Region availability](../../../cognito/latest/developerguide/cognito-user-pools-pinpoint-integration.md#cognito-user-pools-find-region-mappings).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationArn" : String,
  "ApplicationId" : String,
  "ExternalId" : String,
  "RoleArn" : String,
  "UserDataShared" : Boolean
}

```

### YAML

```yaml

  ApplicationArn: String
  ApplicationId: String
  ExternalId: String
  RoleArn: String
  UserDataShared: Boolean

```

## Properties

`ApplicationArn`

The Amazon Resource Name (ARN) of an Amazon Pinpoint project that you want to connect to
your user pool app client. Amazon Cognito publishes events to the Amazon Pinpoint project that
`ApplicationArn` declares. You can also configure your application to
pass an endpoint ID in the `AnalyticsMetadata` parameter of sign-in
operations. The endpoint ID is information about the destination for push
notifications

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationId`

Your Amazon Pinpoint project ID.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-fA-F]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The [external ID](../../../iam/latest/userguide/id-roles-create-for-user-externalid.md) of the role that Amazon Cognito assumes to send
analytics data to Amazon Pinpoint.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of an AWS Identity and Access Management role that has the permissions required for Amazon Cognito to publish
events to Amazon Pinpoint analytics.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserDataShared`

If `UserDataShared` is `true`, Amazon Cognito includes user data in the
events that it publishes to Amazon Pinpoint analytics.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::UserPoolClient

RefreshTokenRotation

All content copied from https://docs.aws.amazon.com/.
