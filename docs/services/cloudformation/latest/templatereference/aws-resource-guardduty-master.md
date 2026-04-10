This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Master

You can use the `AWS::GuardDuty::Master` resource in a GuardDuty
member account to accept an invitation from a GuardDuty administrator account. The
invitation to the member account must be sent prior to using the
`AWS::GuardDuty::Master` resource to accept the administrator account's
invitation. You can invite a member account by using the `InviteMembers`
operation of the GuardDuty API, or by creating an
`AWS::GuardDuty::Member` resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::Master",
  "Properties" : {
      "DetectorId" : String,
      "InvitationId" : String,
      "MasterId" : String
    }
}

```

### YAML

```yaml

Type: AWS::GuardDuty::Master
Properties:
  DetectorId: String
  InvitationId: String
  MasterId: String

```

## Properties

`DetectorId`

The unique ID of the detector of the GuardDuty member account.

To find the `detectorId` in the current Region, see the
Settings page in the GuardDuty console, or run the [ListDetectors](../../../../reference/guardduty/latest/apireference/api-listdetectors.md) API.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InvitationId`

The ID of the invitation that is sent to the account designated as a member account. You
can find the invitation ID by running the [ListInvitations](../../../../reference/guardduty/latest/apireference/api-listinvitations.md) in
the _GuardDuty API Reference_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MasterId`

The AWS account ID of the account designated as the GuardDuty
administrator account.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique ID of the GuardDuty administrator account,
such as `111122223333`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Declare a Master Resource

To declare a GuardDuty `Master` resource:

#### JSON

```json

"GDMaster": {
    "Type" : "AWS::GuardDuty::Master",
    "Properties" : {
        "DetectorId" : "a12abc34d567e8fa901bc2d34e56789f0",
        "MasterId" : "111122223333",
        "InvitationId" : "84b097800250d17d1872b34c4daadcf5"
    }
}
```

#### YAML

```yaml

GDMaster:
    Type: AWS::GuardDuty::Master
    Properties:
        DetectorId: "a12abc34d567e8fa901bc2d34e56789f0"
        MasterId: "111122223333"
        InvitationId: "84b097800250d17d1872b34c4daadcf5"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagItem

AWS::GuardDuty::Member

All content copied from https://docs.aws.amazon.com/.
