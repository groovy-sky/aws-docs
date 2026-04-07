This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Member

You can use the `AWS::GuardDuty::Member` resource to add an AWS account as a GuardDuty member account to the current GuardDuty administrator account. If the value of the `Status` property is
not provided or is set to `Created`, a member account is created but not
invited. If the value of the `Status` property is set to `Invited`, a
member account is created and invited. An `AWS::GuardDuty::Member` resource must
be created with the `Status` property set to `Invited` before the
`AWS::GuardDuty::Master` resource can be created in a GuardDuty
member account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::Member",
  "Properties" : {
      "DetectorId" : String,
      "DisableEmailNotification" : Boolean,
      "Email" : String,
      "MemberId" : String,
      "Message" : String,
      "Status" : String
    }
}

```

### YAML

```yaml

Type: AWS::GuardDuty::Member
Properties:
  DetectorId: String
  DisableEmailNotification: Boolean
  Email: String
  MemberId: String
  Message: String
  Status: String

```

## Properties

`DetectorId`

The ID of the detector associated with the GuardDuty service to add the member
to.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisableEmailNotification`

Specifies whether or not to disable email notification for the member account that you
invite.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Email`

The email address associated with the member account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemberId`

The AWS account ID of the account to designate as a member.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Message`

The invitation message that you want to send to the accounts that you're inviting to
GuardDuty as members.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

You can use the `Status` property to update the status of the relationship
between the member account and its administrator account. Valid values are
`Created` and `Invited` when using an
`AWS::GuardDuty::Member` resource. If the value for this property is not
provided or set to `Created`, a member account is created but not invited. If
the value of this property is set to `Invited`, a member account is created and
invited.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique ID of the GuardDuty member account, such
as `111122223333`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Declare a Member Resource

The following example shows how to declare a GuardDuty `Member` resource:

#### JSON

```json

"GDmaster": {
    "Type": "AWS::GuardDuty::Member",
    "Properties": {
        "Status": "Invited",
        "MemberId": "555555555555",
        "Email": "guardduty-member@amazon.com",
        "Message": "You are invited to enable Amazon Guardduty.",
        "DetectorId": "a12abc34d567e8fa901bc2d34e56789f0",
        "DisableEmailNotification": true
        }
}
```

#### YAML

```yaml

      Type: AWS::GuardDuty::Member
      Properties:
            Status: Invited
            MemberId: 555555555555
            Email: guardduty-member@amazon.com
            Message: You are invited to enable Amazon Guardduty.
            DetectorId: a12abc34d567e8fa901bc2d34e56789f0
            DisableEmailNotification: true

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GuardDuty::Master

AWS::GuardDuty::PublishingDestination
