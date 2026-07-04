#!/usr/bin/env python3
"""
Generate a refined suffix weight ladder diagram with modern styling.
Features: rounded corners, gradient fills, shadows, elegant typography.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

# ─── Data ───
suffixes = [
    ("dev",      50,  "#64748b"),
    ("snapshot", 60,  "#64748b"),
    ("nightly",  70,  "#64748b"),
    ("alpha",    100, "#ef4444"),
    ("beta",     200, "#f59e0b"),
    ("milestone",300, "#f97316"),
    ("rc",       400, "#06b6d4"),
    ("final",    500, "#10b981"),
    ("release",  500, "#10b981"),
    ("ga",       500, "#10b981"),
    ("sp",       600, "#0ea5e9"),
    ("patch",    700, "#0ea5e9"),
    ("post",     800, "#3b82f6"),
]

def draw_shadow(ax, x, y, w, h, radius=0.12, offset=0.05, alpha=0.1):
    shadow = FancyBboxPatch(
        (x + offset, y - offset), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#000000", edgecolor="none", alpha=alpha, zorder=3
    )
    ax.add_patch(shadow)

def draw_gradient_box(ax, x, y, w, h, color, radius=0.12, alpha=0.92):
    box = FancyBboxPatch(
        (x, y), w, h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor=color, edgecolor="none", alpha=alpha, zorder=5
    )
    ax.add_patch(box)
    highlight_h = h * 0.35
    highlight = FancyBboxPatch(
        (x, y + h - highlight_h), w, highlight_h,
        boxstyle=f"round,pad=0,rounding_size={radius}",
        facecolor="#ffffff", edgecolor="none", alpha=0.18, zorder=6
    )
    ax.add_patch(highlight)

# ─── Drawing ───
fig, ax = plt.subplots(figsize=(22, 10))
ax.set_xlim(-1, 23)
ax.set_ylim(-2, 6)
ax.axis("off")
fig.patch.set_facecolor("#fafbfc")

# Background grid
for i in range(-1, 24, 2):
    ax.axvline(i, color="#f1f5f9", lw=0.5, zorder=0)
for i in range(-2, 7, 2):
    ax.axhline(i, color="#f1f5f9", lw=0.5, zorder=0)

# Title
ax.text(11, 5.3, "Suffix Weight Order", ha="center", va="center", fontsize=24, fontweight="bold", color="#0f172a", family="sans-serif")
ax.text(11, 4.6, "Comparison priority: VersionNumbers → Suffix → PublicTime → Raw string", ha="center", va="center", fontsize=11, color="#64748b", family="sans-serif")

# Main axis line — gradient effect
ax.plot([0.8, 21.5], [1.2, 1.2], color="#cbd5e1", lw=3, zorder=1, solid_capstyle='round')

# Axis labels with modern typography
ax.text(0.3, 1.2, "Lighter", ha="center", va="center", fontsize=10, color="#94a3b8", fontweight="medium")
ax.text(22, 1.2, "Heavier", ha="center", va="center", fontsize=10, color="#94a3b8", fontweight="medium")

# Arrow at end
ax.annotate("", xy=(21.8, 1.2), xytext=(21.3, 1.2),
    arrowprops=dict(arrowstyle="->,head_width=0.5,head_length=0.6", color="#94a3b8", lw=3), zorder=2)

# Group by weight for positioning
seen_weights = {}
for suffix, weight, color in suffixes:
    if weight not in seen_weights:
        seen_weights[weight] = []
    seen_weights[weight].append((suffix, color))

# Draw suffix boxes
min_w = 50
max_w = 800
x_start = 1.5
x_end = 21.2
total_range = x_end - x_start

y_base = 1.9
box_w = 1.4
box_h = 0.48
stack_gap = 0.55

for weight_idx, (weight, items) in enumerate(sorted(seen_weights.items())):
    x = x_start + (weight - min_w) / (max_w - min_w) * total_range
    color = items[0][1]

    for stack_idx, (suffix, s_color) in enumerate(items):
        y = y_base + stack_idx * stack_gap

        # Shadow
        draw_shadow(ax, x - box_w / 2, y - box_h / 2, box_w, box_h, radius=0.12, offset=0.05, alpha=0.1)

        # Gradient box
        draw_gradient_box(ax, x - box_w / 2, y - box_h / 2, box_w, box_h, s_color, radius=0.12)
        ax.text(x, y, suffix, ha="center", va="center", fontsize=10, fontweight="bold", color="white", zorder=7)

        # Connection line to axis (dashed)
        ax.plot([x, x], [1.2, y - box_h / 2 - 0.05], color=s_color, lw=1.5, alpha=0.4, zorder=2, linestyle="--")

    # Weight label on axis (modern style)
    ax.text(x, 0.65, str(weight), ha="center", va="center", fontsize=9, color="#475569", fontweight="bold")
    # Tick mark
    ax.plot([x, x], [1.05, 1.35], color="#94a3b8", lw=2, zorder=3, solid_capstyle='round')

# Legend — modern pill-style badges
legend_items = [
    ("Pre-release", "#64748b"),
    ("Alpha/Beta", "#ef4444"),
    ("RC", "#06b6d4"),
    ("Stable", "#10b981"),
    ("Post-release", "#0ea5e9"),
    ("Post-stable", "#3b82f6"),
]
legend_x = 1.5
legend_y = -0.6
pill_w = 1.6
pill_h = 0.35
pill_gap = 0.4

for i, (label, color) in enumerate(legend_items):
    lx = legend_x + i * (pill_w + pill_gap)

    # Colored pill
    pill = FancyBboxPatch(
        (lx, legend_y - pill_h / 2), pill_w, pill_h,
        boxstyle="round,pad=0,rounding_size=0.17",
        facecolor=color, edgecolor="none", alpha=0.85, zorder=5
    )
    ax.add_patch(pill)
    ax.text(lx + pill_w / 2, legend_y, label, ha="center", va="center", fontsize=9, color="white", zorder=6, fontweight="medium")

# Example box — elegant card style
example_y = -1.5
example_box = FancyBboxPatch(
    (4, example_y - 0.35), 14, 0.7,
    boxstyle="round,pad=0,rounding_size=0.15",
    facecolor="#f8fafc", edgecolor="#e2e8f0", lw=1.5, zorder=5
)
ax.add_patch(example_box)
ax.text(11, example_y, 'Example: 1.0.0-alpha < 1.0.0-beta < 1.0.0-rc < 1.0.0 < 1.0.0-sp1',
    ha="center", va="center", fontsize=11, color="#475569", fontweight="medium", zorder=6)

# Decorative corner accents
corner_size = 0.8
corner_color = "#e2e8f0"
ax.plot([-0.5, -0.5], [5 - corner_size, 5], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([-0.5, -0.5 + corner_size], [5, 5], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([22.5, 22.5], [-1 + corner_size, -1], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([22.5, 22.5 - corner_size], [-1, -1], color=corner_color, lw=3, solid_capstyle='round')

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/suffix-weight.png", dpi=150, bbox_inches="tight", facecolor="#fafbfc")
plt.close()
print("suffix-weight.png saved")