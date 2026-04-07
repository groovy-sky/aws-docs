# S3 Encryption Client Migration (3. _x_ to 4. _x_)

**Note:** If you're using version 2. _x_ of the Amazon S3 Encryption Client for Go and want to migrate to version 4. _x_, you must first migrate to version 3. _x_. See [S3 Encryption Client Migration (2.x to 3.x)](https://docs.aws.amazon.com/amazon-s3-encryption-client/latest/developerguide/go-v3-migration.html).

Version 4. _x_ of the Amazon S3 Encryption Client for Go introduces AES GCM with Key Commitment (ALG\_AES\_256\_GCM\_HKDF\_SHA512\_COMMIT\_KEY) and Commitment Policies to enhance security by protecting against data key tampering in Instruction Files. This migration guide explains the two-phase approach to safely upgrade from 3. _x_ to 4. _x_ while maintaining backward compatibility during the transition.

## Migration Overview

Migrating from version 3. _x_ to version 4. _x_ of the Amazon S3 Encryption Client for Go requires a two-phase approach to ensure compatibility and security:

1. **Phase 1: Update existing 3. _x_ clients to read 4. _x_ formats**

First, update all existing 3. _x_ clients in your environment to a version that can read objects encrypted with 4. _x_ algorithms and commitment policies (Amazon S3 Encryption Client for Go version 3.2.0 or greater). This ensures that when you start encrypting with 4. _x_, your existing applications can still decrypt the new objects.

2. **Phase 2: Migrate encryption and decryption clients to 4. _x_**

After all clients can read 4. _x_ formats, migrate your encryption and decryption operations to use 4. _x_ clients with the appropriate Commitment Policy. This phase introduces the enhanced security features while maintaining backward compatibility with existing encrypted objects.

This phased approach prevents compatibility issues and ensures that all encrypted objects remain accessible throughout the migration process.

## Understanding 4. _x_ Concepts

Version 4. _x_ introduces two key security concepts that enhance protection against data key tampering:

### Commitment Policy

Commitment Policy controls how the encryption client handles key commitment during encryption and decryption operations. There are three Commitment Policies:

`FORBID_ENCRYPT_ALLOW_DECRYPT`

**Encryption:** Encrypts without commitment.

**Decryption:** Allows decryption of both committing and non-committing objects.

**Security:** Does not enforce commitment and may allow for tampering of data keys in Instruction Files. Use only during migration for backward compatibility.

**Version Compatibility:** Objects encrypted with this policy can be read by 2. _x_, 3. _x_, and 4. _x_ clients.

`REQUIRE_ENCRYPT_ALLOW_DECRYPT`

**Encryption:** Encrypts with commitment (uses ALG\_AES\_256\_GCM\_HKDF\_SHA512\_COMMIT\_KEY algorithm).

**Decryption:** Allows decryption of both committing and non-committing objects.

**Security:** New objects are protected against tampering in Instruction Files. Old objects remain readable, but are not protected against data key tampering.

**Version Compatibility:** Objects encrypted with this policy can only be read by 3. _x_ clients (Amazon S3 Encryption Client for Go version 3.2.0 or greater) and 4. _x_ clients.

`REQUIRE_ENCRYPT_REQUIRE_DECRYPT` (Default for 4. _x_)

**Encryption:** Encrypts with commitment (uses ALG\_AES\_256\_GCM\_HKDF\_SHA512\_COMMIT\_KEY algorithm).

**Decryption:** Only allows decryption of objects encrypted with key commitment.

**Security:** Strict commitment enforcement provides protection against tampered data keys.

**Version Compatibility:** Objects encrypted with this policy can only be read by 3. _x_ clients (version 3.2.0 or greater) and 4. _x_ clients. This policy will reject non-committing objects during decryption.

### AES GCM with Key Commitment

AES GCM with Key Commitment (ALG\_AES\_256\_GCM\_HKDF\_SHA512\_COMMIT\_KEY) is the new encryption algorithm suite introduced in version 4. _x_ that protects against data key tampering in Instruction Files by cryptographically binding the key to its intended use.

**Version Compatibility:** Objects encrypted with ALG\_AES\_256\_GCM\_HKDF\_SHA512\_COMMIT\_KEY can only be decrypted by 3. _x_ clients (version 3.2.0 or greater) or 4. _x_ clients. 2. _x_ and earlier 3. _x_ clients (v3.1.0 and earlier) cannot read these objects.

## Update Existing Clients

Before migrating to 4. _x_ encryption, you must first update all existing 3. _x_ clients to a version that can read 4. _x_ encrypted objects. This ensures compatibility when you begin encrypting with 4. _x_.

### Build and Install the Latest SDK Version

Update your Go module dependencies to use the latest version of the Amazon S3 Encryption Client for Go 3. _x_ that includes support for reading 4. _x_ messages:

**Update Go modules:**

```bash

go get github.com/aws/amazon-s3-encryption-client-go/v3@latest
```

**Update your go.mod file if needed:**

```go

module your-application

go 1.24

require (
    // The `x`s must be replaced with the specific latest version
    github.com/aws/amazon-s3-encryption-client-go/v3 v3.x.x
    // Add other dependencies as needed
)
```

### Build, Install, and Deploy Applications

After updating your dependencies, rebuild and deploy your applications:

1. **Clean and rebuild:**

```bash

go mod tidy
go build ./...
```

2. **Run your tests:**

```bash

go test ./...
```

3. **Deploy updated applications:** Deploy the updated applications to all environments where S3 Encryption Client is used. Ensure all applications can successfully decrypt existing V3 encrypted objects before proceeding to the next phase.

## Migrate to V4

After all clients in your environment can read 4. _x_ formats, you can migrate your encryption and decryption operations to use 4. _x_ clients. The following examples demonstrate the transition from 3. _x_ to 4. _x_ clients.

### Client Migration Examples

The following examples show how to migrate from 3. _x_ to 4. _x_ clients using different Commitment Policies during the transition:

**Pre-migration (3. _x_ Client)**

```go

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	// V3 S3 Encryption Client imports
	"github.com/aws/amazon-s3-encryption-client-go/v3/client"
	"github.com/aws/amazon-s3-encryption-client-go/v3/materials"
)

func V3EncryptionExample() error {
	ctx := context.Background()

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"))
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// Create KMS and S3 clients
	kmsClient := kms.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	// Create V3 encryption materials
	kmsKeyId := "arn:aws:kms:us-west-2:123456789012:key/12345678-1234-1234-1234-123456789012"
	cmm, err := materials.NewCryptographicMaterialsManager(
		materials.NewKmsKeyring(kmsClient, kmsKeyId))
	if err != nil {
		return fmt.Errorf("failed to create CMM: %v", err)
	}

	// Create V3 S3 encryption client
	s3EncryptionClient, err := client.New(s3Client, cmm)
	if err != nil {
		return fmt.Errorf("failed to create encryption client: %v", err)
	}

	// Encrypt and upload object
	encryptionContext := map[string]string{
		"purpose": "example",
		"department": "engineering",
	}

	_, err = s3EncryptionClient.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-object"),
		Body:   strings.NewReader("Hello, World!"),
		Metadata: encryptionContext,
	})

	return err
}
```

**During Migration (4. _x_ Client with FORBID\_ENCRYPT\_ALLOW\_DECRYPT Policy)**

Update your Go module dependencies to use the latest version of the Amazon S3 Encryption Client for Go 4. _x_:

**Upgrade Go modules:**

```bash

go get github.com/aws/amazon-s3-encryption-client-go/v4@latest
```

**Update your go.mod file if needed:**

```go

module your-application

go 1.24

require (
    // The `x`s must be replaced with the specific latest version
    github.com/aws/amazon-s3-encryption-client-go/v4 v4.x.x
    // Add other dependencies as needed
)
```

After updating your dependencies, make code changes as described below.
This should result in no functional changes to your application.

```go

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	// V4 S3 Encryption Client imports
	"github.com/aws/amazon-s3-encryption-client-go/v4/client"
	"github.com/aws/amazon-s3-encryption-client-go/v4/commitment"
	"github.com/aws/amazon-s3-encryption-client-go/v4/materials"
)

func V4ForbidAllowExample() error {
	ctx := context.Background()

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"))
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// Create KMS and S3 clients
	kmsClient := kms.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	// Create V4 encryption materials with KMS keyring
	kmsKeyId := "arn:aws:kms:us-west-2:123456789012:key/12345678-1234-1234-1234-123456789012"
	cmm, err := materials.NewCryptographicMaterialsManager(
		materials.NewKmsKeyring(kmsClient, kmsKeyId))
	if err != nil {
		return fmt.Errorf("failed to create CMM: %v", err)
	}

	// Create V4 S3 encryption client with FORBID_ENCRYPT_ALLOW_DECRYPT policy
	// This maintains V3 compatibility during migration
	s3EncryptionClient, err := client.New(s3Client, cmm, func(options *client.EncryptionClientOptions) {
		// This MUST be explicitly configured to FORBID_ENCRYPT_ALLOW_DECRYPT.
		// While FORBID_ENCRYPT_ALLOW_DECRYPT is the default for v3 clients,
		// v4 clients default to REQUIRE_ENCRYPT_REQUIRE_DECRYPT.
		// This configuration ensures identical behavior to a v3 client.
		options.CommitmentPolicy = commitment.FORBID_ENCRYPT_ALLOW_DECRYPT
	})
	if err != nil {
		return fmt.Errorf("failed to create V4 encryption client: %v", err)
	}

	// Encrypt and upload object (uses legacy algorithms for V3 compatibility)
	encryptionContext := map[string]string{
		"purpose": "example",
		"department": "engineering",
	}

	_, err = s3EncryptionClient.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-object"),
		Body:   strings.NewReader("Hello, World!"),
		Metadata: encryptionContext,
	})

	return err
}
```

**During migration (4. _x_ Client with REQUIRE\_ENCRYPT\_ALLOW\_DECRYPT Policy)**

After deploying the FORBID\_ENCRYPT\_ALLOW\_DECRYPT changes, make code changes as described below.
This will cause your application to start writing objects encrypted with key committing algorithms.

```go

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	// V4 S3 Encryption Client imports
	"github.com/aws/amazon-s3-encryption-client-go/v4/client"
    "github.com/aws/amazon-s3-encryption-client-go/v4/commitment"
	"github.com/aws/amazon-s3-encryption-client-go/v4/materials"
)

func V4RequireAllowExample() error {
	ctx := context.Background()

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"))
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// Create KMS and S3 clients
	kmsClient := kms.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	// Create V4 encryption materials with KMS keyring
	kmsKeyId := "arn:aws:kms:us-west-2:123456789012:key/12345678-1234-1234-1234-123456789012"
	cmm, err := materials.NewCryptographicMaterialsManager(
		materials.NewKmsKeyring(kmsClient, kmsKeyId))
	if err != nil {
		return fmt.Errorf("failed to create CMM: %v", err)
	}

	// Create V4 S3 encryption client with REQUIRE_ENCRYPT_ALLOW_DECRYPT commitment policy
	// Uses REQUIRE_ENCRYPT_ALLOW_DECRYPT to ensure new messages are encrypted with key commitment
	s3EncryptionClient, err := client.New(s3Client, cmm, func(options *client.EncryptionClientOptions) {
		// Migration note: The commitment policy has been updated to REQUIRE_ENCRYPT_ALLOW_DECRYPT.
		// This change causes the client to start writing objects encrypted with key committing algorithms.
		// The client will continue to be able to read objects encrypted with either
		// key committing or non-key committing algorithms.
		options.CommitmentPolicy = commitment.REQUIRE_ENCRYPT_ALLOW_DECRYPT
	})
	if err != nil {
		return fmt.Errorf("failed to create V4 encryption client: %v", err)
	}

	// Encrypt and upload object (uses AES_GCM_KC with key commitment)
	encryptionContext := map[string]string{
		"purpose": "example",
		"department": "engineering",
	}

	_, err = s3EncryptionClient.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-object"),
		Body:   strings.NewReader("Hello, World!"),
		Metadata: encryptionContext,
	})
	if err != nil {
		return fmt.Errorf("failed to put object: %v", err)
	}

	// Decrypt and download object
	result, err := s3EncryptionClient.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-object"),
	})
	if err != nil {
		return fmt.Errorf("failed to get object: %v", err)
	}
	defer result.Body.Close()

	// Read decrypted content
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	fmt.Printf("Decrypted content: %s\n", string(body))
	return nil
}
```

**Post-migration (4. _x_ Client with default commitment policy)**

After deploying the REQUIRE\_ENCRYPT\_ALLOW\_DECRYPT changes, make code changes as described below.
This will cause your application to stop reading objects encrypted without key committing algorithms.
Before deploying this change, ensure all existing objects are now encrypted with key commiting algorithms.

```go

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	// V4 S3 Encryption Client imports
	"github.com/aws/amazon-s3-encryption-client-go/v4/client"
	"github.com/aws/amazon-s3-encryption-client-go/v4/materials"
)

func V4KeyCommitmentExample() error {
	ctx := context.Background()

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"))
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// Create KMS and S3 clients
	kmsClient := kms.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	// Create V4 encryption materials with KMS keyring
	kmsKeyId := "arn:aws:kms:us-west-2:123456789012:key/12345678-1234-1234-1234-123456789012"
	cmm, err := materials.NewCryptographicMaterialsManager(
		materials.NewKmsKeyring(kmsClient, kmsKeyId))
	if err != nil {
		return fmt.Errorf("failed to create CMM: %v", err)
	}

	// Create V4 S3 encryption client with default commitment policy (REQUIRE_ENCRYPT_REQUIRE_DECRYPT)
	// Uses REQUIRE_ENCRYPT_ALLOW_DECRYPT to ensure all messages read or written are encrypted with key commitment
	s3EncryptionClient, err := client.New(s3Client, cmm)
	if err != nil {
		return fmt.Errorf("failed to create V4 encryption client: %v", err)
	}

	// Encrypt and upload object (uses AES_GCM_KC with key commitment)
	encryptionContext := map[string]string{
		"purpose": "example",
		"department": "engineering",
	}

	_, err = s3EncryptionClient.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-object"),
		Body:   strings.NewReader("Hello, World!"),
		Metadata: encryptionContext,
	})
	if err != nil {
		return fmt.Errorf("failed to put object: %v", err)
	}

	// Decrypt and download object
	result, err := s3EncryptionClient.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("my-object"),
	})
	if err != nil {
		return fmt.Errorf("failed to get object: %v", err)
	}
	defer result.Body.Close()

	// Read decrypted content
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	fmt.Printf("Decrypted content: %s\n", string(body))
	return nil
}
```

## Additional Examples

The following examples demonstrate specific 4. _x_ configuration scenarios for different migration and operational needs.

### Enable Legacy Support for Reading 1. _x_/2. _x_ Objects

During migration, you may need to read objects encrypted with legacy algorithms. Configure your 4. _x_ client to support backward compatibility:

```go

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	// V4 S3 Encryption Client imports
	"github.com/aws/amazon-s3-encryption-client-go/v4/client"
	"github.com/aws/amazon-s3-encryption-client-go/v4/commitment"
	"github.com/aws/amazon-s3-encryption-client-go/v4/materials"
)

func V4WithLegacySupportExample() error {
	ctx := context.Background()

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"))
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// Create KMS and S3 clients
	kmsClient := kms.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	// Create encryption materials
	kmsKeyId := "arn:aws:kms:us-west-2:123456789012:key/12345678-1234-1234-1234-123456789012"
	cmm, err := materials.NewCryptographicMaterialsManager(
		materials.NewKmsKeyring(kmsClient, kmsKeyId, func(options *materials.KeyringOptions) {
			// Enable legacy wrapping algorithms for V2/V3 compatibility
			options.EnableLegacyWrappingAlgorithms = true
		}))
	if err != nil {
		return fmt.Errorf("failed to create CMM: %v", err)
	}

	// Create V4 client with legacy support enabled
	s3EncryptionClient, err := client.New(s3Client, cmm, func(options *client.EncryptionClientOptions) {
		// Allow reading objects encrypted with legacy algorithms
		options.EnableLegacyUnauthenticatedModes = true
		// Use REQUIRE_ENCRYPT_ALLOW_DECRYPT to encrypt with commitment but read legacy objects
		options.CommitmentPolicy = commitment.REQUIRE_ENCRYPT_ALLOW_DECRYPT
	})
	if err != nil {
		return fmt.Errorf("failed to create encryption client: %v", err)
	}

	// This client can now read objects encrypted with legacy algorithms and encrypt new objects with commitment
	result, err := s3EncryptionClient.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String("my-bucket"),
		Key:    aws.String("legacy-encrypted-object"),
	})
	if err != nil {
		return fmt.Errorf("failed to decrypt legacy object: %v", err)
	}
	defer result.Body.Close()

	return nil
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Examples

Migrate from 2.x to 3.x
