# ListStacks

Returns the summary information for stacks whose status matches the specified
`StackStatusFilter`. Summary information for stacks that have been deleted is
kept for 90 days after the stack is deleted. If no `StackStatusFilter` is
specified, summary information for all stacks is returned (including existing stacks and
stacks that have been deleted).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**StackStatusFilter.member.N**

Stack status to use as a filter. Specify one or more stack status codes to list only
stacks with the specified status codes. For a complete list of stack status codes, see the
`StackStatus` parameter of the [Stack](api-stack.md) data type.

Type: Array of strings

Valid Values: `CREATE_IN_PROGRESS | CREATE_FAILED | CREATE_COMPLETE | ROLLBACK_IN_PROGRESS | ROLLBACK_FAILED | ROLLBACK_COMPLETE | DELETE_IN_PROGRESS | DELETE_FAILED | DELETE_COMPLETE | UPDATE_IN_PROGRESS | UPDATE_COMPLETE_CLEANUP_IN_PROGRESS | UPDATE_COMPLETE | UPDATE_FAILED | UPDATE_ROLLBACK_IN_PROGRESS | UPDATE_ROLLBACK_FAILED | UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS | UPDATE_ROLLBACK_COMPLETE | REVIEW_IN_PROGRESS | IMPORT_IN_PROGRESS | IMPORT_COMPLETE | IMPORT_ROLLBACK_IN_PROGRESS | IMPORT_ROLLBACK_FAILED | IMPORT_ROLLBACK_COMPLETE`

Required: No

## Response Elements

The following elements are returned by the service.

**NextToken**

If the output exceeds 1 MB in size, a string that identifies the next page of stacks. If
no additional page exists, this value is null.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**StackSummaries.member.N**

A list of `StackSummary` structures that contains information about the
specified stacks.

Type: Array of [StackSummary](api-stacksummary.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### ListStacks

This example illustrates one usage of ListStacks.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=ListStacks
 &StackStatusFilter.member.1=CREATE_IN_PROGRESS
 &StackStatusFilter.member.2=DELETE_COMPLETE
 &Version=2010-05-15
 &SignatureVersion=2
 &SignatureMethod=HmacSHA256
 &Timestamp=2010-07-27T22%3A26%3A28.000Z
 &AWSAccessKeyId=[AWS Access KeyID]
 &Signature=[Signature]
```

#### Sample Response

```

<ListStacksResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <ListStacksResult>
    <StackSummaries>
      <member>
        <StackId>
          arn:aws:cloudformation:us-east-1:1234567:stack/TestCreate1/aaaaa
        </StackId>
        <StackStatus>CREATE_IN_PROGRESS</StackStatus>
        <StackName>vpc1</StackName>
        <CreationTime>2011-05-23T15:47:44Z</CreationTime>
        <TemplateDescription>
          Creates one EC2 instance and a load balancer.
        </TemplateDescription>
        <ResourceTypes>
          <member>AWS::EC2::Instance</member>
          <member>AWS::ElasticLoadBalancing::LoadBalancer</member>
        </ResourceTypes>
      </member>
      <member>
        <StackId>
          arn:aws:cloudformation:us-east-1:1234567:stack/TestDelete2/bbbbb
        </StackId>
        <StackStatus>DELETE_COMPLETE</StackStatus>
        <DeletionTime>2011-03-10T16:20:51Z</DeletionTime>
        <StackName>WP1</StackName>
        <CreationTime>2011-03-05T19:57:58Z</CreationTime>
        <TemplateDescription>
          A simple basic CloudFormation Template.
        </TemplateDescription>
        <ResourceTypes>
          <member>AWS::EC2::Instance</member>
        </ResourceTypes>
      </member>
    </StackSummaries>
  </ListStacksResult>
  <ResponseMetadata>
    <RequestId>b9b4b068-3a41-11e5-94eb-example</RequestId>
  </ResponseMetadata>
</ListStacksResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/liststacks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/liststacks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListStackResources

ListStackSetAutoDeploymentTargets
