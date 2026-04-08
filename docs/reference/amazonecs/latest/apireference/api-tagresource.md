# TagResource

Associates the specified tags to a resource with the specified
`resourceArn`. If existing tags on a resource aren't specified in the
request parameters, they aren't changed. When a resource is deleted, the tags that are
associated with that resource are deleted as well.

## Request Syntax

```nohighlight

{
   "resourceArn": "string",
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[resourceArn](#API_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the resource to add tags to. Currently, the
supported resources are Amazon ECS capacity providers, tasks, services, task
definitions, clusters, and container instances.

In order to tag a service that has the following ARN format, you need to migrate the
service to the long ARN. For more information, see [Migrate an Amazon\
ECS short service ARN to a long ARN](../../../../services/amazonecs/latest/developerguide/service-arn-migration.md) in the _Amazon Elastic_
_Container Service Developer Guide_.

`arn:aws:ecs:region:aws_account_id:service/service-name`

After the migration is complete, the service has the long ARN format, as shown below.
Use this ARN to tag the service.

`arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name`

If you try to tag a service with a short ARN, you receive an
`InvalidParameterException` error.

Type: String

Required: Yes

**[tags](#API_TagResource_RequestSyntax)**

The tags to add to the resource. A tag is an array of key-value pairs.

The following basic restrictions apply to tags:

- Maximum number of tags per resource - 50

- For each resource, each tag key must be unique, and each tag key can have only
one value.

- Maximum key length - 128 Unicode characters in UTF-8

- Maximum value length - 256 Unicode characters in UTF-8

- If your tagging schema is used across multiple services and resources,
remember that other services may have restrictions on allowed characters.
Generally allowed characters are: letters, numbers, and spaces representable in
UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case-sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase
combination of such as a prefix for either keys or values as it is reserved for
AWS use. You cannot edit or delete tag keys or values with
this prefix. Tags with this prefix do not count against your tags per resource
limit.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ClientException**

These errors are usually caused by a client action. This client action might be using
an action or resource on behalf of a user that doesn't have permissions to use the
action or resource. Or, it might be specifying an identifier that isn't valid.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ClusterNotFoundException**

The specified cluster wasn't found. You can view your available clusters with [ListClusters](api-listclusters.md). Amazon ECS clusters are Region specific.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**InvalidParameterException**

The specified parameter isn't valid. Review the available parameters for the API
request.

For more information about service event errors, see [Amazon ECS\
service event messages](../../../../services/amazonecs/latest/developerguide/service-event-messages-list.md).

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource wasn't found.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server issue.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 500

## Examples

In the following example or examples, the Authorization header contents
( `AUTHPARAMS`) must be replaced with an AWS Signature
Version 4 signature. For more information, see [Signature\
Version 4 Signing Process](../../../../general/general/latest/gr/signature-version-4.md) in the _AWS_
_General Reference_.

You only need to learn how to sign HTTP requests if you intend to create them
manually. When you use the [AWS Command Line Interface](http://aws.amazon.com/cli) or one of the [AWS SDKs](http://aws.amazon.com/tools) to make requests to AWS, these tools automatically sign the requests for you, with the
access key that you specify when you configure the tools. When you use these tools,
you don't have to sign requests yourself.

### Example

This example tags the `test` service in the `dev`
cluster with key `team` and value `dev`.

#### Sample Request

```

POST / HTTP/1.1
Host: ecs.us-west-2.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerServiceV20141113.TagResource
Content-Type: application/x-amz-json-1.1
X-Amz-Date: 20241209T194744Z
Authorization: AUTHPARAMS
Content-Length: 115

{
   "resourceArn":"arn:aws:ecs:us-west-2:012345678910:service/dev/test",
   "tags":[
      {
         "key":"team",
         "value":"dev"
      }
   ]
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f
Content-Type: application/x-amz-json-1.1
Content-Length: 2
Date: Fri, 20 Dec 2024 20:01:34 GMT

{}
```

### Example

This example tags the `dev` cluster with key `team` and
value `dev` and the key `second-key` and value
`dev-key2`.

#### Sample Request

```

POST / HTTP/1.1
Host: ecs.us-west-2.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerServiceV20141113.TagResource
Content-Type: application/x-amz-json-1.1
X-Amz-Date: 20181026T194744Z
Authorization: AUTHPARAMS
Content-Length: 115

{
   "resourceArn":"arn:aws:ecs:us-west-2:012345678910:cluster/dev",
   "tags":[
      {
         "key":"team",
         "value":"dev"
       },
       {
        "key": "second-key",
        "value": "dev-key2"
      }
   ]
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f
Content-Type: application/x-amz-json-1.1
Content-Length: 2
Date: Wed, 19 Oct 2022 20:01:34 GMT

{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecs-2014-11-13/tagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecs-2014-11-13/tagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/tagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecs-2014-11-13/tagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/tagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecs-2014-11-13/tagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecs-2014-11-13/tagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecs-2014-11-13/tagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecs-2014-11-13/tagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/tagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SubmitTaskStateChange

UntagResource
