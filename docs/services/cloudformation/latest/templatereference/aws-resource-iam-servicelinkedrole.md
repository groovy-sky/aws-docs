This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::ServiceLinkedRole

Creates an IAM role that is linked to a specific AWS service. The service controls
the attached policies and when the role can be deleted. This helps ensure that the
service is not broken by an unexpectedly changed or deleted role, which could put your
AWS resources into an unknown state. Allowing the service to control the role helps
improve service stability and proper cleanup when a service and its role are no longer
needed. For more information, see [Using service-linked\
roles](../../../iam/latest/userguide/using-service-linked-roles.md) in the _IAM User Guide_.

To attach a policy to this service-linked role, you must make the request using the
AWS service that depends on this role.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::ServiceLinkedRole",
  "Properties" : {
      "AWSServiceName" : String,
      "CustomSuffix" : String,
      "Description" : String
    }
}

```

### YAML

```yaml

Type: AWS::IAM::ServiceLinkedRole
Properties:
  AWSServiceName: String
  CustomSuffix: String
  Description: String

```

## Properties

`AWSServiceName`

The service principal for the AWS service to which this role is attached. You use a
string similar to a URL but without the http:// in front. For example:
`elasticbeanstalk.amazonaws.com`.

Service principals are unique and case-sensitive. To find the exact service principal
for your service-linked role, see [AWS services\
that work with IAM](../../../iam/latest/userguide/reference-aws-services-that-work-with-iam.md) in the _IAM User Guide_. Look for
the services that have **Yes** in the **Service-Linked Role** column. Choose the **Yes** link to view the service-linked role documentation for that
service.

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomSuffix`

A string that you provide, which is combined with the service-provided prefix to form
the complete role name. If you make multiple requests for the same service, then you
must supply a different `CustomSuffix` for each request. Otherwise the
request fails with a duplicate role name error. For example, you could add
`-1` or `-debug` to the suffix.

Some services do not support the `CustomSuffix` parameter. If you provide
an optional suffix and the operation fails, try the operation again without the
suffix.

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the role.

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u007E\u00A1-\u00FF]*`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `RoleName` created for the service-linked role.
The `CustomSuffix` is appended to the service-provided prefix with an underscore
followed by the `CustomSuffix` to form the complete role name. For example,
`AWSServiceRoleForAutoScaling` or
`AWSServiceRoleForAutoScaling_TestSuffix` if a `CustomSuffix` is
specified.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RoleName`

Returns the friendly name that identifies the role. For example,
`AWSServiceRoleForAutoScaling` or
`AWSServiceRoleForAutoScaling_TestSuffix` if a `CustomSuffix` is
specified.

## Examples

### Service-Linked Role for Auto Scaling

The following example creates a service-linked role that can be assumed by the
Auto Scaling service.

#### JSON

```json

{
    "Description": "SLR resource create test - Auto Scaling",
    "Resources": {
        "BasicSLR": {
            "Type": "AWS::IAM::ServiceLinkedRole",
            "Properties": {
                "AWSServiceName": "autoscaling.amazonaws.com",
                "Description": "Test SLR description",
                "CustomSuffix": "TestSuffix"
            }
        }
    },
    "Outputs": {
        "SLRId": {
            "Value": {
                "Ref": "BasicSLR"
            }
        }
    }
}
```

#### YAML

```yaml

Description: SLR resource create test - Auto Scaling
Resources:
  BasicSLR:
    Type: 'AWS::IAM::ServiceLinkedRole'
    Properties:
      AWSServiceName: autoscaling.amazonaws.com
      Description: Test SLR description
      CustomSuffix: TestSuffix
Outputs:
  SLRId:
    Value: !Ref BasicSLR
```

## See also

- [CreateServiceLinkedRole](../../../../reference/iam/latest/apireference/api-createservicelinkedrole.md) in the _AWS Identity and Access Management API Reference_

- [Using Service-Linked\
Roles](../../../iam/latest/userguide/using-service-linked-roles.md) in the _AWS Identity and Access Management User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IAM::User

All content copied from https://docs.aws.amazon.com/.
