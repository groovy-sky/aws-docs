# Updating the bundle

When a new version of the SDK is released:

1. Navigate to your build project directory

2. Update the SDK packages:

```bash

npm update @amazon-connect/core @amazon-connect/contact @amazon-connect/email
```

3. Rebuild the bundle:

```bash

npm run build
```

4. Copy the new bundle to your website

5. Test your application to verify compatibility

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using in 3P app

Troubleshooting
