#!/usr/bin/env python3
"""
Generate a flat-style access methods hub-spoke diagram — refined.
Color palette: blue-gray, no purple, square corners, generous card spacing.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import math

# ─── Data ───
center_label = "versions-skills"
methods = [
    ("Skills", "13 Claude Code Skills", "SKILL.md files", "Auto-discovered by AI", "#2563eb"),
    ("MCP Server", "Model Context Protocol", "SSE transport layer", "Tool-based access", "#0ea5e9"),
    ("Go SDK", "Go package import", "Type-safe full API", "Composable functions", "#16a34a"),
    ("CLI", "Cobra binary", "Shell / CI/CD ready", "Pipe-friendly I/O", "#ea580c"),
]

# ─── Drawing ───
fig, ax = plt.subplots(figsize=(16, 14))
ax.set_xlim(-8, 8)
ax.set_ylim(-9, 9)
ax.axis("off")
fig.patch.set_facecolor("white")
ax.set_aspect("equal")

# Title
ax.text(0, 8.3, "Access Methods", ha="center", va="center", fontsize=22, fontweight="bold", color="#0f172a")
ax.text(0, 7.6, "4 ways to integrate versions-skills into your workflow", ha="center", va="center", fontsize=11, color="#64748b")

# Center hub
center_w, center_h = 3.4, 1.1
hub = FancyBboxPatch((-center_w / 2, -center_h / 2), center_w, center_h,
    boxstyle="square,pad=0.04", facecolor="#0f172a", edgecolor="#0f172a", lw=1.5, zorder=5)
ax.add_patch(hub)
ax.text(0, 0, center_label, ha="center", va="center", fontsize=13, fontweight="bold", color="white", zorder=6, family="monospace")

# Spoke cards — spread with more space
radius = 5.2
card_w = 4.0
card_h = 3.2
angles = [90, 0, 270, 180]  # top, right, bottom, left

for i, (label, line1, line2, line3, color) in enumerate(methods):
    angle_rad = math.radians(angles[i])
    cx = radius * math.cos(angle_rad)
    cy = radius * math.sin(angle_rad)

    # Connection line (Manhattan style)
    if angles[i] == 90:  # top
        ax.plot([0, 0], [0.55, cy - card_h / 2 + 0.05], color=color, lw=2.5, alpha=0.35, zorder=2)
    elif angles[i] == 0:  # right
        ax.plot([0.0, cx - card_w / 2 + 0.05], [0, 0], color=color, lw=2.5, alpha=0.35, zorder=2)
    elif angles[i] == 270:  # bottom
        ax.plot([0, 0], [-0.55, cy + card_h / 2 - 0.05], color=color, lw=2.5, alpha=0.35, zorder=2)
    elif angles[i] == 180:  # left
        ax.plot([0.0, cx + card_w / 2 - 0.05], [0, 0], color=color, lw=2.5, alpha=0.35, zorder=2)

    # Card background
    card_bg = FancyBboxPatch((cx - card_w / 2, cy - card_h / 2), card_w, card_h,
        boxstyle="square,pad=0.05", facecolor=color, edgecolor=color, lw=1.5, alpha=0.06, zorder=3)
    ax.add_patch(card_bg)

    # Card border
    card_border = FancyBboxPatch((cx - card_w / 2, cy - card_h / 2), card_w, card_h,
        boxstyle="square,pad=0.05", facecolor="none", edgecolor=color, lw=1.5, zorder=4)
    ax.add_patch(card_border)

    # Card header
    header_h = 0.6
    header = FancyBboxPatch((cx - card_w / 2, cy + card_h / 2 - header_h), card_w, header_h,
        boxstyle="square,pad=0", facecolor=color, edgecolor="none", lw=0, alpha=0.92, zorder=5)
    ax.add_patch(header)
    ax.text(cx, cy + card_h / 2 - header_h / 2, label, ha="center", va="center",
        fontsize=12, fontweight="bold", color="white", zorder=6)

    # Card description — one line per feature
    desc_lines = [line1, line2, line3]
    for j, line in enumerate(desc_lines):
        ly = cy + card_h / 2 - header_h - 0.45 - j * 0.55
        ax.text(cx, ly, line, ha="center", va="center", fontsize=10, color="#334155", zorder=6)

# Combo note
combo_box = FancyBboxPatch((-3.5, -8.2), 7.0, 0.6,
    boxstyle="square,pad=0.04", facecolor="#2563eb", edgecolor="#2563eb", lw=0, alpha=0.08, zorder=3)
ax.add_patch(combo_box)
ax.text(0, -7.9, "Recommended: Skills + MCP = maximum AI agent integration", ha="center", va="center",
    fontsize=10, fontweight="bold", color="#2563eb", zorder=4)

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/access-methods.png", dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("access-methods.png saved")
