# Use path extractors

Amazon Ion is a document style file format, but Apache Hive is a flat columnar format. You
can use special Amazon Ion SerDe properties called `path extractors` to map
between the two formats. Path extractors flatten the hierarchical Amazon Ion format, map
Amazon Ion values to Hive columns, and can be used to rename fields.

Athena can generate the extractors for you, but you can also define your own extractors if
necessary.

###### Topics

- [Use Athena generated path extractors](https://docs.aws.amazon.com/athena/latest/ug/ion-serde-generated-path-extractors.html)

- [Specify your own path extractors](https://docs.aws.amazon.com/athena/latest/ug/ion-serde-specifying-your-own-path-extractors.html)

- [Use search paths in path extractors](https://docs.aws.amazon.com/athena/latest/ug/ion-serde-using-search-paths-in-path-extractors.html)

- [Path extractor examples](https://docs.aws.amazon.com/athena/latest/ug/ion-serde-examples.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Ion SerDe property reference

Use Athena generated path extractors
