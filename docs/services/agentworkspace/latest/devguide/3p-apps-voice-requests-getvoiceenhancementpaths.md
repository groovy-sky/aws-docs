---
title: "Get voice enhancement model paths in Connect Customer agent workspace"
---

# Get voice enhancement model paths in Connect Customer agent workspace

Returns the voice enhancements models static assets URL paths.

**Signature**

```

async getVoiceEnhancementPaths(): Promise<VoiceEnhancementPaths>
```

**Usage**

```

voiceClient.getVoiceEnhancementPaths();

// VoiceEnhancementPaths structure
interface VoiceEnhancementPaths {
  processors: string;
  workers: string;
  wasm: string;
  models: string;
}
```

**Permissions required:**

```

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getVoiceEnhancementMode()

isParticipantOnHold()

All content copied from https://docs.aws.amazon.com/.
