This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container

The `AWS::Lightsail::Container` resource specifies a container
service.

A Lightsail container service is a compute resource to which you can
deploy containers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Container",
  "Properties" : {
      "ContainerServiceDeployment" : ContainerServiceDeployment,
      "IsDisabled" : Boolean,
      "Power" : String,
      "PrivateRegistryAccess" : PrivateRegistryAccess,
      "PublicDomainNames" : [ PublicDomainName, ... ],
      "Scale" : Integer,
      "ServiceName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Container
Properties:
  ContainerServiceDeployment:
    ContainerServiceDeployment
  IsDisabled: Boolean
  Power: String
  PrivateRegistryAccess:
    PrivateRegistryAccess
  PublicDomainNames:
    - PublicDomainName
  Scale: Integer
  ServiceName: String
  Tags:
    - Tag

```

## Properties

`ContainerServiceDeployment`

An object that describes the current container deployment of the container
service.

_Required_: No

_Type_: [ContainerServiceDeployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-container-containerservicedeployment.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`IsDisabled`

A Boolean value indicating whether the container service is disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Power`

The power specification of the container service.

The power specifies the amount of RAM, the number of vCPUs, and the base price of the
container service.

_Required_: Yes

_Type_: String

_Allowed values_: `nano | micro | small | medium | large | xlarge`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateRegistryAccess`

An object that describes the configuration for the container service to access private
container image repositories, such as Amazon Elastic Container Registry (Amazon ECR)
private repositories.

For more information, see [Configuring access to an Amazon ECR private repository for an Amazon Lightsail\
container service](https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-container-service-ecr-private-repo-access) in the _Amazon Lightsail Developer_
_Guide_.

_Required_: No

_Type_: [PrivateRegistryAccess](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-container-privateregistryaccess.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicDomainNames`

The public domain name of the container service, such as `example.com` and
`www.example.com`.

You can specify up to four public domain names for a container service. The domain names
that you specify are used when you create a deployment with a container that is configured as the
public endpoint of your container service.

If you don't specify public domain names, then you can use the default domain of the
container service.

###### Important

You must create and validate an SSL/TLS certificate before you can use public domain
names with your container service. Use the [AWS::Lightsail::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-certificate.html) resource to create a certificate for the public
domain names that you want to use with your container service.

_Required_: No

_Type_: Array of [PublicDomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-container-publicdomainname.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scale`

The scale specification of the container service.

The scale specifies the allocated compute nodes of the container service.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

The name of the container service.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9]{1,2}|[a-z0-9][a-z0-9-]+[a-z0-9]$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-container-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ContainerArn`

The Amazon Resource Name (ARN) of the container.

`PrincipalArn`

The principle Amazon Resource Name (ARN) of the role.

`PrivateRegistryAccess.EcrImagePullerRole.PrincipalArn`

The principle Amazon Resource Name (ARN) of the role.

`Url`

The publicly accessible URL of the container service.

If no public endpoint is specified in the current deployment, this URL returns a 404
response.

## Remarks

_Container deployment failures_

The stack will drift if the deployment to your `ContainerService` fails.
This is because the template will have a different deployment compared to the container service.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Container
