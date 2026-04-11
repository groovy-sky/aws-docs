# GetBucketReplication

###### Note

This operation gets an Amazon S3 on Outposts bucket's replication configuration. To get an
S3 bucket's replication configuration, see [GetBucketReplication](api-getbucketreplication.md)
in the _Amazon S3 API Reference_.

Returns the replication configuration of an S3 on Outposts bucket. For more information
about S3 on Outposts, see [Using Amazon S3 on Outposts](../userguide/s3onoutposts.md) in the
_Amazon S3 User Guide_. For information about S3 replication on Outposts
configuration, see [Replicating objects for\
S3 on Outposts](../userguide/s3outpostsreplication.md) in the _Amazon S3 User Guide_.

###### Note

It can take a while to propagate `PUT` or `DELETE` requests for
a replication configuration to all S3 on Outposts systems. Therefore, the replication
configuration that's returned by a `GET` request soon after a
`PUT` or `DELETE` request might return a more recent result
than what's on the Outpost. If an Outpost is offline, the delay in updating the
replication configuration on that Outpost can be significant.

This action requires permissions for the
`s3-outposts:GetReplicationConfiguration` action. The Outposts bucket owner
has this permission by default and can grant it to others. For more information about
permissions, see [Setting up IAM with\
S3 on Outposts](../userguide/s3outpostsiam.md) and [Managing access to\
S3 on Outposts bucket](../userguide/s3outpostsbucketpolicy.md) in the _Amazon S3 User Guide_.

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-getbucketreplication.md#API_control_GetBucketReplication_Examples) section.

If you include the `Filter` element in a replication configuration, you must
also include the `DeleteMarkerReplication`, `Status`, and
`Priority` elements. The response also returns those elements.

For information about S3 on Outposts replication failure reasons, see [Replication failure reasons](../userguide/outposts-replication-eventbridge.md#outposts-replication-failure-codes) in the _Amazon S3 User Guide_.

The following operations are related to `GetBucketReplication`:

- [PutBucketReplication](api-control-putbucketreplication.md)

- [DeleteBucketReplication](api-control-deletebucketreplication.md)

## Request Syntax

```nohighlight

GET /v20180820/bucket/name/replication HTTP/1.1
Host: Bucket.s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetBucketReplication_RequestSyntax)**

Specifies the bucket to get the replication information for.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the bucket accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/bucket/<my-bucket-name>`. For example, to access the bucket `reports` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports`. The value must be URL encoded.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_GetBucketReplication_RequestSyntax)**

The AWS account ID of the Outposts bucket.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetBucketReplicationResult>
   <ReplicationConfiguration>
      <Role>string</Role>
      <Rules>
         <Rule>
            <Bucket>string</Bucket>
            <DeleteMarkerReplication>
               <Status>string</Status>
            </DeleteMarkerReplication>
            <Destination>
               <AccessControlTranslation>
                  <Owner>string</Owner>
               </AccessControlTranslation>
               <Account>string</Account>
               <Bucket>string</Bucket>
               <EncryptionConfiguration>
                  <ReplicaKmsKeyID>string</ReplicaKmsKeyID>
               </EncryptionConfiguration>
               <Metrics>
                  <EventThreshold>
                     <Minutes>integer</Minutes>
                  </EventThreshold>
                  <Status>string</Status>
               </Metrics>
               <ReplicationTime>
                  <Status>string</Status>
                  <Time>
                     <Minutes>integer</Minutes>
                  </Time>
               </ReplicationTime>
               <StorageClass>string</StorageClass>
            </Destination>
            <ExistingObjectReplication>
               <Status>string</Status>
            </ExistingObjectReplication>
            <Filter>
               <And>
                  <Prefix>string</Prefix>
                  <Tags>
                     <S3Tag>
                        <Key>string</Key>
                        <Value>string</Value>
                     </S3Tag>
                  </Tags>
               </And>
               <Prefix>string</Prefix>
               <Tag>
                  <Key>string</Key>
                  <Value>string</Value>
               </Tag>
            </Filter>
            <ID>string</ID>
            <Prefix>string</Prefix>
            <Priority>integer</Priority>
            <SourceSelectionCriteria>
               <ReplicaModifications>
                  <Status>string</Status>
               </ReplicaModifications>
               <SseKmsEncryptedObjects>
                  <Status>string</Status>
               </SseKmsEncryptedObjects>
            </SourceSelectionCriteria>
            <Status>string</Status>
         </Rule>
      </Rules>
   </ReplicationConfiguration>
</GetBucketReplicationResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetBucketReplicationResult](#API_control_GetBucketReplication_ResponseSyntax)**

Root level tag for the GetBucketReplicationResult parameters.

Required: Yes

**[ReplicationConfiguration](#API_control_GetBucketReplication_ResponseSyntax)**

A container for one or more replication rules. A replication configuration must have at
least one rule and you can add up to 100 rules. The maximum size of a replication
configuration is 128 KB.

Type: [ReplicationConfiguration](api-control-replicationconfiguration.md) data type

## Examples

### Sample request to get the replication configuration of an Amazon S3 on Outposts bucket

The following example shows how to get the replication configuration of an
Outposts bucket.

```HTTP

GET /v20180820/bucket/example-outpost-bucket/replication HTTP/1.1
Host: s3-outposts.<Region>.amazonaws.com
x-amz-account-id: example-account-id
x-amz-outpost-id: op-01ac5d28a6a232904
Authorization: signatureValue

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getbucketreplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getbucketreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketPolicy

GetBucketTagging

All content copied from https://docs.aws.amazon.com/.
