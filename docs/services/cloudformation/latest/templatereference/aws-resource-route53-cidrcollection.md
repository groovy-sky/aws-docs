This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::CidrCollection

Creates a CIDR collection in the current AWS account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53::CidrCollection",
  "Properties" : {
      "Locations" : [ Location, ... ],
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53::CidrCollection
Properties:
  Locations:
    - Location
  Name: String

```

## Properties

`Locations`

A complex type that contains information about the list of CIDR locations.

_Required_: No

_Type_: Array of [Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53-cidrcollection-location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a CIDR collection.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_\-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `CidrCollection` ID.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

"The Amazon resource name (ARN) to uniquely identify the AWS resource.

`Id`

The UUID of the CIDR collection.

## Examples

### Create a CIDR collection

The following example creates CIDR collection with two locations that each contain two CIDR blocks.

#### JSON

```json

{
  "MyCidrCollection":{
    "Type":"AWS::Route53::CidrCollection",
    "Properties":{
      "Name":"my-first-cidr-collection",
      "Locations":[
      {
        "LocationName":"location-1",
        "CidrList":[
          "1.1.0.0/24",
          "2.1.0.0/16"
        ]
      },
      {
        "LocationName":"location-2",
        "CidrList":[
          "2002::1234:0:0:0:0:0/48",
          "1002::/32"
        ]
      }
    ]
   }
  }
}
```

#### YAML

```yaml

MyCidrCollection:
  Type: AWS::Route53::CidrCollection
  Properties:
    Name: "my-first-cidr-collection"
    Locations:
      - LocationName: "location-1"
        CidrList:
          - "1.1.0.0/24"
          - "2.1.0.0/16"
      - LocationName: "location-2"
        CidrList:
          - "2002::1234:0:0:0:0:0/48"
          - "1002::/32"
```

## See also

- [CreateCidrCollection](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateCidrCollection.html) in the _Amazon Route 53 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Route 53

Location
