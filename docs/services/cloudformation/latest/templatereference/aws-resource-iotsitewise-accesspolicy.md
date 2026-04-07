This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AccessPolicy

###### Important

The AWS IoT SiteWise Monitor feature will no longer be open to new customers starting November 7, 2025 . If you would like to use the AWS IoT SiteWise Monitor feature,
sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS IoT SiteWise Monitor availability change](https://docs.aws.amazon.com/iot-sitewise/latest/appguide/iotsitewise-monitor-availability-change.html).

Creates an access policy that grants the specified identity (IAM Identity Center user, IAM Identity Center group, or
IAM user) access to the specified AWS IoT SiteWise Monitor portal or project resource.

###### Note

Support for access policies that use an SSO Group as the identity is not supported at this time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::AccessPolicy",
  "Properties" : {
      "AccessPolicyIdentity" : AccessPolicyIdentity,
      "AccessPolicyPermission" : String,
      "AccessPolicyResource" : AccessPolicyResource
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::AccessPolicy
Properties:
  AccessPolicyIdentity:
    AccessPolicyIdentity
  AccessPolicyPermission: String
  AccessPolicyResource:
    AccessPolicyResource

```

## Properties

`AccessPolicyIdentity`

The identity for this access policy. Choose an IAM Identity Center user, an IAM Identity Center group, or an IAM user.

_Required_: Yes

_Type_: [AccessPolicyIdentity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessPolicyPermission`

The permission level for this access policy. Note that a project `ADMINISTRATOR` is also known as a project owner.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessPolicyResource`

The AWS IoT SiteWise Monitor resource for this access policy. Choose either a portal or a project.

_Required_: Yes

_Type_: [AccessPolicyResource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-accesspolicy-accesspolicyresource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AccessPolicyId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccessPolicyArn`

The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the access policy, which has the following format.

`arn:${Partition}:iotsitewise:${Region}:${Account}:access-policy/${AccessPolicyId}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`AccessPolicyId`

The ID of the access policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS IoT SiteWise

AccessPolicyIdentity
