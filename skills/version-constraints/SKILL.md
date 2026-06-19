---
name: version-constraints
description: Check whether a version satisfies a constraint expression (caret, tilde, wildcard, range operators). Use when resolving dependency compatibility, filtering versions by semver ranges, or working with constraint expressions that combine AND/OR logic.
argument-hint: <constraint-expression> <version>
---

# Version Constraints

> **Setup:** See `/installation` for one-time SDK/CLI/MCP install.  
> **Layers:** SDK (Go) → CLI (shell) → MCP (AI tools) — pick your entry point.

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

**Caret `^` semantics:** `^1.2.3` = `>=1.2.3,<2.0.0`; `^0.2.3` = `>=0.2.3,<0.3.0` (more restrictive for 0.x); `^0.0.3` = `>=0.0.3,<0.0.4`.

**Tilde `~` semantics:** `~1.2.3` = `>=1.2.3,<1.3.0` (patch-level changes only).

**Wildcard semantics:** `1.x` = `1.X` = `1.*` = `>=1.0.0,<2.0.0`; `1.2.x` = `>=1.2.0,<1.3.0`.

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

## API Reference

### SDK — Parsing Functions

```go
// Parse a single constraint expression (e.g. ">=1.0.0", "^1.2.3", "~1.2.3")
func ParseConstraint(expr string) (*Constraint, error)

// Parse comma-separated AND constraints (e.g. ">=1.0.0,<2.0.0")
func ParseConstraintSet(expr string) (*ConstraintSet, error)

// Parse ||-separated OR constraint sets (e.g. ">=1.0.0,<2.0.0 || >=3.0.0")
func ParseConstraintUnion(expr string) (*ConstraintUnion, error)
```

### SDK — Matching Methods

```go
// Single constraint matching
func (c *Constraint) Match(v *Version) bool
func (c *Constraint) String() string

// ConstraintSet (AND logic) — all must match
func (cs *ConstraintSet) Match(v *Version) bool
func (cs *ConstraintSet) Satisfies(v *Version) bool   // alias for Match
func (cs *ConstraintSet) Len() int
func (cs *ConstraintSet) String() string

// ConstraintUnion (OR logic) — any set must match
func (cu *ConstraintUnion) Match(v *Version) bool
func (cu *ConstraintUnion) Satisfies(v *Version) bool // alias for Match
func (cu *ConstraintUnion) String() string

// Version-centric convenience methods
func (v *Version) Satisfies(constraint *Constraint) bool
func (v *Version) Matches(expr string) (bool, error)  // parse + match one-liner
```

### SDK — Filter Functions

```go
func FilterByConstraint(versions []*Version, constraint *Constraint) []*Version
func FilterByConstraintSet(versions []*Version, cs *ConstraintSet) []*Version
```

### CLI Commands

```bash
# Check constraint — single type
versions constraint "<expr>" <version> --type single

# Check constraint — ConstraintSet (AND logic, default)
versions constraint "<expr>" <version>

# Check constraint — ConstraintUnion (OR logic)
versions constraint "<expr>" <version> --type union

# Version-centric check (auto-detects type)
versions satisfies <version> "<constraint-expression>"

# Filter by constraint — single
versions filter --constraint "<expr>" --constraint-type single <versions...>

# Filter by constraint — set (default)
versions filter --constraint "<expr>" <versions...>

# Filter by constraint — union
versions filter --constraint "<expr>" --constraint-type union <versions...>
```

**Examples:**
```bash
versions constraint ">=1.0.0,<2.0.0" 1.5.0         # AND logic
versions constraint ">=1.0.0 || >=3.0.0" 3.5.0 --type union
versions satisfies 1.5.0 ">=1.0.0,<2.0.0"
versions filter --constraint ">=1.0.0,<2.0.0" 0.5.0 1.0.0 1.5.0 2.0.0
```

### MCP Tools

| Tool | Arguments | Returns |
|------|-----------|---------|
| `version_constraint_check` | `expression: string`, `version: string`, `type?: string` | `{matches: bool}` |
| `version_filter` | `versions: string[]`, `constraint?: string`, `stable?: bool`, `prerelease?: bool`, `major?: int`, `minor?: int`, `patch?: int`, `prefix?: string`, `suffix?: string` | filtered list |

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
