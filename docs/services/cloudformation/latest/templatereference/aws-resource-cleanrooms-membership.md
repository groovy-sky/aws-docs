This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership

Creates a membership for a specific collaboration identifier and joins the
collaboration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRooms::Membership",
  "Properties" : {
      "CollaborationIdentifier" : String,
      "DefaultJobResultConfiguration" : MembershipProtectedJobResultConfiguration,
      "DefaultResultConfiguration" : MembershipProtectedQueryResultConfiguration,
      "IsMetricsEnabled" : Boolean,
      "JobLogStatus" : String,
      "PaymentConfiguration" : MembershipPaymentConfiguration,
      "QueryLogStatus" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRooms::Membership
Properties:
  CollaborationIdentifier: String
  DefaultJobResultConfiguration:
    MembershipProtectedJobResultConfiguration
  DefaultResultConfiguration:
    MembershipProtectedQueryResultConfiguration
  IsMetricsEnabled: Boolean
  JobLogStatus: String
  PaymentConfiguration:
    MembershipPaymentConfiguration
  QueryLogStatus: String
  Tags:
    - Tag

```

## Properties

`CollaborationIdentifier`

The unique ID for the associated collaboration.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultJobResultConfiguration`

The default job result configuration for the membership.

_Required_: No

_Type_: [MembershipProtectedJobResultConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-membership-membershipprotectedjobresultconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultResultConfiguration`

The default protected query result configuration as specified by the member who can
receive results.

_Required_: No

_Type_: [MembershipProtectedQueryResultConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-membership-membershipprotectedqueryresultconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsMetricsEnabled`

An indicator as to whether Amazon CloudWatch metrics are enabled for the membership.

When `true`, metrics about query execution are collected in Amazon CloudWatch.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobLogStatus`

An indicator as to whether job logging has been enabled or disabled
for the collaboration.

When `ENABLED`, AWS Clean Rooms logs details about jobs run within this
collaboration and those logs can be viewed in Amazon CloudWatch Logs. The default value is
`DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PaymentConfiguration`

The payment responsibilities accepted by the collaboration member.

_Required_: No

_Type_: [MembershipPaymentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-membership-membershippaymentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryLogStatus`

An indicator as to whether query logging has been enabled or disabled for the
membership.

When `ENABLED`, AWS Clean Rooms logs details about queries run within this
collaboration and those logs can be viewed in Amazon CloudWatch Logs. The default value is
`DISABLED`.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An optional label that you can assign to a resource when you create it. Each tag
consists of a key and an optional value, both of which you define. When you use tagging,
you can also use tag-based access control in IAM policies to control access
to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-membership-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `MembershipId`, such as
`a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`. For example:

`{ "Ref": "MyMembership" }`

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the specified membership.

Example:
`arn:aws:cleanrooms:us-east-1:111122223333:membership/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

`CollaborationArn`

Returns the Amazon Resource Name (ARN) of the specified collaboration.

Example:
`arn:aws:cleanrooms:us-east-1:111122223333:collaboration/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

`CollaborationCreatorAccountId`

Returns the unique identifier of the specified collaboration creator account.

Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

`MembershipIdentifier`

Returns the unique identifier of the specified membership.

Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE22222`

## Remarks

If you are a collaboration owner, ensure that you add `"DeletionPolicy:
          Retain"` and `"UpdateReplacePolicy: Retain"` to the
`Membership` resource in your CloudFormation template.

If your `Membership` resource depends on the collaboration explicitly or
implicitly by using a `"Ref"` on the `Collaboration`
( `CollaborationIdentifier: !Ref Collaboration`), CloudFormation tries to delete
the `Membership` resource before the `Collaboration` resource.
However, this attempt will fail because collaboration owners must delete the collaboration
before deleting their membership to the collaboration.

For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) and [UpdateReplacePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html).

## Examples

- [Membership for a collaboration creator](#aws-resource-cleanrooms-membership--examples--Membership_for_a_collaboration_creator)

- [Membership for a non-collaboration creator](#aws-resource-cleanrooms-membership--examples--Membership_for_a_non-collaboration_creator)

### Membership for a collaboration creator

The following an example of a membership for a collaboration creator. See the Remarks
section for notes on the use of the deletion and update replacement policies.

#### JSON

```json

"CollaborationCreatorMembership": {
  "Type" : "AWS::CleanRooms::Membership",
  "DeletionPolicy": "Retain",
  "UpdateReplacePolicy": "Retain",
  "Properties" : {
      "CollaborationIdentifier" : "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
      "QueryLogStatus" : "ENABLED",
    }
  }
```

#### YAML

```yaml

CollaborationCreatorMembership:
  Type: AWS::CleanRooms::Membership
  DeletionPolicy: Retain
  UpdateReplacePolicy: Retain
  Properties:
    CollaborationIdentifier: a1b2c3d4-5678-90ab-cdef-EXAMPLE11111
    QueryLogStatus: ENABLED
```

### Membership for a non-collaboration creator

The following is an example of a membership for a non-collaboration creator.

#### JSON

```json

"ExampleMembership": {
  "Type" : "AWS::CleanRooms::Membership",
  "Properties" : {
      "CollaborationIdentifier" : "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
      "QueryLogStatus" : "ENABLED",
    }
  }
```

#### YAML

```yaml

ExampleMembership:
  Type: AWS::CleanRooms::Membership
  Properties:
    CollaborationIdentifier: a1b2c3d4-5678-90ab-cdef-EXAMPLE11111
    QueryLogStatus: ENABLED
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

MembershipJobComputePaymentConfig
