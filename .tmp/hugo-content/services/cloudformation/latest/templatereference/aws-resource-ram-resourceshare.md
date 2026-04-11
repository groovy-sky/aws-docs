This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RAM::ResourceShare

Creates a resource share. You can provide a list of the Amazon Resource Names (ARNs)
for the resources that you want to share, a list of principals you want to share the
resources with, and the permissions to grant those principals.

###### Note

Sharing a resource makes it available for use by principals outside of the AWS account that created the resource. Sharing doesn't change any
permissions or quotas that apply to the resource in the account that created
it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RAM::ResourceShare",
  "Properties" : {
      "AllowExternalPrincipals" : Boolean,
      "Name" : String,
      "PermissionArns" : [ String, ... ],
      "Principals" : [ String, ... ],
      "ResourceArns" : [ String, ... ],
      "Sources" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RAM::ResourceShare
Properties:
  AllowExternalPrincipals: Boolean
  Name: String
  PermissionArns:
    - String
  Principals:
    - String
  ResourceArns:
    - String
  Sources:
    - String
  Tags:
    - Tag

```

## Properties

`AllowExternalPrincipals`

Specifies whether principals outside your organization in AWS Organizations can be associated
with a resource share. A value of `true` lets you share with individual AWS accounts
that are _not_ in your organization. A value of `false`
only has meaning if your account is a member of an AWS Organization. The default value
is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Specifies the name of the resource share.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PermissionArns`

Specifies the [Amazon Resource Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md) of the AWS RAM permission to associate with the resource share. If you do
not specify an ARN for the permission, AWS RAM automatically attaches the default version
of the permission for each resource type. You can associate only one permission with
each resource type included in the resource share.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principals`

Specifies the principals to associate with the resource share. The possible values
are:

- An AWS account ID

- An Amazon Resource Name (ARN) of an organization in AWS Organizations

- An ARN of an organizational unit (OU) in AWS Organizations

- An ARN of an IAM role

- An ARN of an IAM user

###### Note

Not all resource types can be shared with IAM roles and users. For
more information, see the column **Can share with IAM roles and users** in the tables on [Shareable AWS resources](../../../ram/latest/userguide/shareable.md) in the _AWS Resource Access Manager User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArns`

Specifies a list of one or more ARNs of the resources to associate with the
resource share.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sources`

Specifies from which source accounts the service principal has access to the resources in this resource share.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies one or more tags to attach to the resource share itself. It doesn't attach the tags to
the resources associated with the resource share.

_Required_: No

_Type_: Array of [Tag](aws-properties-ram-resourceshare-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns The ID of the resource share.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the resource share.

`CreationTime`

The date and time when the resource share was created.

`FeatureSet`

Indicates what features are available for this resource share. This parameter can have one of
the following values:

- **STANDARD** – A resource share that supports all
functionality. These resource shares are visible to all principals you share the resource share with.
You can modify these resource shares in AWS RAM using the console or APIs. This resource share might
have been created by AWS RAM, or it might have been **CREATED\_FROM\_POLICY** and then promoted.

- **CREATED\_FROM\_POLICY** – The customer
manually shared a resource by attaching a resource-based policy. That policy did
not match any existing managed permissions, so AWS RAM created this customer managed permission automatically on the
customer's behalf based on the attached policy document. This type of resource share
is visible only to the AWS account that created it. You can't modify it in
AWS RAM unless you promote it. For more information,
see PromoteResourceShareCreatedFromPolicy.

- **PROMOTING\_TO\_STANDARD** – This
resource share was originally `CREATED_FROM_POLICY`, but the customer ran
the PromoteResourceShareCreatedFromPolicy and that operation
is still in progress. This value changes to `STANDARD` when
complete.

`LastUpdatedTime`

The date and time when the resource share was last updated.

`OwningAccountId`

The ID of the AWS account that owns the resource share.

`Status`

The current status of the resource share.

## Examples

### Creating a resource share

The following example demonstrates how to create a resource share.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myresourceshare:
    Type: "AWS::RAM::ResourceShare"
    Properties:
      Name: "My Resource Share"
      ResourceArns:
        - "arn:aws:ec2:us-east-1:123456789012:resource-type/12345678-1234-1234-1234-12345678"
      Principals:
        - "210987654321"
      Tags:
        - Key: "Key1"
          Value: "Value1"
        - Key: "Key2"
          Value: "Value2"

```

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09T00:00:00.000Z",
  "Resources": {
    "myresourceshare": {
      "Type": "AWS::RAM::ResourceShare",
      "Properties": {
        "Name": "My Resource Share",
        "ResourceArns": [
          "arn:aws:ec2:us-east-1:123456789012:resource-type/12345678-1234-1234-1234-12345678"
        ],
        "Principals": [
          "210987654321"
        ],
        "Tags": [
          {
            "Key": "Key1",
            "Value": "Value1"
          },
          {
            "Key": "Key2",
            "Value": "Value2"
          }
        ]
      }
    }
  }
}
```

## See also

- [CreateResourceShare](../../../../reference/ram/latest/apireference/api-createresourceshare.md) in the _AWS Resource Access Manager API Reference_

- [AWS Resource Access Manager User Guide](../../../ram/latest/userguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
