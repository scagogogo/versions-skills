#!/usr/bin/env python3
"""
Generate a layered architecture diagram for versions-skills.
No emoji - uses colored labels.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
from matplotlib.patches import FancyBboxPatch
import numpy as np

fig, ax = plt.subplots(figsize=(18, 10))
ax.set_xlim(0, 18)
ax.set_ylim(0, 10)
ax.axis("off")
fig.patch.set_facecolor("white")

layers = [
    {
        "y": 7.8, "height": 1.4,
        "color": "#8b5cf6", "edge": "#6d28d9",
        "title": "AI Agent Layer",
        "subtitle": "Claude Code / Cursor / Windsurf / VS Code Copilot",
        "items": [
            ("Skills Plugin\n13 SKILL.md files\nslash commands", 4.5),
            ("MCP Server\n21 version_* tools\nJSON responses", 13.5),
        ],
    },
    {
        "y": 5.5, "height": 1.4,
        "color": "#3b82f6", "edge": "#1d4ed8",
        "title": "Interface Layer",
        "subtitle": "Human & Program Access Points",
        "items": [
            ("CLI Binary\n25+ commands\nShell / CI/CD", 4.5),
            ("Go SDK\nFull API\nGo Programs", 13.5),
        ],
    },
    {
        "y": 3.2, "height": 1.4,
        "color": "#10b981", "edge": "#059669",
        "title": "Feature Layer",
        "subtitle": "Core Capabilities",
        "items": [
            ("Parse\nValidate", 1.8),
            ("Compare\nDiff", 3.8),
            ("Sort\nFilter", 5.8),
            ("Group\nPartition", 7.8),
            ("Constraint\nRange", 9.8),
            ("Type Check\nMutate", 11.8),
            ("File I/O\nVisualize", 13.8),
            ("Serialize\nSet Ops", 15.8),
        ],
    },
    {
        "y": 1.0, "height": 1.4,
        "color": "#f59e0b", "edge": "#d97706",
        "title": "Core Library",
        "subtitle": "Go - Zero External Dependencies - Immutable Design",
        "items": [
            ("Version | VersionNumbers | VersionRange | Constraint | VersionBuilder | VersionSlice", 9.0),
        ],
    },
]

for layer in layers:
    # Background bar
    bar = FancyBboxPatch(
        (0.5, layer["y"]), 17, layer["height"],
        boxstyle="round,pad=0.15", facecolor=layer["color"], edgecolor=layer["edge"],
        lw=2, alpha=0.12, zorder=1
    )
    ax.add_patch(bar)

    # Layer title
    ax.text(0.8, layer["y"] + layer["height"] - 0.15, layer["title"],
            fontsize=12, fontweight="bold", color=layer["edge"], zorder=3)

    # Subtitle
    ax.text(0.8, layer["y"] + 0.15, layer["subtitle"],
            fontsize=8.5, color="#64748b", zorder=3)

    # Items
    for label, cx in layer["items"]:
        if len(layer["items"]) > 2:
            w, h = 1.6, 0.8
        elif len(layer["items"]) == 2:
            w, h = 3.5, 0.85
        else:
            w, h = 15, 0.7

        cy = layer["y"] + layer["height"] / 2
        item_box = FancyBboxPatch(
            (cx - w/2, cy - h/2), w, h,
            boxstyle="round,pad=0.08", facecolor=layer["color"],
            edgecolor=layer["edge"], lw=1.5, alpha=0.85, zorder=4
        )
        ax.add_patch(item_box)
        ax.text(cx, cy, label, ha="center", va="center",
                fontsize=7.5 if len(layer["items"]) > 2 else 9,
                fontweight="bold", color="white", zorder=5, linespacing=1.2)

# Arrows between layers
arrow_props = dict(arrowstyle="->", color="#94a3b8", lw=2, connectionstyle="arc3,rad=0")
for i in range(len(layers) - 1):
    y_from = layers[i]["y"]
    y_to = layers[i+1]["y"] + layers[i+1]["height"]
    ax.annotate("", xy=(9, y_to), xytext=(9, y_from), arrowprops=arrow_props)

# Title
ax.text(9, 9.7, "versions-skills  --  Architecture", ha="center", va="center",
        fontsize=20, fontweight="bold", color="#1e293b", family="monospace")

plt.tight_layout()
plt.savefig("/home/cc11001100/github/scagogogo/versions-skills/docs/images/architecture.png",
            dpi=150, bbox_inches="tight", facecolor="white")
plt.close()
print("architecture.png saved")
