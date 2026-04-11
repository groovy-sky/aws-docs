This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup EndpointDetails

The security details and endpoint information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsGroundStationAgentEndpoint" : AwsGroundStationAgentEndpoint,
  "Endpoint" : DataflowEndpoint,
  "SecurityDetails" : SecurityDetails
}

```

### YAML

```yaml

  AwsGroundStationAgentEndpoint:
    AwsGroundStationAgentEndpoint
  Endpoint:
    DataflowEndpoint
  SecurityDetails:
    SecurityDetails

```

## Properties

`AwsGroundStationAgentEndpoint`

An agent endpoint.

_Required_: No

_Type_: [AwsGroundStationAgentEndpoint](aws-properties-groundstation-dataflowendpointgroup-awsgroundstationagentendpoint.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Endpoint`

Information about the endpoint such as name and the endpoint address.

_Required_: No

_Type_: [DataflowEndpoint](aws-properties-groundstation-dataflowendpointgroup-dataflowendpoint.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityDetails`

The role ARN, and IDs for security groups and subnets.

_Required_: No

_Type_: [SecurityDetails](aws-properties-groundstation-dataflowendpointgroup-securitydetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Create a EndpointDetails

The following example creates Ground Station `EndpointDetails`

#### JSON

```json

{
  "EndpointDetails": [
    {
      "SecurityDetails": {
        "SubnetIds": [
          "subnet-6782e71e"
        ],
        "SecurityGroupIds": [
          "sg-6979fe18"
        ],
        "RoleArn": "arn:aws:iam::012345678910:role/groundstation-service-role-AWSServiceRoleForAmazonGroundStation-EXAMPLEBQ4PI"
      },
      "Endpoint": {
        "Name": "myEndpoint",
        "Address": {
          "Name": "172.10.0.2",
          "Port": 44720
        },
        "Mtu": 1500
      }
    }
  }
}
```

#### YAML

```yaml

EndpointDetails:
  - SecurityDetails:
      SubnetIds:
        - subnet-12345678
      SecurityGroupIds:
        - sg-87654321
      RoleArn: arn:aws:iam::012345678910:role/groundstation-service-role-AWSServiceRoleForAmazonGroundStation-EXAMPLEABCDE
Endpoint:
  Name: myEndpoint
  Address:
    Name: 172.10.0.2
    Port: 44720
  Mtu: 1500

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataflowEndpoint

IntegerRange

All content copied from https://docs.aws.amazon.com/.
