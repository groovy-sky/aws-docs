This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaStore::Container

The AWS::MediaStore::Container resource specifies a storage container to hold objects. A container is similar to a bucket in Amazon S3.

When you create a container using CloudFormation, the template manages data for five API actions: creating a container, setting access
logging, updating the default container policy, adding a cross-origin resource sharing (CORS) policy, and adding an object lifecycle policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaStore::Container",
  "Properties" : {
      "AccessLoggingEnabled" : Boolean,
      "ContainerName" : String,
      "CorsPolicy" : [ CorsRule, ... ],
      "LifecyclePolicy" : String,
      "MetricPolicy" : MetricPolicy,
      "Policy" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaStore::Container
Properties:
  AccessLoggingEnabled: Boolean
  ContainerName: String
  CorsPolicy:
    - CorsRule
  LifecyclePolicy: String
  MetricPolicy:
    MetricPolicy
  Policy: String
  Tags:
    - Tag

```

## Properties

`AccessLoggingEnabled`

The state of access logging on the container. This value is `false` by default, indicating that AWS Elemental MediaStore does not send access logs to Amazon CloudWatch Logs. When you enable access logging on the container, MediaStore changes this value to `true`, indicating that the service delivers access logs for objects stored in that container to CloudWatch Logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerName`

The name for the container. The name must be from 1 to 255 characters. Container
names must be unique to your AWS account within a specific region. As an example, you could
create a container named `movies` in every region, as long as you don’t have an
existing container with that name.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CorsPolicy`

###### Important

End of support notice: On November 13, 2025, AWS
will discontinue support for AWS Elemental MediaStore. After November 13, 2025, you will
no longer be able to access the AWS Elemental MediaStore console or AWS Elemental MediaStore resources.
For more information, visit this
[blog post](https://aws.amazon.com/blogs/media/support-for-aws-elemental-mediastore-ending-soon).

Sets the cross-origin resource sharing (CORS) configuration on a container so that
the container can service cross-origin requests. For example, you might want to enable a
request whose origin is http://www.example.com to access your AWS Elemental MediaStore
container at my.example.container.com by using the browser's XMLHttpRequest
capability.

To enable CORS on a container, you attach a CORS policy to the container. In the CORS
policy, you configure rules that identify origins and the HTTP methods that can be executed
on your container. The policy can contain up to 398,000 characters. You can add up to 100
rules to a CORS policy. If more than one rule applies, the service uses the first
applicable rule listed.

To learn more about CORS, see [Cross-Origin Resource Sharing (CORS) in AWS Elemental MediaStore](../../../mediastore/latest/ug/cors-policy.md).

_Required_: No

_Type_: Array of [CorsRule](aws-properties-mediastore-container-corsrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecyclePolicy`

###### Important

End of support notice: On November 13, 2025, AWS
will discontinue support for AWS Elemental MediaStore. After November 13, 2025, you will
no longer be able to access the AWS Elemental MediaStore console or AWS Elemental MediaStore resources.
For more information, visit this
[blog post](https://aws.amazon.com/blogs/media/support-for-aws-elemental-mediastore-ending-soon).

Writes an object lifecycle policy to a container. If the container already has an object lifecycle policy, the service replaces the existing policy with the new policy. It takes up to 20 minutes for the change to take effect.

For information about how to construct an object lifecycle policy, see [Components of an Object Lifecycle Policy](../../../mediastore/latest/ug/policies-object-lifecycle-components.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricPolicy`

The metric policy that is associated with the container. A metric policy allows AWS Elemental MediaStore to send metrics to Amazon CloudWatch. In the policy, you must indicate whether you want MediaStore to send container-level metrics. You can also include rules to define groups of objects that you want MediaStore to send object-level metrics for.

To view examples of how to construct a metric policy for your use case, see [Example Metric Policies](../../../mediastore/latest/ug/policies-metric-examples.md).

_Required_: No

_Type_: [MetricPolicy](aws-properties-mediastore-container-metricpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policy`

Creates an access policy for the specified container to restrict the users and
clients that can access it. For information about the data that is included in an access
policy, see the [AWS Identity and\
Access Management User Guide](../../../iam/index.md).

For this release of the REST API, you can create only one policy for a container. If
you enter `PutContainerPolicy` twice, the second command modifies the existing
policy.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediastore-container-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the container.

For example: `{ "Ref": "myContainer" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Endpoint`

The DNS endpoint of the container. Use the endpoint to identify the specific
container when sending requests to the data plane. The service assigns this value when the
container is created. Once the value has been assigned, it does not change.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Elemental MediaStore

CorsRule

All content copied from https://docs.aws.amazon.com/.
