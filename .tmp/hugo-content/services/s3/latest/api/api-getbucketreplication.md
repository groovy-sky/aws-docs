# GetBucketReplication

###### Note

This operation is not supported for directory buckets.

Returns the replication configuration of a bucket.

###### Note

It can take a while to propagate the put or delete a replication configuration to all Amazon S3
systems. Therefore, a get request soon after put or delete can return a wrong result.

For information about replication configuration, see [Replication](../dev/replication.md) in the
_Amazon S3 User Guide_.

This action requires permissions for the `s3:GetReplicationConfiguration` action. For
more information about permissions, see [Using Bucket Policies and User\
Policies](../dev/using-iam-policies.md).

If you include the `Filter` element in a replication configuration, you must also include
the `DeleteMarkerReplication` and `Priority` elements. The response also returns
those elements.

For information about `GetBucketReplication` errors, see [List of replication-related\
error codes](errorresponses.md#ReplicationErrorCodeList)

The following operations are related to `GetBucketReplication`:

- [PutBucketReplication](api-putbucketreplication.md)

- [DeleteBucketReplication](api-deletebucketreplication.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?replication HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketReplication_RequestSyntax)**

The bucket name for which to get the replication information.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketReplication_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ReplicationConfiguration>
   <Role>string</Role>
   <Rule>
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
            <Tag>
               <Key>string</Key>
               <Value>string</Value>
            </Tag>
            ...
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
   ...
</ReplicationConfiguration>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ReplicationConfiguration](#API_GetBucketReplication_ResponseSyntax)**

Root level tag for the ReplicationConfiguration parameters.

Required: Yes

**[Role](#API_GetBucketReplication_ResponseSyntax)**

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating
objects. For more information, see [How to Set Up Replication](../dev/replication-how-setup.md) in the
_Amazon S3 User Guide_.

Type: String

**[Rule](#API_GetBucketReplication_ResponseSyntax)**

A container for one or more replication rules. A replication configuration must have at least one
rule and can contain a maximum of 1,000 rules.

Type: Array of [ReplicationRule](api-replicationrule.md) data types

## Examples

### Sample Request: Retrieve replication configuration information

The following GET request retrieves information about the replication configuration set for the
`amzn-s3-demo-bucket` bucket:

```

            GET /?replication HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Date: Tue, 10 Feb 2015 00:17:21 GMT
            Authorization: authorization string

```

### Sample Response

The following response shows that replication is enabled on the bucket. The empty prefix
indicates that Amazon S3 will replicate all objects that are created in the
`amzn-s3-demo-bucket` bucket. The `Destination` element identifies the
target bucket where Amazon S3 creates the object replicas, and the storage class (STANDARD\_IA) that Amazon S3
uses when creating replicas.

Amazon S3 assumes the specified IAM role to replicate objects on behalf of the bucket owner, which
is the AWS account that created the bucket.

```

            HTTP/1.1 200 OK
            x-amz-id-2: ITnGT1y4RyTmXa3rPi4hklTXouTf0hccUjo0iCPjz6FnfIutBj3M7fPGlWO2SEWp
            x-amz-request-id: 51991C342example
            Date: Tue, 10 Feb 2015 00:17:23 GMT
            Server: AmazonS3
            Content-Length: contentlength

            <?xml version="1.0" encoding="UTF-8"?>
            <ReplicationConfiguration>
              <Role>arn:aws:iam::35667example:role/CrossRegionReplicationRoleForS3</Role>
             <Rule>
               <ID>rule1</ID>
               <Status>Enabled</Status>
               <Priority>1</Priority>
               <DeleteMarkerReplication>
                   <Status>Disabled</Status>
               </DeleteMarkerReplication>
               <Filter>
                  <And>
                       <Prefix>TaxDocs</Prefix>
                       <Tag>
                         <Key>key1</Key>
                         <Value>value1</Value>
                       </Tag>
                       <Tag>
                         <Key>key1</Key>
                        <Value>value1</Value>
                      </Tag>
                  </And>
                </Filter>
               <Destination>
                  <Bucket>arn:aws:s3:::exampletargetbucket</Bucket>
               </Destination>
              </Rule>
            </ReplicationConfiguration>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/getbucketreplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/getbucketreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketPolicyStatus

GetBucketRequestPayment

All content copied from https://docs.aws.amazon.com/.
