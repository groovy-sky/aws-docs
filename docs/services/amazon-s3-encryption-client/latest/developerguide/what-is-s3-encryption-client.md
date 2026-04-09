# What is the Amazon S3 Encryption Client?

###### Note

This documentation describes the Amazon S3 Encryption Client version 3. _x_ and newer, which is an independent library.
For information about previous versions of the Amazon S3 Encryption Client, see the AWS SDK Developer Guide for your programming language.

The Amazon S3 Encryption Client is a client-side encryption library that enables you to encrypt an object
locally to ensure its security before passing it to [Amazon Simple Storage Service](../../../s3/latest/userguide/welcome.md)
(Amazon S3). Amazon S3 receives your object already encrypted; it does not play a role in encrypting or
decrypting it. After you instantiate the Amazon S3 Encryption Client, your objects are automatically encrypted and
decrypted as part of your Amazon S3 `PutObject` and `GetObject` requests. The
Amazon S3 Encryption Client is provided free of charge under the Apache 2.0 license.

The Amazon S3 Encryption Client is supported in the following programming languages and platforms. This guide
focuses on version 4.x of the [Amazon S3 Encryption Client for Java](java.md) and [Amazon S3 Encryption Client for Go](go.md).
For more information on the remaining language implementations, see their respective AWS SDK
Developer Guides.

- [AWS SDK for C++](../../../../reference/sdk-for-cpp/latest/developer-guide.md)

- [AWS SDK for Go](https://github.com/aws/amazon-s3-encryption-client-go)

- [AWS SDK for Java](https://github.com/aws/amazon-s3-encryption-client-java)

- [AWS SDK for .NET](../../../../reference/sdk-for-net/latest/developer-guide/kms-keys-s3-encryption.md)

- [AWS SDK for Ruby](../../../../reference/sdk-for-ruby/v3/api/aws/s3/encryptionv2/client.md)

- [AWS SDK for PHP](../../../../reference/sdk-for-php/latest/developer-guide/s3-encryption-client.md)

The Amazon S3 Encryption Client provides:

**A default implementation that adheres to cryptography best practices**

By default, the Amazon S3 Encryption Client generates a unique data key for each object that it encrypts. This
follows the cryptography best practice of using unique data keys for each encryption
operation.

The Amazon S3 Encryption Client encrypts your objects using a secure, authenticated, symmetric key
algorithm.

**A framework for protecting data keys with wrapping keys**

The Amazon S3 Encryption Client protects the data keys that encrypt your objects by encrypting them under a
[wrapping key](concepts.md#wrapping-key). With the Amazon S3 Encryption Client, you define a wrapping key
by passing the key to the Amazon S3 Encryption Client, which it uses to optimize its settings.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Terms and concepts

All content copied from https://docs.aws.amazon.com/.
