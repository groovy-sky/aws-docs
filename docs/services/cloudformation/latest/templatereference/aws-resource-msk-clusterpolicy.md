This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::ClusterPolicy

Create or update cluster policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::ClusterPolicy",
  "Properties" : {
      "ClusterArn" : String,
      "Policy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::MSK::ClusterPolicy
Properties:
  ClusterArn: String
  Policy: Json

```

## Properties

`ClusterArn`

The Amazon Resource Name (ARN) that uniquely identifies the cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[\w-]+:kafka:[\w-]+:\d+:cluster.*\Z`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

Resource policy for the cluster. The maximum size supported for a resource-based policy document is 20 KB.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the cluster for the resource policy.

For Amazon MSK cluster policy `MyClusterPolicy`, `Ref` returns the cluster ARN for the resource policy whose logical ID is `MyClusterPolicy`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CurrentVersion`

The current version of the policy attached to the specified cluster.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConnectivityTls

AWS::MSK::Configuration

All content copied from https://docs.aws.amazon.com/.
