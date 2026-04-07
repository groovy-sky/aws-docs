This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SnapshotBlockPublicAccess

Specifies the state of the _block public access for snapshots_
setting for the Region. For more information, see [Block public access for snapshots](https://docs.aws.amazon.com/ebs/latest/userguide/block-public-access-snapshots.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::SnapshotBlockPublicAccess",
  "Properties" : {
      "State" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::SnapshotBlockPublicAccess
Properties:
  State: String

```

## Properties

`State`

The mode in which to enable block public access for snapshots for the Region.
Specify one of the following values:

- `block-all-sharing` \- Prevents all public sharing of snapshots in
the Region. Users in the account will no longer be able to request new public
sharing. Additionally, snapshots that are already publicly shared are treated as
private and they are no longer publicly available.

###### Note

If you enable block public access for snapshots in `block-all-sharing`
mode, it does not change the permissions for snapshots that are already publicly shared.
Instead, it prevents these snapshots from be publicly visible and publicly accessible.
Therefore, the attributes for these snapshots still indicate that they are publicly
shared, even though they are not publicly available.

- `block-new-sharing` \- Prevents only new public sharing of snapshots
in the Region. Users in the account will no longer be able to request new public
sharing. However, snapshots that are already publicly shared, remain publicly
available.

_Required_: Yes

_Type_: String

_Allowed values_: `block-all-sharing | block-new-sharing`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the AWS account.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the AWS account.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::SecurityGroupVpcAssociation

AWS::EC2::SpotFleet
