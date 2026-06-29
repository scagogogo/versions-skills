#!/usr/bin/env python3
"""
Generate a flat-style capability tree diagram — refined version.
Color palette: blue-gray, no purple, square corners, clean spacing.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch

# ─── Data ───
tree = {
    "versions-skills": [
        ("Parse & Validate", ["NewVersion / MustParse", "Coerce (from string)", "Validate / ValidateSemver", "IsSemver / IsValid"]),
        ("Compare", ["CompareTo / IsNewerThan", "IsOlderThan / Equals", "IsBetween / Diff"]),
        ("Sort & Filter", ["SortVersionSlice", "SortVersionStringSlice", "Filter / FilterByStable", "FilterByConstraint / Unique"]),
        ("Group", ["GroupByMajor", "GroupByMinor", "Group (custom depth)", "SortedVersionGroups"]),
        ("Constraints", ["ParseConstraint", "ParseConstraintSet", "ParseConstraintUnion", "NegateConstraint", "Satisfies / Matches"]),
        ("Range Query", ["NewClosedRange", "NewOpenRange", "Contains / Filter", "O(log n) search"]),
        ("Type Check", ["IsStable / IsPrerelease", "IsAlpha / IsBeta / IsRC", "IsDev / IsSnapshot / IsNightly", "IsFinal / IsGA / IsSP / IsPost"]),
        ("Mutate & Build", ["BumpMajor / BumpMinor / BumpPatch", "WithPrefix / WithSuffix", "WithMajor / WithMinor", "VersionBuilder (fluent)"]),
        ("File I/O", ["ReadFromFile", "WriteToFile", "ReadFromReader"]),
        ("Visualization", ["VisualizeVersions", "VisualizeGroups"]),
        ("Serialization", ["JSON Marshal/Unmarshal", "Text Marshal/Unmarshal", "SQL Scan / Value"]),
        ("Set Operations", ["Min / Max / LatestStable", "Contains / IndexOf", "Difference / Intersection", "Union / Partition"]),
    ],
}

# Flat palette — blue-gray, no purple
cat_colors = [
    "#2563eb", "#0ea5e9", "#16a34a", "#0891b2",
    "#ea580c", "#dc2626", "#0d9488", "#d97706",
    "#059669", "#0284c7", "#15803d", "#0369a1",
]

# ─── Layout computation ───
def compute_positions(tree_data, x_start=0, y_center=0, level_gap=4.5, leaf_h=0.56):
    positions = {}
    root = list(tree_data.keys())[0]
    children = tree_data[root]

    total_leaves = sum(max(1, len(subs)) for _, subs in children)
    total_height = total_leaves * leaf_h
    positions[root] = (x_start, y_center)

    y_cursor = y_center + total_height / 2 - leaf_h / 2
    x_l1 = x_start + level_gap
    x_l2 = x_start + 2 * level_gap

    for cat_label, sub_items in children:
        if not sub_items:
            positions[cat_label] = (x_l1, y_cursor)
            y_cursor -= leaf_h
        else:
            sub_ys = []
            for sub_label in sub_items:
                positions[sub_label] = (x_l2, y_cursor)
                sub_ys.append(y_cursor)
                y_cursor -= leaf_h
            positions[cat_label] = (x_l1, sum(sub_ys) / len(sub_ys))
    return positions


# ─── Drawing ───
fig, ax = plt.subplots(figsize=(26, 16))
ax.set_xlim(-3, 24)
ax.set_ylim(-10, 5)
ax.axis("off")
fig.patch.set_facecolor("white")

positions = compute_positions(tree, x_start=0, y_center=-2.5, level_gap=4.8, leaf_h=0.60)
root = list(tree.keys())[0]
children = tree[root]

# Connections — root to category
rx, ry = positions[root]
for i, (cat_label, sub_items) in enumerate(children):
    cx, cy = positions[cat_label]
    color = cat_colors[i % len(cat_colors)]
    # Horizontal then vertical (Manhattan routing)
    mid_x = (rx + 1.5 + cx - 1.6) / 2
    ax.plot([rx + 1.5, mid_x], [ry, ry], color=color, lw=1.6, alpha=0.35, zorder=2)
    ax.plot([mid_x, mid_x], [ry, cy], color=color, lw=1.6, alpha=0.35, zorder=2)
    ax.plot([mid_x, cx - 1.6], [cy, cy], color=color, lw=1.6, alpha=0.35, zorder=2)

    # Category to sub items
    for sub_label in sub_items:
        sx, sy = positions[sub_label]
        mid_x2 = (cx + 1.5 + sx - 1.5) / 2
        ax.plot([cx + 1.5, mid_x2], [cy, cy], color=color, lw=0.8, alpha=0.2, zorder=2)
        ax.plot([mid_x2, mid_x2], [cy, sy], color=color, lw=0.8, alpha=0.2, zorder=2)
        ax.plot([mid_x2, sx - 1.5], [sy, sy], color=color, lw=0.8, alpha=0.2, zorder=2)

# Root node
rx, ry = positions[root]
root_box = FancyBboxPatch((rx - 1.6, ry - 0.32), 3.2, 0.64,
    boxstyle="square,pad=0.04", facecolor="#0f172a", edgecolor="#0f172a", lw=1.5, zorder=5)
ax.add_patch(root_box)
ax.text(rx, ry, root, ha="center", va="center", fontsize=14, fontweight="bold", color="white", zorder=6, family="monospace")

# Category & sub nodes
for i, (cat_label, sub_items) in enumerate(children):
    cx, cy = positions[cat_label]
    color = cat_colors[i % len(cat_colors)]

    # Category box
    cat_box = FancyBboxPatch((cx - 1.6, cy - 0.26), 3.2, 0.52,
        boxstyle="square,pad=0.04", facecolor=color, edgecolor=color, lw=1, alpha=0.92, zorder=5)
    ax.add_patch(cat_box)
    ax.text(cx, cy, cat_label, ha="center", va="center", fontsize=10, fontweight="bold", color="white", zorder=6)

    # Sub item boxes
    for sub_label in sub_items:
        sx, sy = positions[sub_label]
        sub_box = FancyBboxPatch((sx - 1.5, sy - 0.18), 3.0, 0.36,
            boxstyle="square,pad=0.03", facecolor="white", edgecolor=color, lw=1, zorder=5)
        ax.add_patch(sub_box)
        ax.text(sx, sy, sub_label, ha="center", va="center", fontsize=8.5, color="#334155", zorder=6, family="monospace")

# Title
ax.text(12, 4.5, "Capability Map", ha="center", va="center", fontsize=22, fontweight="bold", color="#0f172a")
ax.text(12, 3.7, "12 domains · 3-level tree: Category → Sub-system → Specific API", ha="center", va="center", fontsize=11, color="#64748b")

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/capability-tree.png", dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("capability-tree.png saved")
