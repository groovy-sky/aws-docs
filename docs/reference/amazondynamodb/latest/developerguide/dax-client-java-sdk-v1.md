# Using DAX with AWS SDK for Java 1.x

Follow this procedure to run the Java sample for Amazon DynamoDB Accelerator (DAX) on your Amazon EC2
instance.

###### Note

These instructions are for applications using AWS SDK for Java 1.x. For applications using
AWS SDK for Java 2.x, see [Java and DAX](../../../../services/dynamodb/latest/developerguide/dax-client-run-application-java.md).

###### To run the Java sample for DAX

1. Install the Java Development Kit (JDK).

```nohighlight

sudo yum install -y java-devel
```

2. Download the AWS SDK for Java ( `.zip` file), and then extract
    it.

```nohighlight

wget http://sdk-for-java.amazonwebservices.com/latest/aws-java-sdk.zip

unzip aws-java-sdk.zip
```

3. Download the latest version of the DAX Java client ( `.jar`
    file).

```nohighlight

wget http://dax-sdk.s3-website-us-west-2.amazonaws.com/java/DaxJavaClient-latest.jar
```

###### Note

The client for the DAX SDK for Java is available on Apache Maven. For more
information, see [Using the client as an Apache Maven dependency](#DAXClient.Maven).

4. Set your `CLASSPATH` variable. In this example, replace
    `sdkVersion` with the actual version
    number of the AWS SDK for Java (for example, `1.11.112`).

```nohighlight

export SDKVERSION=sdkVersion

export CLASSPATH=$(pwd)/TryDax/java:$(pwd)/DaxJavaClient-latest.jar:$(pwd)/aws-java-sdk-$SDKVERSION/lib/aws-java-sdk-$SDKVERSION.jar:$(pwd)/aws-java-sdk-$SDKVERSION/third-party/lib/*
```

5. Download the sample program source code ( `.zip` file).

```nohighlight

wget http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/samples/TryDax.zip
```

When the download is complete, extract the source files.

```nohighlight

unzip TryDax.zip
```

6. Navigate to the Java code directory and compile the code as follows.

```nohighlight

cd TryDax/java/
javac TryDax*.java
```

7. Run the program.

```nohighlight

java TryDax
```

You should see output similar to the following.

```nohighlight

Creating a DynamoDB client

Attempting to create table; please wait...
Successfully created table.  Table status: ACTIVE
Writing data to the table...
Writing 10 items for partition key: 1
Writing 10 items for partition key: 2
Writing 10 items for partition key: 3
Writing 10 items for partition key: 4
Writing 10 items for partition key: 5
Writing 10 items for partition key: 6
Writing 10 items for partition key: 7
Writing 10 items for partition key: 8
Writing 10 items for partition key: 9
Writing 10 items for partition key: 10

Running GetItem, Scan, and Query tests...
First iteration of each test will result in cache misses
Next iterations are cache hits

GetItem test - partition key 1 and sort keys 1-10
   	Total time: 136.681 ms - Avg time: 13.668 ms
   	Total time: 122.632 ms - Avg time: 12.263 ms
   	Total time: 167.762 ms - Avg time: 16.776 ms
   	Total time: 108.130 ms - Avg time: 10.813 ms
   	Total time: 137.890 ms - Avg time: 13.789 ms
Query test - partition key 5 and sort keys between 2 and 9
   	Total time: 13.560 ms - Avg time: 2.712 ms
   	Total time: 11.339 ms - Avg time: 2.268 ms
   	Total time: 7.809 ms - Avg time: 1.562 ms
   	Total time: 10.736 ms - Avg time: 2.147 ms
   	Total time: 12.122 ms - Avg time: 2.424 ms
Scan test - all items in the table
   	Total time: 58.952 ms - Avg time: 11.790 ms
   	Total time: 25.507 ms - Avg time: 5.101 ms
   	Total time: 37.660 ms - Avg time: 7.532 ms
   	Total time: 26.781 ms - Avg time: 5.356 ms
   	Total time: 46.076 ms - Avg time: 9.215 ms

Attempting to delete table; please wait...
Successfully deleted table.
```

Take note of the timing information—the number of milliseconds required for
    the `GetItem`, `Query`, and `Scan` tests.

8. In the previous step, you ran the program against the DynamoDB endpoint. Now run the
    program again, but this time, the `GetItem`, `Query`, and
    `Scan` operations are processed by your DAX cluster.

To determine the endpoint for your DAX cluster, choose one of the
    following:

- **Using the DynamoDB console** — Choose your DAX
cluster. The cluster endpoint is shown on the console, as in the following
example.

```nohighlight

dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
```

- **Using the AWS CLI** — Enter the following
command.

```nohighlight

aws dax describe-clusters --query "Clusters[*].ClusterDiscoveryEndpoint"
```

The cluster endpoint is shown in the output, as in the following
example.

```json

{
      "Address": "my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com",
      "Port": 8111,
      "URL": "dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com"
}
```

Now run the program again, but this time, specify the cluster endpoint as a
command line parameter.

```nohighlight

java TryDax dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
```

Look at the rest of the output, and take note of the timing information. The
elapsed times for `GetItem`, `Query`, and `Scan`
should be significantly lower with DAX than with DynamoDB.

For more information about this program, see the following sections:

- [TryDax.java](../../../../services/dynamodb/latest/developerguide/dax-client-run-application-java-trydax.md)

- [TryDaxHelper.java](../../../../services/dynamodb/latest/developerguide/dax-client-run-application-java-trydaxhelper.md)

- [TryDaxTests.java](../../../../services/dynamodb/latest/developerguide/dax-client-run-application-java-trydaxtests.md)

## Using the client as an Apache Maven dependency

Follow these steps to use the client for the DAX SDK for Java in your application as a
dependency.

###### To use the client as a Maven dependency

1. Download and install Apache Maven. For more information, see [Downloading Apache Maven](https://maven.apache.org/download.cgi)
    and [Installing Apache\
    Maven](https://maven.apache.org/install.html).

2. Add the client Maven dependency to your application's Project Object Model
    (POM) file. In this example, replace `x.x.x.x` with the actual
    version number of the client (for example, `1.0.200704.0`).

```xml

<!--Dependency:-->
<dependencies>
       <dependency>
        <groupId>com.amazonaws</groupId>
        <artifactId>amazon-dax-client</artifactId>
        <version>x.x.x.x</version>
       </dependency>
</dependencies>
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS SDK for Java 1.x examples

TryDax.java

All content copied from https://docs.aws.amazon.com/.
