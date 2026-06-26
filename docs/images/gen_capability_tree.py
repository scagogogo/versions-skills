#!/usr/bin/env python3
"""
Generate a left-to-right capability tree diagram (fishbone/mind-map style)
for the versions-skills project. No emoji - uses colored icons instead.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
from matplotlib.patches import FancyBboxPatch
import numpy as np

# ─── Data: capability tree ───
tree = {
    "versions-skills": [
        ("Parse & Validate", [
            ("NewVersion / MustParse", []),
            ("Coerce (extract from string)", []),
            ("Validate / ValidateSemver", []),
            ("IsSemver / IsValid", []),
        ]),
        ("Compare", [
            ("CompareTo / IsNewerThan", []),
            ("IsOlderThan / Equals", []),
            ("IsBetween / Diff", []),
        ]),
        ("Sort & Filter", [
            ("SortVersionSlice", []),
            ("SortVersionStringSlice", []),
            ("Filter / FilterByStable", []),
            ("FilterByConstraint / Unique", []),
        ]),
        ("Group", [
            ("GroupByMajor", []),
            ("GroupByMinor", []),
            ("Group (custom depth)", []),
            ("SortedVersionGroups", []),
        ]),
        ("Constraints", [
            ("ParseConstraint", []),
            ("ParseConstraintSet", []),
            ("ParseConstraintUnion", []),
            ("NegateConstraint", []),
            ("Satisfies / Matches", []),
        ]),
        ("Range Query", [
            ("NewClosedRange / NewOpenRange", []),
            ("Contains / Filter", []),
            ("O(log n) binary search", []),
        ]),
        ("Type Check", [
            ("IsStable / IsPrerelease", []),
            ("IsAlpha / IsBeta / IsRC", []),
            ("IsDev / IsSnapshot / IsNightly", []),
            ("IsFinal / IsGA / IsSP / IsPost", []),
        ]),
        ("Mutate & Build", [
            ("BumpMajor / BumpMinor / BumpPatch", []),
            ("WithPrefix / WithSuffix", []),
            ("WithMajor / WithMinor / WithPatch", []),
            ("VersionBuilder (fluent)", []),
        ]),
        ("File I/O", [
            ("ReadVersionsFromFile", []),
            ("WriteVersionsToFile", []),
            ("ReadVersionsFromReader", []),
        ]),
        ("Visualization", [
            ("VisualizeVersions", []),
            ("VisualizeVersionGroups", []),
        ]),
        ("Serialization", [
            ("JSON Marshal/Unmarshal", []),
            ("Text Marshal/Unmarshal", []),
            ("SQL Scan / Value", []),
        ]),
        ("Set Operations", [
            ("Min / Max / LatestStable", []),
            ("Contains / IndexOf", []),
            ("Difference / Intersection", []),
            ("Union / Partition", []),
        ]),
    ],
}

# Category colors for visual distinction
cat_colors = [
    "#3b82f6",  # Parse - blue
    "#8b5cf6",  # Compare - purple
    "#10b981",  # Sort - green
    "#06b6d4",  # Group - cyan
    "#f59e0b",  # Constraints - amber
    "#ef4444",  # Range - red
    "#ec4899",  # Type Check - pink
    "#f97316",  # Mutate - orange
    "#14b8a6",  # File I/O - teal
    "#22c55e",  # Visualization - lime
    "#6366f1",  # Serialization - indigo
    "#a855f7",  # Set Ops - violet
]

# ─── Layout engine ───
def compute_positions(tree_data, x_start=0, y_center=0, level_width=3.5, leaf_height=0.55):
    """Compute x, y positions for each node in a left-to-right tree."""
    positions = {}
    root = list(tree_data.keys())[0]
    children = tree_data[root]

    total_leaves = sum(
        max(1, len(sub_children)) if sub_children else 1
        for _, sub_children in children
    )
    total_height = total_leaves * leaf_height

    positions[root] = (x_start, y_center)

    y_cursor = y_center + total_height / 2 - leaf_height / 2
    x_level1 = x_start + level_width
    x_level2 = x_start + 2 * level_width

    for cat_label, sub_items in children:
        if not sub_items:
            positions[cat_label] = (x_level1, y_cursor)
            y_cursor -= leaf_height
        else:
            sub_ys = []
            for sub_label, _ in sub_items:
                positions[sub_label] = (x_level2, y_cursor)
                sub_ys.append(y_cursor)
                y_cursor -= leaf_height
            cat_y = sum(sub_ys) / len(sub_ys)
            positions[cat_label] = (x_level1, cat_y)

    return positions


def draw_connections(ax, tree_data, positions, level_width=3.5):
    """Draw lines connecting parent to children."""
    root = list(tree_data.keys())[0]
    children = tree_data[root]
    root_x, root_y = positions[root]

    for i, (cat_label, sub_items) in enumerate(children):
        cat_x, cat_y = positions[cat_label]
        color = cat_colors[i % len(cat_colors)]

        # Root -> Category (thicker, colored)
        ax.plot([root_x + 1.6, root_x + level_width - 1.55],
                [root_y, cat_y],
                color=color, lw=2.0, alpha=0.6, zorder=2)

        for sub_label, _ in sub_items:
            sub_x, sub_y = positions[sub_label]
            ax.plot([cat_x + 1.55, sub_x - 1.55],
                    [cat_y, sub_y],
                    color=color, lw=1.0, alpha=0.35, zorder=2)


def draw_nodes(ax, positions, tree_data):
    """Draw rounded-rect boxes for each node."""
    root = list(tree_data.keys())[0]
    children = tree_data[root]

    # Root node
    rx, ry = positions[root]
    root_box = FancyBboxPatch(
        (rx - 1.6, ry - 0.32), 3.2, 0.64,
        boxstyle="round,pad=0.12", facecolor="#1e293b", edgecolor="#0f172a", lw=2.5, zorder=5
    )
    ax.add_patch(root_box)
    ax.text(rx, ry, root, ha="center", va="center", fontsize=14, fontweight="bold",
            color="white", zorder=6, family="monospace")

    # Category nodes
    for i, (cat_label, sub_items) in enumerate(children):
        cx, cy = positions[cat_label]
        color = cat_colors[i % len(cat_colors)]
        cat_box = FancyBboxPatch(
            (cx - 1.55, cy - 0.24), 3.1, 0.48,
            boxstyle="round,pad=0.08", facecolor=color,
            edgecolor=color, lw=1.5, alpha=0.9, zorder=5
        )
        ax.add_patch(cat_box)
        ax.text(cx, cy, cat_label, ha="center", va="center", fontsize=9.5,
                fontweight="bold", color="white", zorder=6)

        for sub_label, _ in sub_items:
            sx, sy = positions[sub_label]
            sub_box = FancyBboxPatch(
                (sx - 1.55, sy - 0.18), 3.1, 0.36,
                boxstyle="round,pad=0.06", facecolor="white",
                edgecolor=color, lw=1.2, zorder=5
            )
            ax.add_patch(sub_box)
            ax.text(sx, sy, sub_label, ha="center", va="center", fontsize=7.8,
                    color="#334155", zorder=6, family="monospace")


# ─── Main ───
fig, ax = plt.subplots(figsize=(24, 16))
ax.set_xlim(-2.5, 24)
ax.set_ylim(-9, 5)
ax.axis("off")
fig.patch.set_facecolor("white")

positions = compute_positions(tree, x_start=0, y_center=-2, level_width=4.0, leaf_height=0.62)
draw_connections(ax, tree, positions, level_width=4.0)
draw_nodes(ax, positions, tree)

# Title
ax.text(12, 4.2, "versions-skills  --  Capability Map", ha="center", va="center",
        fontsize=22, fontweight="bold", color="#1e293b", family="monospace")
ax.text(12, 3.5, "Parse | Compare | Sort | Group | Constraints | Range | Check | Mutate | File | Visualize | Serialize | SetOps",
        ha="center", va="center", fontsize=10, color="#64748b")

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/capability-tree.png",
            dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("capability-tree.png saved")
