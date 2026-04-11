This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Route

Creates an AWS Migration Hub Refactor Spaces route. The account owner of
the service resource is always the environment owner, regardless of which account creates the
route. Routes target a service in the application. If an application does not have any routes,
then the first route must be created as a `DEFAULT` `RouteType`.

When created, the default route defaults to an active state so state is not a required
input. However, like all other state values the state of the default route can be updated
after creation, but only when all other routes are also inactive. Conversely, no route can be
active without the default route also being active.

###### Note

In the `AWS::RefactorSpaces::Route` resource, you can only update the
`ActivationState` property, which resides under the `UriPathRoute`
and `DefaultRoute` properties. All other properties associated with the
`AWS::RefactorSpaces::Route` cannot be updated, even though the property
description might indicate otherwise. Updating all other properties will result in the
replacement of Route.

When you create a route, Refactor Spaces configures the Amazon API Gateway to send
traffic to the target service as follows:

- **URL Endpoints**

If the service has a URL endpoint, and the endpoint resolves to a private IP address,
Refactor Spaces routes traffic using the API Gateway VPC link. If a service
endpoint resolves to a public IP address, Refactor Spaces routes traffic over the public
internet. Services can have HTTP or HTTPS URL endpoints. For HTTPS URLs, publicly-signed
certificates are supported. Private Certificate Authorities (CAs) are permitted only if
the CA's domain is also publicly resolvable.

Refactor Spaces automatically resolves the public Domain Name System (DNS) names that
are set in `CreateService:UrlEndpoint ` when you create a service. The DNS names
resolve when the DNS time-to-live (TTL) expires, or every 60 seconds for TTLs less than 60
seconds. This periodic DNS resolution ensures that the route configuration remains
up-to-date.

**One-time health check**

A one-time health check is performed on the service when either the route is updated
from inactive to active, or when it is created with an active state. If the health check
fails, the route transitions the route state to `FAILED`, an error code of
`SERVICE_ENDPOINT_HEALTH_CHECK_FAILURE` is provided, and no traffic is sent
to the service.

For private URLs, a target group is created on the Network Load Balancer and the load
balancer target group runs default target health checks. By default, the health check is
run against the service endpoint URL. Optionally, the health check can be performed
against a different protocol, port, and/or path using the [CreateService:UrlEndpoint](../../../../reference/migrationhub-refactor-spaces/latest/apireference/api-createservice.md#migrationhubrefactorspaces-CreateService-request-UrlEndpoint) parameter. All other health check settings for the
load balancer use the default values described in the [Health\
checks for your target groups](../../../elasticloadbalancing/latest/application/target-group-health-checks.md) in the _Elastic Load Balancing_
_guide_. The health check is considered successful if at least one target
within the target group transitions to a healthy state.

- **AWS Lambda function endpoints**

If the service has an AWS Lambda function endpoint, then Refactor Spaces
configures the Lambda function's resource policy to allow the application's
API Gateway to invoke the function.

The Lambda function state is checked. If the function is not active, the
function configuration is updated so that Lambda resources are provisioned. If
the Lambda state is `Failed`, then the route creation fails. For
more information, see the [GetFunctionConfiguration's State response parameter](../../../lambda/latest/dg/api-getfunctionconfiguration.md#SSS-GetFunctionConfiguration-response-State) in the _AWS Lambda Developer Guide_.

A check is performed to determine that a Lambda function with the specified ARN
exists. If it does not exist, the health check fails. For public URLs, a connection is
opened to the public endpoint. If the URL is not reachable, the health check fails.

**Environments without a network bridge**

When you create environments without a network bridge ( [CreateEnvironment:NetworkFabricType](../../../../reference/migrationhub-refactor-spaces/latest/apireference/api-createenvironment.md#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType) is `NONE)` and you use your own
networking infrastructure, you need to configure [VPC to VPC connectivity](../../../whitepapers/latest/aws-vpc-connectivity-options/amazon-vpc-to-amazon-vpc-connectivity-options.md) between your network and the application proxy VPC. Route
creation from the application proxy to service endpoints will fail if your network is not
configured to connect to the application proxy VPC. For more information, see [Create\
a route](../../../migrationhub-refactor-spaces/latest/userguide/getting-started-create-role.md) in the _Refactor Spaces User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RefactorSpaces::Route",
  "Properties" : {
      "ApplicationIdentifier" : String,
      "DefaultRoute" : DefaultRouteInput,
      "EnvironmentIdentifier" : String,
      "RouteType" : String,
      "ServiceIdentifier" : String,
      "Tags" : [ Tag, ... ],
      "UriPathRoute" : UriPathRouteInput
    }
}

```

### YAML

```yaml

Type: AWS::RefactorSpaces::Route
Properties:
  ApplicationIdentifier: String
  DefaultRoute:
    DefaultRouteInput
  EnvironmentIdentifier: String
  RouteType: String
  ServiceIdentifier: String
  Tags:
    - Tag
  UriPathRoute:
    UriPathRouteInput

```

## Properties

`ApplicationIdentifier`

The unique identifier of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^app-([0-9A-Za-z]{10}$)`

_Minimum_: `14`

_Maximum_: `14`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultRoute`

Configuration for the default route type.

_Required_: No

_Type_: [DefaultRouteInput](aws-properties-refactorspaces-route-defaultrouteinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentIdentifier`

The unique identifier of the environment.

_Required_: Yes

_Type_: String

_Pattern_: `^env-([0-9A-Za-z]{10}$)`

_Minimum_: `14`

_Maximum_: `14`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteType`

The route type of the route.

_Required_: Yes

_Type_: String

_Allowed values_: `DEFAULT | URI_PATH`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceIdentifier`

The unique identifier of the service.

_Required_: Yes

_Type_: String

_Pattern_: `^svc-([0-9A-Za-z]{10}$)`

_Minimum_: `14`

_Maximum_: `14`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the route.

_Required_: No

_Type_: Array of [Tag](aws-properties-refactorspaces-route-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UriPathRoute`

The configuration for the URI path route type.

_Required_: No

_Type_: [UriPathRouteInput](aws-properties-refactorspaces-route-uripathrouteinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a composite ID following this format:
`<EnvironmentId>|<ApplicationId>|<RouteId>`, for example,
`env-1234654123|app-1234654123|rte-1234654123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the route.

`PathResourceToId`

A mapping of Amazon API Gateway path resources to resource IDs.

`RouteIdentifier`

The unique identifier of the route.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DefaultRouteInput

All content copied from https://docs.aws.amazon.com/.
