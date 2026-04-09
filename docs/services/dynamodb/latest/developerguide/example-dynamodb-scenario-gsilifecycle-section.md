# Manage DynamoDB Global Secondary Indexes using AWS Command Line Interface v2

The following code example shows how to manage the complete lifecycle of Global Secondary Indexes.

- Create a table with a Global Secondary Index.

- Add a new GSI to an existing table.

- Update (increase) GSI warm throughput.

- Query data using GSIs.

- Delete a GSI.

Bash

**AWS CLI with Bash script**

Create a table with a Global Secondary Index.

```bash

# Create a table with a GSI
aws dynamodb create-table \
    --table-name MusicCollection \
    --attribute-definitions \
        AttributeName=Artist,AttributeType=S \
        AttributeName=SongTitle,AttributeType=S \
        AttributeName=AlbumTitle,AttributeType=S \
    --key-schema \
        AttributeName=Artist,KeyType=HASH \
        AttributeName=SongTitle,KeyType=RANGE \
    --billing-mode PAY_PER_REQUEST \
    --global-secondary-indexes \
        "IndexName=AlbumIndex,\
        KeySchema=[{AttributeName=AlbumTitle,KeyType=HASH}],\
        Projection={ProjectionType=ALL}"

```

Add a new (on-demand) GSI to an existing table.

```bash

# Add a new GSI to an existing table
aws dynamodb update-table \
    --table-name MusicCollection \
    --attribute-definitions \
        AttributeName=Genre,AttributeType=S \
    --global-secondary-index-updates \
        "[{\"Create\":{\"IndexName\":\"GenreIndex\",\
        \"KeySchema\":[{\"AttributeName\":\"Genre\",\"KeyType\":\"HASH\"}],\
        \"Projection\":{\"ProjectionType\":\"ALL\"}}}]"

```

Update (increase) GSI warm throughput.

```bash

# Increase the warm throughput of a GSI (default values are 12k reads, 4k writes)
aws dynamodb update-table \
    --table-name MusicCollection \
    --global-secondary-index-updates \
        "[{\"Update\":{\"IndexName\":\"AlbumIndex\",\
        \"WarmThroughput\":{\"ReadUnitsPerSecond\":15000,\"WriteUnitsPerSecond\":6000}}}]"

```

Query data using GSIs.

```bash

# Query the AlbumIndex GSI
aws dynamodb query \
    --table-name MusicCollection \
    --index-name AlbumIndex \
    --key-condition-expression "AlbumTitle = :album" \
    --expression-attribute-values '{":album":{"S":"Let It Be"}}'

# Query the GenreIndex GSI
aws dynamodb query \
    --table-name MusicCollection \
    --index-name GenreIndex \
    --key-condition-expression "Genre = :genre" \
    --expression-attribute-values '{":genre":{"S":"Jazz"}}'

```

Delete a GSI.

```bash

# Delete a GSI from a table
aws dynamodb update-table \
    --table-name MusicCollection \
    --global-secondary-index-updates \
        "[{\"Delete\":{\"IndexName\":\"GenreIndex\"}}]"

```

- For API details, see the following topics in _AWS CLI Command Reference_.

- [CreateTable](../../../goto/aws-cli/dynamodb-2012-08-10/createtable.md)

- [DeleteTable](../../../goto/aws-cli/dynamodb-2012-08-10/deletetable.md)

- [Query](../../../goto/aws-cli/dynamodb-2012-08-10/query.md)

- [UpdateTable](../../../goto/aws-cli/dynamodb-2012-08-10/updatetable.md)

For a complete list of AWS SDK developer guides and code examples, see
[Using DynamoDB with an AWS SDK](../../../../reference/amazondynamodb/latest/developerguide/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Invoke a Lambda function from a browser

Manage resource-based policies

All content copied from https://docs.aws.amazon.com/.
