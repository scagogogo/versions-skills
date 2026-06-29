#!/usr/bin/env python3
"""
Generate a flat-style 4-layer architecture diagram.
Color palette: blue-gray, no purple, square corners.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch

# ─── Data ───
layers = [
    ("AI Agent Layer", ["Claude Code", "Cursor", "Windsurf", "Cline", "Aider"], "#2563eb"),
    ("Access Layer",    ["Skills (13)", "MCP Server", "Go SDK", "CLI"],           "#0ea5e9"),
    ("Core Library",    ["Parse", "Compare", "Sort", "Group", "Constraint", "Range", "TypeCheck", "Mutate", "FileIO", "Visualize", "Serialize", "SetOps"], "#16a34a"),
    ("Foundation",      ["Go Runtime", "Immutable Design", "Zero External Deps (core)"], "#0891b2"),
]

# ─── Drawing ───
fig, ax = plt.subplots(figsize=(20, 10))
ax.set_xlim(0, 20)
ax.set_ylim(0, 10)
ax.axis("off")
fig.patch.set_facecolor("white")

# Title
ax.text(10, 9.5, "Architecture", ha="center", va="center", fontsize=20, fontweight="bold", color="#0f172a")
ax.text(10, 9.0, "4-layer design: AI Agent → Access → Core → Foundation", ha="center", va="center", fontsize=10, color="#64748b")

layer_y = 7.2
layer_gap = 2.1
box_h = 1.5
margin_x = 0.8

for layer_idx, (layer_name, items, color) in enumerate(layers):
    y = layer_y - layer_idx * layer_gap
    x = margin_x
    w = 20 - 2 * margin_x

    # Layer background
    layer_bg = FancyBboxPatch((x, y - box_h / 2), w, box_h,
        boxstyle="square,pad=0.04", facecolor=color, edgecolor=color, lw=0, alpha=0.08, zorder=1)
    ax.add_patch(layer_bg)

    # Layer border (left accent)
    accent = FancyBboxPatch((x, y - box_h / 2), 0.12, box_h,
        boxstyle="square,pad=0", facecolor=color, edgecolor="none", lw=0, alpha=0.9, zorder=2)
    ax.add_patch(accent)

    # Layer label
    ax.text(x + 0.4, y + box_h / 2 - 0.28, layer_name, va="center", fontsize=11, fontweight="bold", color=color, zorder=3)

    # Item boxes
    n = len(items)
    item_margin = 0.6
    item_gap = 0.25
    total_item_w = w - 2 * item_margin
    item_w = (total_item_w - (n - 1) * item_gap) / n
    item_h = 0.50

    for i, item in enumerate(items):
        ix = x + item_margin + i * (item_w + item_gap)
        iy = y - box_h / 2 + 0.22
        item_box = FancyBboxPatch((ix, iy), item_w, item_h,
            boxstyle="square,pad=0.03", facecolor="white", edgecolor=color, lw=1.2, zorder=4)
        ax.add_patch(item_box)
        ax.text(ix + item_w / 2, iy + item_h / 2, item, ha="center", va="center",
            fontsize=8.5, color="#334155", zorder=5, family="monospace")

    # Arrow to next layer
    if layer_idx < len(layers) - 1:
        arrow_y_start = y - box_h / 2 - 0.05
        arrow_y_end = y - layer_gap + box_h / 2 + 0.05
        ax.annotate("", xy=(10, arrow_y_end), xytext=(10, arrow_y_start),
            arrowprops=dict(arrowstyle="->", color="#94a3b8", lw=1.5), zorder=3)

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/architecture.png", dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("architecture.png saved")
