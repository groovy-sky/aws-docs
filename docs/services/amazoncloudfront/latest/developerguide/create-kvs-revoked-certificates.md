# Step 1: Create a KeyValueStore for revoked certificates

Create a KeyValueStore to store revoked certificate serial numbers that your
Connection Function can check during mTLS connections.

First, prepare your revoked certificate serial numbers in JSON format:

```json

{
  "data": [
    {
      "key": "ABC123DEF456",
      "value": ""
    },
    {
      "key": "789XYZ012GHI",
      "value": ""
    }
  ]
}
```

Upload this JSON file to an S3 bucket, then create the KeyValueStore:

```nohighlight

aws s3 cp revoked-serials.json s3://your-bucket-name/revoked-serials.json
aws cloudfront create-key-value-store \
  --name revoked-serials-kvs \
  --import-source '{
    "SourceType": "S3",
    "SourceARN": "arn:aws:s3:::your-bucket-name/revoked-serials.json"
  }'
```

Wait for the KeyValueStore to finish provisioning. Check the status with:

```nohighlight

aws cloudfront get-key-value-store --name "revoked-serials-kvs"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Implement certificate revocation for mutual TLS (viewer) with CloudFront Functions and KeyValueStore

Step 2: Create the revocation Connection Function

All content copied from https://docs.aws.amazon.com/.
