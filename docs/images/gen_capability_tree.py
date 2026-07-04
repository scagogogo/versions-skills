#!/usr/bin/env python3
"""
Generate a refined capability tree diagram with modern styling.
Features: rounded corners, gradient fills, shadows, elegant typography.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch, Rectangle
import numpy as np

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

# Refined color palette — modern blue-gray, no purple
cat_colors = [
    "#3b82f6",  # blue
    "#06b6d4",  # cyan
    "#10b981",  # emerald
    "#0ea5e9",  # sky
    "#f59e0b",  # amber
    "#ef4444",  # red
    "#14b8a6",  # teal
    "#f97316",  # orange
    "#22c55e",  # green
    "#0284c7",  # deeper sky
    "#16a34a",  # deeper green
    "#0369a1",  # deep blue
]

def draw_shadow(ax, x, y, w, h, radius=0.15, offset=0.08, alpha=0.12):
    """Draw a subtle shadow under a box."""
    shadow = FancyBboxPatch(
        (x + offset, y - offset), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#000000", edgecolor="none", alpha=alpha, zorder=3
    )
    ax.add_patch(shadow)

def draw_gradient_box(ax, x, y, w, h, color, radius=0.15, alpha_start=0.95, alpha_end=0.75):
    """Draw a box with gradient-like effect using multiple layers."""
    # Main box
    box = FancyBboxPatch(
        (x, y), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor=color, edgecolor="none", alpha=alpha_start, zorder=5
    )
    ax.add_patch(box)
    # Subtle top highlight
    highlight_h = h * 0.3
    highlight = FancyBboxPatch(
        (x, y + h - highlight_h), w, highlight_h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#ffffff", edgecolor="none", alpha=0.15, zorder=6
    )
    ax.add_patch(highlight)

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
fig, ax = plt.subplots(figsize=(28, 18))
ax.set_xlim(-3, 26)
ax.set_ylim(-11, 6)
ax.axis("off")
fig.patch.set_facecolor("#fafbfc")

positions = compute_positions(tree, x_start=0, y_center=-2.5, level_gap=5.2, leaf_h=0.65)
root = list(tree.keys())[0]
children = tree[root]

# Background pattern - subtle grid
for i in range(-3, 27, 2):
    ax.axvline(i, color="#f1f5f9", lw=0.5, zorder=0)
for i in range(-11, 7, 2):
    ax.axhline(i, color="#f1f5f9", lw=0.5, zorder=0)

# Connections — smooth curved lines
rx, ry = positions[root]
for i, (cat_label, sub_items) in enumerate(children):
    cx, cy = positions[cat_label]
    color = cat_colors[i % len(cat_colors)]

    # Draw curved connection using bezier-like path
    mid_x = (rx + 1.8 + cx - 1.8) / 2

    # Horizontal from root
    ax.plot([rx + 1.8, mid_x], [ry, ry], color=color, lw=2, alpha=0.3, zorder=2, solid_capstyle='round')
    # Vertical drop
    ax.plot([mid_x, mid_x], [ry, cy], color=color, lw=2, alpha=0.3, zorder=2, solid_capstyle='round')
    # Horizontal to category
    ax.plot([mid_x, cx - 1.8], [cy, cy], color=color, lw=2, alpha=0.3, zorder=2, solid_capstyle='round')

    # Category to sub items
    for sub_label in sub_items:
        sx, sy = positions[sub_label]
        mid_x2 = (cx + 1.7 + sx - 1.6) / 2
        ax.plot([cx + 1.7, mid_x2], [cy, cy], color=color, lw=1, alpha=0.2, zorder=2, solid_capstyle='round')
        ax.plot([mid_x2, mid_x2], [cy, sy], color=color, lw=1, alpha=0.2, zorder=2, solid_capstyle='round')
        ax.plot([mid_x2, sx - 1.6], [sy, sy], color=color, lw=1, alpha=0.2, zorder=2, solid_capstyle='round')

# Root node — elegant dark with gradient effect
rx, ry = positions[root]
draw_shadow(ax, rx - 1.8, ry - 0.38, 3.6, 0.76, radius=0.2, offset=0.1)
draw_gradient_box(ax, rx - 1.8, ry - 0.38, 3.6, 0.76, "#0f172a", radius=0.2)
ax.text(rx, ry, root, ha="center", va="center", fontsize=15, fontweight="bold", color="white", zorder=7, family="sans-serif")

# Category & sub nodes
for i, (cat_label, sub_items) in enumerate(children):
    cx, cy = positions[cat_label]
    color = cat_colors[i % len(cat_colors)]

    # Category box with shadow and gradient
    draw_shadow(ax, cx - 1.8, cy - 0.32, 3.6, 0.64, radius=0.15, offset=0.06)
    draw_gradient_box(ax, cx - 1.8, cy - 0.32, 3.6, 0.64, color, radius=0.15)
    ax.text(cx, cy, cat_label, ha="center", va="center", fontsize=11, fontweight="600", color="white", zorder=7)

    # Sub item boxes — clean white with colored border
    for sub_label in sub_items:
        sx, sy = positions[sub_label]

        # Subtle shadow
        draw_shadow(ax, sx - 1.6, sy - 0.22, 3.2, 0.44, radius=0.1, offset=0.04, alpha=0.08)

        # White box with colored border
        sub_box = FancyBboxPatch(
            (sx - 1.6, sy - 0.22), 3.2, 0.44,
            boxstyle="round,pad=0,rounding_size=0.1",
            facecolor="white", edgecolor=color, lw=1.5, zorder=5
        )
        ax.add_patch(sub_box)

        # Monospace text for API names
        ax.text(sx, sy, sub_label, ha="center", va="center", fontsize=9, color="#334155", zorder=6, family="monospace", weight="medium")

# Title with modern typography
ax.text(12, 5.2, "Capability Map", ha="center", va="center", fontsize=26, fontweight="bold", color="#0f172a", family="sans-serif")
ax.text(12, 4.3, "12 domains · 3-level tree: Category → Sub-system → Specific API", ha="center", va="center", fontsize=12, color="#64748b", family="sans-serif")

# Decorative elements — subtle corner accents
corner_size = 0.8
corner_color = "#e2e8f0"
# Top left
ax.plot([-2.5, -2.5], [5.5 - corner_size, 5.5], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([-2.5, -2.5 + corner_size], [5.5, 5.5], color=corner_color, lw=3, solid_capstyle='round')
# Bottom right
ax.plot([25, 25], [-10 + corner_size, -10], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([25, 25 - corner_size], [-10, -10], color=corner_color, lw=3, solid_capstyle='round')

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/capability-tree.png", dpi=150, bbox_inches="tight", facecolor="#fafbfc")
plt.close()
print("capability-tree.png saved")
