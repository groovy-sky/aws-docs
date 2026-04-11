This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::EventBusPolicy

Running `PutPermission` permits the specified AWS account or AWS organization
to put events to the specified _event bus_. Amazon EventBridge rules in your account are triggered by these events arriving to an event bus in your
account.

For another account to send events to your account, that external account must have an
EventBridge rule with your account's event bus as a target.

To enable multiple AWS accounts to put events to your event bus, run
`PutPermission` once for each of these accounts. Or, if all the accounts are
members of the same AWS organization, you can run `PutPermission`
once specifying `Principal` as "\*" and specifying the AWS
organization ID in `Condition`, to grant permissions to all accounts in that
organization.

If you grant permissions using an organization, then accounts in that organization must
specify a `RoleArn` with proper permissions when they use `PutTarget` to
add your account's event bus as a target. For more information, see [Sending and\
Receiving Events Between AWS Accounts](../../../eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.md) in the _Amazon EventBridge User Guide_.

The permission policy on the event bus cannot exceed 10 KB in size.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Events::EventBusPolicy",
  "Properties" : {
      "EventBusName" : String,
      "Statement" : Json,
      "StatementId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Events::EventBusPolicy
Properties:
  EventBusName: String
  Statement: Json
  StatementId: String

```

## Properties

`EventBusName`

The name of the event bus associated with the rule. If you omit this, the default event
bus is used.

_Required_: No

_Type_: String

_Pattern_: `[\.\-_A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Statement`

A JSON string that describes the permission policy statement. You can include a
`Policy` parameter in the request instead of using the `StatementId`,
`Action`, `Principal`, or `Condition` parameters.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatementId`

An identifier string for the external account that you are granting permissions to. If you
later want to revoke the permission for this external account, specify this
`StatementId` when you run [RemovePermission](../../../../reference/eventbridge/latest/apireference/api-removepermission.md).

###### Note

Each `StatementId` must be unique.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-_]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the event bus policy ID, such as
`EventBusPolicy-1aBCdeFGh2J3`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Grant Permission to One Account](#aws-resource-events-eventbuspolicy--examples--Grant_Permission_to_One_Account)

- [Grant Permission to an Organization](#aws-resource-events-eventbuspolicy--examples--Grant_Permission_to_an_Organization)

- [Deny policy using multiple principals and actions](#aws-resource-events-eventbuspolicy--examples--Deny_policy_using_multiple_principals_and_actions)

- [Grant Permission to an Organization using a custom event bus](#aws-resource-events-eventbuspolicy--examples--Grant_Permission_to_an_Organization_using_a_custom_event_bus)

### Grant Permission to One Account

The following example grants permission to one AWS account with an account ID of
`111122223333`.

#### JSON

```json

"SampleEventBusPolicy": {
    "Type": "AWS::Events::EventBusPolicy",
    "Properties": {
        "StatementId": "MyStatement",
        "Statement": {
            "Effect": "Allow",
            "Principal" : {"AWS" : "arn:aws:iam::111122223333:root"},
            "Action": "events:PutEvents",
            "Resource": "arn:aws:events:us-east-1:111122223333:event-bus/default"
        }
    }
}
```

#### YAML

```yaml

SampleEventBusPolicy:
    Type: AWS::Events::EventBusPolicy
    Properties:
        StatementId: "MyStatement"
        Statement:
            Effect: "Allow"
            Principal:
                AWS: "arn:aws:iam::111122223333:root"
            Action: "events:PutEvents"
            Resource: "arn:aws:events:us-east-1:111122223333:event-bus/default"
```

### Grant Permission to an Organization

The following example grants permission to all AWS accounts in the organization with
an organization ID of `o-1234567890`.

#### JSON

```json

"SampleEventBusPolicy": {
    "Type": "AWS::Events::EventBusPolicy",
    "Properties": {
        "StatementId": "MyStatement",
        "Statement": {
            "Effect": "Allow",
            "Principal" : "*",
            "Action": "events:PutEvents",
            "Resource": "arn:aws:events:us-east-1:111122223333:event-bus/default",
            "Condition": {
                "StringEquals": {"aws:PrincipalOrgID": "o-1234567890"}
             }
        }
    }
}
```

#### YAML

```yaml

SampleEventBusPolicy:
    Type: AWS::Events::EventBusPolicy
    Properties:
        StatementId: "MyStatement"
        Statement:
            Effect: "Allow"
            Principal: "*"
            Action: "events:PutEvents"
            Resource: "arn:aws:events:us-east-1:111122223333:event-bus/default"
            Condition:
                StringEquals:
                    "aws:PrincipalOrgID": "o-1234567890"
```

### Deny policy using multiple principals and actions

The following example demonstrates a deny policy statement using multiple principals
and actions.

#### JSON

```json

"SampleDenyEventBusPolicy": {
    "Type": "AWS::Events::EventBusPolicy",
    "Properties": {
        "StatementId": "MyDenyStatement",
        "Statement": {
            "Effect": "Deny",
            "Principal" :
                {"AWS" : ["arn:aws:iam::111122223333:user/alice", "arn:aws:iam::111122223333:user/bob"]},
            "Action": [
                "events:PutEvents",
                "events:PutRule"
            ],
            "Resource": "arn:aws:events:us-east-1:111122223333:event-bus/default"
        }
    }
}
```

#### YAML

```yaml

SampleDenyEventBusPolicy:
    Type: AWS::Events::EventBusPolicy
    Properties:
        StatementId: "MyDenyStatement"
        Statement:
            Effect: "Deny"
            Principal:
                AWS:
                    - "arn:aws:iam::111122223333:user/alice"
                    - "arn:aws:iam::111122223333:user/bob"
            Action:
                - "events:PutEvents"
                - "events:PutRule"
            Resource: "arn:aws:events:us-east-1:111122223333:event-bus/default"
```

### Grant Permission to an Organization using a custom event bus

The following example grants permission to all AWS accounts in the organization with an organization ID of `o-1234567890` using a custom event bus.

#### JSON

```json

"SampleCustomEventBus": {
    "Type": "AWS::Events::EventBus",
    "Properties": {
        "Name": "MyCustomEventBus"
     }
},
"SampleCustomEventBusPolicy": {
    "Type": "AWS::Events::EventBusPolicy",
    "Properties": {
        "EventBusName": {
            "Ref": "SampleCustomEventBus"
        },
        "StatementId": "MyCustomEventBusStatement",
        "Statement": {
            "Effect": "Allow",
            "Principal" : "*",
            "Action": "events:PutEvents",
            "Resource": {
                "Fn::GetAtt": [
                    "SampleCustomEventBus",
                    "Arn"
                ]
            },
            "Condition": {
                "StringEquals": {"aws:PrincipalOrgID": "o-1234567890"}
            }
        }
    }
}
```

#### YAML

```yaml

SampleCustomEventBus:
    Type: AWS::Events::EventBus
    Properties:
        Name: "MyCustomEventBus"

SampleCustomEventBusPolicy:
    Type: AWS::Events::EventBusPolicy
    Properties:
        EventBusName:
            Ref: "SampleCustomEventBus"
        StatementId: "MyCustomEventBusStatement"
        Statement:
            Effect: "Allow"
            Principal: "*"
            Action: "events:PutEvents"
            Resource: !GetAtt "SampleCustomEventBus.Arn"
            Condition:
                StringEquals:
                    "aws:PrincipalOrgID": "o-1234567890"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Events::Rule

All content copied from https://docs.aws.amazon.com/.
