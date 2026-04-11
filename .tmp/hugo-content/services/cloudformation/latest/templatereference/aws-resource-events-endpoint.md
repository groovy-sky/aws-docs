This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Endpoint

A global endpoint used to improve your application's availability by making it
regional-fault tolerant. For more information about global endpoints, see [Making\
applications Regional-fault tolerant with global endpoints and event replication](../../../eventbridge/latest/userguide/eb-global-endpoints.md) in
the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Events::Endpoint",
  "Properties" : {
      "Description" : String,
      "EventBuses" : [ EndpointEventBus, ... ],
      "Name" : String,
      "ReplicationConfig" : ReplicationConfig,
      "RoleArn" : String,
      "RoutingConfig" : RoutingConfig
    }
}

```

### YAML

```yaml

Type: AWS::Events::Endpoint
Properties:
  Description: String
  EventBuses:
    - EndpointEventBus
  Name: String
  ReplicationConfig:
    ReplicationConfig
  RoleArn: String
  RoutingConfig:
    RoutingConfig

```

## Properties

`Description`

A description for the endpoint.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBuses`

The event buses being used by the endpoint.

_Exactly_: `2`

_Required_: Yes

_Type_: Array of [EndpointEventBus](aws-properties-events-endpoint-endpointeventbus.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the endpoint.

_Required_: No

_Type_: String

_Pattern_: `^[\.\-_A-Za-z0-9]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicationConfig`

Whether event replication was enabled or disabled for this endpoint. The default state is
`ENABLED` which means you must supply a `RoleArn`. If you don't have a
`RoleArn` or you don't want event replication enabled, set the state to
`DISABLED`.

_Required_: No

_Type_: [ReplicationConfig](aws-properties-events-endpoint-replicationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role used by event replication for the endpoint.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:iam::\d{12}:role\/[\w+=,.@/-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingConfig`

The routing configuration of the endpoint.

_Required_: Yes

_Type_: [RoutingConfig](aws-properties-events-endpoint-routingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns Endpoint ID, such as
`mystack-Endpoint-ABCDEFGHIJK`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The ARN of the endpoint.

`EndpointId`

The ID of the endpoint.

`EndpointUrl`

The URL of the endpoint.

`State`

The main Region of the endpoint.

`StateReason`

The reason the endpoint is in its current state.

## Examples

- [Create a global endpoint with event replication](#aws-resource-events-endpoint--examples--Create_a_global_endpoint_with_event_replication)

- [Create a global endpoint without event replication](#aws-resource-events-endpoint--examples--Create_a_global_endpoint_without_event_replication)

### Create a global endpoint with event replication

The following example creates a global endpoint with event replication enabled.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "This template will create an Endpoint resource",
  "Resources": {
    "SampleEndpoint": {
      "Type": "AWS::Events::Endpoint",
      "Properties": {
        "Name": "CreateExampleEndpoint",
        "RoutingConfig": {
          "FailoverConfig": {
            "Primary": {
              "HealthCheck": {
                  "arn:aws:route53:::healthcheck/0123456789abc"
              }
            },
            "Secondary": {
              "Route": {
                  "us-east-1"
              }
            }
          }
        },
        "ReplicationConfig": {
          "State": "ENABLED"
        },
        "RoleArn": {
          "arn:aws:iam::123456789012:role/EndpointReplicationRole"
        },
        "EventBuses": [
          {
            "EventBusArn": {
              "arn:aws:events:us-west-2:123456789012:event-bus/ExampleEventBus"
            }
          },
          {
            "EventBusArn": {
              "arn:aws:events:us-east-1:123456789012:event-bus/ExampleEventBus"
            }
          }
        ]
      }
    }
  },
  "Outputs": {
    "SampleEndpointName": {
      "Value": {
        "Ref": "SampleEndpoint"
      }
    }
  }
}
```

#### YAML

```yaml

  AWSTemplateFormatVersion: "2010-09-09"
  Description: "This template will create an Endpoint resource"
  Resources:
    SampleEndpoint:
      Type: AWS::Events::Endpoint
      Properties:
        Name: CreateExampleEndpoint
        RoutingConfig:
          FailoverConfig:
            Primary:
              HealthCheck: "arn:aws:route53:::healthcheck/0123456789abc"
            Secondary:
              Route: "us-east-1"
        ReplicationConfig:
          State: "ENABLED"
        RoleArn: "arn:aws:iam::123456789012:role/EndpointReplicationRole"
        EventBuses:
          - EventBusArn: "arn:aws:events:us-west-2:123456789012:event-bus/ExampleEventBus"
          - EventBusArn: "arn:aws:events:us-east-1:123456789012:event-bus/ExampleEventBus"

  Outputs:
    SampleEndpointName:
      Value: !Ref SampleEndpoint
```

### Create a global endpoint without event replication

The following example creates a global endpoint with event replication disabled.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "This template will create an Endpoint resource",
  "Resources": {
    "SampleEndpoint": {
      "Type": "AWS::Events::Endpoint",
      "Properties": {
        "Name": "CreateExampleEndpoint",
        "RoutingConfig": {
          "FailoverConfig": {
            "Primary": {
              "HealthCheck": {
                  "arn:aws:route53:::healthcheck/0123456789abc"
              }
            },
            "Secondary": {
              "Route": {
                  "us-east-1"
              }
            }
          }
        },
        "ReplicationConfig": {
          "State": "DISABLED"
        },
        "EventBuses": [
          {
            "EventBusArn": {
              "arn:aws:events:us-west-2:123456789012:event-bus/ExampleEventBus"
            }
          },
          {
            "EventBusArn": {
              "arn:aws:events:us-east-1:123456789012:event-bus/ExampleEventBus"
            }
          }
        ]
      }
    }
  },
  "Outputs": {
    "SampleEndpointName": {
      "Value": {
        "Ref": "SampleEndpoint"
      }
    }
  }
}
```

#### YAML

```yaml

  AWSTemplateFormatVersion: "2010-09-09"
  Description: "This template will create an Endpoint resource"
  Resources:
    SampleEndpoint:
      Type: AWS::Events::Endpoint
      Properties:
        Name: CreateExampleEndpoint
        RoutingConfig:
          FailoverConfig:
            Primary:
              HealthCheck: "arn:aws:route53:::healthcheck/0123456789abc"
            Secondary:
              Route: "us-east-1"
        ReplicationConfig:
          State: "DISABLED"
        EventBuses:
          - EventBusArn: "arn:aws:events:us-west-2:123456789012:event-bus/ExampleEventBus"
          - EventBusArn: "arn:aws:events:us-east-1:123456789012:event-bus/ExampleEventBus"

  Outputs:
    SampleEndpointName:
      Value: !Ref SampleEndpoint
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceParameters

EndpointEventBus

All content copied from https://docs.aws.amazon.com/.
