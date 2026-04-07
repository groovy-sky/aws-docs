This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ResourcePolicy

Retrieves information about the specified resource policy. The resource policy is an IAM
policy created on behalf of the resource owner when they share a resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::ResourcePolicy",
  "Properties" : {
      "Policy" : Json,
      "ResourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::ResourcePolicy
Properties:
  Policy: Json
  ResourceArn: String

```

## Properties

`Policy`

The Amazon Resource Name (ARN) of the service network or service.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

An IAM policy.

_Required_: Yes

_Type_: String

_Pattern_: `^arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:((servicenetwork/sn)|(service/svc)|(resourceconfiguration/rcfg))-[0-9a-z]{17}$`

_Minimum_: `20`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the resource policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::VpcLattice::Rule
