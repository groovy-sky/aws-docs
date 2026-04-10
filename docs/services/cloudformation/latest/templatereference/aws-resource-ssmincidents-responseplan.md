This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan

The `AWS::SSMIncidents::ResponsePlan` resource specifies the details of the
response plan that are used when creating an incident.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMIncidents::ResponsePlan",
  "Properties" : {
      "Actions" : [ Action, ... ],
      "ChatChannel" : ChatChannel,
      "DisplayName" : String,
      "Engagements" : [ String, ... ],
      "IncidentTemplate" : IncidentTemplate,
      "Integrations" : [ Integration, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSMIncidents::ResponsePlan
Properties:
  Actions:
    - Action
  ChatChannel:
    ChatChannel
  DisplayName: String
  Engagements:
    - String
  IncidentTemplate:
    IncidentTemplate
  Integrations:
    - Integration
  Name: String
  Tags:
    - Tag

```

## Properties

`Actions`

The actions that the response plan starts at the beginning of an incident.

_Required_: No

_Type_: Array of [Action](aws-properties-ssmincidents-responseplan-action.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChatChannel`

The Amazon Q Developer in chat applications chat channel used for collaboration during an
incident.

_Required_: No

_Type_: [ChatChannel](aws-properties-ssmincidents-responseplan-chatchannel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The human readable name of the response plan.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engagements`

The Amazon Resource Name (ARN) for the contacts and escalation plans that the response
plan engages during an incident.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncidentTemplate`

Details used to create an incident when using this response plan.

_Required_: Yes

_Type_: [IncidentTemplate](aws-properties-ssmincidents-responseplan-incidenttemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Integrations`

Information about third-party services integrated into the response plan.

_Required_: No

_Type_: Array of [Integration](aws-properties-ssmincidents-responseplan-integration.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the response plan.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-ssmincidents-responseplan-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a response plan

The following example creates a response plan.

###### Note

This example also demonstrates creating a replication set. We recommend
creating a replication set and response plan using a single template. This
template also incorporates resources from [AWS Systems Manager Incident Manager Contacts](../userguide/aws-ssmcontacts.md).

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Description": "Sample AWS CloudFormation template to create a response plan and replication set (JSON).",
   "Parameters": {
      "ContactEmail": {
         "Type": "String",
         "Description": "The email address to use to engage the contact channel."
      },
      "ChatbotSnsArn": {
         "Type": "String",
         "Description": "The Amazon Simple Notification Service (SNS) targets that AWS Chatbot uses to provide the chat channel with updates about an incident."
      },
      "NotificationTargetArn": {
         "Type": "String",
         "Description": "The SNS topic used by AWS Chatbot to send notifications to the incidents chat channel.",
         "AllowedPattern": "^arn:aws(-cn|-us-gov)?:[a-z0-9-]*:[a-z0-9-]*:([0-9]{12})?:.+$"
      },
      "SsmRoleArn": {
         "Type": "String",
         "Description": "The Amazon Resource Name (ARN) of the role that the Automation runbook assumes when running commands.",
         "AllowedPattern": "^arn:aws(-cn|-us-gov)?:iam::([0-9]{12})?:role/.+$"
      },
      "SsmAutomationDocument": {
         "Type": "String",
         "Description": "The Systems Manager Automation runbook to use during an incident.",
         "AllowedPattern": "^[a-zA-Z0-9_\\-.:/]{3,128}$"
      }
   },
   "Resources": {
      "MyReplicationSet": {
         "Type": "AWS::SSMIncidents::ReplicationSet",
         "Properties": {
            "Regions": [
               {
                  "RegionName": {
                     "Ref": "AWS::Region"
                  }
               }
            ],
            "Tags": [
               {
                  "Key": "MyReplicationSetTagKey",
                  "Value": "MyReplicationSetTagValue"
               }
            ]
         }
      },
      "MyIncidentsResponsePlan": {
         "Type": "AWS::SSMIncidents::ResponsePlan",
         "Properties": {
            "Name": "MyResponsePlan",
            "ChatChannel": {
               "ChatbotSns": "ChatbotSnsArn"
            },
            "Engagements": "MyEscalationPlanContact",
            "Actions": [
               {
                  "SsmAutomation": {
                     "RoleArn": "SsmRoleArn",
                     "DocumentName": "SsmAutomationDocument",
                     "DocumentVersion": "1",
                     "TargetAccount": "IMPACTED_ACCOUNT",
                     "Parameters": [
                        {
                           "Key": "Key1",
                           "Values": [
                              "Value1"
                           ]
                        }
                     ],
                     "DynamicParameters": [
                        {
                           "Key": "Key2",
                           "Value": {
                              "Variable": "INCIDENT_RECORD_ARN"
                           }
                        },
                        {
                           "Key": "Key3",
                           "Value": {
                              "Variable": "INVOLVED_RESOURCES"
                           }
                        }
                     ]
                  }
               }
            ],
            "IncidentTemplate": {
               "Title": "MyIncident",
               "Impact": 1,
               "Summary": "My incident summary.",
               "DedupeString": "MyDedupeString",
               "NotificationTargets": [
                  {
                     "SnsTopicArn": "NotificationTargetArn"
                  }
               ],
               "IncidentTags": [
                  {
                     "Key": "MyIncidentTagKey",
                     "Value": "MyIncidentTagValue"
                  }
               ]
            },
            "Tags": [
               {
                  "Key": "MyResponsePlanTagKey",
                  "Value": "MyResponsePlanTagValue"
               }
            ]
         },
         "DependsOn": "MyReplicationSet"
      },
      "MyContactChannelEmail": {
         "Type": "AWS::SSMContacts::ContactChannel",
         "Properties": {
            "ContactId": "MySSMContact",
            "ChannelName": "MyEmailChannel",
            "ChannelType": "EMAIL",
            "ChannelAddress": "contact-email-channel@example.com"
         },
         "DependsOn": "MyReplicationSet"
      },
      "MySSMContact": {
         "Type": "AWS::SSMContacts::Contact",
         "Properties": {
            "Alias": "my-ssm-contact",
            "DisplayName": "MySSMContact",
            "Type": "PERSONAL"
         },
         "DependsOn": "MyReplicationSet"
      },
      "MySSMContactEngagementPlan": {
         "Type": "AWS::SSMContacts::Plan",
         "Properties": {
            "ContactId": "MySSMContact",
            "Stages": [
               {
                  "DurationInMinutes": 10,
                  "Targets": [
                     {
                        "ChannelTargetInfo": {
                           "ChannelId": "MyContactChannelEmail",
                           "RetryIntervalInMinutes": 3
                        }
                     }
                  ]
               }
            ]
         },
         "DependsOn": "MyReplicationSet"
      },
      "MyEscalationPlanContact": {
         "Type": "AWS::SSMContacts::Contact",
         "Properties": {
            "Alias": "my-escalation-plan",
            "DisplayName": "MyEscalationPlan",
            "Type": "ESCALATION"
         },
         "DependsOn": "MyReplicationSet"
      },
      "MyEscalationPlan": {
         "Type": "AWS::SSMContacts::Plan",
         "Properties": {
            "ContactId": "MyEscalationPlanContact",
            "Stages": [
               {
                  "DurationInMinutes": 0,
                  "Targets": [
                     {
                        "ContactTargetInfo": {
                           "ContactId": "MySSMContact",
                           "IsEssential": true
                        }
                     }
                  ]
               }
            ]
         },
         "DependsOn": "MyReplicationSet"
      }
   },
   "Outputs": {
      "ReplicationSetArn": {
         "Value": "MyReplicationSet"
      },
      "Region": {
         "Value": "AWS::Region"
      },
      "SSMContactArn": {
         "Value": "MySSMContact"
      },
      "EscalationPlanArn": {
         "Value": "MyEscalationPlanContact"
      },
      "ResponsePlanArn": {
         "Value": "MyIncidentsResponsePlan"
      }
   }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: 2010-09-09
Description: "Sample AWS CloudFormation template to create a response plan and replication set (YAML)."
Parameters:
  ContactEmail:
    Type: String
    Description: The email address to use to engage the contact channel.
  ChatbotSnsArn:
    Type: String
    Description: The Amazon Simple Notification Service (SNS) targets that AWS
      Chatbot uses to provide the chat channel with updates about an incident.
  NotificationTargetArn:
    Type: String
    Description: The SNS topic used by AWS Chatbot to send notifications to the
      incidents chat channel.
    AllowedPattern: ^arn:aws(-cn|-us-gov)?:[a-z0-9-]*:[a-z0-9-]*:([0-9]{12})?:.+$
  SsmRoleArn:
    Type: String
    Description: The Amazon Resource Name (ARN) of the role that the Automation
      runbook assumes when running commands.
    AllowedPattern: ^arn:aws(-cn|-us-gov)?:iam::([0-9]{12})?:role/.+$
  SsmAutomationDocument:
    Type: String
    Description: The Systems Manager Automation runbook to use during an incident.
    AllowedPattern: ^[a-zA-Z0-9_\-.:/]{3,128}$
Resources:
  MyReplicationSet:
    Type: AWS::SSMIncidents::ReplicationSet
    Properties:
      Regions:
        - RegionName:
            Ref: AWS::Region
      Tags:
        - Key: MyReplicationSetTagKey
          Value: MyReplicationSetTagValue
  MyIncidentsResponsePlan:
    Type: AWS::SSMIncidents::ResponsePlan
    Properties:
      Name: MyResponsePlan
      ChatChannel:
        ChatbotSns: ChatbotSnsArn
      Engagements: MyEscalationPlanContact
      Actions:
        - SsmAutomation:
            RoleArn: SsmRoleArn
            DocumentName: SsmAutomationDocument
            DocumentVersion: "1"
            TargetAccount: IMPACTED_ACCOUNT
            Parameters:
              - Key: Key1
                Values:
                  - Value1
            DynamicParameters:
              - Key: Key2
                Value:
                  Variable: INCIDENT_RECORD_ARN
              - Key: Key3
                Value:
                  Variable: INVOLVED_RESOURCES
      IncidentTemplate:
        Title: MyIncident
        Impact: 1
        Summary: "My incident summary."
        DedupeString: MyDedupeString
        NotificationTargets:
          - SnsTopicArn: NotificationTargetArn
        IncidentTags:
          - Key: MyIncidentTagKey
            Value: MyIncidentTagValue
      Tags:
        - Key: MyResponsePlanTagKey
          Value: MyResponsePlanTagValue
    DependsOn: MyReplicationSet
  MyContactChannelEmail:
    Type: AWS::SSMContacts::ContactChannel
    Properties:
      ContactId: MySSMContact
      ChannelName: MyEmailChannel
      ChannelType: EMAIL
      ChannelAddress: contact-email-channel@example.com
    DependsOn: MyReplicationSet
  MySSMContact:
    Type: AWS::SSMContacts::Contact
    Properties:
      Alias: my-ssm-contact
      DisplayName: MySSMContact
      Type: PERSONAL
    DependsOn: MyReplicationSet
  MySSMContactEngagementPlan:
    Type: AWS::SSMContacts::Plan
    Properties:
      ContactId: MySSMContact
      Stages:
        - DurationInMinutes: 10
          Targets:
            - ChannelTargetInfo:
                ChannelId: MyContactChannelEmail
                RetryIntervalInMinutes: 3
    DependsOn: MyReplicationSet
  MyEscalationPlanContact:
    Type: AWS::SSMContacts::Contact
    Properties:
      Alias: my-escalation-plan
      DisplayName: MyEscalationPlan
      Type: ESCALATION
    DependsOn: MyReplicationSet
  MyEscalationPlan:
    Type: AWS::SSMContacts::Plan
    Properties:
      ContactId: MyEscalationPlanContact
      Stages:
        - DurationInMinutes: 0
          Targets:
            - ContactTargetInfo:
                ContactId: MySSMContact
                IsEssential: true
    DependsOn: MyReplicationSet
Outputs:
  ReplicationSetArn:
    Value: MyReplicationSet
  Region:
    Value: AWS::Region
  SSMContactArn:
    Value: MySSMContact
  EscalationPlanArn:
    Value: MyEscalationPlanContact
  ResponsePlanArn:
    Value: MyIncidentsResponsePlan
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Action

All content copied from https://docs.aws.amazon.com/.
