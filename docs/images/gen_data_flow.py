#!/usr/bin/env python3
"""
Generate a flat-style data flow diagram.
Color palette: blue-gray, no purple, square corners.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch

# ─── Data ───
stages = [
    ("INPUT",   ["Raw String", "File", "Reader", "Constraint Expr"], "#0f172a"),
    ("PARSE",   ["NewVersion()", "MustParse()", "Coerce()", "ParseConstraint()"], "#2563eb"),
    ("PROCESS", ["CompareTo", "Sort", "Group", "Filter", "Satisfies"], "#16a34a"),
    ("QUERY",   ["Range Search", "Set Ops", "Type Check", "Latest Stable"], "#0891b2"),
    ("OUTPUT",  ["Version Obj", "Sorted Slice", "Group Map", "Bool / Int"], "#ea580c"),
]

# ─── Drawing ───
fig, ax = plt.subplots(figsize=(22, 8))
ax.set_xlim(0, 22)
ax.set_ylim(0, 8)
ax.axis("off")
fig.patch.set_facecolor("white")

# Title
ax.text(11, 7.5, "Data Flow", ha="center", va="center", fontsize=20, fontweight="bold", color="#0f172a")
ax.text(11, 7.0, "5-stage pipeline: Input → Parse → Process → Query → Output", ha="center", va="center", fontsize=10, color="#64748b")

stage_w = 3.6
stage_h = 5.0
gap = 0.45
start_x = 0.6
start_y = 1.2

for idx, (stage_name, items, color) in enumerate(stages):
    x = start_x + idx * (stage_w + gap)
    y = start_y

    # Stage background
    bg = FancyBboxPatch((x, y), stage_w, stage_h,
        boxstyle="square,pad=0.04", facecolor=color, edgecolor=color, lw=1.5, alpha=0.06, zorder=1)
    ax.add_patch(bg)

    # Stage header
    header = FancyBboxPatch((x, y + stage_h - 0.65), stage_w, 0.65,
        boxstyle="square,pad=0", facecolor=color, edgecolor="none", lw=0, alpha=0.92, zorder=2)
    ax.add_patch(header)
    ax.text(x + stage_w / 2, y + stage_h - 0.325, stage_name, ha="center", va="center",
        fontsize=11, fontweight="bold", color="white", zorder=3)

    # Items
    n = len(items)
    item_h = 0.42
    item_gap = 0.12
    item_start_y = y + stage_h - 0.65 - 0.25 - item_h

    for i, item in enumerate(items):
        iy = item_start_y - i * (item_h + item_gap)
        ix = x + 0.2
        iw = stage_w - 0.4
        item_box = FancyBboxPatch((ix, iy), iw, item_h,
            boxstyle="square,pad=0.02", facecolor="white", edgecolor=color, lw=1, zorder=4)
        ax.add_patch(item_box)
        ax.text(ix + iw / 2, iy + item_h / 2, item, ha="center", va="center",
            fontsize=8.5, color="#334155", zorder=5, family="monospace")

    # Arrow to next stage
    if idx < len(stages) - 1:
        arrow_x_start = x + stage_w + 0.02
        arrow_x_end = x + stage_w + gap - 0.02
        arrow_y = y + stage_h / 2
        ax.annotate("", xy=(arrow_x_end, arrow_y), xytext=(arrow_x_start, arrow_y),
            arrowprops=dict(arrowstyle="->", color="#94a3b8", lw=2), zorder=6)

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/data-flow.png", dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("data-flow.png saved")
