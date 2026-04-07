This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Shield::ProactiveEngagement

Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.

To enable proactive engagement, you must be subscribed to the [Business Support plan](https://aws.amazon.com/premiumsupport/business-support) or the [Enterprise Support plan](https://aws.amazon.com/premiumsupport/enterprise-support).

**Configure `AWS::Shield::ProactiveEngagement` for one account**

To configure this resource through CloudFormation, you must be subscribed to AWS Shield Advanced. You can subscribe
through the [Shield Advanced console](https://console.aws.amazon.com/wafv2/shieldv2) and through
the APIs. For more information, see
[Subscribe to AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/enable-ddos-prem.html).

See example templates for Shield Advanced in CloudFormation at [aws-samples/aws-shield-advanced-examples](https://github.com/aws-samples/aws-shield-advanced-examples).

**Configure Shield Advanced using AWS CloudFormation and AWS Firewall Manager**

You might be able to use Firewall Manager with AWS CloudFormation to configure Shield Advanced across multiple accounts and protected resources. To do this, your accounts must be part of an organization in AWS Organizations. You can use Firewall Manager to configure Shield Advanced protections for any resource types except for Amazon Route 53 or AWS Global Accelerator.

For an example of this, see the one-click configuration guidance published by the AWS technical community at
[One-click deployment of Shield Advanced](https://youtu.be/LCA3FwMk_QE).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Shield::ProactiveEngagement",
  "Properties" : {
      "EmergencyContactList" : [ EmergencyContact, ... ],
      "ProactiveEngagementStatus" : String
    }
}

```

### YAML

```yaml

Type: AWS::Shield::ProactiveEngagement
Properties:
  EmergencyContactList:
    - EmergencyContact
  ProactiveEngagementStatus: String

```

## Properties

`EmergencyContactList`

The list of email addresses and phone numbers that the Shield Response Team (SRT) can use to
contact you for escalations to the SRT and to initiate proactive customer support, plus any relevant notes.

To enable proactive engagement, the contact list must include at least one phone number.

If you provide more than one contact, in the notes, indicate the circumstances under which each contact should be used. Include primary and secondary contact designations, and provide the hours of availability and time zones for each contact.

Example contact notes:

- This is a hotline that's staffed 24x7x365. Please work with the responding analyst and they will get the appropriate person on the call.

- Please contact the secondary phone number if the hotline doesn't respond within 5 minutes.

_Required_: Yes

_Type_: Array of [EmergencyContact](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-shield-proactiveengagement-emergencycontact.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProactiveEngagementStatus`

Specifies whether proactive engagement is enabled or disabled.

Valid values:

`ENABLED` \- The Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.

`DISABLED` \- The SRT will not proactively notify contacts about escalations or to initiate proactive customer support.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the account that submitted the template.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The ID of the account that submitted the template.

## Examples

### Enable proactive engagement and define contacts

The following shows an example proactive engagement configuration with proactive engagement enabled and with two emergency contacts.

#### YAML

```yaml

Resources:
  TestProactiveEngagement:
    DeletionPolicy: Delete
    Type: AWS::Shield::ProactiveEngagement
    Properties:
      ProactiveEngagementStatus: ENABLED
      EmergencyContactList:
        - EmailAddress: !Sub 'dev-on-duty@example.com'
          ContactNotes: !Sub 'Dev On Duty'
          PhoneNumber: '+10000000001'
        - EmailAddress: !Sub 'security@example.com'
          ContactNotes: !Sub 'Security Team'
          PhoneNumber: '+10000000002'
```

#### JSON

```json

{
    "Resources": {
        "TestProactiveEngagement": {
            "DeletionPolicy": "Delete",
            "Type": "AWS::Shield::ProactiveEngagement",
            "Properties": {
                "ProactiveEngagementStatus": "ENABLED",
                "EmergencyContactList": [
                    {
                        "EmailAddress": {
                            "Fn::Sub": "dev-on-duty@example.com"
                        },
                        "ContactNotes": {
                            "Fn::Sub": "Dev On Duty"
                        },
                        "PhoneNumber": "+10000000001"
                    },
                    {
                        "EmailAddress": {
                            "Fn::Sub": "security@example.com"
                        },
                        "ContactNotes": {
                            "Fn::Sub": "Security Team"
                        },
                        "PhoneNumber": "+10000000002"
                    }
                ]
            }
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Shield::DRTAccess

EmergencyContact
