This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancing::LoadBalancer AccessLoggingPolicy

Specifies where and how access logs are stored for your Classic Load Balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmitInterval" : Integer,
  "Enabled" : Boolean,
  "S3BucketName" : String,
  "S3BucketPrefix" : String
}

```

### YAML

```yaml

  EmitInterval: Integer
  Enabled: Boolean
  S3BucketName: String
  S3BucketPrefix: String

```

## Properties

`EmitInterval`

The interval for publishing the access logs. You can specify an interval of either 5 minutes or 60 minutes.

Default: 60 minutes

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether access logs are enabled for the load balancer.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

The name of the Amazon S3 bucket where the access logs are stored.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketPrefix`

The logical hierarchy you created for your Amazon S3 bucket, for example `my-bucket-prefix/prod`.
If the prefix is not provided, the log is placed at the root level of the bucket.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ModifyLoadBalancerAttributes](https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_ModifyLoadBalancerAttributes.html)
in the _Elastic Load Balancing API Reference (version 2012-06-01)_

- [Access Logs](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/access-log-collection.html)
in the _User Guide for Classic Load Balancers_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElasticLoadBalancing::LoadBalancer

AppCookieStickinessPolicy
