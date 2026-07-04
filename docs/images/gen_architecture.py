#!/usr/bin/env python3
"""
Generate a refined 4-layer architecture diagram with modern styling.
Features: rounded corners, gradient fills, shadows, elegant typography.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

# ─── Data ───
layers = [
    ("AI Agent Layer", ["Claude Code", "Cursor", "Windsurf", "Cline", "Aider"], "#3b82f6"),
    ("Access Layer",    ["Skills (13)", "MCP Server", "Go SDK", "CLI"],           "#06b6d4"),
    ("Core Library",    ["Parse", "Compare", "Sort", "Group", "Constraint", "Range", "TypeCheck", "Mutate", "FileIO", "Visualize", "Serialize", "SetOps"], "#10b981"),
    ("Foundation",      ["Go Runtime", "Immutable Design", "Zero External Deps (core)"], "#0284c7"),
]

def draw_shadow(ax, x, y, w, h, radius=0.15, offset=0.08, alpha=0.12):
    shadow = FancyBboxPatch(
        (x + offset, y - offset), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#000000", edgecolor="none", alpha=alpha, zorder=3
    )
    ax.add_patch(shadow)

def draw_gradient_box(ax, x, y, w, h, color, radius=0.15, alpha=0.95):
    box = FancyBboxPatch(
        (x, y), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor=color, edgecolor="none", alpha=alpha, zorder=5
    )
    ax.add_patch(box)
    highlight_h = h * 0.3
    highlight = FancyBboxPatch(
        (x, y + h - highlight_h), w, highlight_h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#ffffff", edgecolor="none", alpha=0.15, zorder=6
    )
    ax.add_patch(highlight)

# ─── Drawing ───
fig, ax = plt.subplots(figsize=(22, 12))
ax.set_xlim(0, 22)
ax.set_ylim(0, 12)
ax.axis("off")
fig.patch.set_facecolor("#fafbfc")

# Background grid
for i in range(0, 23, 2):
    ax.axvline(i, color="#f1f5f9", lw=0.5, zorder=0)
for i in range(0, 13, 2):
    ax.axhline(i, color="#f1f5f9", lw=0.5, zorder=0)

# Title
ax.text(11, 11.3, "Architecture", ha="center", va="center", fontsize=24, fontweight="bold", color="#0f172a", family="sans-serif")
ax.text(11, 10.6, "4-layer design: AI Agent → Access → Core → Foundation", ha="center", va="center", fontsize=11, color="#64748b", family="sans-serif")

layer_y = 8.5
layer_gap = 2.3
box_h = 1.7
margin_x = 1.0

for layer_idx, (layer_name, items, color) in enumerate(layers):
    y = layer_y - layer_idx * layer_gap
    x = margin_x
    w = 22 - 2 * margin_x

    # Layer background with subtle fill
    layer_bg = FancyBboxPatch(
        (x, y - box_h / 2), w, box_h,
        boxstyle="round,pad=0,rounding_size=0.2",
        facecolor=color, edgecolor="none", lw=0, alpha=0.06, zorder=1
    )
    ax.add_patch(layer_bg)

    # Layer left accent bar (gradient-like)
    accent = FancyBboxPatch(
        (x, y - box_h / 2), 0.15, box_h,
        boxstyle="round,pad=0,rounding_size=0.05",
        facecolor=color, edgecolor="none", lw=0, alpha=0.9, zorder=2
    )
    ax.add_patch(accent)

    # Layer label
    ax.text(x + 0.5, y + box_h / 2 - 0.3, layer_name, va="center", fontsize=12, fontweight="bold", color=color, zorder=3, family="sans-serif")

    # Item boxes with rounded corners and shadows
    n = len(items)
    item_margin = 0.7
    item_gap = 0.3
    total_item_w = w - 2 * item_margin - 0.3
    item_w = (total_item_w - (n - 1) * item_gap) / n
    item_h = 0.6

    for i, item in enumerate(items):
        ix = x + item_margin + 0.3 + i * (item_w + item_gap)
        iy = y - box_h / 2 + 0.25

        # Shadow
        draw_shadow(ax, ix, iy, item_w, item_h, radius=0.1, offset=0.05, alpha=0.1)

        # White box with colored border
        item_box = FancyBboxPatch(
            (ix, iy), item_w, item_h,
            boxstyle="round,pad=0,rounding_size=0.1",
            facecolor="white", edgecolor=color, lw=1.5, zorder=4
        )
        ax.add_patch(item_box)
        ax.text(ix + item_w / 2, iy + item_h / 2, item, ha="center", va="center",
            fontsize=9, color="#334155", zorder=5, family="monospace", weight="medium")

    # Arrow to next layer (modern style)
    if layer_idx < len(layers) - 1:
        arrow_y_start = y - box_h / 2 - 0.08
        arrow_y_end = y - layer_gap + box_h / 2 + 0.08
        ax.annotate("", xy=(11, arrow_y_end), xytext=(11, arrow_y_start),
            arrowprops=dict(arrowstyle="->,head_width=0.4,head_length=0.5", color="#94a3b8", lw=2), zorder=3)

# Decorative corner accents
corner_size = 0.8
corner_color = "#e2e8f0"
ax.plot([0.5, 0.5], [11 - corner_size, 11], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([0.5, 0.5 + corner_size], [11, 11], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([21.5, 21.5], [1 + corner_size, 1], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([21.5, 21.5 - corner_size], [1, 1], color=corner_color, lw=3, solid_capstyle='round')

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/architecture.png", dpi=150, bbox_inches="tight", facecolor="#fafbfc")
plt.close()
print("architecture.png saved")
