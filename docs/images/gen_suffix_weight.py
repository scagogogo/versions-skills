#!/usr/bin/env python3
"""
Generate a flat-style suffix weight ladder diagram.
Color palette: blue-gray, no purple, square corners.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch

# ─── Data ───
suffixes = [
    ("dev",      50,  "#64748b"),
    ("snapshot", 60,  "#64748b"),
    ("nightly",  70,  "#64748b"),
    ("alpha",    100, "#dc2626"),
    ("beta",     200, "#ea580c"),
    ("milestone",300, "#d97706"),
    ("rc",       400, "#0891b2"),
    ("final",    500, "#16a34a"),
    ("release",  500, "#16a34a"),
    ("ga",       500, "#16a34a"),
    ("sp",       600, "#0ea5e9"),
    ("patch",    700, "#0ea5e9"),
    ("post",     800, "#2563eb"),
]

# ─── Drawing ───
fig, ax = plt.subplots(figsize=(20, 8))
ax.set_xlim(-1, 21)
ax.set_ylim(-1.5, 5)
ax.axis("off")
fig.patch.set_facecolor("white")

# Title
ax.text(10, 4.5, "Suffix Weight Order", ha="center", va="center", fontsize=20, fontweight="bold", color="#0f172a")
ax.text(10, 3.9, "Comparison priority: VersionNumbers → Suffix → PublicTime → Raw string", ha="center", va="center", fontsize=10, color="#64748b")

# Horizontal axis line
ax.plot([0.5, 20], [1.0, 1.0], color="#cbd5e1", lw=2, zorder=1)

# Axis labels
ax.text(0.2, 1.0, "Lighter", ha="center", va="center", fontsize=9, color="#94a3b8")
ax.text(20.3, 1.0, "Heavier", ha="center", va="center", fontsize=9, color="#94a3b8")

# Arrow at end
ax.annotate("", xy=(20.2, 1.0), xytext=(19.8, 1.0),
    arrowprops=dict(arrowstyle="->", color="#94a3b8", lw=2), zorder=2)

# Suffix boxes
min_w = 50
max_w = 800
x_start = 1.0
x_end = 19.5
total_range = x_end - x_start

# Group by weight for positioning
seen_weights = {}
for suffix, weight, color in suffixes:
    if weight not in seen_weights:
        seen_weights[weight] = []
    seen_weights[weight].append((suffix, color))

# Draw by unique weights, stacking same-weight items vertically
y_base = 1.6
box_w = 1.3
box_h = 0.40
stack_gap = 0.45

for weight_idx, (weight, items) in enumerate(sorted(seen_weights.items())):
    x = x_start + (weight - min_w) / (max_w - min_w) * total_range
    color = items[0][1]  # Use color from first item

    for stack_idx, (suffix, s_color) in enumerate(items):
        y = y_base + stack_idx * stack_gap

        # Box
        s_box = FancyBboxPatch((x - box_w / 2, y - box_h / 2), box_w, box_h,
            boxstyle="square,pad=0.03", facecolor=s_color, edgecolor=s_color, lw=1.2, alpha=0.88, zorder=5)
        ax.add_patch(s_box)
        ax.text(x, y, suffix, ha="center", va="center", fontsize=9, fontweight="bold", color="white", zorder=6)

        # Tick mark connecting to axis
        ax.plot([x, x], [1.0, y - box_h / 2 - 0.02], color=s_color, lw=1, alpha=0.4, zorder=2)

    # Weight label on axis
    ax.text(x, 0.55, str(weight), ha="center", va="center", fontsize=8, color="#64748b", fontweight="bold")
    ax.plot([x, x], [0.9, 1.1], color="#94a3b8", lw=1.5, zorder=3)

# Legend
legend_items = [
    ("Pre-release", "#64748b"),
    ("Alpha/Beta", "#dc2626"),
    ("RC", "#0891b2"),
    ("Stable", "#16a34a"),
    ("Post-release", "#0ea5e9"),
    ("Post-stable", "#2563eb"),
]
legend_x = 1.0
legend_y = -0.5
for i, (label, color) in enumerate(legend_items):
    lx = legend_x + i * 3.2
    box = FancyBboxPatch((lx, legend_y - 0.12), 0.3, 0.24,
        boxstyle="square,pad=0", facecolor=color, edgecolor=color, lw=0, alpha=0.88, zorder=5)
    ax.add_patch(box)
    ax.text(lx + 0.5, legend_y, label, ha="left", va="center", fontsize=8.5, color="#475569")

# Example
ax.text(10, -1.1, 'Example: 1.0.0-alpha < 1.0.0-beta < 1.0.0-rc < 1.0.0 < 1.0.0-sp1',
    ha="center", va="center", fontsize=10, color="#475569",
    bbox=dict(boxstyle="square,pad=0.3", facecolor="#f8fafc", edgecolor="#e2e8f0"))

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/suffix-weight.png", dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("suffix-weight.png saved")
