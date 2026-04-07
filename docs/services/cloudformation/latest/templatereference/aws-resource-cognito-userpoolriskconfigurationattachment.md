This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment

The `AWS::Cognito::UserPoolRiskConfigurationAttachment` resource sets the
risk configuration that is used for Amazon Cognito advanced security features.

You can specify risk configuration for a single client (with a specific
`clientId`) or for all clients (by setting the `clientId` to
`ALL`). If you specify `ALL`, the default configuration is
used for every client that has had no risk configuration set previously. If you specify
risk configuration for a particular client, it no longer falls back to the
`ALL` configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolRiskConfigurationAttachment",
  "Properties" : {
      "AccountTakeoverRiskConfiguration" : AccountTakeoverRiskConfigurationType,
      "ClientId" : String,
      "CompromisedCredentialsRiskConfiguration" : CompromisedCredentialsRiskConfigurationType,
      "RiskExceptionConfiguration" : RiskExceptionConfigurationType,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolRiskConfigurationAttachment
Properties:
  AccountTakeoverRiskConfiguration:
    AccountTakeoverRiskConfigurationType
  ClientId: String
  CompromisedCredentialsRiskConfiguration:
    CompromisedCredentialsRiskConfigurationType
  RiskExceptionConfiguration:
    RiskExceptionConfigurationType
  UserPoolId: String

```

## Properties

`AccountTakeoverRiskConfiguration`

The settings for automated responses and notification templates for adaptive
authentication with threat protection.

_Required_: No

_Type_: [AccountTakeoverRiskConfigurationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfigurationtype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The app client where this configuration is applied. When this parameter isn't present,
the risk configuration applies to all user pool app clients that don't have
client-level settings.

_Required_: Yes

_Type_: String

_Pattern_: `[\w+]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CompromisedCredentialsRiskConfiguration`

Settings for compromised-credentials actions and authentication types with threat
protection in full-function `ENFORCED` mode.

_Required_: No

_Type_: [CompromisedCredentialsRiskConfigurationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RiskExceptionConfiguration`

Exceptions to the risk evaluation configuration, including always-allow and
always-block IP address ranges.

_Required_: No

_Type_: [RiskExceptionConfigurationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpoolriskconfigurationattachment-riskexceptionconfigurationtype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool that has the risk configuration applied.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the physicalResourceId, which is
“UserPoolRiskConfigurationAttachment-UserPoolId-ClientId". For example:

`{ "Ref":
                “UserPoolRiskConfigurationAttachment-us-east-1_FAKEPOOLID-2asc123fakeclientidajjulj6bh”
                }`

For the Amazon Cognito risk configuration attachment
`UserPoolRiskConfigurationAttachment-us-east-1_FAKEPOOLID-2asc123fakeclientidajjulj6bh`,
Ref returns the name of the risk configuration attachment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating a new risk configuration attachment for a user pool

The following example sets risk configurations in the referenced user pool and
client.

#### JSON

```json

{
   "UserPoolRiskConfiguration":{
      "Type":"AWS::Cognito::UserPoolRiskConfigurationAttachment",
      "Properties":{
         "UserPoolId":{
            "Ref":"UserPool"
         },
         "ClientId":{
            "Ref":"Client"
         },
         "AccountTakeoverRiskConfiguration":{
            "Actions":{
               "HighAction":{
                  "EventAction":"MFA_REQUIRED",
                  "Notify":true,

               },
               "MediumAction":{
                  "EventAction":"MFA_IF_CONFIGURED",
                  "Notify":true
               },
               "LowAction":{
                  "EventAction":{
                     "Ref":"EventAction"
                  },
                  "Notify":false
               }
            },
            "NotifyConfiguration":{
               "BlockEmail":{
                  "HtmlBody":"html body",
                  "Subject":"Your account got blocked",
                  "TextBody":"Your account got blocked"
               },
               "MfaEmail":{
                  "HtmlBody":"html body",
                  "Subject":"Your account needs MFA verification",
                  "TextBody":"Your account needs MFA verification"
               },
               "NoActionEmail":{
                  "HtmlBody":{
                     "Ref":"HtmlBody"
                  },
                  "Subject":{
                     "Ref":"Subject"
                  },
                  "TextBody":{
                     "Ref":"TextBody"
                  },

               },
               "From":"your-from-email@amazon.com",
               "SourceArn":{
                  "Ref":"SourceArn"
               },
               "ReplyTo":"your-reply-to@amazon.com"
            }
         },
         "CompromisedCredentialsRiskConfiguration":{
            "Actions":{
               "EventAction":"BLOCK"
            },
            "EventFilter":[
               {
                  "Ref":"EventFilter"
               },

            ]
         },
         "RiskExceptionConfiguration":{
            "BlockedIPRangeList":[
               "198.0.0.1"
            ],
            "SkippedIPRangeList":[
               "198.0.0.1"
            ]
         }
      }
   }
}
```

#### YAML

```yaml

UserPoolRiskConfiguration:
  Type: AWS::Cognito::UserPoolRiskConfigurationAttachment
  Properties:
    UserPoolId: !Ref UserPool
    ClientId: !Ref Client
    AccountTakeoverRiskConfiguration:
      Actions:
        HighAction:
          EventAction: "MFA_REQUIRED"
          Notify: True
        MediumAction:
          EventAction: "MFA_IF_CONFIGURED"
          Notify: True
        LowAction:
          EventAction: !Ref LowEventAction
          Notify: False
      NotifyConfiguration:
        BlockEmail:
          HtmlBody: "html body"
          Subject: "Your account got blocked"
          TextBody: "Your account got blocked"
        MfaEmail:
          HtmlBody: "html body"
          Subject: "Your account needs MFA verification"
          TextBody: "Your account needs MFA verification"
        NoActionEmail:
          HtmlBody: !Ref HtmlBody
          Subject: !Ref Subject
          TextBody: !Ref TextBody
        From: "your-from-email@amazon.com"
        SourceArn: !Ref SourceArn
        ReplyTo: "your-reply-to@amazon.com"
    CompromisedCredentialsRiskConfiguration:
      Actions:
        EventAction: "BLOCK"
        EventFilter: - !Ref EventFilter
    RiskExceptionConfiguration:
        BlockedIPRangeList:
          - "198.0.0.1"
        SkippedIPRangeList:
          - "198.0.0.1"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceServerScopeType

AccountTakeoverActionsType
