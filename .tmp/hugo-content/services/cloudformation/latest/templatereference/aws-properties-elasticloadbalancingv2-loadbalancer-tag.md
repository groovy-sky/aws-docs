This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::LoadBalancer Tag

Information about a tag.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key of the tag.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the tag.

_Required_: No

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

The following example creates a Network Load Balancer with two tags.

#### YAML

```yaml

myLoadBalancer:
    Type: AWS::ElasticLoadBalancingV2::LoadBalancer
    Properties:
      Name: my-nlb
      Type: network
      Scheme: internal
      Subnets:
        - !Ref subnet-AZ1
        - !Ref subnet-AZ2
      SecurityGroups:
        - !Ref mySecurityGroup
      Tags:
        - Key: "department"
          Value: "123"
        - Key: "project"
          Value: "lima"
```

#### JSON

```json

{
    "myLoadBalancer": {
        "Type": "AWS::ElasticLoadBalancingV2::LoadBalancer",
        "Properties": {
            "Name": "my-alb",
            "Type": "network",
            "Scheme": "internal",
            "Subnets": [
                {
                    "Ref": "subnet-AZ1"
                },
                {
                    "Ref": "subnet-AZ2"
                }
            ],
            "SecurityGroups": [
                {
                    "Ref": "mySecurityGroup"
                }
            ],
            "Tags": [
                {
                    "Key": "department",
                    "Value": "123"
                },
                {
                    "Key": "project",
                    "Value": "lima"
                }
            ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubnetMapping

AWS::ElasticLoadBalancingV2::TargetGroup

All content copied from https://docs.aws.amazon.com/.
