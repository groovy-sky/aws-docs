This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::DataflowEndpointGroup

Creates a Dataflow Endpoint Group request.

Dataflow endpoint groups contain a list of endpoints.
When the name of a dataflow endpoint group is specified in a mission profile, the Ground Station service will connect to the endpoints and flow data during a contact.

For more information about dataflow endpoint groups, see [Dataflow Endpoint Groups](../../../ground-station/latest/ug/dataflowendpointgroups.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GroundStation::DataflowEndpointGroup",
  "Properties" : {
      "ContactPostPassDurationSeconds" : Integer,
      "ContactPrePassDurationSeconds" : Integer,
      "EndpointDetails" : [ EndpointDetails, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GroundStation::DataflowEndpointGroup
Properties:
  ContactPostPassDurationSeconds: Integer
  ContactPrePassDurationSeconds: Integer
  EndpointDetails:
    - EndpointDetails
  Tags:
    - Tag

```

## Properties

`ContactPostPassDurationSeconds`

Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a `POSTPASS` state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the `POSTPASS` state.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Maximum_: `480`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContactPrePassDurationSeconds`

Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a `PREPASS` state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the `PREPASS` state.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Maximum_: `480`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointDetails`

List of Endpoint Details, containing address and port for each endpoint.

All dataflow endpoints within a single dataflow endpoint group must be of the same type.
You cannot mix AWS Ground Station Agent endpoints with Dataflow endpoints in the same group.
If your use case requires both types of endpoints, you must create separate dataflow endpoint
groups for each type.

_Required_: Yes

_Type_: [Array](aws-properties-groundstation-dataflowendpointgroup-endpointdetails.md) of [EndpointDetails](aws-properties-groundstation-dataflowendpointgroup-endpointdetails.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags assigned to a resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-groundstation-dataflowendpointgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the dataflow endpoint group. For example:

`{ "Ref": "DataflowEndpointGroup" }`

For the Ground Station dataflow endpoint group, `Ref` returns the ARN of the dataflow endpoint group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the dataflow endpoint group, such as
`arn:aws:groundstation:us-east-2:012345678910:dataflow-endpoint-group/9940bf3b-d2ba-427e-9906-842b5e5d2296`.

`Id`

UUID of a dataflow endpoint group.

## Examples

### Create a DataflowEndpointGroup

The following example creates a Ground Station `DataflowEndpointGroup`

#### JSON

```json

{
  "Resources": {
    "myDataflowEndpointGroup": {
      "Type": "AWS::GroundStation::DataflowEndpointGroup",
      "Properties": {
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
        ]
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  myDataflowEndpointGroup:
    Type: AWS::GroundStation::DataflowEndpointGroup
    Properties:
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

UplinkSpectrumConfig

AwsGroundStationAgentEndpoint

All content copied from https://docs.aws.amazon.com/.
