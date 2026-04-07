This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryControl::RoutingControl

Creates a routing control in Amazon Route 53 Application Recovery Controller. Routing control states are maintained
on the highly reliable cluster data plane.

To get or update the state of the routing control, you must specify a cluster endpoint, which is an
endpoint URL and an AWS Region. For more information, see [Code examples](https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples.html)
in the Amazon Route 53 Application Recovery Controller Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryControl::RoutingControl",
  "Properties" : {
      "ClusterArn" : String,
      "ControlPanelArn" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryControl::RoutingControl
Properties:
  ClusterArn: String
  ControlPanelArn: String
  Name: String

```

## Properties

`ClusterArn`

The Amazon Resource Name (ARN) of the cluster that hosts the routing control.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9:\/_-]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ControlPanelArn`

The Amazon Resource Name (ARN) of the control panel that includes the routing control.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9:\/_-]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the routing control. You can use any non-white space character in the name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `RoutingControlArn` object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RoutingControlArn`

The Amazon Resource Name (ARN) of the routing control.

`Status`

The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING\_DELETION.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Route53RecoveryControl::SafetyRule
