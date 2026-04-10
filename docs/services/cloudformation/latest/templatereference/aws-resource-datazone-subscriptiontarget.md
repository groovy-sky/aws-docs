This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::SubscriptionTarget

The `AWS::DataZone::SubscriptionTarget` resource specifies an Amazon DataZone subscription target.
Subscription targets enable you to access the data to which you have subscribed in your projects. A subscription
target specifies the location (for example, a database or a schema) and the required permissions (for example, an IAM
role) that Amazon DataZone can use to establish a connection with the source data and to create the necessary grants
so that members of the Amazon DataZone project can start querying the data to which they have subscribed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::SubscriptionTarget",
  "Properties" : {
      "ApplicableAssetTypes" : [ String, ... ],
      "AuthorizedPrincipals" : [ String, ... ],
      "DomainIdentifier" : String,
      "EnvironmentIdentifier" : String,
      "ManageAccessRole" : String,
      "Name" : String,
      "Provider" : String,
      "SubscriptionTargetConfig" : [ SubscriptionTargetForm, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::SubscriptionTarget
Properties:
  ApplicableAssetTypes:
    - String
  AuthorizedPrincipals:
    - String
  DomainIdentifier: String
  EnvironmentIdentifier: String
  ManageAccessRole: String
  Name: String
  Provider: String
  SubscriptionTargetConfig:
    - SubscriptionTargetForm
  Type: String

```

## Properties

`ApplicableAssetTypes`

The asset types included in the subscription target.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizedPrincipals`

The authorized principals included in the subscription target.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The ID of the Amazon DataZone domain in which subscription target is created.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentIdentifier`

The ID of the environment in which subscription target is created.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManageAccessRole`

The manage access role that is used to create the subscription target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the subscription target.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Provider`

The provider of the subscription target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriptionTargetConfig`

The configuration of the subscription target.

_Required_: Yes

_Type_: Array of [SubscriptionTargetForm](aws-properties-datazone-subscriptiontarget-subscriptiontargetform.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the subscription target.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId`,
`EnvironmentId` and the `SubscriptionTargetId` that uniquely identify the subscription target.
For example: `{ "Ref": "MySubscriptionTarget" }` for the resource with the logical ID
`MySubscriptionTarget`, `Ref` returns
`DomainId|EnvironmentId|SubscriptionTargetId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when the subscription target was created.

`CreatedBy`

The Amazon DataZone user who created the subscription target.

`DomainId`

The identifier of the Amazon DataZone domain in which the subscription target exists.

`EnvironmentId`

The identifier of the environment of the subscription target.

`Id`

The identifier of the subscription target.

`ProjectId`

The identifier of the project specified in the subscription target.

`UpdatedAt`

The timestamp of when the subscription target was updated.

`UpdatedBy`

The Amazon DataZone user who updated the subscription target.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTagParameter

SubscriptionTargetForm

All content copied from https://docs.aws.amazon.com/.
