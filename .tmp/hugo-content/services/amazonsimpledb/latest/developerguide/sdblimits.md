# Limits

Following is a table that describes current limits within Amazon SimpleDB.

Parameter Restriction Domain size 10 GB per domain Domain size 1 billion attributes per domain Domain name 3-255 characters (a-z, A-Z, 0-9, '\_', '-', and '.') Domains per account 250 Attribute name-value pairs per item 256 Attribute name length 1024 bytes Attribute value length 1024 bytes Item name length 1024 bytes Attribute name, attribute value, and item name allowed characters

All UTF-8 characters that are valid in XML documents.

Control characters and any sequences that are not valid in XML are returned
Base64-encoded. For more information, see [Working with XML-Restricted Characters](invalidcharacters.md).

Attributes per `PutAttributes` operation256 Attributes requested per `Select` operation 256 Items per `BatchDeleteAttributes` operation25Items per `BatchPutAttributes` operation25Maximum items in `Select` response2500Maximum query execution time5 secondsMaximum number of unique attributes per `Select` expression

20

Maximum number of comparisons per `Select` expression

20

Maximum response size for `Select`1MBExports per domain5 in a rolling 24-hour windowExports per AWS account25 in a rolling 24-hour window

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Consistency

Data Set Partitioning

All content copied from https://docs.aws.amazon.com/.
