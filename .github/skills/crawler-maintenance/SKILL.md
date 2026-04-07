---
name: crawler-maintenance
description: 'Maintain aws-docs crawler behavior and keep crawler logic documentation in sync. Use for crawler bug fixes, crawl flow changes, extraction/mapping updates, and seed or mode behavior changes.'
argument-hint: 'Describe the crawler change to implement and validate'
user-invocable: true
---

# Crawler Maintenance

## Goal

Implement crawler-related code changes safely while keeping the crawler logic source of truth synchronized.

## Source of Truth

- Crawler logic must be documented in [Crawler Logic Reference](../../../internal/INSTRUCTION.md).
- Any crawler code change that affects behavior must update that file in the same change.

## When to Use

Use this skill when work touches any of these areas:

- `internal/crawl/` logic (queueing, extraction, conversion, mapping, URL filtering)
- `internal/app/run.go` crawler wiring or run modes
- `internal/config/config.go` crawler runtime configuration behavior
- `internal/store/store.go` crawl metadata behavior
- `internal/write/services_index.go` behavior tied to crawl output interpretation

## Procedure

1. Identify behavior change scope.
2. Read affected code paths end-to-end before editing.
3. Implement the minimal code change needed.
4. Update [internal/INSTRUCTION.md](../../../internal/INSTRUCTION.md) for every behavior, flow, or rule change.
5. Add or update tests for the changed behavior when feasible.
6. Run task-specific verification commands appropriate for the change scope.
7. Summarize:
- What changed in code
- What changed in documentation
- What was validated

## Completion Checks

- Code behavior and documentation are aligned.
- [internal/INSTRUCTION.md](../../../internal/INSTRUCTION.md) reflects the latest crawler logic.
- Task-specific tests or validation commands relevant to the change have run successfully.
- No unrelated files are modified by this workflow unless explicitly requested.

## Documentation Format

- No fixed changelog section format is required in [internal/INSTRUCTION.md](../../../internal/INSTRUCTION.md).
- Keep updates concise and place them in the most relevant existing section.

## Ambiguity Handling

If a requested crawler change is unclear, clarify these first:

- Expected crawl behavior after change
- Backward compatibility requirements
- Whether documentation wording should be concise or exhaustive
