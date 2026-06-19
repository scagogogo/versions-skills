---
name: version-constraints
description: Check whether a version satisfies a constraint expression (caret, tilde, wildcard, range operators). Use when resolving dependency compatibility, filtering versions by semver ranges, or working with constraint expressions that combine AND/OR logic.
argument-hint: <constraint-expression> <version>
---

# Version Constraints

> **Prerequisite:** See `/installation` skill for SDK/CLI/MCP setup.

## When to Use

- You need to check if a version satisfies a constraint expression (e.g., `">=1.0.0,<2.0.0"`)
- You are resolving dependency version compatibility
- You need to filter versions by semver ranges (caret `^`, tilde `~`, wildcards `x/*`)
- You are working with constraint expressions that combine AND (comma) / OR (`||`) logic

## Decision Tree

```
Need to check version against constraints?
├─ Single constraint (e.g., ">=1.0.0")?       → ParseConstraint() + Constraint.Match()
├─ AND logic (comma-separated, e.g., ">=1.0.0,<2.0.0")?
│                                              → ParseConstraintSet() / default CLI
├─ OR logic (||-separated, e.g., ">=1.0.0 || >=3.0.0")?
│                                              → ParseConstraintUnion() / --type union
├─ Quick one-liner (parse + match)?            → Version.Matches(expr)
├─ Version-centric check?                      → Version.Satisfies(constraint)
└─ Filter a collection by constraint?          → FilterByConstraint() / versions filter --constraint
```

## Constraint Operators Reference

| Operator | Meaning | Example | Matches |
|----------|---------|---------|---------|
| `=` | Exact match | `=1.2.3` | Only 1.2.3 |
| `!=` | Not equal | `!=1.2.3` | Anything except 1.2.3 |
| `>` | Greater than | `>1.2.3` | 1.2.4, 2.0.0, ... |
| `>=` | Greater or equal | `>=1.2.3` | 1.2.3, 1.2.4, ... |
| `<` | Less than | `<2.0.0` | 1.9.9, 1.0.0, ... |
| `<=` | Less or equal | `<=2.0.0` | 2.0.0, 1.9.9, ... |
| `^` | Caret (compatible) | `^1.2.3` | >=1.2.3, <2.0.0 |
| `~` | Tilde (approximate) | `~1.2.3` | >=1.2.3, <1.3.0 |
| `x/X/*` | Wildcard | `1.x` | >=1.0.0, <2.0.0 |

**Combining:** comma `,` = AND, double-pipe `||` = OR. Spaces around operators are NOT supported: use `>=1.0.0` not `>= 1.0.0`.

## Task Patterns

### Check if a version satisfies a range constraint

**Goal:** Is `"1.5.0"` within `">=1.0.0,<2.0.0"`?

**SDK approach:**
```go
cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
v := versions.NewVersion("1.5.0")
ok := cs.Satisfies(v) // true

// Or the one-liner:
ok, _ := versions.NewVersion("1.5.0").Matches(">=1.0.0,<2.0.0")
```

**CLI approach:**
```bash
versions constraint ">=1.0.0,<2.0.0" 1.5.0        # AND logic (default)
versions satisfies 1.5.0 ">=1.0.0,<2.0.0"          # version-centric
```

**MCP approach:**
```json
{"tool": "version_constraint_check", "arguments": {"expression": ">=1.0.0,<2.0.0", "version": "1.5.0"}}
```

### Check if a version satisfies OR logic constraints

**Goal:** Is `"3.5.0"` within `">=1.0.0,<2.0.0 || >=3.0.0"`?

**SDK approach:**
```go
cu, _ := versions.ParseConstraintUnion(">=1.0.0,<2.0.0 || >=3.0.0")
ok := cu.Satisfies(versions.NewVersion("3.5.0"))  // true
ok = cu.Satisfies(versions.NewVersion("2.5.0"))   // false
```

**CLI approach:**
```bash
versions constraint ">=1.0.0 || >=3.0.0" 3.5.0 --type union    # match (3.5.0 >= 3.0.0)
versions constraint ">=1.0.0 || >=3.0.0" 2.5.0 --type union    # no match
```

**MCP approach:**
```json
{"tool": "version_constraint_check", "arguments": {"expression": ">=1.0.0 || >=3.0.0", "version": "3.5.0", "type": "union"}}
```

### Use caret constraint

**Goal:** Check if `"1.9.9"` is compatible with `"^1.2.3"` (matches `>=1.2.3, <2.0.0`).

**SDK approach:**
```go
c, _ := versions.ParseConstraint("^1.2.3")
c.Match(versions.NewVersion("1.9.9"))  // true
c.Match(versions.NewVersion("2.0.0"))  // false

// Caret with 0.x is more restrictive:
c2, _ := versions.ParseConstraint("^0.2.3")  // >=0.2.3, <0.3.0
```

**CLI approach:**
```bash
versions constraint "^1.2.3" 1.9.9 --type single    # true
versions constraint "^1.2.3" 2.0.0 --type single    # false
```

**MCP approach:**
```json
{"tool": "version_constraint_check", "arguments": {"expression": "^1.2.3", "version": "1.9.9", "type": "single"}}
```

### Use tilde constraint

**Goal:** Check if `"1.2.9"` is within patch-level range `"~1.2.3"` (matches `>=1.2.3, <1.3.0`).

**SDK approach:**
```go
c, _ := versions.ParseConstraint("~1.2.3")
c.Match(versions.NewVersion("1.2.9"))  // true
c.Match(versions.NewVersion("1.3.0"))  // false
```

**CLI approach:**
```bash
versions constraint "~1.2.3" 1.2.9 --type single    # true
versions constraint "~1.2.3" 1.3.0 --type single    # false
```

**MCP approach:**
```json
{"tool": "version_constraint_check", "arguments": {"expression": "~1.2.3", "version": "1.2.9", "type": "single"}}
```

### Filter a version list by constraint

**Goal:** From `["0.5.0", "1.0.0", "1.5.0", "2.0.0"]`, keep only versions matching `">=1.0.0,<2.0.0"`.

**SDK approach:**
```go
cs, _ := versions.ParseConstraintSet(">=1.0.0,<2.0.0")
versionList := versions.NewVersions("0.5.0", "1.0.0", "1.5.0", "2.0.0")
filtered := versions.FilterByConstraintSet(versionList, cs)
// filtered = [1.0.0, 1.5.0]
```

**CLI approach:**
```bash
versions filter --constraint ">=1.0.0,<2.0.0" 0.5.0 1.0.0 1.5.0 2.0.0
```

**MCP approach:**
```json
{"tool": "version_filter", "arguments": {"versions": ["0.5.0", "1.0.0", "1.5.0", "2.0.0"], "constraint": ">=1.0.0,<2.0.0"}}
```

## Cross-References

- [[version-check]] — for simpler boolean type checks (IsStable, IsBeta, etc.)
- [[version-comparison]] — for direct pairwise version comparison
- [[version-range-query]] — for range queries on version collections
- [[version-parsing]] — for parsing version strings before constraint checking

## Important Notes

- **Caret `^` follows npm semver semantics**: `^1.2.3` = `>=1.2.3,<2.0.0`, but `^0.2.3` = `>=0.2.3,<0.3.0` (more restrictive for 0.x)
- **Tilde `~` allows patch-level changes only**: `~1.2.3` matches `1.2.x` but NOT `1.3.0`
- **Wildcards `x/X/*` are equivalent**: `1.x` = `1.X` = `1.*`
- **ConstraintSet (comma) = AND, ConstraintUnion (`||`) = OR**
- **Spaces around operators are NOT supported**: use `>=1.0.0` not `>= 1.0.0`
- **ParseConstraintSet wraps errors** with context (`parse constraint "BAD": ...`); ParseConstraintUnion does NOT
- **CLI constraint type defaults to "set" (AND)**; use `--type single` for single constraint, `--type union` for OR logic
