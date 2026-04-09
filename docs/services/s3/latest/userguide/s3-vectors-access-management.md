# Identity and Access management in S3 Vectors

Access management in S3 Vectors follows AWS security best practices, providing multiple
layers of control to ensure that only authorized users and applications can access your
vector data. The service integrates with IAM and supports both identity-based and
resource-based policies, giving you flexibility in how you structure and manage permissions
across your organization.

## Authenticating and authorizing requests

S3 Vectors uses AWS standard authentication and authorization mechanisms to secure
access to vector buckets and their contents. Every request to S3 Vectors must be
authenticated using valid AWS credentials, and the service evaluates permissions based
on the combination of identity-based policies, resource-based policies, and any
applicable service control policies.

The authentication process begins when a client makes a request to S3 Vectors using
AWS credentials (access keys, temporary credentials from AWS STS, or IAM roles). The
service validates these credentials and then evaluates the permissions associated with
the authenticated identity against the requested action and target resource. This
evaluation process considers multiple policy types and applies the principle of least
privilege to determine whether the request should be allowed or denied.

Authorization in S3 Vectors operates at multiple levels of granularity. You can
control access at the vector bucket level, individual vector index level, or even
specific operations within an index. This hierarchical permission model allows you to
implement sophisticated access control schemes that align with your organizational
structure and data governance requirements.

## Resource types defined for vector buckets

S3 Vectors defines specific resource types that can be referenced in IAM policies and
resource-based policies. Understanding these resource types is essential for creating
effective access control policies that provide the right level of access to the right
users and applications.

The following table describes the resource types available in S3 Vectors.

Resource types available in S3 VectorsResource typeARN formatDescriptionVectorBucketarn:aws:s3vectors: `region`: `123456789012`:bucket/ `bucket-name`Represents a vector bucket and is used for bucket-level operations
such as creating, deleting, or configuring the bucketIndexarn:aws:s3vectors: `region`: `123456789012`:bucket/ `bucket-name`/index/ `index-name`Represents a vector index within a bucket and is used for
index-specific operations such as querying vectors or managing index
contents

## Policy actions for vector buckets

S3 Vectors provides a comprehensive set of policy actions that correspond to the
various operations you can perform on vector buckets and indexes. These actions are
designed to provide fine-grained control over who can perform specific operations,
allowing you to implement the principle of least privilege effectively.

The following table lists all available policy actions for S3 Vectors
resources.

Policy actions for S3 Vectors resourcesResource typeAPI operationsPolicy actionsDescription of policy actionsAccess levelCondition keysAccount[ListVectorBuckets](../api/api-s3vectorbuckets-listvectorbuckets.md)s3vectors:ListVectorBucketsGrants permission to list all vector buckets in the account and
regionListVectorBucket[CreateVectorBucket](../api/api-s3vectorbuckets-createvectorbucket.md)s3vectors:CreateVectorBucketGrants permission to create a new vector bucket with specified
configurationWrites3vectors:sseType, s3vectors:kmsKeyArnVectorBucket[GetVectorBucket](../api/api-s3vectorbuckets-getvectorbucket.md)s3vectors:GetVectorBucketGrants permission to retrieve vector bucket attributes and
configurationReadVectorBucket[DeleteVectorBucket](../api/api-s3vectorbuckets-deletevectorbucket.md)s3vectors:DeleteVectorBucketGrants permission to delete an empty vector bucketWriteVectorBucket[ListIndexes](../api/api-s3vectorbuckets-listindexes.md)s3vectors:ListIndexesGrants permission to list all indexes within a
vector bucketListVectorBucket[PutVectorBucketPolicy](../api/api-s3vectorbuckets-putvectorbucketpolicy.md)s3vectors:PutVectorBucketPolicyGrants permission to apply or update a resource-based policy on a
vector bucketPermissions managementVectorBucket[GetVectorBucketPolicy](../api/api-s3vectorbuckets-getvectorbucketpolicy.md)s3vectors:GetVectorBucketPolicyGrants permission to retrieve the resource-based policy attached to a
vector bucketReadVectorBucket[DeleteVectorBucketPolicy](../api/api-s3vectorbuckets-deletevectorbucketpolicy.md)s3vectors:DeleteVectorBucketPolicyGrants permission to remove the resource-based policy from a
vector bucketPermissions managementIndex[CreateIndex](../api/api-s3vectorbuckets-createindex.md)s3vectors:CreateIndexGrants permission to create a new vector index with specified
dimensions and metadata configurationWriteIndex[GetIndex](../api/api-s3vectorbuckets-getindex.md)s3vectors:GetIndexGrants permission to retrieve vector index attributes and
configurationReadIndex[DeleteIndex](../api/api-s3vectorbuckets-deleteindex.md)s3vectors:DeleteIndexGrants permission to delete a vector index and all its
contentsWriteIndex[QueryVectors](../api/api-s3vectorbuckets-queryvectors.md)(Required) s3vectors:QueryVectors

Grants permission to perform similarity queries against
vectors in an index.

**With `s3vectors:QueryVectors`**
**only**, you can retrieve vector keys of approximate
nearest neighbors and their computed distances from the query
vector. This permission is sufficient only when you don't set any
metadata filters and don't request metadata (by
keeping the `returnMetadata`
parameter set to false or not specified).

Read(Conditionally required): s3vectors:GetVectors

Required if you set metadata filters, set `returnMetadata` to true in your request.

**With both `s3vectors:QueryVectors`**
**and `s3vectors:GetVectors`**, you can filter
results by using metadata criteria and retrieve vector keys along with
their associated data, metadata, and computed distances from the query
vector.

ReadIndex[PutVectors](../api/api-s3vectorbuckets-putvectors.md)s3vectors:PutVectorsGrants permission to add or update vectors in an indexWriteIndex[GetVectors](../api/api-s3vectorbuckets-getvectors.md)s3vectors:GetVectorsGrants permission to retrieve specific vectors and their metadata
by vector keyReadIndex[ListVectors](../api/api-s3vectorbuckets-listvectors.md)(Required) s3vectors:ListVectors

Grants permission to list vector keys in an index.

**With `s3vectors:ListVectors` only**,
you can list vector keys when both `returnData` and
`returnMetadata` parameters are false or not
specified.

Read(Conditionally required): s3vectors:GetVectors

Required if you set either `returnData` or
`returnMetadata` parameter to true in your request.

**With both `s3vectors:ListVectors`**
**and `s3vectors:GetVectors`**, you can retrieve
vector keys along with their associated data and metadata by setting
`returnData` and `returnMetadata` to
true.

ReadIndex[DeleteVectors](../api/api-s3vectorbuckets-deletevectors.md)s3vectors:DeleteVectorsGrants permission to delete specific vectors from an indexWrite

These actions can be combined in various ways to create policies that match your
specific access requirements. For example, you might create a read-only policy that
includes `s3vectors:GetVectorBucket`, `s3vectors:ListIndexes`,
`s3vectors:QueryVectors`, and `s3vectors:GetVectors` actions,
or a policy that includes query and vector retrieval permissions but excludes
administrative actions like creating or deleting indexes.

## Condition keys for vector buckets

Condition keys for vector bucketsCondition keysDescriptionType1s3vectors:sseTypeFilters access by server-side encryption type **Valid values:** `AES256 | aws:kms`String2s3vectors:kmsKeyArnFilters access by the AWS AWS KMS key ARN for the key used to encrypt
a vector bucketARN

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Security in S3 Vectors

S3 Vectors identity-based policy examples

All content copied from https://docs.aws.amazon.com/.
