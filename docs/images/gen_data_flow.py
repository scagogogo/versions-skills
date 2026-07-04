#!/usr/bin/env python3
"""
Generate a refined data flow diagram with modern styling.
Features: rounded corners, gradient fills, shadows, elegant typography.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

# ─── Data ───
stages = [
    ("INPUT",   ["Raw String", "File", "Reader", "Constraint Expr"], "#0f172a"),
    ("PARSE",   ["NewVersion()", "MustParse()", "Coerce()", "ParseConstraint()"], "#3b82f6"),
    ("PROCESS", ["CompareTo", "Sort", "Group", "Filter", "Satisfies"], "#10b981"),
    ("QUERY",   ["Range Search", "Set Ops", "Type Check", "Latest Stable"], "#06b6d4"),
    ("OUTPUT",  ["Version Obj", "Sorted Slice", "Group Map", "Bool / Int"], "#f97316"),
]

def draw_shadow(ax, x, y, w, h, radius=0.15, offset=0.06, alpha=0.1):
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
fig, ax = plt.subplots(figsize=(24, 10))
ax.set_xlim(0, 24)
ax.set_ylim(0, 10)
ax.axis("off")
fig.patch.set_facecolor("#fafbfc")

# Background grid
for i in range(0, 25, 2):
    ax.axvline(i, color="#f1f5f9", lw=0.5, zorder=0)
for i in range(0, 11, 2):
    ax.axhline(i, color="#f1f5f9", lw=0.5, zorder=0)

# Title
ax.text(12, 9.2, "Data Flow", ha="center", va="center", fontsize=24, fontweight="bold", color="#0f172a", family="sans-serif")
ax.text(12, 8.5, "5-stage pipeline: Input → Parse → Process → Query → Output", ha="center", va="center", fontsize=11, color="#64748b", family="sans-serif")

stage_w = 4.0
stage_h = 5.5
gap = 0.6
start_x = 0.8
start_y = 1.5

for idx, (stage_name, items, color) in enumerate(stages):
    x = start_x + idx * (stage_w + gap)
    y = start_y

    # Stage background — subtle rounded card
    draw_shadow(ax, x, y, stage_w, stage_h, radius=0.2, offset=0.08, alpha=0.08)
    bg = FancyBboxPatch(
        (x, y), stage_w, stage_h,
        boxstyle="round,pad=0,rounding_size=0.2",
        facecolor=color, edgecolor=color, lw=1.5, alpha=0.05, zorder=1
    )
    ax.add_patch(bg)

    # Stage header — gradient effect
    header = FancyBboxPatch(
        (x, y + stage_h - 0.8), stage_w, 0.8,
        boxstyle="round,pad=0,rounding_size=0.15",
        facecolor=color, edgecolor="none", lw=0, alpha=0.95, zorder=2
    )
    ax.add_patch(header)
    # Header highlight
    highlight = FancyBboxPatch(
        (x, y + stage_h - 0.3), stage_w, 0.3,
        boxstyle="round,pad=0,rounding_size=0.1",
        facecolor="#ffffff", edgecolor="none", alpha=0.18, zorder=3
    )
    ax.add_patch(highlight)
    ax.text(x + stage_w / 2, y + stage_h - 0.4, stage_name, ha="center", va="center",
        fontsize=12, fontweight="bold", color="white", zorder=4, family="sans-serif")

    # Items — clean white cards
    n = len(items)
    item_h = 0.5
    item_gap = 0.18
    item_start_y = y + stage_h - 0.8 - 0.3 - item_h

    for i, item in enumerate(items):
        iy = item_start_y - i * (item_h + item_gap)
        ix = x + 0.25
        iw = stage_w - 0.5

        # Shadow
        draw_shadow(ax, ix, iy, iw, item_h, radius=0.1, offset=0.03, alpha=0.08)

        # White box with colored border
        item_box = FancyBboxPatch(
            (ix, iy), iw, item_h,
            boxstyle="round,pad=0,rounding_size=0.1",
            facecolor="white", edgecolor=color, lw=1.5, zorder=4
        )
        ax.add_patch(item_box)
        ax.text(ix + iw / 2, iy + item_h / 2, item, ha="center", va="center",
            fontsize=9, color="#334155", zorder=5, family="monospace", weight="medium")

    # Modern arrow to next stage
    if idx < len(stages) - 1:
        arrow_x_start = x + stage_w + 0.05
        arrow_x_end = x + stage_w + gap - 0.05
        arrow_y = y + stage_h / 2
        ax.annotate("", xy=(arrow_x_end, arrow_y), xytext=(arrow_x_start, arrow_y),
            arrowprops=dict(arrowstyle="->,head_width=0.4,head_length=0.6", color="#94a3b8", lw=2.5), zorder=6)

# Decorative corner accents
corner_size = 0.8
corner_color = "#e2e8f0"
ax.plot([0.5, 0.5], [9 - corner_size, 9], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([0.5, 0.5 + corner_size], [9, 9], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([23.5, 23.5], [1 + corner_size, 1], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([23.5, 23.5 - corner_size], [1, 1], color=corner_color, lw=3, solid_capstyle='round')

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/data-flow.png", dpi=150, bbox_inches="tight", facecolor="#fafbfc")
plt.close()
print("data-flow.png saved")