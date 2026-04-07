This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Detective::MemberInvitation

The `AWS::Detective::MemberInvitation` resource is an Amazon Detective
resource type that creates an invitation to join a Detective behavior graph. The
administrator account can choose whether to send an email notification of the invitation
to the root user email address of the AWS account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Detective::MemberInvitation",
  "Properties" : {
      "DisableEmailNotification" : Boolean,
      "GraphArn" : String,
      "MemberEmailAddress" : String,
      "MemberId" : String,
      "Message" : String
    }
}

```

### YAML

```yaml

Type: AWS::Detective::MemberInvitation
Properties:
  DisableEmailNotification: Boolean
  GraphArn: String
  MemberEmailAddress: String
  MemberId: String
  Message: String

```

## Properties

`DisableEmailNotification`

Whether to send an invitation email to the member account. If set to true, the member account does not receive an invitation email.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GraphArn`

The ARN of the behavior graph to invite the account to contribute data to.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:detective:(([a-z]+-)+[0-9]+):[0-9]{12}:graph:[0-9a-f]{32}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemberEmailAddress`

The root user email address of the invited account. If the email address provided is
not the root user email address for the provided account, the invitation creation
fails.

_Required_: Yes

_Type_: String

_Pattern_: `.*@.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemberId`

The AWS account identifier of the invited account

_Required_: Yes

_Type_: String

_Pattern_: `[0-9]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Message`

Customized text to include in the invitation email message.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the behavior graph and the member account
identifier, separated by a pipe character ('\|').

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Sending a behavior graph invitation to a member account](#aws-resource-detective-memberinvitation--examples--Sending_a_behavior_graph_invitation_to_a_member_account)

- [Blocking the email notification of an invitation to a member account](#aws-resource-detective-memberinvitation--examples--Blocking_the_email_notification_of_an_invitation_to_a_member_account)

### Sending a behavior graph invitation to a member account

This example shows how to declare a new
`AWS:Detective:MemberInvitation` resource to create a new
invitation to a member account and send an email notification.

#### JSON

```json

"MemberInvitation": {
    "Type": "AWS::Detective::MemberInvitation",
    "Properties": {
        "GraphArn": "arn:aws:detective:us-east-1:111122223333:graph:027c7c4610ea4aacaf0b883093cab899",
        "MemberId": "444455556666",
        "MemberEmailAddress": "mmajor@example.com",
        "Message": "This is Paul Santos. I need to add your account to the data we use for security investigation in Detective. If you have any questions, contact me at psantos@example.com."
     }
}
```

#### YAML

```yaml

MemberInvitation:
    Type: AWS::Detective::MemberInvitation
    Properties:
      GraphArn: "arn:aws:detective:us-east-1:111122223333:graph:027c7c4610ea4aacaf0b883093cab899"
      MemberId: 444455556666
      MemberEmailAddress: mmajor@example.com
      Message: This is Paul Santos. I need to add your account to the data we use for security investigation in Detective. If you have any questions, contact me at psantos@example.com.
```

### Blocking the email notification of an invitation to a member account

This example shows how to declare a new
`AWS:Detective:MemberInvitation` resource to create a new
invitation to a member account. The email notification is blocked.

#### JSON

```json

"MemberInvitation": {
    "Type": "AWS::Detective::MemberInvitation",
    "Properties": {
        "GraphArn": "arn:aws:detective:us-east-1:111122223333:graph:027c7c4610ea4aacaf0b883093cab899",
        "MemberId": "444455556666",
        "MemberEmailAddress": "mmajor@example.com",
        "DisableEmailNotification": "true"
     }
}
```

#### YAML

```yaml

MemberInvitation:
    Type: AWS::Detective::MemberInvitation
    Properties:
      GraphArn: "arn:aws:detective:us-east-1:111122223333:graph:027c7c4610ea4aacaf0b883093cab899"
      MemberId: 444455556666
      MemberEmailAddress: mmajor@example.com
      DisableEmailNotification: true
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Detective::OrganizationAdmin
