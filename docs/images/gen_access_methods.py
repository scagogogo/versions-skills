#!/usr/bin/env python3
"""
Generate a diagram showing the 4 access methods (Skills, CLI, SDK, MCP)
and their relationship to the core capabilities.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch
import numpy as np

fig, ax = plt.subplots(figsize=(16, 12))
ax.set_xlim(0, 16)
ax.set_ylim(0, 12)
ax.axis("off")
fig.patch.set_facecolor("white")

# Central hub
cx, cy = 8, 6
hub = plt.Circle((cx, cy), 1.8, facecolor="#1e293b", edgecolor="#0f172a", lw=3, zorder=5)
ax.add_patch(hub)
ax.text(cx, cy + 0.4, "versions-skills", ha="center", va="center",
        fontsize=13, fontweight="bold", color="white", zorder=6, family="monospace")
ax.text(cx, cy - 0.3, "Core Library", ha="center", va="center",
        fontsize=10, color="#94a3b8", zorder=6)

# 4 Access methods
access_methods = [
    {
        "x": 3.5, "y": 10, "color": "#8b5cf6", "edge": "#6d28d9",
        "title": "SKILLS",
        "desc": "Claude Code Plugin",
        "detail": "13 SKILL.md files\nslash commands\ndomain knowledge",
    },
    {
        "x": 12.5, "y": 10, "color": "#3b82f6", "edge": "#1d4ed8",
        "title": "MCP SERVER",
        "desc": "AI Agent Protocol",
        "detail": "21 version_* tools\nJSON responses\nany MCP client",
    },
    {
        "x": 3.5, "y": 2, "color": "#10b981", "edge": "#059669",
        "title": "CLI",
        "desc": "Shell & CI/CD",
        "detail": "25+ commands\npipeline friendly\ncross-platform",
    },
    {
        "x": 12.5, "y": 2, "color": "#f59e0b", "edge": "#d97706",
        "title": "Go SDK",
        "desc": "Go Programs",
        "detail": "Full API\ntype-safe\nzero dependencies",
    },
]

for am in access_methods:
    w, h = 4.0, 2.8
    card = FancyBboxPatch(
        (am["x"] - w/2, am["y"] - h/2), w, h,
        boxstyle="round,pad=0.15", facecolor=am["color"],
        edgecolor=am["edge"], lw=2, alpha=0.9, zorder=4
    )
    ax.add_patch(card)

    ax.text(am["x"], am["y"] + 0.85, am["title"], ha="center", va="center",
            fontsize=14, fontweight="bold", color="white", zorder=5, family="monospace")
    ax.text(am["x"], am["y"] + 0.3, am["desc"], ha="center", va="center",
            fontsize=9, color="white", alpha=0.8, zorder=5)
    ax.text(am["x"], am["y"] - 0.55, am["detail"], ha="center", va="center",
            fontsize=8, color="white", alpha=0.75, zorder=5, linespacing=1.3)

    # Connection line to center
    ax.annotate(
        "", xy=(cx, cy), xytext=(am["x"], am["y"]),
        arrowprops=dict(arrowstyle="-|>", color=am["color"], lw=2.5,
                        alpha=0.5, connectionstyle="arc3,rad=0"),
        zorder=2
    )

# Capabilities ring
capabilities = [
    "Parse", "Compare", "Sort", "Filter", "Group",
    "Constraint", "Range", "Check", "Mutate", "Build",
    "File I/O", "Visualize", "Serialize", "Set Ops",
]

n = len(capabilities)
radius = 3.5
for i, cap in enumerate(capabilities):
    angle = 2 * np.pi * i / n - np.pi / 2
    x = cx + radius * np.cos(angle)
    y = cy + radius * np.sin(angle)

    dot = plt.Circle((x, y), 0.35, facecolor="#f1f5f9", edgecolor="#94a3b8",
                      lw=1.5, zorder=3)
    ax.add_patch(dot)
    ax.text(x, y, cap, ha="center", va="center", fontsize=7,
            fontweight="bold", color="#475569", zorder=4)

    ax.plot([cx, x], [cy, y], color="#e2e8f0", lw=1, zorder=1)

# Title
ax.text(8, 11.5, "versions-skills  --  Access Methods & Capabilities", ha="center",
        va="center", fontsize=18, fontweight="bold", color="#1e293b", family="monospace")

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/access-methods.png",
            dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("access-methods.png saved")
