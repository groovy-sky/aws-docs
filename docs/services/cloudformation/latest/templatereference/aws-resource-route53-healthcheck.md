This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::HealthCheck

The `AWS::Route53::HealthCheck` resource is a Route 53 resource type that contains settings for
a Route 53 health check.

For information about associating health checks with records, see
[HealthCheckId](../../../../reference/route53/latest/apireference/api-resourcerecordset.md#Route53-Type-ResourceRecordSet-HealthCheckId)
in
[ChangeResourceRecordSets](../../../../reference/route53/latest/apireference/api-changeresourcerecordsets.md).

###### Note

You can't create a health check with simple routing.

**ELB Load Balancers**

If you're registering EC2 instances with an Elastic Load Balancing (ELB) load balancer, do not create Amazon Route 53 health checks for the
EC2 instances. When you register an EC2 instance with a load balancer, you configure settings for an ELB health check, which performs a
similar function to a Route 53 health check.

**Private Hosted Zones**

You can associate health checks with failover records in a private hosted zone. Note the following:

- Route 53 health checkers are outside the VPC. To check the health of an endpoint within a VPC by IP address, you must
assign a public IP address to the instance in the VPC.

- You can configure a health checker to check the health of an external resource that the instance relies on, such as a
database server.

- You can create a CloudWatch metric, associate an alarm with the metric, and then create a health check that is based on the
state of the alarm. For example, you might create a CloudWatch metric that checks the status of the Amazon EC2 `StatusCheckFailed` metric,
add an alarm to the metric, and then create a health check that is based on the state of the alarm. For information about creating
CloudWatch metrics and alarms by using the CloudWatch console, see the [Amazon CloudWatch User Guide](../../../amazoncloudwatch/latest/developerguide/whatiscloudwatch.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53::HealthCheck",
  "Properties" : {
      "HealthCheckConfig" : HealthCheckConfig,
      "HealthCheckTags" : [ HealthCheckTag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53::HealthCheck
Properties:
  HealthCheckConfig:
    HealthCheckConfig
  HealthCheckTags:
    - HealthCheckTag

```

## Properties

`HealthCheckConfig`

A complex type that contains detailed information about one health check.

For the values to enter for `HealthCheckConfig`, see
[HealthCheckConfig](../../../../reference/route53/latest/apireference/api-healthcheckconfig.md)

_Required_: Yes

_Type_: [HealthCheckConfig](aws-properties-route53-healthcheck-healthcheckconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckTags`

The `HealthCheckTags` property describes key-value pairs that are associated with an `AWS::Route53::HealthCheck` resource.

_Required_: No

_Type_: Array of [HealthCheckTag](aws-properties-route53-healthcheck-healthchecktag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the health check ID, such as `e0a123b4-4dba-4650-935e-example`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`HealthCheckId`

The identifier that Amazon Route 53 assigned to the health check when you created it. When you add or update a resource record set, you use this value to specify which health check to use. The value can be up to 64 characters long.

## Examples

### Create health check

The following example creates an Amazon Route 53 health check that sends HTTP requests to the specified endpoint.

#### JSON

```json

{
   "myHealthCheck": {
      "Type": "AWS::Route53::HealthCheck",
      "Properties": {
         "HealthCheckConfig": {
            "IPAddress": "192.0.2.44",
            "Port": "80",
            "Type": "HTTP",
            "ResourcePath": "/example/index.html",
            "FullyQualifiedDomainName": "example.com",
            "RequestInterval": "30",
            "FailureThreshold": "3"
         },
         "HealthCheckTags": [
            {
               "Key": "SampleKey1",
               "Value": "SampleValue1"
            },
            {
               "Key": "SampleKey2",
               "Value": "SampleValue2"
            }
         ]
      }
   }
}
```

#### YAML

```yaml

myHealthCheck:
  Type: 'AWS::Route53::HealthCheck'
  Properties:
    HealthCheckConfig:
      IPAddress: 192.0.2.44
      Port: 80
      Type: HTTP
      ResourcePath: '/example/index.html'
      FullyQualifiedDomainName: example.com
      RequestInterval: 30
      FailureThreshold: 3
    HealthCheckTags:
      -
        Key: SampleKey1
        Value: SampleValue1
      -
        Key: SampleKey2
        Value: SampleValue2
```

## See also

- [CreateHealthCheck](../../../../reference/route53/latest/apireference/api-createhealthcheck.md) in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Route53::DNSSEC

AlarmIdentifier

All content copied from https://docs.aws.amazon.com/.
