#!/usr/bin/env python3
"""
Generate a refined access methods hub-spoke diagram with modern styling.
Features: rounded corners, gradient fills, shadows, elegant typography.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch, Circle
import math

# ─── Data ───
center_label = "versions-skills"
methods = [
    ("Skills", "13 Claude Code Skills", "SKILL.md files", "Auto-discovered by AI", "#3b82f6"),
    ("MCP Server", "Model Context Protocol", "SSE transport layer", "Tool-based access", "#06b6d4"),
    ("Go SDK", "Go package import", "Type-safe full API", "Composable functions", "#10b981"),
    ("CLI", "Cobra binary", "Shell / CI/CD ready", "Pipe-friendly I/O", "#f97316"),
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
fig, ax = plt.subplots(figsize=(18, 16))
ax.set_xlim(-9, 9)
ax.set_ylim(-10, 10)
ax.axis("off")
fig.patch.set_facecolor("#fafbfc")
ax.set_aspect("equal")

# Background grid
for i in range(-9, 10, 2):
    ax.axvline(i, color="#f1f5f9", lw=0.5, zorder=0)
for i in range(-10, 11, 2):
    ax.axhline(i, color="#f1f5f9", lw=0.5, zorder=0)

# Title
ax.text(0, 9.2, "Access Methods", ha="center", va="center", fontsize=26, fontweight="bold", color="#0f172a", family="sans-serif")
ax.text(0, 8.4, "4 ways to integrate versions-skills into your workflow", ha="center", va="center", fontsize=12, color="#64748b", family="sans-serif")

# Center hub — elegant dark with glow effect
center_w, center_h = 3.8, 1.2

# Glow effect (multiple shadows)
for i, alpha in enumerate([0.05, 0.03, 0.02]):
    glow = FancyBboxPatch(
        (-center_w / 2 - 0.15 * (i + 1), -center_h / 2 - 0.15 * (i + 1)),
        center_w + 0.3 * (i + 1), center_h + 0.3 * (i + 1),
        boxstyle="round,pad=0,rounding_size=0.25",
        facecolor="#3b82f6", edgecolor="none", alpha=alpha, zorder=3
    )
    ax.add_patch(glow)

draw_shadow(ax, -center_w / 2, -center_h / 2, center_w, center_h, radius=0.2, offset=0.1, alpha=0.15)
draw_gradient_box(ax, -center_w / 2, -center_h / 2, center_w, center_h, "#0f172a", radius=0.2)
ax.text(0, 0, center_label, ha="center", va="center", fontsize=14, fontweight="bold", color="white", zorder=7, family="monospace")

# Spoke cards — radial layout with modern cards
radius = 5.8
card_w = 4.4
card_h = 3.6
angles = [90, 0, 270, 180]  # top, right, bottom, left

for i, (label, line1, line2, line3, color) in enumerate(methods):
    angle_rad = math.radians(angles[i])
    cx = radius * math.cos(angle_rad)
    cy = radius * math.sin(angle_rad)

    # Connection line — smooth curve with glow
    if angles[i] == 90:  # top
        ax.plot([0, 0], [0.6, cy - card_h / 2 + 0.1], color=color, lw=3, alpha=0.4, zorder=2, solid_capstyle='round')
    elif angles[i] == 0:  # right
        ax.plot([0.4, cx - card_w / 2 + 0.1], [0, 0], color=color, lw=3, alpha=0.4, zorder=2, solid_capstyle='round')
    elif angles[i] == 270:  # bottom
        ax.plot([0, 0], [-0.6, cy + card_h / 2 - 0.1], color=color, lw=3, alpha=0.4, zorder=2, solid_capstyle='round')
    elif angles[i] == 180:  # left
        ax.plot([-0.4, cx + card_w / 2 - 0.1], [0, 0], color=color, lw=3, alpha=0.4, zorder=2, solid_capstyle='round')

    # Card background — subtle fill
    draw_shadow(ax, cx - card_w / 2, cy - card_h / 2, card_w, card_h, radius=0.2, offset=0.08, alpha=0.1)

    card_bg = FancyBboxPatch(
        (cx - card_w / 2, cy - card_h / 2), card_w, card_h,
        boxstyle="round,pad=0,rounding_size=0.2",
        facecolor=color, edgecolor=color, lw=1.5, alpha=0.04, zorder=4
    )
    ax.add_patch(card_bg)

    # Card border
    card_border = FancyBboxPatch(
        (cx - card_w / 2, cy - card_h / 2), card_w, card_h,
        boxstyle="round,pad=0,rounding_size=0.2",
        facecolor="none", edgecolor=color, lw=2, zorder=5
    )
    ax.add_patch(card_border)

    # Card header — gradient effect
    header_h = 0.7
    draw_gradient_box(ax, cx - card_w / 2, cy + card_h / 2 - header_h, card_w, header_h, color, radius=0.15)
    ax.text(cx, cy + card_h / 2 - header_h / 2, label, ha="center", va="center",
        fontsize=13, fontweight="bold", color="white", zorder=7, family="sans-serif")

    # Card description — clean typography
    desc_lines = [line1, line2, line3]
    for j, line in enumerate(desc_lines):
        ly = cy + card_h / 2 - header_h - 0.55 - j * 0.6
        ax.text(cx, ly, line, ha="center", va="center", fontsize=10, color="#334155", zorder=7, family="sans-serif", weight="medium")

# Combo recommendation — elegant card
combo_w = 8.0
combo_h = 0.8
draw_shadow(ax, -combo_w / 2, -9.2, combo_w, combo_h, radius=0.15, offset=0.05, alpha=0.08)
combo_box = FancyBboxPatch(
    (-combo_w / 2, -9.2), combo_w, combo_h,
    boxstyle="round,pad=0,rounding_size=0.15",
    facecolor="#3b82f6", edgecolor="none", alpha=0.08, zorder=4
)
ax.add_patch(combo_box)
ax.text(0, -8.8, "Recommended: Skills + MCP = maximum AI agent integration", ha="center", va="center",
    fontsize=11, fontweight="bold", color="#3b82f6", zorder=5, family="sans-serif")

# Decorative corner accents
corner_size = 1.0
corner_color = "#e2e8f0"
ax.plot([-8.5, -8.5], [9 - corner_size, 9], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([-8.5, -8.5 + corner_size], [9, 9], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([8.5, 8.5], [-9 + corner_size, -9], color=corner_color, lw=3, solid_capstyle='round')
ax.plot([8.5, 8.5 - corner_size], [-9, -9], color=corner_color, lw=3, solid_capstyle='round')

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/access-methods.png", dpi=150, bbox_inches="tight", facecolor="#fafbfc")
plt.close()
print("access-methods.png saved")