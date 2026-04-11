# Limitations and restrictions

Amazon S3 Vectors has certain limitations and restrictions that you should be aware of when planning your vector storage and search applications.

- Vector buckets per AWS Region in an account: 10,000

- Vector indexes per vector bucket: 10,000

- Vectors per vector index: Up to 2 billion

- Dimension value per vector: 1 to 4,096

- Total metadata per vector: Up to 40 KB (filterable + non-filterable)

- Total metadata keys per vector: Up to 50

- Filterable metadata per vector: Up to 2 KB

- Non-filterable metadata keys per vector index: Up to 10

- Combined PutVectors and DeleteVectors requests per second per vector index: Up to 1,000

- Combined vectors inserted and deleted per second per vector index: Up to 2,500

- Request payload size: Up to 20 MiB

- Vectors per [PutVectors](../api/api-s3vectorbuckets-putvectors.md) API call: Up to 500

- Vectors per [DeleteVectors](../api/api-s3vectorbuckets-deletevectors.md) API call: Up to 500

- Vectors per [GetVectors](../api/api-s3vectorbuckets-getvectors.md) API call: Up to 100

- Top-K results per [QueryVectors](../api/api-s3vectorbuckets-queryvectors.md) request: Up to 100

- Vectors listed per page in a [ListVectors](../api/api-s3vectorbuckets-listvectors.md) response: Up to 1,000

- Vector buckets listed per page in a [ListVectorBuckets](../api/api-s3vectorbuckets-listvectorbuckets.md) response: Up to 500.

- Vector indexes listed per page in a [ListIndexes](../api/api-s3vectorbuckets-listindexes.md) response: Up to 500.

- Segment count for parallel listing in a ListVectors API call: Up to 16

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Metadata filtering

S3 Vectors best practices

All content copied from https://docs.aws.amazon.com/.
