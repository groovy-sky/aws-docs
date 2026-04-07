# Amazon S3 Encryption Client for Java

###### Note

This documentation describes the Amazon S3 Encryption Client version 3. _x_ and newer, which is an independent library.
For information about previous versions of the Amazon S3 Encryption Client, see the AWS SDK Developer Guide for your programming language.

This topic explains how to install and use the Amazon S3 Encryption Client for Java. For details about programming
with the Amazon S3 Encryption Client for Java, see the [amazon-s3-encryption-client-java](https://github.com/aws/amazon-s3-encryption-client-java) repository on GitHub.

###### Topics

- [Prerequisites](#java-prerequisites)

- [Installation](#java-installation)

- [Amazon S3 Encryption Client for Java examples](https://docs.aws.amazon.com/amazon-s3-encryption-client/latest/developerguide/java-examples.html)

- [Asynchronous programming in the Amazon S3 Encryption Client for Java](https://docs.aws.amazon.com/amazon-s3-encryption-client/latest/developerguide/using-s3ec-async.html)

- [S3 Encryption Client Migration (V3 to V4)](https://docs.aws.amazon.com/amazon-s3-encryption-client/latest/developerguide/java-v4-migration.html)

- [S3 Encryption Client Migration (V2 to V3)](https://docs.aws.amazon.com/amazon-s3-encryption-client/latest/developerguide/java-v3-migration.html)

## Prerequisites

Before you install the Amazon S3 Encryption Client for Java, be sure you have the following
prerequisites.

**A Java development environment**

You will need Java 8 or later. On the Oracle website, go to [Java SE Downloads](https://www.oracle.com/java/technologies/downloads), and then download and install the Java SE
Development Kit (JDK).

If you use the Oracle JDK, you must also download and install the [Java Cryptography Extension (JCE) Unlimited Strength Jurisdiction\
Policy Files](http://www.oracle.com/java/technologies/javase-jce8-downloads.html).

**AWS SDK for Java 2.x**

The Amazon S3 Encryption Client for Java requires the Amazon S3 and AWS KMS modules of the AWS SDK for Java 2.x. You can
install the entire SDK or just the Amazon S3 and AWS KMS modules.

For information about updating your version of the AWS SDK for Java, see [Migrating from version 1.x to 2.x of\
the AWS SDK for Java](https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/migration.html).

To install the AWS SDK for Java, use Apache Maven.

- To [import the entire AWS SDK for Java](https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/setup-project-maven.html#build-the-entire-sdk-into-your-project) as a dependency, declare
it in your `pom.xml` file.

- To create a dependency for the Amazon S3 module in the AWS SDK for Java,
follow the instructions for [specifying particular modules](https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/setup-project-maven.html#modules-dependencies). Set the
`groupId` to `software.amazon.awssdk` and
the `artifactID` to `s3`.

- To create a dependency for the AWS KMS module in the AWS SDK for Java, follow
the instructions for [specifying particular modules](https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/setup-project-maven.html#modules-dependencies). Set the `groupId` to
`software.amazon.awssdk` and the `artifactId` to
`kms`.

## Installation

You can install the latest version of the Amazon S3 Encryption Client for Java in the
following ways.

**Manually**

To install the Amazon S3 Encryption Client for Java, clone or download the [amazon-s3-encryption-client-java](https://github.com/aws/amazon-s3-encryption-client-java)
GitHub repository.

**Using Apache Maven**

The Amazon S3 Encryption Client for Java is available through [Apache Maven](https://maven.apache.org/) with the following dependency definition. Install
the latest version offered.

```xml

<dependency>
  <groupId>software.amazon.encryption.s3</groupId>
  <artifactId>amazon-s3-encryption-client-java</artifactId>
  <version>4.x</version>
</dependency>
```

### Optional dependencies

**Multipart upload (high-level API)**

To perform multipart uploads with the [high-level API](https://docs.aws.amazon.com/amazon-s3-encryption-client/latest/developerguide/java-examples.html#highlevel-multipart-upload), create
dependencies for the AWS CRT-based Amazon S3 client. For help creating these
dependencies, see [Add dependencies to use the AWS CRT-based Amazon S3 client](https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/crt-based-s3-client.html#crt-based-s3-client-depend) in the
_AWS SDK for Java 2.x Developer Guide_.

For more information on multipart uploads in the Amazon S3 Encryption Client, see [Multipart upload](https://docs.aws.amazon.com/amazon-s3-encryption-client/latest/developerguide/java-examples.html#multipart-upload).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Programming languages

Examples
